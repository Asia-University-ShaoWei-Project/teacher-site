package database

import (
	"context"
	"teacher-site/model"
)

func (db *DB) DeleteInformation(ctx context.Context, id uint) error {
	if err := db.orm.Delete(&model.Informations{}, id).Error; err != nil {
		return err
	}
	return nil
}
