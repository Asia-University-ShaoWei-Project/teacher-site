package mock

import (
	"teacher-site/domain"
	"teacher-site/pkg/util"

	"gorm.io/gorm"
)

const (
	Unknown       = "unknown"
	TeacherDomain = "domain"
	UserId        = "user_id"
	UserPassword  = "password"
	UserNameZh    = "名字"
	UserNameUs    = "name"
	Email         = "mock@asia.edu.tw"
)

func GenerateAuth() domain.Auths {
	return domain.Auths{
		UserId:       UserId,
		UserPassword: util.GeneralHashPassword(Ctx, UserPassword, Conf.Secure.Salt, Conf.Secure.HashCost),
		Salt:         string(util.GeneralSalt(Ctx, Conf.Secure.SaltSize)),
		Teacher: domain.Teachers{
			Domain:  TeacherDomain,
			Email:   Email,
			NameZh:  UserNameZh,
			NameUs:  UserNameUs,
			Infos:   []domain.Infos{},
			Courses: []domain.Courses{},
		}}
}
func GetUserToken(db *gorm.DB) (string, error) {
	auth := domain.Auths{UserId: UserId}
	result := db.Find(&auth)

	return auth.Token, result.Error
}
