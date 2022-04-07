package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"teacher-site/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func VerifyJWT(ctx context.Context, srv service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		field := strings.Split(bearerToken, " ")
		if len(field) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		bearerToken = field[1]
		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return srv.GetJWTSecure, nil
		})
		if err != nil {
			srv.Error(err)
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
