package database

import (
	"context"
	"crypto/rand"
	"teacher-site/logsrv"
	"teacher-site/mock"
	"teacher-site/model"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

var (
	ctx    = context.Background()
	logger = logsrv.NewLogrus(ctx)
	db     = NewSqlite("./", logger)
	conf   = model.NewMockServiceConfig()
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
	hashPassword, salt := generalHashPassword()

	data := model.Auths{
		UserID:       mock.UserID,
		UserPassword: hashPassword,
		Salt:         salt,
		Teacher: model.Teachers{
			Domain:    mock.Domain,
			Email:     mock.Email,
			NameZH:    mock.UserName,
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
func generalHashPassword() (string, string) {
	var salt = make([]byte, conf.SaltSize)
	rand.Read(salt[:])
	saltPassword := append([]byte(password), salt...)
	hashPassword, _ := bcrypt.GenerateFromPassword(saltPassword, conf.HashCost)
	return string(hashPassword), string(salt)
}
