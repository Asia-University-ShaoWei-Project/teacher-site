package database

import (
	"context"
	"teacher-site/model"

	"gorm.io/gorm"
)

func (db *DB) CreateInfo(ctx context.Context, domain string, req *model.ReqCreateInfo) (model.ResCreateInfo, error) {
	var (
		info     model.Infos
		bulletin model.InfoBulletinBoards
		res      model.ResCreateInfo
	)
	err := db.orm.Transaction(func(tx *gorm.DB) error {
		if err := updateInfoLastModified(tx, domain, &info); err != nil {
			return err
		}
		// create info bulletin
		bulletin = model.InfoBulletinBoards{InfoID: info.AutoModel.ID, Content: req.Content}
		result := tx.Create(&bulletin)
		if err := checkErrAndExist(result); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return res, err
	}
	res = model.ResCreateInfo{
		ID:   bulletin.AutoModel.ID,
		Date: bulletin.AutoModel.CreatedAT.Format(model.BulletinDateFormat),
	}
	return res, nil

}

func (db *DB) GetInfo(ctx context.Context, domain string) ([]model.ResGetInfo, error) {
	var infoBulletin []model.ResGetInfo
	db.orm.Table("info_bulletin_boards i").
		Select("i.id, DATE(i.created_at) AS date, i.content").
		Joins("JOIN infos ON i.info_id = infos.id").
		Joins("JOIN teachers t ON teacher_id = t.domain").
		Where("t.domain=? AND i.deleted_at IS NULL", domain).Find(&infoBulletin)
	return infoBulletin, nil
}

func (db *DB) GetInfoLastUpdated(ctx context.Context, domain string) (string, error) {
	var info model.Infos
	if err := db.orm.Where("teacher_id=?", domain).Find(&info).Error; err != nil {
		return "", err
	}
	return info.LastModified, nil
}

// todo: checkErrAndExist is require?
func (db *DB) UpdateInfo(ctx context.Context, domain string, req *model.ReqUpdateInfoBulletin) error {
	err := db.orm.Transaction(func(tx *gorm.DB) error {
		var info model.Infos
		if err := updateInfoLastModified(tx, domain, &info); err != nil {
			return err
		}
		// update the content of infos_bulletin
		// bulletin := model.InfoBulletinBoards{
		// AutoModel: model.AutoModel{ID: req.BulletinID},
		// Content:   req.Content}
		result := tx.Model(&model.InfoBulletinBoards{}).Where("id=? AND info_id=?", req.BulletinID, info.AutoModel.ID).Update("content", req.Content)
		if err := checkErrAndExist(result); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (db *DB) DeleteInfo(ctx context.Context, domain string, req *model.ReqDeleteInfo) error {
	err := db.orm.Transaction(func(tx *gorm.DB) error {
		var info model.Infos
		if err := updateInfoLastModified(tx, domain, &info); err != nil {
			return err
		}
		if err := tx.Where("id=? AND info_id=?", req.BulletinID, info.AutoModel.ID).Delete(&model.InfoBulletinBoards{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
func updateInfoLastModified(tx *gorm.DB, domain string, info *model.Infos) error {
	var err error
	// check domain is existed
	result := tx.Model(&info).Where(`teacher_id=?`, domain).Find(&info)
	err = checkErrAndExist(result)
	if err != nil {
		return err
	}
	// update info modified time
	result = tx.Model(&info).Update("last_modified", newLastModifiedTime())
	err = checkErrAndExist(result)
	if err != nil {
		return err
	}
	return nil
}
