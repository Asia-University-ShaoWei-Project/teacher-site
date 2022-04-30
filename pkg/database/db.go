package database

import (
	"strconv"
	"teacher-site/config"
	"time"

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
