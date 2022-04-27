package repository

import (
	"context"
	"fmt"
	"teacher-site/config"
	"teacher-site/domain"

	"gorm.io/gorm"
)

type DbRepository struct {
	db   *gorm.DB
	conf *config.DB
}

func NewDbRepository(db *gorm.DB, conf *config.DB) domain.PageDbRepository {
	return &DbRepository{
		db:   db,
		conf: conf,
	}
}

func (r *DbRepository) GetTeachers(ctx context.Context, limit, offset int) ([]domain.TeacherResponse, error) {
	var teachers []domain.TeacherResponse
	result := r.db.Model(&domain.Teachers{}).
		Select("domain", "name_zh", "name_us").
		Limit(limit).
		Offset(offset).
		Find(&teachers)
	return teachers, checkErrAndRecord(result)
}

func (r *DbRepository) GetTeacherByDomain(ctx context.Context, teacherDomain string) (domain.Teachers, error) {
	teacher := domain.Teachers{Domain: teacherDomain}
	result := r.db.Model(&teacher).
		Select("email", "name_zh", "name_us", "office", "call", "education").
		Find(&teacher)
	return teacher, checkErrAndRecord(result)
}
func (r *DbRepository) CheckAuthByIdAndToken(ctx context.Context, userId, token string) error {
	// todo: optimize: find id(Pk) -> check token(Srv)
	var auth domain.Auths
	result := r.db.Model(&auth).
		Where("user_id=? AND token=?", userId, token).
		Find(&auth)
	return checkErrAndRecord(result)
}
func checkErrAndRecord(result *gorm.DB) error {
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil

}
