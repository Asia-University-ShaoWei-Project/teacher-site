package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var errInvalidBearerToken = errors.New("invalid bearer token")

func IsTeacher(ctx context.Context, c *gin.Context, secure []byte) bool {
	authHeader := getAuthorization(ctx, c)
	if err := verifyJwtValid(ctx, secure, authHeader); err != nil {
		return false
	}
	return true
}
func getAuthorization(ctx context.Context, c *gin.Context) string {
	return c.GetHeader("Authorization")
}

func VerifyAuth(ctx context.Context, secure []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := getAuthorization(ctx, c)
		if err := verifyJwtValid(ctx, secure, authHeader); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func verifyJwtValid(ctx context.Context, secure []byte, authHeader string) error {

	bearerToken, err := getBearerToken(ctx, authHeader)
	if err != nil {
		return err
	}
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secure, nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// expire example: https://github.com/golang-jwt/jwt/blob/main/map_claims_test.go
		// todo: expire handle! claims["exp"]
		fmt.Println(claims["exp"])
		return nil
	}
	return jwt.ErrInvalidKey
}

func getBearerToken(ctx context.Context, authHeader string) (string, error) {
	field := strings.Split(authHeader, " ")
	if len(field) != 2 {
		return "", errInvalidBearerToken
	}
	return field[1], nil
}
