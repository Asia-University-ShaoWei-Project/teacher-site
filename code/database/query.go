package database

import (
	"context"
	"teacher-site/model"

	"gorm.io/gorm"
)

func (db *DB) GetCourseWithContent(ctx context.Context, id uint) *model.Courses {
	var course *model.Courses
	// todo error handle
	db.orm.Where("id=?", id).Find(course)
	db.orm.Model(course).Association("BulletinBoard").Find(course.BulletinBoard)
	db.orm.Model(course).Association("Slide").Find(course.Slide)
	db.orm.Model(course).Association("Homework").Find(course.Homework)
	return course
}
func (db *DB) GetInit(ctx context.Context, domain string) (*model.Init, error) {
	var infos []model.Informations
	var coursesName []model.Courses
	var result *gorm.DB
	infoSQL := `
	SELECT info, created_date
	FROM informations i
	INNER JOIN teachers t
	ON i.teacher_id = t.domain
	WHERE t.domain = ?
	`
	coursesSQL := `
	SELECT c.id, c.name_zh, c.name_us
	FROM courses c
	INNER JOIN teachers t
	ON c.teacher_id = t.domain
	WHERE t.domain=?`
	result = db.orm.Raw(infoSQL, domain).Scan(&infos)
	if result.Error != nil {
		db.log.Error(result.Error)
		return &model.Init{}, result.Error
	}
	result = db.orm.Raw(coursesSQL, domain).Scan(&coursesName)
	if result.Error != nil {
		db.log.Error(result.Error)
		return &model.Init{}, result.Error
	}
	init := &model.Init{
		Courses:      coursesName,
		Informations: infos,
	}
	return init, nil
}

func (db *DB) GetAuth(ctx context.Context, auth *model.Auths) error {
	result := db.orm.Where("user_id=?", auth.UserID).Find(auth)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (db *DB) FindDomain(ctx context.Context, domain *string) error {
	result := db.orm.Where("domain=?", *domain).Find(&model.Teachers{})
	if result.Error != nil {
		db.log.Error(result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		db.log.Warn(gorm.ErrRecordNotFound)
		return gorm.ErrRecordNotFound
	}
	return nil
}

// todo testing
func (db *DB) VerifyAuthAndGetTeacher(ctx context.Context, auth *model.BindAuth) (*model.Teachers, error) {
	// sql := `
	// SELECT t.domain
	// FROM teachers t
	// INNER JOIN auths a
	// ON t.auth_id = a.user_id
	// WHERE a.user_id =? AND a.user_password=?`
	var teacher *model.Teachers
	err := db.orm.Where(auth).Association("Teacher").Find(teacher)
	return teacher, err
}
