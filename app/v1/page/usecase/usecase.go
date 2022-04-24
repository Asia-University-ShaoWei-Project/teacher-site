package usecase

import (
	"context"
	"teacher-site/config"
	"teacher-site/domain"

	"github.com/sirupsen/logrus"
)

type Usecase struct {
	DbRepository    domain.PageDbRepository
	CacheRepository domain.PageCacheRepository
	conf            *config.Config
	log             *logrus.Logger
}

func NewUsecase(dbRepo domain.PageDbRepository, cacheRepo domain.PageCacheRepository, conf *config.Config, logger *logrus.Logger) domain.PageUsecase {
	return &Usecase{
		DbRepository:    dbRepo,
		CacheRepository: cacheRepo,
		conf:            conf,
		log:             logger,
	}
}

// todo
func (u *Usecase) TeacherList(ctx context.Context, req *domain.TeacherListRequest) (domain.TeacherListResponse, error) {
	var res domain.TeacherListResponse
	// todo: get teachers -> make a list to response for teacher-list html
	limit := u.conf.Limit.TeacherListPageCount
	offset := (limit * req.Page) - limit
	teachers, err := u.DbRepository.GetTeachers(ctx, limit, offset)
	if err != nil {
		return res, err
	}
	res = domain.TeacherListResponse{List: teachers}
	return res, nil

}
func (u *Usecase) Home(ctx context.Context, req *domain.HomeRequest) (domain.HomeResponse, error) {
	//todo the domain is exist
	var res domain.HomeResponse
	teacher, err := u.DbRepository.GetTeacherByDomain(ctx, req.Domain)
	if err != nil {
		return res, err
	}
	res = domain.HomeResponse{
		Email:     teacher.Email,
		NameZh:    teacher.NameZH,
		NameUs:    teacher.NameUS,
		Office:    teacher.Office,
		Call:      teacher.Call,
		Education: teacher.Education,
	}
	return res, nil
}
func (u *Usecase) Login(ctx context.Context) {
	//todo: usecase to check the userID and token

}
