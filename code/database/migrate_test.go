package database

import (
	"context"
	"teacher-site/logsrv"
	"teacher-site/model"
	"testing"
)

var (
	ctx    = context.Background()
	logger = logsrv.NewLogrus(ctx)
	conf   = model.NewTMPConfig()

	db = NewSqlite(".", logger)
)

func TestMigrate(t *testing.T) {
	db.Migrate(
		ctx,
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
	db.Create(ctx, &data)
}
