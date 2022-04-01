package database

import (
	"teacher-site/message"
	"teacher-site/model"

	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

type DB struct {
	orm *gorm.DB
}
type Databaseer interface {
	Create(value interface{}) error
	Migrate(dst ...interface{}) error
	FindDomain(domain *string) error
	// VerifyAuthAndGetTeacher(auth *model.BindAuth) (*model.Teachers, error)
	CreateInformation(obj *model.Informations) error
	CreateUser(obj *model.Auths) error
	GetInit(domain string) (*model.Init, error)
	GetCourseWithContent(id uint) *model.Courses
	UpdateInformation(info *model.Informations) error
	UpdateUserToken(auth *model.Auths, token string) error
	DeleteInformation(id uint) error
}

func NewSqlite(path string) Databaseer {
	file := path + "/" + "sqlite.db"
	orm, err := gorm.Open(sqlite.Open(file))
	if err != nil {
		panic(err)
	}
	return &DB{
		orm: orm,
	}
}
func (db *DB) Migrate(dst ...interface{}) error {
	return db.orm.AutoMigrate(dst...)
}
func (db *DB) Create(value interface{}) error {
	return db.orm.Create(value).Error
}

func (db *DB) FindDomain(domain *string) error {
	result := db.orm.Where("domain=?", &domain).Find(&model.Teachers{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return message.ErrQueryNotFound
	}
	return nil
}
