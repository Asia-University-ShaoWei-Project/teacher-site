package delivery

import (
	"context"
	"fmt"
	"net/http"
	"teacher-site/config"
	"teacher-site/domain"
	mw "teacher-site/middleware"

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
	// /api/v1/rikki/info/1/bulletin
	r.POST("/:infoId/bulletin", mw.VerifyAuth(ctx, conf.Jwt.Secret), handler.Create(ctx))
	r.PUT("/:infoId/bulletin/:bulletinId", mw.VerifyAuth(ctx, conf.Jwt.Secret), handler.Update(ctx))
	r.DELETE("/:infoId/bulletin/:bulletinId", mw.VerifyAuth(ctx, conf.Jwt.Secret), handler.Delete(ctx))
	// bulletin := r.Group("/:infoId/bulletin", mw.VerifyAuth(ctx, conf.Jwt.Secure))
	// {
	// 	bulletin.POST("/", handler.Create(ctx))
	// 	bulletin.PUT("/:bulletinId", handler.Update(ctx))
	// 	bulletin.DELETE("/:bulletinId", handler.Delete(ctx))
	// }
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
		if err != nil {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.JSON(http.StatusOK, &gin.H{
			"data": res,
		})
	}
}

func (h *Handler) Update(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bind domain.UpdateInfoBulletinRequest
		// bind teacherDomain, infoId and bulletinId
		if err := c.ShouldBindUri(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		// bind bulletin content
		if err := c.ShouldBindJSON(&bind); err != nil {
			fmt.Println(err)

			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.Update(ctx, &bind)
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
