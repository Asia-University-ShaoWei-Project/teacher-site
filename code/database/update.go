package database

import (
	"teacher-site/message"
	"teacher-site/model"

	"gorm.io/gorm"
)

func (db *DB) UpdateInformation(info *model.Informations) error {
	var originInfo *model.Informations
	err := db.orm.Transaction(func(tx *gorm.DB) error {
		result := db.orm.Where("id = ?", info.ID).Find(originInfo)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return message.ErrQueryNotFound
		}
		if err := db.orm.Model(originInfo).Updates(info).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (db *DB) UpdateUserToken(auth *model.Auths, token string) error {
	err := db.orm.Transaction(func(tx *gorm.DB) error {
		if err := db.orm.Model(auth).Update("token", token).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
