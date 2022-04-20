package util

import (
	"teacher-site/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwt(conf *config.Jwt, userID string) (string, error) {
	exp := time.Now().Add(conf.TokenExpireTime * time.Minute).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userID,
		"exp":  exp,
	})
	token, err := claims.SignedString(conf.Secure)
	return token, err
}
