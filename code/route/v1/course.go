package v1

import (
	"context"
	"net/http"
	"teacher-site/model"
	"teacher-site/service"

	"github.com/gin-gonic/gin"
)

func CreateCourse(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusCreated)
	}
}
func UpdateCourse(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
func DeleteCourse(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	}
}
func GetCourse(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			bindCourse model.BindCourse
			course     model.Courses
		)
		if err := c.ShouldBindUri(&bindCourse); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		err := srv.GetCourse(ctx, &bindCourse, &course)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, &gin.H{
			"data": course,
		})
	}
}
