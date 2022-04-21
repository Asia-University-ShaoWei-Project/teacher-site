package database

import (
	"teacher-site/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(path string, conf *config.DB) *gorm.DB {
	file := path + "/sqlite_dev.db"
	db, err := gorm.Open(sqlite.Open(file))
	if err != nil {
		panic(err)
	}
	return db
}
