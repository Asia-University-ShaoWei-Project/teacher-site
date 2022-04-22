package middleware

import (
	"context"
	"fmt"
	"net/http"
	"teacher-site/pkg/util"

	"github.com/gin-gonic/gin"
)

func IsTeacher(ctx context.Context, c *gin.Context, secret []byte) bool {
	bearerToken, err := util.GetBearerToken(ctx, c)
	if err != nil {
		fmt.Println("not have auth header")
		return false
	}
	if err := verifyJwtValid(ctx, bearerToken, secret); err != nil {
		fmt.Println("not a teacher")

		return false
	}
	fmt.Println("is a teacher")

	return true
}

func VerifyAuth(ctx context.Context, secret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {

		bearerToken, err := util.GetBearerToken(ctx, c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err := verifyJwtValid(ctx, bearerToken, secret); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func verifyJwtValid(ctx context.Context, bearerToken string, secret []byte) error {
	_, err := util.ParseJwt(ctx, bearerToken, secret)
	if err != nil {
		return err
	}
	return nil
}
