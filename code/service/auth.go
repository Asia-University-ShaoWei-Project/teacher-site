package service

import (
	"context"
	"crypto/rand"
	"teacher-site/model"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//ref: https://github.com/golang-jwt/jwt/blob/main/example_test.go
func (srv *Service) Login(ctx context.Context, bindAuth *model.BindAuth) error {
	var err error
	auth := &model.Auths{UserID: bindAuth.UserID}
	if err = srv.db.GetAuth(ctx, auth); err != nil {
		return err
	}
	saltPassword := []byte(bindAuth.UserPassword + auth.Salt)
	if err = bcrypt.CompareHashAndPassword([]byte(auth.UserPassword), saltPassword); err != nil {
		return err
	}
	return nil
}

func (srv *Service) UpdateJwtToken(ctx context.Context, bindAuth *model.BindAuth) error {
	auth := &model.Auths{
		UserID: bindAuth.UserID,
	}
	jwtToken, err := newJwtToken(ctx, auth.UserID, srv.conf.TokenExpireTime, srv.conf.JWTSecure)
	if err != nil {
		return err
	}
	if err = srv.db.UpdateUserToken(ctx, auth, jwtToken); err != nil {
		return err
	}
	if err = srv.cache.SetTokenWithDomain(ctx, jwtToken, srv.domain); err != nil {
		return err
		// todo: handle this
	}
	return nil

}
func newJwtToken(ctx context.Context, userID string, duration time.Duration, secure []byte) (string, error) {
	exp := time.Now().Add(duration * time.Minute).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userID,
		"exp":  exp,
	})
	token, err := claims.SignedString(secure)
	return token, err
}

// bcrypt: https://github.com/golang/crypto/blob/master/bcrypt/bcrypt_test.go
func (srv *Service) Register(ctx context.Context, bindRegister *model.BindRegister) error {
	var err error
	if err = srv.checkUserAndDomain(ctx, bindRegister.UserID, bindRegister.Domain); err != nil {
		return err
	}
	salt, err := generateSalt(ctx, srv.conf.SaltSize)
	if err != nil {
		return err
	}
	password, err := generatePassword(ctx, srv.conf.HashCost, salt, bindRegister.UserPassword)
	if err != nil {
		return err
	}
	token, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	auth := &model.Auths{
		UserID:       bindRegister.UserID,
		UserPassword: password,
		Token:        token.String(),
		Salt:         string(salt),
		Teacher: model.Teachers{
			Domain: bindRegister.Domain,
			NameZH: bindRegister.NameZH,
			Email:  bindRegister.Email,
		},
	}
	return srv.db.CreateUser(ctx, auth)
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
func (srv *Service) GetJWTSecure(ctx context.Context) []byte {
	return srv.conf.JWTSecure
}
func (srv *Service) checkUserAndDomain(ctx context.Context, userID, domain string) error {
	var err error
	err = srv.db.UserIsExist(ctx, userID)
	if err != nil {
		return err
	}
	err = srv.db.DomainIsExist(ctx, domain)
	if err != nil {
		return err
	}
	return nil
}

// todo: logout
