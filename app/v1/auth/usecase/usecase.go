package usecase

import (
	"context"
	"teacher-site/config"
	"teacher-site/domain"
	"teacher-site/pkg/util"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Usecase struct {
	DbRepository    domain.AuthDbRepository
	CacheRepository domain.AuthCacheRepository
	conf            *config.Config
	log             *logrus.Logger
}

func NewUsecase(dbRepo domain.AuthDbRepository, cacheRepo domain.AuthCacheRepository, conf *config.Config, logger *logrus.Logger) domain.AuthUsecase {
	return &Usecase{
		DbRepository:    dbRepo,
		CacheRepository: cacheRepo,
		conf:            conf,
		log:             logger,
	}
}
func (u *Usecase) Login(ctx context.Context, req *domain.LoginRequest) (domain.LoginResponse, error) {
	var res domain.LoginResponse

	account, err := u.DbRepository.GetAccountByUserId(ctx, req.UserId)
	if err != nil {
		u.log.Error(err)
		return res, err
	}

	// compare password
	saltPassword := append([]byte(req.UserPassword), account.Salt...)
	if err = bcrypt.CompareHashAndPassword([]byte(account.UserPassword), saltPassword); err != nil {
		u.log.Error(err)
		return res, err
	}
	teacher, err := u.DbRepository.GetTeacherDomainByUserId(ctx, account.UserId)
	if err != nil {
		u.log.Error(err)
		return res, err
	}
	// generate new token for the header of client(authorization)
	jwtReq := domain.JwtInfoRequest{UserId: account.UserId, Domain: teacher.Domain}
	token, err := util.GenerateJwt(u.conf.Jwt, &jwtReq)
	if err != nil {
		// todo: try again
		u.log.Error(err)
	}
	res = domain.LoginResponse{
		Token:  token,
		Domain: teacher.Domain,
	}
	if err = u.DbRepository.UpdateTokenByUserId(ctx, account.UserId, token); err != nil {
		// todo: error handle of update token
		u.log.Error(err)
	}
	return res, nil
}

func (u *Usecase) Logout(ctx context.Context, id string) error {
	return u.DbRepository.DeleteTokenById(ctx, id)
}
func (u *Usecase) Register(ctx context.Context, req *domain.RegisterRequest) error {
	if err := u.DbRepository.CheckUserExistByUserIdAndDomain(ctx, req.UserId, req.Domain); err != nil {
		return err
	}
	salt := util.GeneralSalt(ctx, u.conf.Secure.SaltSize)
	hashPassword := util.GeneralHashPassword(ctx, req.UserPassword, salt, u.conf.Secure.HashCost)

	auth := &domain.Auths{
		UserId:       req.UserId,
		UserPassword: hashPassword,
		Salt:         salt,
		Teacher: domain.Teachers{
			Domain:  req.Domain,
			NameZh:  req.NameZh,
			Email:   req.Email,
			Infos:   []domain.Infos{},
			Courses: []domain.Courses{},
		},
	}
	if err := u.DbRepository.CreateTeacher(ctx, auth); err != nil {
		return err
	}
	path := u.conf.Server.Path.GetDocPath()
	if err := util.CreateDirByTeacherDomain(ctx, path, auth.Teacher.Domain); err != nil {
		u.log.Error(err)
	}
	return nil
}
