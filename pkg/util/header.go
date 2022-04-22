package util

import (
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	headerTokenKey = "Authorization"
)

var errInvalidBearerToken = errors.New("invalid bearer token")

func AddBearerHeader(c *gin.Context, token string) {
	c.Header(headerTokenKey, "Bearer "+token)
}

func RemoveBearerHeader(c *gin.Context) {
	c.Header(headerTokenKey, "")
}

func GetBearerToken(ctx context.Context, c *gin.Context) (string, error) {
	auth := c.GetHeader("Authorization")

	field := strings.Split(auth, " ")
	if len(field) != 2 {
		return "", errInvalidBearerToken
	}
	return field[1], nil
}
