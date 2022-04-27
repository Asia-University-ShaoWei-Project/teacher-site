package util

import (
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeaderKey = "Authorization"
)

var errInvalidBearerToken = errors.New("invalid bearer token")

func AddBearerHeader(c *gin.Context, token string) {
	c.Header(authHeaderKey, "Bearer "+token)
}

func DeleteAuthHeader(c *gin.Context) {
	c.Header(authHeaderKey, "")
}

func GetBearerToken(ctx context.Context, c *gin.Context) (string, error) {
	auth := c.GetHeader(authHeaderKey)

	field := strings.Split(auth, " ")
	if len(field) != 2 {
		return "", errInvalidBearerToken
	}
	return field[1], nil
}
