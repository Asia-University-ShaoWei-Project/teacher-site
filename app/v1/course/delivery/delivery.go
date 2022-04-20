package delivery

import (
	"context"
	"net/http"
	"teacher-site/config"
	"teacher-site/domain"
	mw "teacher-site/middleware"

	"github.com/gin-gonic/gin"
)

type InfoHandler struct {
	Usecase domain.InfoUsecase
	conf    *config.Jwt
}

func NewInfoHandler(ctx context.Context, r *gin.RouterGroup, usecase domain.InfoUsecase, conf *config.Jwt) {
	handler := &InfoHandler{
		Usecase: usecase,
		conf:    conf,
	}
	info := r.Group("/info")
	{
		info.GET("/", handler.Get(ctx))
		auth := info.Group("/", mw.VerifyAuth(ctx, conf.Secure))
		{
			auth.POST("/", handler.Create(ctx))
			auth.PUT("/:id", handler.Update(ctx))
			auth.DELETE("/:id", handler.Delete(ctx))
		}
	}
}

func (i *InfoHandler) Create(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bind domain.ReqCreateInfo
		if err := c.ShouldBindUri(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBindJSON(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		bulletin, err := i.Usecase.Create(ctx, &bind)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res := domain.ResCreateInfo{
			ID:   bulletin.AutoModel.ID,
			Date: bulletin.AutoModel.CreatedAT.Format(domain.BulletinDateFormat),
		}
		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}
func (i *InfoHandler) Get(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bind domain.ReqGetInfo
		if err := c.ShouldBindUri(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := i.Usecase.Get(ctx, &bind)
		if err != nil {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.JSON(http.StatusOK, &gin.H{
			"data": res,
		})
	}
}

func (i *InfoHandler) Update(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bind domain.ReqUpdateInfoBulletin
		if err := c.ShouldBindUri(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBindJSON(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := i.Usecase.Update(ctx, &bind)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
		c.Status(http.StatusOK)
	}
}

func (i *InfoHandler) Delete(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bind domain.ReqDeleteInfo
		if err := c.ShouldBindUri(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := i.Usecase.Delete(ctx, &bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Status(http.StatusNoContent)
	}
}
