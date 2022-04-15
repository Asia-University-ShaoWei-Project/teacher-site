package v1

import (
	"context"
	"net/http"
	"strconv"
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
			bind model.BindInfo
			err  error
		)
		if err = c.ShouldBindJSON(&bind); err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		bind.ID, err = CovertID(c.Param("id"))
		if err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		err = srv.UpdateInfo(ctx, &bind)
		if err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Status(http.StatusOK)
	}
}

func DeleteInfo(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			bind model.BindInfo
			err  error
		)
		id, err := CovertID(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		bind = model.BindInfo{ID: id}
		if err = srv.DeleteInfo(ctx, &bind); err != nil {
			srv.Info(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func CovertID(id string) (uint, error) {
	_id, err := strconv.ParseUint(id, 10, 32)
	return uint(_id), err
}
