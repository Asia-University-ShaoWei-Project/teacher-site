package v1

import (
	"context"
	"net/http"
	"teacher-site/service"

	"github.com/gin-gonic/gin"
)

func CreateSlide(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// data := srv.GetInitData(ctx)
		c.Status(http.StatusCreated)
	}
}
func UpdateSlide(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
func DeleteSlide(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	}
}
