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
func (r *DbRepository) GetContentByCourseId(ctx context.Context, courseId uint) (domain.GetCourseContentResponse, error) {
	return domain.GetCourseContentResponse{}, nil
}
func (r *DbRepository) GetLastModifiedByCourseId(ctx context.Context, courseId uint) (domain.Courses, error) {
	return domain.Courses{}, nil
}

// cache
type CacheRepository struct{}

func NewCacheRepository() domain.CourseCacheRepository {
	return &CacheRepository{}
}
