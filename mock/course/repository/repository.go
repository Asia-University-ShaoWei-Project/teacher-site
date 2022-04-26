package repository

import (
	"context"
	"teacher-site/domain"
)

// rdbms
type DbRepository struct{}

func NewDbRepository() domain.CourseDbRepository {
	return &DbRepository{}
}

func (r *DbRepository) GetByTeacherDomain(ctx context.Context, teacherDomain string) ([]domain.CourseResponse, error) {
	return []domain.CourseResponse{}, nil
}

// cache
type CacheRepository struct{}

func NewCacheRepository() domain.CourseCacheRepository {
	return &CacheRepository{}
}
