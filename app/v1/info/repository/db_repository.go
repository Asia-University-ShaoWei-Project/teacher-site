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

type DbRepository struct {
	db   *gorm.DB
	conf *config.DB
}

func NewDbRepository(db *gorm.DB, conf *config.DB) domain.InfoDbRepository {
	return &DbRepository{
		db:   db,
		conf: conf,
	}
}

func (r *DbRepository) Create(ctx context.Context, req *domain.CreateInfoBulletinRequest) (domain.InfoBulletinBoards, error) {
	var bulletin domain.InfoBulletinBoards
	err := r.db.Transaction(func(tx *gorm.DB) error {
		info := domain.Infos{AutoModel: domain.AutoModel{Id: req.InfoId}}
		if err := updateInfoLastModified(tx, &info); err != nil {
			return err
		}
		// create info bulletin
		bulletin = domain.InfoBulletinBoards{InfoId: info.AutoModel.Id, Content: req.Content}
		result := tx.Create(&bulletin)
		if err := checkErrAndExist(result); err != nil {
			return err
		}
		return nil
	})
	return bulletin, err
}

func (r *DbRepository) GetByTeacherDomain(ctx context.Context, teacherDomain string) (domain.Infos, error) {
	var info domain.Infos
	result := r.db.Model(&info).Where("teacher_id=? AND deleted_at IS NULL", teacherDomain).Find(&info)
	err := checkErrAndExist(result)
	return info, err
}
func (r *DbRepository) GetBulletinsByInfoId(ctx context.Context, id uint) ([]domain.InfoBulletinResponse, error) {
	var infoBulletin []domain.InfoBulletinResponse
	result := r.db.Table("info_bulletin_boards ib").
		Select("ib.id, DATE(ib.created_at) AS date, ib.content").
		Joins("JOIN infos i ON ib.info_id = i.id").
		Where("i.id=? AND ib.deleted_at IS NULL", id).
		Order("ib.created_at desc").
		Find(&infoBulletin)
	err := checkErrAndExist(result)
	return infoBulletin, err
}

// todo: delete?
func (r *DbRepository) GetLastModified(ctx context.Context, id uint) (string, error) {
	var info domain.Infos
	result := r.db.Find(&info, id)
	err := checkErrAndExist(result)
	return info.LastModified, err
}

func (r *DbRepository) Update(ctx context.Context, req *domain.UpdateInfoBulletinRequest) (domain.Infos, error) {
	info := domain.Infos{AutoModel: domain.AutoModel{Id: req.InfoId}}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := updateInfoLastModified(tx, &info); err != nil {
			return err
		}
		// update the content of infos_bulletin
		// bulletin := domain.InfoBulletinBoards{
		// AutoModel: domain.AutoModel{Id: req.BulletinId},
		// Content:   req.Content}
		result := tx.Model(&domain.InfoBulletinBoards{}).
			Where("id=? AND info_id=?", req.BulletinId, info.AutoModel.Id).
			Update("content", req.Content)
		if err := checkErrAndExist(result); err != nil {
			return err
		}
		return nil
	})
	return info, err
}
func (r *DbRepository) Delete(ctx context.Context, req *domain.DeleteInfoBulletinRequest) (domain.Infos, error) {
	info := domain.Infos{AutoModel: domain.AutoModel{Id: req.InfoId}}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := updateInfoLastModified(tx, &info); err != nil {
			return err
		}
		fmt.Println(`after time:`, info.LastModified)

		result := tx.Where("id=? AND info_id=?", req.BulletinId, info.AutoModel.Id).Delete(&domain.InfoBulletinBoards{})

		if err := checkErrAndExist(result); err != nil {
			return err
		}
		return nil
	})
	return info, err
}
func updateInfoLastModified(tx *gorm.DB, info *domain.Infos) error {
	result := tx.Model(&info).Update("last_modified", newLastModifiedTime())
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
