package delivery

import (
	"context"
	"fmt"
	"net/http"
	"teacher-site/config"
	"teacher-site/domain"
	mw "teacher-site/middleware"
	"teacher-site/pkg/message"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Usecase domain.InfoUsecase
	conf    *config.Config
}

func NewHandler(ctx context.Context, r *gin.RouterGroup, usecase domain.InfoUsecase, conf *config.Config) {
	handler := &Handler{
		Usecase: usecase,
		conf:    conf,
	}
	r.GET("/bulletin", handler.Get(ctx))
	auth := r.Group("/:infoId/bulletin", mw.VerifyAuth(ctx, conf.Jwt.Secret))
	{
		auth.POST("", handler.Create(ctx))
		bulletin := auth.Group("/:bulletinId")
		{
			bulletin.PUT("", handler.Update(ctx))
			bulletin.DELETE("", handler.Delete(ctx))
		}
	}
}

func (h *Handler) Create(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.CreateInfoBulletinRequest
		// bind teacher domain and info id
		if err := c.ShouldBindUri(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		// bind bulletin content
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		res, err := h.Usecase.Create(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data": res,
		})
	}
}
func (h *Handler) Get(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.GetInfoBulletinRequest
		// bind teacher domain and last_modified
		if err := c.ShouldBindUri(&req); err != nil {
			// if err := c.ShouldBindUri(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.ShouldBindQuery(&req)
		res, err := h.Usecase.Get(ctx, &req)
		if err == message.ErrUnnecessaryUpdate {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, &gin.H{
			"data": res,
		})
	}
}

func (h *Handler) Update(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uriReq domain.UpdateInfoBulletinUriRequest
			req    domain.UpdateInfoBulletinRequest
		)
		if err := c.ShouldBindUri(&uriReq); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		req.SetupUri(&uriReq)

		res, err := h.Usecase.Update(ctx, &req)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		// todo: add a 409 status code(try again)
		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}

func (h *Handler) Delete(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {

		var bind domain.DeleteInfoBulletinRequest
		// bind teacherDomain, infoId and bulletinId
		if err := c.ShouldBindUri(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.Delete(ctx, &bind)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}
