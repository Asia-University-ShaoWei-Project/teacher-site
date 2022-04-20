package middleware

import (
	"net/http"
	"teacher-site/domain"

	"github.com/gin-gonic/gin"
)

func CheckTeacherDomain() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindUri(&domain.TeacherDomainRequest{}); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}
}
