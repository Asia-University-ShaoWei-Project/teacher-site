package repository

import (
	"context"
	"fmt"
	"strconv"
	"teacher-site/config"
	"teacher-site/domain"
	"time"

	"gorm.io/gorm"
)

type dbRepository struct {
	db   *gorm.DB
	conf *config.DB
}

func NewInfoRepository(db *gorm.DB, conf *config.DB) domain.InfoDbRepository {
	return &dbRepository{
		db:   db,
		conf: conf,
	}
}

func (r *dbRepository) Create(ctx context.Context, req *domain.CreateInfoBulletinRequest) (domain.InfoBulletinBoards, error) {
	var bulletin domain.InfoBulletinBoards
	err := r.db.Transaction(func(tx *gorm.DB) error {
		info := domain.Infos{TeacherID: req.TeacherDomain}
		if err := updateInfoLastModified(tx, &info); err != nil {
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
	return bulletin, err
}

func (r *dbRepository) GetByTeacherDomain(ctx context.Context, teacherDomain string) (domain.Infos, error) {
	var info domain.Infos
	result := r.db.Model(&info).Where("teacher=? AND deleted_at IS NULL", teacherDomain).Find(&info)
	err := checkErrAndExist(result)
	return info, err
}
func (r *dbRepository) GetBulletinByInfoId(ctx context.Context, id uint) ([]domain.InfoBulletinResponse, error) {
	var infoBulletin []domain.InfoBulletinResponse
	r.db.Table("info_bulletin_boards ib").
		Select("ib.id, DATE(ib.created_at) AS date, ib.content").
		Joins("JOIN infos i ON ib.info_id = i.id").
		Where("i.id=? AND ib.deleted_at IS NULL", id).Find(&infoBulletin)
	return infoBulletin, nil
}

// todo: delete?
func (r *dbRepository) GetLastModified(ctx context.Context, teacherDomain string) (string, error) {
	var info domain.Infos
	err := r.db.Where("teacher_id=?", teacherDomain).Find(&info).Error
	return info.LastModified, err
}

func (r *dbRepository) Update(ctx context.Context, req *domain.UpdateInfoBulletinRequest) (domain.Infos, error) {
	info := domain.Infos{TeacherID: req.TeacherDomain}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := updateInfoLastModified(tx, &info); err != nil {
			return err
		}
		// update the content of infos_bulletin
		// bulletin := domain.InfoBulletinBoards{
		// AutoModel: domain.AutoModel{ID: req.BulletinID},
		// Content:   req.Content}
		result := tx.Model(&domain.InfoBulletinBoards{}).
			Where("id=? AND info_id=?", req.BulletinID, info.AutoModel.ID).
			Update("content", req.Content)
		if err := checkErrAndExist(result); err != nil {
			return err
		}
		return nil
	})
	return info, err
}
func (r *dbRepository) Delete(ctx context.Context, req *domain.DeleteInfoBulletinRequest) (domain.Infos, error) {
	info := domain.Infos{TeacherID: req.TeacherDomain}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := updateInfoLastModified(tx, &info); err != nil {
			return err
		}
		fmt.Println(info)

		result := tx.Where("id=? AND info_id=?", req.BulletinID, info.AutoModel.ID).Delete(&domain.InfoBulletinBoards{})

		if err := checkErrAndExist(result); err != nil {
			return err
		}
		return nil
	})
	return info, err
}
func updateInfoLastModified(tx *gorm.DB, info *domain.Infos) error {
	result := tx.Model(&info).
		Where(`teacher_id=?`, info.TeacherID).
		Update("last_modified", newLastModifiedTime())
	return checkErrAndExist(result)
}

// unix time
func newLastModifiedTime() string {
	now := time.Now()
	return strconv.FormatInt(now.Unix(), 10)
}

// result.Err or gorm Not Found
func checkErrAndExist(result *gorm.DB) error {
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
