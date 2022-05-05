package usecase

import (
	"context"
	"fmt"
	"strconv"
	"teacher-site/domain"
	"teacher-site/mock"

	"gorm.io/gorm"
)

type Usecase struct{}

func NewUsecase() domain.PageUsecase {
	return &Usecase{}
}

func (u *Usecase) TeacherList(ctx context.Context, req *domain.TeacherListRequest) (domain.TeacherListResponse, error) {
	var _i string
	teachers := make([]domain.TeacherResponse, 0, 10)
	for i := 0; i < cap(teachers); i++ {
		_i = strconv.Itoa(i)
		teachers = append(teachers, domain.TeacherResponse{
			Domain: "domain_" + _i,
			NameZh: "name_zh_" + _i,
			NameUs: "name_uh_" + _i,
		})
	}
	fmt.Println(teachers)
	res := domain.TeacherListResponse{
		Teachers: teachers,
	}
	fmt.Println(res)

	return res, nil
}

func (u *Usecase) Home(ctx context.Context, req *domain.HomeRequest) (domain.HomeResponse, error) {
	var res domain.HomeResponse
	if req.Domain == mock.Unknown {
		return res, gorm.ErrRecordNotFound
	}
	res = domain.HomeResponse{NameZh: mock.UserNameZh}
	return res, nil
}

// todo
func (u *Usecase) Login(ctx context.Context, userId, token string) error {
	return nil
}
