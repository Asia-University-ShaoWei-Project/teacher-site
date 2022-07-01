package database

import (
	"fmt"
	"strconv"
	"teacher-site/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(path string, conf *config.DB) *gorm.DB {
	file := path + "/" + conf.Filename
	db, err := gorm.Open(sqlite.Open(file))
	if err != nil {
		panic(err)
	}
	return db
}
func NewDbByPostgres(conf *config.DB) *gorm.DB {
	conn := "host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf(conn, conf.User, conf.Password, conf.Database, conf.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
func CheckErrAndExist(result *gorm.DB) error {
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func NewLastModifiedTime() string {
	now := time.Now()
	return strconv.FormatInt(now.Unix(), 10)
}
