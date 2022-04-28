package repository

import (
	"context"
	"teacher-site/config"
	"teacher-site/domain"

	"gorm.io/gorm"
)

type DbRepository struct {
	db   *gorm.DB
	conf *config.DB
}

func NewDbRepository(db *gorm.DB, conf *config.DB) domain.AuthDbRepository {
	return &DbRepository{
		db:   db,
		conf: conf,
	}
}

func (r *DbRepository) GetAccountByUserId(ctx context.Context, id string) (domain.Auths, error) {
	auth := domain.Auths{UserId: id}
	result := r.db.Find(&auth)
	err := checkErrAndExist(result)
	return auth, err
}
func (r *DbRepository) GetTeacherDomainByUserId(ctx context.Context, id string) (domain.Teachers, error) {
	var teacher domain.Teachers
	result := r.db.Model(&teacher).Select("domain").Where("auth_id", id).Find(&teacher)
	err := checkErrAndExist(result)
	return teacher, err
}

func (r *DbRepository) UpdateTokenByUserId(ctx context.Context, id, token string) error {
	auth := domain.Auths{UserId: id}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&auth).Update("token", token)
		if err := checkErrAndExist(result); err != nil {
			return err
		}
		return nil
	})

	return err
}

func (r *DbRepository) DeleteTokenById(ctx context.Context, id string) error {
	auth := domain.Auths{UserId: id}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&auth).Update("token", "")
		if err := checkErrAndExist(result); err != nil {
			return err
		}
		return nil
	})
	return err
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
