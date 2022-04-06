package database

import (
	"context"

	"teacher-site/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

type DB struct {
	orm *gorm.DB
	log *log.Logger
}
type Databaseer interface {
	Create(ctx context.Context, value interface{}) error
	Migrate(ctx context.Context, dst ...interface{}) error
	FindDomain(ctx context.Context, domain *string) error
	// VerifyAuthAndGetTeacher(auth *model.BindAuth) (*model.Teachers, error)
	CreateInformation(ctx context.Context, obj *model.Informations) error
	CreateUser(ctx context.Context, obj *model.Auths) error
	GetAuth(ctx context.Context, auth *model.Auths) error
	GetInit(ctx context.Context, domain string) (*model.Init, error)
	GetCourseWithContent(ctx context.Context, id uint) *model.Courses
	UpdateInformation(ctx context.Context, info *model.Informations) error
	UpdateUserToken(ctx context.Context, auth *model.Auths, token string) error
	DeleteInformation(ctx context.Context, id uint) error
}

func NewSqlite(path string, logger *log.Logger) Databaseer {
	file := path + "/" + "sqlite.db"
	orm, err := gorm.Open(sqlite.Open(file))
	if err != nil {
		panic(err)
	}
	return &DB{
		orm: orm,
		log: logger,
	}
}
func (db *DB) Migrate(ctx context.Context, dst ...interface{}) error {
	return db.orm.AutoMigrate(dst...)
}
func (db *DB) Create(ctx context.Context, value interface{}) error {
	return db.orm.Create(value).Error
}
