package service

import (
	"context"
	"teacher-site/cache"
	"teacher-site/database"
	"teacher-site/model"

	log "github.com/sirupsen/logrus"
)

type Service struct {
	domain string
	db     database.Databaseer
	cache  cache.Cacheer
	log    *log.Logger // log *zap.Logger
	conf   *model.ServiceConfig
}
type Servicer interface {
	GetJwtSecure(ctx context.Context) []byte
	GetInit(ctx context.Context, init *model.Init) error

	GetCourse(ctx context.Context, courseBind *model.BindCourse, course *model.Courses) error

	CreateInfo(ctx context.Context, reqInfo *model.BindInfo) error
	UpdateInfo(ctx context.Context, reqInfo *model.BindInfo) error
	DeleteInfo(ctx context.Context, reqInfo *model.BindInfo) error

	//? auth
	Login(ctx context.Context, bindAuth *model.BindAuth) error
	NewJwtToken(ctx context.Context, bindAuth *model.BindAuth) (string, error)
	UpdateJwtToken(ctx context.Context, token string, bindAuth *model.BindAuth) error
	Register(ctx context.Context, bindRegister *model.BindRegister) error

	// todo
	// Logout(ctx context.Context) error
	DomainIsExist(ctx context.Context, domain string) error
	SetDomain(ctx context.Context, domain string)
	// ? temporary
	Info(value ...interface{})
	Error(value ...interface{})
}

// TODO: receive logger parameter
func NewService(db database.Databaseer, cache cache.Cacheer, logger *log.Logger, conf *model.ServiceConfig) Servicer {
	srv := &Service{
		db:    db,
		cache: cache,
		log:   logger,
		conf:  conf,
	}
	return srv
}
func (srv *Service) SetDomain(ctx context.Context, domain string) {
	srv.domain = domain
}
func (srv *Service) DomainIsExist(ctx context.Context, domain string) error {
	return srv.db.DomainIsExist(ctx, domain)
}
func (srv *Service) Info(value ...interface{}) {
	srv.log.Info(value...)
}
func (srv *Service) Error(value ...interface{}) {
	srv.log.Error(value...)
}
