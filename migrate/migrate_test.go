package database

import (
	"crypto/rand"
	"teacher-site/config"
	"teacher-site/domain"
	"teacher-site/mock"
	"teacher-site/pkg/database"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

var (
	conf = config.New()
	db   = database.NewDB("../pkg/database", conf.DB)
)

func TestMigrate(t *testing.T) {
	db.AutoMigrate(
		&domain.Auths{},
		&domain.Teachers{},
		&domain.Infos{},
		&domain.InfoBulletinBoards{},
		&domain.Courses{},
		&domain.BulletinBoards{},
		&domain.Slides{},
		&domain.Homeworks{},
	)
}

func TestInsertAll(t *testing.T) {
	hashPassword, salt := generalHashPassword(mock.UserPassword)

	data := domain.Auths{
		UserID:       mock.UserID,
		UserPassword: hashPassword,
		Salt:         salt,
		Teacher: domain.Teachers{
			Domain:    mock.TeacherDomain,
			Email:     mock.Email,
			NameZH:    mock.UserName,
			NameUS:    "",
			Office:    "",
			Call:      "",
			Education: "",
			Infos: []domain.Infos{
				{
					BulletinBoards: []domain.InfoBulletinBoards{
						{Content: ""},
					},
				},
			},
			Courses: []domain.Courses{
				{
					NameZH: "",
					NameUS: "",
					BulletinBoard: []domain.BulletinBoards{
						{Content: ""},
					},
					Slide: []domain.Slides{
						{Chapter: "", File: domain.File{Title: "", Type: ""}},
					},
					Homework: []domain.Homeworks{
						{Number: "", File: domain.File{Title: "", Type: ""}},
					},
				},
			},
		},
	}
	db.Create(&data)
}
func generalHashPassword(password string) (string, string) {
	var salt = make([]byte, conf.Secure.SaltSize)
	rand.Read(salt[:])
	saltPassword := append([]byte(password), salt...)
	hashPassword, _ := bcrypt.GenerateFromPassword(saltPassword, conf.Secure.HashCost)
	return string(hashPassword), string(salt)
}
