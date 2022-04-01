package database

import "teacher-site/model"

func (db *DB) DeleteInformation(id uint) error {
	if err := db.orm.Delete(&model.Informations{}, id).Error; err != nil {
		return err
	}
	return nil
}
