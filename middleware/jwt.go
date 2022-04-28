package middleware

import (
	"context"
	"net/http"
	"teacher-site/domain"
	"teacher-site/pkg/util"

	"github.com/gin-gonic/gin"
)

func IsTeacher(ctx context.Context, c *gin.Context, secret []byte) bool {
	bearerToken, err := util.GetBearerToken(ctx, c)
	if err != nil {
		return false
	}

	if _, err := util.ParseJwt(ctx, bearerToken, secret); err != nil {
		return false
	}
	return true
}

func VerifyAuth(ctx context.Context, secret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.TeacherDomainRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		bearerToken, err := util.GetBearerToken(ctx, c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, err := util.ParseJwt(ctx, bearerToken, secret)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		teacherDomain := util.GetJwtUserDomain(claims)
		if req.TeacherDomain != teacherDomain {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}
}
