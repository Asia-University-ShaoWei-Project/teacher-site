package repository

import (
	"context"
	"teacher-site/config"
	"teacher-site/domain"

	"gorm.io/gorm"
)

type dbRepository struct {
	db   *gorm.DB
	conf *config.DB
}

func NewAuthRepository(db *gorm.DB, conf *config.DB) domain.AuthDbRepository {
	return &dbRepository{
		db:   db,
		conf: conf,
	}
}

func (r *dbRepository) GetAccountByUserId(ctx context.Context, id string) (domain.Auths, error) {
	auth := domain.Auths{UserID: id}
	result := r.db.Find(&auth)
	err := checkErrAndExist(result)
	return auth, err
}

func (r *dbRepository) UpdateTokenByUserId(ctx context.Context, id, token string) error {
	auth := domain.Auths{UserID: id}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&auth).Update("token", token)
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
