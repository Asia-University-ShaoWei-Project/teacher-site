package service

import (
	"context"
	"teacher-site/cache"
	"teacher-site/database"
	"teacher-site/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	domain string
	db     database.Databaseer
	cache  cache.Cacheer
	// log *zap.Logger
	log *log.Logger
}
type Servicer interface {
	GetInit(ctx context.Context) (model.Init, error)
	// todo: course
	GetCourse(ctx context.Context, courseBind *model.BindCourse) (model.Courses, error)
	// CreateCourse(ctx context.Context, id model.CourseID) string
	// UpdateCourse(ctx context.Context, id model.CourseID) string
	// todo edit
	CreateInfo(ctx context.Context, reqInfo *model.BindInfo) error
	UpdateInfo(ctx context.Context, reqInfo *model.BindInfo) error
	DeleteInfo(ctx context.Context, reqInfo *model.BindInfo) error

	//? other
	IsExistDomain(ctx context.Context, domain *string) error
	SetDomain(ctx context.Context, domain *string)
	//? auth
	Login(ctx context.Context, auth *model.BindAuth, cfg *model.Config) (string, error)
	Register(ctx context.Context, bindRegister *model.BindRegister, cfg *model.Config) bool
	// todo
	// Logout(ctx context.Context) error
	// todo, gin.context
	// SetAuthBearerHeader(ctx context.Context, c *gin.Context, token string)
	// ? temporary
	Debug(value ...interface{})
}

// TODO: receive logger parameter
func NewService(db database.Databaseer, cache cache.Cacheer, logger *log.Logger) Servicer {
	srv := &Service{
		// TODO: logger
		db:    db,
		cache: cache,
		log:   logger,
	}
	// srv.newLog()
	return srv
}

func (srv *Service) SetDomain(ctx context.Context, domain *string) {
	srv.log.Info("my domain = ", *domain)
	srv.domain = *domain
}

func (srv *Service) IsExistDomain(ctx context.Context, domain *string) error {
	return srv.db.FindDomain(domain)
}

// func (srv *Service) newLog() error {
// 	return srv.cache.LSet(LogKey, 0, "").Err()
// }
// func (srv *Service) setLog(err error) {
// 	fmt.Println(err)
// 	srv.cache.RPush(LogKey, err.Error())
// }
func (srv *Service) SetAuthBearerHeader(ctx context.Context, c *gin.Context, token string) {
	// hash: auth <token> <domain>
	bearer := "bearer " + token
	c.Header("Authorization", bearer)
}
func (srv *Service) Debug(value ...interface{}) {
	srv.log.Debug(value...)
}
