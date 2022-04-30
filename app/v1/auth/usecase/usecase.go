package usecase

import (
	"context"
	"teacher-site/config"
	"teacher-site/domain"
	"teacher-site/pkg/util"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	DbRepository    domain.AuthDbRepository
	CacheRepository domain.AuthCacheRepository
	conf            *config.Config
	log             *logrus.Logger
}

func NewUsecase(dbRepo domain.AuthDbRepository, cacheRepo domain.AuthCacheRepository, conf *config.Config, logger *logrus.Logger) domain.AuthUsecase {
	return &AuthUsecase{
		DbRepository:    dbRepo,
		CacheRepository: cacheRepo,
		conf:            conf,
		log:             logger,
	}
}
func (auth *AuthUsecase) Login(ctx context.Context, req *domain.LoginRequest) (domain.LoginResponse, error) {
	var res domain.LoginResponse

	account, err := auth.DbRepository.GetAccountByUserId(ctx, req.UserId)
	if err != nil {
		auth.log.Error(err)
		return res, err
	}

	// compare password
	saltPassword := []byte(req.UserPassword + account.Salt)
	if err = bcrypt.CompareHashAndPassword([]byte(account.UserPassword), saltPassword); err != nil {
		auth.log.Error(err)
		return res, err
	}
	teacher, err := auth.DbRepository.GetTeacherDomainByUserId(ctx, account.UserId)
	if err != nil {
		auth.log.Error(err)
		return res, err
	}
	// generate new token for the header of client(authorization)
	jwtReq := domain.JwtInfoRequest{UserId: account.UserId, Domain: teacher.Domain}
	token, err := util.GenerateJwt(auth.conf.Jwt, &jwtReq)
	if err != nil {
		// todo: try again
		auth.log.Error(err)
	}
	res = domain.LoginResponse{
		Token:  token,
		Domain: teacher.Domain,
	}
	if err = auth.DbRepository.UpdateTokenByUserId(ctx, account.UserId, token); err != nil {
		// todo: error handle of update token
		auth.log.Error(err)
	}
	return res, nil
}

func (auth *AuthUsecase) Logout(ctx context.Context, id string) error {
	return auth.DbRepository.DeleteTokenById(ctx, id)
}
