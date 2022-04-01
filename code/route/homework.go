package route

import (
	"context"
	"net/http"
	"teacher-site/service"

	"github.com/gin-gonic/gin"
)

func CreateHomework(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusCreated)
	}
}
func UpdateHomework(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
func DeleteHomework(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	}
}
