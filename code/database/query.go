package database

import (
	"context"
	"teacher-site/model"

	"gorm.io/gorm"
)

func (db *DB) GetCourseLastUpdated(ctx context.Context, courseID uint) (string, error) {
	var course model.Courses
	result := db.orm.Where("id=?", courseID).Find(&course)
	err := checkErrAndRecord(result)
	return course.LastUpdated, err
}
func (db *DB) GetCourseContent(ctx context.Context, courseID uint) (model.Courses, error) {
	var course model.Courses
	if err := db.orm.Where("id=?", courseID).Find(&course).Error; err != nil {
		return course, err
	}
	db.orm.Model(&course).Association("BulletinBoard").Find(&course.BulletinBoard)
	db.orm.Model(&course).Association("Slide").Find(&course.Slide)
	db.orm.Model(&course).Association("Homework").Find(&course.Homework)
	return course, nil
}
func (db *DB) GetInit(ctx context.Context, init *model.Init, domain string) error {
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
		return result.Error
	}
	result = db.orm.Raw(coursesSQL, domain).Scan(&coursesName)
	if result.Error != nil {
		return result.Error
	}
	*init = model.Init{
		Courses:      coursesName,
		Informations: infos,
	}
	return nil
}

func (db *DB) GetAuth(ctx context.Context, auth *model.Auths) error {
	result := db.orm.Where("user_id=?", auth.UserID).Find(auth)
	err := checkErrAndRecord(result)
	return err
}
func (db *DB) GetAuthToken(ctx context.Context, auth *model.Auths) (string, error) {
	result := db.orm.Where("user_id=?", auth.UserID).Find(auth)
	err := checkErrAndRecord(result)
	return auth.Token, err
}
func (db *DB) DomainIsExist(ctx context.Context, domain string) error {
	result := db.orm.Where("domain=?", domain).Find(&model.Teachers{})
	err := checkErrAndRecord(result)
	if err != nil {
		return err
	}
	return nil
}
func (db *DB) UserIsExist(ctx context.Context, userID string) error {

	result := db.orm.Where("user_id=?", userID).Find(&model.Auths{})
	err := checkErrAndRecord(result)
	if err != nil {
		return err
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
func checkErrAndRecord(result *gorm.DB) error {
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
