package route

import (
	"context"
	"net/http"
	"teacher-site/model"
	"teacher-site/service"

	"github.com/gin-gonic/gin"
)

func Login(ctx context.Context, srv service.Servicer, cfg *model.Config) gin.HandlerFunc {
	var bind *model.BindAuth
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(bind); err != nil {
			srv.Debug(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		token, err := srv.Login(ctx, bind, cfg)
		if err != nil {
			srv.Debug(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		addBearerHeader(c, token)
		c.Status(http.StatusOK)
	}
}
func addBearerHeader(c *gin.Context, token string) {
	c.Request.Header.Add("Authorization", `Bearer `+token)
}
func Register(ctx context.Context, srv service.Servicer, cfg *model.Config) gin.HandlerFunc {
	var bind *model.BindRegister
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(bind); err != nil {
			srv.Debug(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if ok := srv.Register(ctx, bind, cfg); ok {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// todo: logout
// func Logout(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 	}
// }
