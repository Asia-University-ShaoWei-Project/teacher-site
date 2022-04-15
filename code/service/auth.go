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

var errIsExisted = errors.New("the value was existed")

//ref: https://github.com/golang-jwt/jwt/blob/main/example_test.go
func (srv *Service) Login(ctx context.Context, bind *model.BindAuth) error {
	var err error
	auth := &model.Auths{UserID: bind.UserID}
	if err = srv.db.GetAuth(ctx, auth); err != nil {
		// todo: handle error
		return err
	}
	saltPassword := []byte(bind.UserPassword + auth.Salt)
	if err = bcrypt.CompareHashAndPassword([]byte(auth.UserPassword), saltPassword); err != nil {
		// todo: handle error
		return err
	}
	return nil
}

func (srv *Service) UpdateJwtToken(ctx context.Context, token string, bind *model.BindAuth) error {
	var err error
	auth := &model.Auths{UserID: bind.UserID}
	if err = srv.db.UpdateUserToken(ctx, auth, token); err != nil {
		// todo: error handle
		return err
	}
	if err = srv.cache.SetTokenWithDomain(ctx, token, srv.domain); err != nil {
		// todo: error handle
		return err
	}
	return nil

}
func (srv *Service) NewJwtToken(ctx context.Context, bind *model.BindAuth) (string, error) {
	exp := time.Now().Add(srv.conf.TokenExpireTime * time.Minute).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": bind.UserID,
		"exp":  exp,
	})
	token, err := claims.SignedString(srv.conf.JwtSecure)
	return token, err
}

// bcrypt: https://github.com/golang/crypto/blob/master/bcrypt/bcrypt_test.go
func (srv *Service) Register(ctx context.Context, bind *model.BindRegister) error {
	var err error
	if err = srv.checkUserAndDomain(ctx, bind.UserID, bind.Domain); err != nil {
		return err
	}
	salt, err := generateSalt(ctx, srv.conf.SaltSize)
	if err != nil {
		// todo: handle error
		return err
	}
	password, err := generatePassword(ctx, srv.conf.HashCost, salt, bind.UserPassword)
	if err != nil {
		// todo: handle error
		return err
	}
	token, err := uuid.NewUUID()
	if err != nil {
		// todo: handle error
		return err
	}
	auth := &model.Auths{
		UserID:       bind.UserID,
		UserPassword: password,
		Token:        token.String(),
		Salt:         string(salt),
		Teacher: model.Teachers{
			Domain: bind.Domain,
			NameZH: bind.NameZH,
			Email:  bind.Email,
		},
	}
	return srv.db.CreateUser(ctx, auth)
}

func (srv *Service) checkUserAndDomain(ctx context.Context, userID, domain string) error {
	var err error
	err = srv.db.UserIsExist(ctx, userID)
	if err == nil {
		return errIsExisted
	}
	err = srv.db.DomainIsExist(ctx, domain)
	if err == nil {
		return errIsExisted
	}
	return nil
}
func generateSalt(ctx context.Context, saltSize int) ([]byte, error) {
	var salt = make([]byte, saltSize)
	_, err := rand.Read(salt[:])
	return salt, err
}
func generatePassword(ctx context.Context, cost int, salt []byte, password string) (string, error) {
	saltPassword := append([]byte(password), salt...)
	hashPassword, err := bcrypt.GenerateFromPassword(saltPassword, cost)
	return string(hashPassword), err
}
func (srv *Service) GetJwtSecure(ctx context.Context) []byte {
	return srv.conf.JwtSecure
}

// todo: logout
