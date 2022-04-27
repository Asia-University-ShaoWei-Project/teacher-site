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

func NewDbRepository(db *gorm.DB, conf *config.DB) domain.CourseDbRepository {
	return &DbRepository{
		db:   db,
		conf: conf,
	}
}
func (r *DbRepository) GetByTeacherDomain(ctx context.Context, teacherDomain string) ([]domain.CourseResponse, error) {
	var coursesRes []domain.CourseResponse
	result := r.db.Model(&domain.Courses{}).
		Select("id", "name_zh", "name_us").
		Where("teacher_id=?", teacherDomain).
		Find(&coursesRes)
	err := checkErrAndExist(result)
	return coursesRes, err
}
func (r *DbRepository) GetContentByCourseId(ctx context.Context, courseId uint) (domain.GetCourseContentResponse, error) {
	var courseRes domain.GetCourseContentResponse
	course := domain.Courses{AutoModel: domain.AutoModel{Id: courseId}}

	result := r.db.Model(&course).
		Select("id, last_modified").
		Where("id=?", courseId).
		Find(&courseRes)
	if err := checkErrAndExist(result); err != nil {
		return courseRes, err
	}

	r.db.Model(&course).Select("id", "DATE(created_at) AS date", "content").Association("BulletinBoard").Find(&courseRes.BulletinBoard)
	r.db.Model(&course).Association("Slide").Find(&courseRes.Slide)
	r.db.Model(&course).Association("Homework").Find(&courseRes.Homework)

	return courseRes, nil
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

func (r *DbRepository) GetLastModifiedByCourseId(ctx context.Context, courseId uint) (domain.Courses, error) {
	var course domain.Courses
	result := r.db.Find(&course, courseId)
	err := checkErrAndExist(result)
	return course, err
}
