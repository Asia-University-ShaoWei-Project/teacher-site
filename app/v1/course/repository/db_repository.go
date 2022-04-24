package repository

import (
	"context"
	"strconv"
	"teacher-site/config"
	"teacher-site/domain"
	"time"

	"gorm.io/gorm"
)

type DbRepository struct {
	db   *gorm.DB
	conf *config.DB
}

func NewInfoRepository(db *gorm.DB, conf *config.DB) domain.InfoDbRepository {
	return &DbRepository{
		db:   db,
		conf: conf,
	}
}
func (r *DbRepository) Create(ctx context.Context, req *domain.ReqCreateInfo) (domain.InfoBulletinBoards, error) {
	var (
		info     domain.Infos
		bulletin domain.InfoBulletinBoards
	)
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := updateInfoLastModified(tx, req.TeacherDomain, &info); err != nil {
			return err
		}
		// create info bulletin
		bulletin = domain.InfoBulletinBoards{InfoID: info.AutoModel.ID, Content: req.Content}
		result := tx.Create(&bulletin)
		if err := checkErrAndExist(result); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return bulletin, err
	}

	return bulletin, nil

}
func (r *DbRepository) Get(ctx context.Context, req *domain.ReqGetInfo) ([]domain.GetInfoBulletin, error) {
	var infoBulletin []domain.GetInfoBulletin
	r.db.Table("info_bulletin_boards ib").
		Select("ib.id, DATE(ib.created_at) AS date, ib.content").
		Joins("JOIN infos i ON ib.info_id = i.id").
		Joins("JOIN teachers t ON teacher_id = t.domain").
		Where("t.domain=? AND ib.deleted_at IS NULL", req.TeacherDomain).Find(&infoBulletin)
	return infoBulletin, nil
}

// func (r *Repository) GetLastModified(ctx context.Context) (string, error)

// 	var info domain.Infos
// 	if err := r.db.Where("teacher_id=?", domain).Find(&info).Error; err != nil {
// 		return "", err
// 	}
// 	return info.LastModified, nil
// }

// todo: checkErrAndExist is require?
func (r *DbRepository) Update(ctx context.Context, req *domain.ReqUpdateInfoBulletin) (domain.InfoBulletinBoards, error) {
	var bulletin domain.InfoBulletinBoards
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var info domain.Infos
		if err := updateInfoLastModified(tx, req.TeacherDomain, &info); err != nil {
			return err
		}
		// update the content of infos_bulletin
		// bulletin := domain.InfoBulletinBoards{
		// AutoModel: domain.AutoModel{ID: req.BulletinID},
		// Content:   req.Content}
		result := tx.Model(&bulletin).Where("id=? AND info_id=?", req.BulletinID, info.AutoModel.ID).Update("content", req.Content)
		if err := checkErrAndExist(result); err != nil {
			return err
		}
		return nil
	})
	return bulletin, err
}
func (r *DbRepository) Delete(ctx context.Context, req *domain.ReqDeleteInfo) error {

	err := r.db.Transaction(func(tx *gorm.DB) error {
		var info domain.Infos
		if err := updateInfoLastModified(tx, req.TeacherDomain, &info); err != nil {
			return err
		}
		if err := tx.Where("id=? AND info_id=?", req.BulletinID, info.AutoModel.ID).Delete(&domain.InfoBulletinBoards{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
func updateInfoLastModified(tx *gorm.DB, teacherDomain string, info *domain.Infos) error {
	var err error
	// check domain is existed
	result := tx.Model(&info).Where(`teacher_id=?`, teacherDomain).Find(&info)
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
func checkErrAndExist(result *gorm.DB) error {
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func newLastModifiedTime() string {
	now := time.Now()
	return strconv.FormatInt(now.Unix(), 10)
}
