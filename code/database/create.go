package database

import (
	"teacher-site/model"

	"gorm.io/gorm"
)

func (db *DB) CreateInformation(obj *model.Informations) error {
	err := db.orm.Transaction(func(tx *gorm.DB) error {
		if err := db.orm.Omit("id").Create(obj).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
func (db *DB) CreateUser(obj *model.Auths) error {
	err := db.orm.Transaction(func(tx *gorm.DB) error {
		if err := db.orm.Omit("id").Create(obj).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
