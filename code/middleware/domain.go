package middleware

import (
	"context"
	"net/http"
	"teacher-site/model"
	"teacher-site/service"

	"github.com/gin-gonic/gin"
)

func SetupServiceDomain(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var domain *model.BindDomain
		if err := c.ShouldBindUri(&domain); err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := srv.IsExistDomain(ctx, domain.Domain); err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		srv.SetDomain(ctx, domain.Domain)
	}
}
