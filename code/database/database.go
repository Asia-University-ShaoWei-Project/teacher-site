package database

import (
	"context"
	"os"

	"teacher-site/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
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
	// VerifyAuthAndGetTeacher(auth *model.BindAuth) (*model.Teachers, error)
	CreateInformation(ctx context.Context, obj *model.Informations) error
	CreateUser(ctx context.Context, obj *model.Auths) error

	GetAuth(ctx context.Context, auth *model.Auths) error
	GetInit(ctx context.Context, init *model.Init, domain string) error
	GetCourseLastUpdated(ctx context.Context, courseID uint) (string, error)
	GetCourseContent(ctx context.Context, courseID uint) (model.Courses, error)

	UpdateInformation(ctx context.Context, info *model.Informations) error
	UpdateUserToken(ctx context.Context, auth *model.Auths, token string) error

	DeleteInformation(ctx context.Context, id uint) error

	DomainIsExist(ctx context.Context, domain string) error
	UserIsExist(ctx context.Context, userID string) error
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

// var (
// 	host     = "localhost"
// 	user     = "server"
// 	password = "password"
// 	dbName   = "main"
// 	port     = "5432"
// 	sslMode  = "disable"
// 	timeZone = "Asia/Taipei"
// )

func NewPostgres() Databaseer {
	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
	// 	host, user, password, dbName, port, sslMode, timeZone,
	// )
	// postgres://qgkqtgplobxbnm:fdfddd8d7be467e67fe028ff70202cfe185420a2adb4c0ac9bd51bcebf8e687f@ec2-34-207-12-160.compute-1.amazonaws.com:5432/dbedrq41b75uae
	dsn := os.Getenv("key string")
	orm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DB{
		orm: orm,
	}
}

func (db *DB) Migrate(ctx context.Context, dst ...interface{}) error {
	return db.orm.AutoMigrate(dst...)
}
func (db *DB) Create(ctx context.Context, value interface{}) error {
	return db.orm.Create(value).Error
}
