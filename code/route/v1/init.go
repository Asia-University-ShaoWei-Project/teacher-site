package v1

import (
	"context"
	"net/http"
	mw "teacher-site/middleware"
	"teacher-site/model"
	"teacher-site/service"

	"github.com/gin-gonic/gin"
)

func GetInit(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var init model.Init
		err := srv.GetInit(ctx, &init)
		if err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		auth := mw.IsTeacher(ctx, c, srv)
		c.JSON(http.StatusOK, gin.H{"auth": auth, "data": init})
	}
}
