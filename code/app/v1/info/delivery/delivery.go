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
	conf    *config.Config
}

func NewInfoHandler(ctx context.Context, r *gin.RouterGroup, usecase domain.InfoUsecase, conf *config.Config) {
	handler := &InfoHandler{
		Usecase: usecase,
		conf:    conf,
	}
	r.GET("/", handler.Get(ctx))
	auth := r.Group("/", mw.VerifyAuth(ctx, conf.Jwt.Secure))
	{
		auth.POST("/", handler.Create(ctx))
		auth.PUT("/:id", handler.Update(ctx))
		auth.DELETE("/:id", handler.Delete(ctx))
	}
}

func (i *InfoHandler) Create(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bind domain.CreateInfoBulletinRequest
		if err := c.ShouldBindUri(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBindJSON(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := i.Usecase.Create(ctx, &bind)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}
func (i *InfoHandler) Get(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bind domain.GetInfoBulletinRequest
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
		var bind domain.UpdateInfoBulletinRequest
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
		// todo: add a 409 status code(try again)
		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}

func (i *InfoHandler) Delete(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bind domain.DeleteInfoBulletinRequest
		if err := c.ShouldBindUri(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := i.Usecase.Delete(ctx, &bind)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}
