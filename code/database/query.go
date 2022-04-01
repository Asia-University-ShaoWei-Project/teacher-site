package database

import "teacher-site/model"

func (db *DB) GetCourseWithContent(id uint) *model.Courses {
	var course *model.Courses
	// todo error handle
	db.orm.Where("id=?", id).Find(course)
	db.orm.Model(course).Association("BulletinBoard").Find(course.BulletinBoard)
	db.orm.Model(course).Association("Slide").Find(course.Slide)
	db.orm.Model(course).Association("Homework").Find(course.Homework)
	return course
}
func (db *DB) GetInit(domain string) (*model.Init, error) {
	var infos *[]model.Informations
	var coursesName *[]model.Courses
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

	db.orm.Raw(infoSQL, domain).Scan(infos)
	db.orm.Raw(coursesSQL, domain).Scan(coursesName)
	init := &model.Init{
		Courses:      *coursesName,
		Informations: *infos,
	}
	return init, nil
}

// todo testing
func (db *DB) VerifyAuthAndGetTeacher(auth *model.BindAuth) (*model.Teachers, error) {
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
