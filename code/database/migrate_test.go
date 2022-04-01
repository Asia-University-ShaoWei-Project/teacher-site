package database

import (
	"teacher-site/model"
	"testing"
)

var db = NewSqlite(".")

func TestMigrate(t *testing.T) {
	db.Migrate(
		&model.Auths{},
		&model.Teachers{},
		&model.Informations{},
		&model.Courses{},
		&model.BulletinBoard{},
		&model.Slides{},
		&model.Homeworks{},
	)
}

func TestInsertAll(t *testing.T) {

	data := model.Auths{
		UserID:       "",
		UserPassword: "",
		Teacher: model.Teachers{
			Domain:    "",
			Email:     "",
			NameZH:    "",
			NameUS:    "",
			Office:    "",
			Call:      "",
			Education: "",
			Informations: []model.Informations{
				{
					Info: "",
				},
			},
			Courses: []model.Courses{
				{
					NameZH: "",
					NameUS: "",
					BulletinBoard: []model.BulletinBoard{
						{Info: ""},
					},
					Slide: []model.Slides{
						{Chapter: "", File: model.File{Title: "", Type: ""}},
					},
					Homework: []model.Homeworks{
						{Number: "", File: model.File{Title: "", Type: ""}},
					},
				},
			},
		},
	}
	db.Create(&data)
}
