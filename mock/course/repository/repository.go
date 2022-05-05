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
func (r *DbRepository) CreateBulletin(ctx context.Context, bulletin *domain.BulletinBoards) (string, error) {
	return "", nil
}
func (r *DbRepository) CreateSlide(ctx context.Context, slide *domain.Slides) (string, error) {
	return "", nil
}
func (r *DbRepository) CreateHomework(ctx context.Context, homework *domain.Homeworks) (string, error) {
	return "", nil
}

func (r *DbRepository) GetByTeacherDomain(ctx context.Context, teacherDomain string) ([]domain.CourseResponse, error) {
	return []domain.CourseResponse{}, nil
}
func (r *DbRepository) GetContentByCourseId(ctx context.Context, courseId uint) (domain.GetCourseContentResponse, error) {
	return domain.GetCourseContentResponse{}, nil
}
func (r *DbRepository) GetLastModifiedByCourseId(ctx context.Context, courseId uint) (string, error) {
	return "", nil
}

func (r *DbRepository) UpdateBulletinById(ctx context.Context, bulletin *domain.BulletinBoards) (string, error) {
	return "", nil
}
func (r *DbRepository) UpdateSlideById(ctx context.Context, slide *domain.Slides) (string, error) {
	return "", nil
}
func (r *DbRepository) UpdateHomeworkById(ctx context.Context, homework *domain.Homeworks) (string, error) {
	return "", nil
}

func (r *DbRepository) DeleteBulletinById(ctx context.Context, bulletin *domain.BulletinBoards) (string, error) {
	return "", nil
}
func (r *DbRepository) DeleteSlideById(ctx context.Context, slide *domain.Slides) (string, error) {
	return "", nil
}
func (r *DbRepository) DeleteHomeworkById(ctx context.Context, homework *domain.Homeworks) (string, error) {
	return "", nil
}

func (r *DbRepository) CheckByDomainAndCourseId(ctx context.Context, course *domain.Courses) error {
	return nil
}

type CacheRepository struct{}

func NewCacheRepository() domain.CourseCacheRepository {
	return &CacheRepository{}
}

func (c *CacheRepository) GetLastModifiedByCourseId(ctx context.Context, courseId uint) (string, error) {
	return "", nil
}
func (c *CacheRepository) UpdateLastModifiedByCourseId(ctx context.Context, courseId uint, lastModified string) error {
	return nil
}
