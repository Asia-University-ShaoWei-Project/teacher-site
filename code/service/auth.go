package service

import (
	"context"
	"teacher-site/model"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// todo: get environment variable by viper
var jwtSecure = []byte(`secure`)

func (srv *Service) Login(ctx context.Context, bindAuth *model.BindAuth, cfg *model.Config) (string, error) {
	var err error
	auth := &model.Auths{
		UserID:       bindAuth.UserID,
		UserPassword: bindAuth.UserPassword,
	}
	// todo: encrypt the password
	// password := encrypt(auth.password)

	// teacher, err := srv.db.VerifyAuthAndGetTeacher(auth)
	// if err != nil {
	// 	return "", err
	// }
	token, err := srv.newToken(ctx, auth.UserID)
	if err != nil {
		srv.log.Error(err)
		return "", err
	}
	if err = srv.db.UpdateUserToken(auth, token); err != nil {
		// setTokenInDB(ctx, token, auth.UserID)
		srv.log.Error(err)
		return "", err
	}
	if err = srv.cache.SetToken(token, srv.domain); err != nil {
		srv.log.Error(err)
		return "", err
	}
	return "", err
}
func (srv *Service) newToken(ctx context.Context, userID string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"expire":  time.Now().Add(1 * time.Minute).Unix(),
		// "expire": time.Now().Add(20 * time.Minute).Unix(),
	})
	token, err := claims.SignedString(jwtSecure)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (srv *Service) Register(ctx context.Context, bindRegister *model.BindRegister, cfg *model.Config) bool {
	//todo: check user_id and domain exist
	saltPassword := []byte(bindRegister.UserPassword + cfg.PasswordSecure)
	hashPassword, err := bcrypt.GenerateFromPassword(saltPassword, cfg.HashCost)
	if err != nil {
		srv.log.Error(err)
		return false
	}
	data := &model.Auths{
		UserID:       bindRegister.UserID,
		UserPassword: string(hashPassword),
		Teacher: model.Teachers{
			Domain: bindRegister.Domain,
			NameZH: bindRegister.NameZH,
			Email:  bindRegister.Email,
		},
	}
	srv.db.CreateUser(data)
	return true
}

// todo: logout
// func (srv *Service) Logout(ctx context.Context) error {}
