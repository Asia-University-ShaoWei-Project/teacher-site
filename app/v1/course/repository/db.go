package repository

import (
	"context"
	"teacher-site/config"
	"teacher-site/domain"
	utildb "teacher-site/pkg/database"

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

// ===== CREATE =====

func (r *DbRepository) CreateBulletin(ctx context.Context, bulletin *domain.BulletinBoards) (string, error) {
	course := domain.Courses{AutoModel: domain.AutoModel{Id: bulletin.CourseId}}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := updateCourseLastModified(tx, &course); err != nil {
			return err
		}
		result := tx.Create(&bulletin)
		if err := utildb.CheckErrAndExist(result); err != nil {
			return err
		}
		return nil
	})
	return course.LastModified, err
}

func (r *DbRepository) CreateSlide(ctx context.Context, slide *domain.Slides) (string, error) {
	course := domain.Courses{AutoModel: domain.AutoModel{Id: slide.CourseId}}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := updateCourseLastModified(tx, &course); err != nil {
			return err
		}
		result := tx.Create(&slide)
		if err := utildb.CheckErrAndExist(result); err != nil {
			return err
		}
		return nil
	})
	return course.LastModified, err
}

func (r *DbRepository) CreateHomework(ctx context.Context, homework *domain.Homeworks) (string, error) {
	course := domain.Courses{AutoModel: domain.AutoModel{Id: homework.CourseId}}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := updateCourseLastModified(tx, &course); err != nil {
			return err
		}
		result := tx.Create(&homework)
		if err := utildb.CheckErrAndExist(result); err != nil {
			return err
		}
		return nil
	})
	return course.LastModified, err
}

// ===== GET =====
func (r *DbRepository) GetByTeacherDomain(ctx context.Context, teacherDomain string) ([]domain.CourseResponse, error) {
	var coursesRes []domain.CourseResponse
	result := r.db.Model(&domain.Courses{}).
		Select("id", "name_zh", "name_us").
		Where("teacher_id=?", teacherDomain).
		Find(&coursesRes)
	err := utildb.CheckErrAndExist(result)
	return coursesRes, err
}
func (r *DbRepository) GetContentByCourseId(ctx context.Context, courseId uint) (domain.GetCourseContentResponse, error) {
	var courseRes domain.GetCourseContentResponse
	course := domain.Courses{AutoModel: domain.AutoModel{Id: courseId}}

	result := r.db.Model(&course).
		Select("id, last_modified").
		Where("id=?", courseId).
		Find(&courseRes)
	if err := utildb.CheckErrAndExist(result); err != nil {
		return courseRes, err
	}

	r.db.Model(&course).Select("id", "DATE(created_at) AS date", "content").Order("created_at DESC").Association("BulletinBoard").Find(&courseRes.BulletinBoard)
	r.db.Model(&course).Select("id, chapter, title, name").Association("Slide").Find(&courseRes.Slide)
	r.db.Model(&course).Select("id, number, title, name").Association("Homework").Find(&courseRes.Homework)

	return courseRes, nil
}

func (r *DbRepository) GetLastModifiedByCourseId(ctx context.Context, courseId uint) (string, error) {
	var course domain.Courses
	result := r.db.Select("last_modified").Find(&course, courseId)
	err := utildb.CheckErrAndExist(result)
	return course.LastModified, err
}

// ===== UPDATE =====

func (r *DbRepository) UpdateBulletinById(ctx context.Context, bulletin *domain.BulletinBoards) (string, error) {
	course := domain.Courses{AutoModel: domain.AutoModel{Id: bulletin.CourseId}}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := updateCourseLastModified(tx, &course); err != nil {
			return err
		}
		result := tx.Model(&bulletin).
			Where("id=? AND course_id=?", bulletin.AutoModel.Id, bulletin.CourseId).
			Update("content", bulletin.Content)
		err := utildb.CheckErrAndExist(result)
		return err
	})
	return course.LastModified, err
}

func (r *DbRepository) UpdateSlideById(ctx context.Context, slide *domain.Slides) (string, error) {
	course := domain.Courses{AutoModel: domain.AutoModel{Id: slide.CourseId}}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := updateCourseLastModified(tx, &course); err != nil {
			return err
		}
		result := tx.Model(&slide).
			Where("id=? AND course_id=?", slide.AutoModel.Id, slide.CourseId).
			Updates(&slide)
		err := utildb.CheckErrAndExist(result)
		return err
	})
	return course.LastModified, err
}

func (r *DbRepository) UpdateHomeworkById(ctx context.Context, homework *domain.Homeworks) (string, error) {
	course := domain.Courses{AutoModel: domain.AutoModel{Id: homework.CourseId}}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := updateCourseLastModified(tx, &course); err != nil {
			return err
		}
		result := tx.Model(&homework).
			Where("id=? AND course_id=?", homework.AutoModel.Id, homework.CourseId).
			Updates(&homework)
		err := utildb.CheckErrAndExist(result)
		return err
	})
	return course.LastModified, err
}

// ===== DELETE =====

func (r *DbRepository) DeleteBulletinById(ctx context.Context, bulletin *domain.BulletinBoards) (string, error) {
	course := domain.Courses{AutoModel: domain.AutoModel{Id: bulletin.CourseId}}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := updateCourseLastModified(tx, &course); err != nil {
			return err
		}
		// todo: should i to find the slide by this step?
		result := tx.Model(&bulletin).
			Where(`id=? AND course_id=?`, bulletin.AutoModel.Id, bulletin.CourseId).
			Find(&bulletin)
		if err := utildb.CheckErrAndExist(result); err != nil {
			return err
		}
		result = tx.Model(&bulletin).Delete(&bulletin)

		return utildb.CheckErrAndExist(result)
	})
	return course.LastModified, err
}

func (r *DbRepository) DeleteSlideById(ctx context.Context, slide *domain.Slides) (string, error) {
	course := domain.Courses{AutoModel: domain.AutoModel{Id: slide.CourseId}}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := updateCourseLastModified(tx, &course); err != nil {
			return err
		}
		// todo: should i to find the slide by this step?
		result := tx.Model(&slide).
			Where(`id=? AND course_id=?`, slide.AutoModel.Id, slide.CourseId).
			Find(&slide)
		if err := utildb.CheckErrAndExist(result); err != nil {
			return err
		}
		result = tx.Model(&slide).Delete(&slide)

		return utildb.CheckErrAndExist(result)
	})
	return course.LastModified, err
}

func (r *DbRepository) DeleteHomeworkById(ctx context.Context, homework *domain.Homeworks) (string, error) {
	course := domain.Courses{AutoModel: domain.AutoModel{Id: homework.CourseId}}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := updateCourseLastModified(tx, &course); err != nil {
			return err
		}
		// todo: should i to find the slide by this step?
		result := tx.Model(&homework).
			Where(`id=? AND course_id=?`, homework.AutoModel.Id, homework.CourseId).
			Find(&homework)
		if err := utildb.CheckErrAndExist(result); err != nil {
			return err
		}
		result = tx.Model(&homework).Delete(&homework)

		return utildb.CheckErrAndExist(result)
	})
	return course.LastModified, err
}

func (r *DbRepository) CheckByDomainAndCourseId(ctx context.Context, course *domain.Courses) error {
	result := r.db.Model(&course).Where(`id=? AND teacher_id=?`, course.AutoModel.Id, course.TeacherId).Find(&course)
	return utildb.CheckErrAndExist(result)
}

func updateCourseLastModified(tx *gorm.DB, course *domain.Courses) error {
	lastModifiedTime := utildb.NewLastModifiedTime()
	result := tx.Model(&course).Update("last_modified", lastModifiedTime)
	return utildb.CheckErrAndExist(result)
}
