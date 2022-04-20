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

type InfoHandler struct {
	Usecase domain.InfoUsecase
	conf    *config.Config
}

func NewInfoHandler(ctx context.Context, r *gin.RouterGroup, usecase domain.InfoUsecase, conf *config.Config) {
	handler := &InfoHandler{
		Usecase: usecase,
		conf:    conf,
	}
	r.GET("/bulletin", handler.Get(ctx))
	r.POST("/:info_id/bulletin", mw.VerifyAuth(ctx, conf.Jwt.Secure), handler.Create(ctx))
	r.PUT("/:info_id/bulletin/:bulletin_id", mw.VerifyAuth(ctx, conf.Jwt.Secure), handler.Update(ctx))
	r.DELETE("/:info_id/bulletin/:bulletin_id", mw.VerifyAuth(ctx, conf.Jwt.Secure), handler.Delete(ctx))
	// bulletin := r.Group("/:info_id/bulletin", mw.VerifyAuth(ctx, conf.Jwt.Secure))
	// {
	// 	bulletin.POST("/", handler.Create(ctx))
	// 	bulletin.PUT("/:bulletin_id", handler.Update(ctx))
	// 	bulletin.DELETE("/:bulletin_id", handler.Delete(ctx))
	// }
}

func (i *InfoHandler) Create(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bind domain.CreateInfoBulletinRequest
		// bind teacher domain
		if err := c.BindUri(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		// bind bulletin content
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
		// bind teacher domain and last_modified
		if err := c.ShouldBindUri(&bind); err != nil {
			// if err := c.ShouldBindUri(&bind); err != nil {
			fmt.Println("uri, ", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.ShouldBindQuery(&bind)
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
		// bind teacher_domain, info_id and bulletin_id
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
		res, err := i.Usecase.Update(ctx, &bind)
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

func (i *InfoHandler) Delete(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bind domain.DeleteInfoBulletinRequest
		// bind teacher_domain, info_id and bulletin_id
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
