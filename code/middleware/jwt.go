package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"teacher-site/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var errInvalidBearerToken = errors.New("invalid bearer token")

func VerifyJWT(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken, err := getBearerToken(ctx, c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return srv.GetJWTSecure, nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// expire example: https://github.com/golang-jwt/jwt/blob/main/map_claims_test.go
			// todo: expire handle! claims["exp"]
			fmt.Println(claims["expire"])
			return
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
func getBearerToken(ctx context.Context, c *gin.Context) (string, error) {
	auth := c.GetHeader("Authorization")
	field := strings.Split(auth, " ")
	if len(field) != 2 {
		return "", errInvalidBearerToken
	}
	return field[1], nil
}
