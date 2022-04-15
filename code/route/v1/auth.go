package v1

import (
	"context"
	"net/http"
	"teacher-site/model"
	"teacher-site/service"

	"github.com/gin-gonic/gin"
)

func Login(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	var bind model.BindAuth
	return func(c *gin.Context) {
		var (
			token string
			err   error
		)
		if err = c.ShouldBindJSON(&bind); err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err = srv.Login(ctx, &bind); err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		token, err = srv.NewJwtToken(ctx, &bind)
		if err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if err = srv.UpdateJwtToken(ctx, token, &bind); err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		addBearerHeader(c, token)
		c.Status(http.StatusOK)
	}
}
func addBearerHeader(c *gin.Context, token string) {
	c.Request.Header.Add("Authorization", `Bearer `+token)
}
func Register(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	var bind model.BindRegister
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(&bind); err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := srv.Register(ctx, &bind); err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusNotImplemented)
			return
		}
		c.AbortWithStatus(http.StatusOK)
	}
}

// todo: logout
// func Logout(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 	}
// }
