package mock

import (
	"teacher-site/domain"

	"gorm.io/gorm"
)

const (
	Unknown       = "unknown"
	TeacherDomain = "domain"
	UserId        = "user_id"
	UserPassword  = "password"
	UserName      = "name"
	Email         = "mock@asia.edu.tw"
)

func GetUserToken(db *gorm.DB) (string, error) {
	auth := domain.Auths{UserId: UserId}
	result := db.Find(&auth)

	return auth.Token, result.Error
}
