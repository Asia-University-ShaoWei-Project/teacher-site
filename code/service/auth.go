package service

import (
	"context"
	"crypto/rand"
	"errors"
	"teacher-site/model"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	errAuthNotMatch = errors.New("fail to authorization")
)

//ref: https://github.com/golang-jwt/jwt/blob/main/example_test.go
func (srv *Service) LoginAndGetNewToken(ctx context.Context, bindAuth *model.BindAuth) (string, error) {
	var err error
	auth := &model.Auths{UserID: bindAuth.UserID}
	if err = srv.db.GetAuth(ctx, auth); err != nil {
		srv.log.Error(err)
		return "", err
	}
	saltPassword := []byte(bindAuth.UserPassword + auth.Salt)
	if bcrypt.CompareHashAndPassword([]byte(auth.UserPassword), saltPassword) != nil {
		srv.log.Error("The password not match the user password")
		return "", errAuthNotMatch
	}
	token, err := newToken(ctx, auth.Token, srv.conf.TokenExpireTime, srv.conf.JWTSecure)
	if err != nil {
		srv.log.Error(err)
		return "", err
	}
	if err = srv.db.UpdateUserToken(ctx, auth, token); err != nil {
		srv.log.Error(err)
		return "", err
	}
	if err = srv.cache.SetToken(srv.domain, token); err != nil {
		srv.log.Error(err)
		// todo: handle this
	}
	return token, nil
}

func newToken(ctx context.Context, userToken string, duration time.Duration, secure []byte) (string, error) {
	exp := time.Now().Add(duration * time.Minute).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_token": userToken,
		"exp":        exp,
	})
	token, err := claims.SignedString(secure)
	return token, err
}

// bcrypt: https://github.com/golang/crypto/blob/master/bcrypt/bcrypt_test.go
func (srv *Service) Register(ctx context.Context, bindRegister *model.BindRegister) bool {
	//todo: check user_id and domain exist
	var err error
	salt, err := generateSalt(srv.conf.SaltSize)
	if err != nil {
		srv.Error(err)
		return false
	}
	saltPassword := append([]byte(bindRegister.UserPassword), salt...)
	hashPassword, err := bcrypt.GenerateFromPassword(saltPassword, srv.conf.HashCost)
	if err != nil {
		srv.log.Error(err)
		return false
	}
	token, err := uuid.NewUUID()
	if err != nil {
		srv.Error(err)
		return false
	}
	auth := &model.Auths{
		UserID:       bindRegister.UserID,
		UserPassword: string(hashPassword),
		Token:        token.String(),
		Salt:         string(salt),
		Teacher: model.Teachers{
			Domain: bindRegister.Domain,
			NameZH: bindRegister.NameZH,
			Email:  bindRegister.Email,
		},
	}
	srv.db.CreateUser(ctx, auth)
	return true
}
func generateSalt(saltSize int) ([]byte, error) {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])
	return salt, err
}
func (srv *Service) GetJWTSecure(ctx context.Context) []byte {
	return srv.conf.JWTSecure
}

// todo: logout
