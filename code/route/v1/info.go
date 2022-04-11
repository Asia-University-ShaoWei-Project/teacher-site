package v1

import (
	"context"
	"net/http"
	"teacher-site/model"
	"teacher-site/service"

	"github.com/gin-gonic/gin"
)

func CreateInfo(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			infoBind model.BindInfo
			err      error
		)

		if err = c.ShouldBindJSON(infoBind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		err = srv.CreateInfo(ctx, &infoBind)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Status(http.StatusCreated)
	}
}
func UpdateInfo(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			infoBind model.BindInfo
			err      error
		)
		if err = c.ShouldBindJSON(&infoBind); err != nil {
			srv.Info(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		err = srv.UpdateInfo(ctx, &infoBind)
		if err != nil {
			srv.Info(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Status(http.StatusOK)
	}
}

func DeleteInfo(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			infoBind model.BindInfo
			err      error
		)
		if err = c.ShouldBindJSON(&infoBind); err != nil {
			srv.Info(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err = srv.DeleteInfo(ctx, &infoBind); err != nil {
			srv.Info(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Status(http.StatusNoContent)
	}
}
