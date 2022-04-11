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
		var bindDomain model.BindDomain
		if err := c.ShouldBindUri(&bindDomain); err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		err := srv.DomainIsExist(ctx, bindDomain.Domain)
		if err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		srv.SetDomain(ctx, bindDomain.Domain)
	}
}
