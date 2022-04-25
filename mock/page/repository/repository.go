package repository

import (
	"context"
	"teacher-site/domain"
	"teacher-site/mock"

	"gorm.io/gorm"
)

// var (
// 	db = database.NewDB("../../../../pkg/database", mock.Conf.DB)
// )

// rdbms
type DbRepository struct{}

func NewDbRepository() domain.PageDbRepository {
	return &DbRepository{}
}

// todo
func (r *DbRepository) GetTeachers(ctx context.Context, limit, offset int) ([]domain.TeacherListRow, error) {
	return []domain.TeacherListRow{}, nil
}

func (r *DbRepository) GetTeacherByDomain(ctx context.Context, teacherDomain string) (domain.Teachers, error) {
	var teacher domain.Teachers
	if teacherDomain == mock.Unknown {
		return teacher, gorm.ErrRecordNotFound
	}
	return teacher, nil
}

// todo
func (r *DbRepository) CheckAuthByIdAndToken(ctx context.Context, userId, token string) error {
	return nil
}

// *
type CacheRepository struct{}

func NewCacheRepository() domain.AuthCacheRepository {
	return &CacheRepository{}
}
