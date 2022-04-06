package v1

import (
	"context"
	"net/http"
	"teacher-site/service"

	"github.com/gin-gonic/gin"
)

func GetInit(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := srv.GetInit(ctx)
		if err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
