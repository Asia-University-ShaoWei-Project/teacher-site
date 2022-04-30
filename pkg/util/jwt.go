package util

import (
	"context"
	"fmt"
	"teacher-site/config"
	"teacher-site/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	JwtUserKey    = "user"
	JwtUserDomain = "domain"
	JwtExpireKey  = "exp"
)

func GenerateJwt(conf *config.Jwt, req *domain.JwtInfoRequest) (string, error) {
	exp := time.Now().Add(conf.TokenExpireTime * time.Minute).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		JwtUserKey:    req.UserId,
		JwtUserDomain: req.Domain,
		JwtExpireKey:  exp,
	})
	token, err := claims.SignedString(conf.Secret)
	return token, err
}

func ParseJwt(ctx context.Context, bearerToken string, secret []byte) (jwt.MapClaims, error) {
	var claims jwt.MapClaims

	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return claims, err
	}
	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// expire example: https://github.com/golang-jwt/jwt/blob/main/map_claims_test.go
		// todo: expire handle! claims["exp"]
		return claims, nil
	}
	return claims, jwt.ErrInvalidKey
}

func GetJwtUser(claims jwt.MapClaims) string {
	return claims[JwtUserKey].(string)
}
func GetJwtUserDomain(claims jwt.MapClaims) string {
	return claims[JwtUserDomain].(string)
}
