package usecase

import (
	"context"
	"teacher-site/domain"
	"teacher-site/mock"

	"gorm.io/gorm"
)

type Usecase struct{}

func NewUsecase() domain.PageUsecase {
	return &Usecase{}
}

func (u *Usecase) TeacherList(ctx context.Context, req *domain.TeacherListRequest) (domain.TeacherListResponse, error) {
	return domain.TeacherListResponse{}, nil
}

func (u *Usecase) Home(ctx context.Context, req *domain.HomeRequest) (domain.HomeResponse, error) {
	var res domain.HomeResponse
	if req.Domain == mock.Unknown {
		return res, gorm.ErrRecordNotFound
	}
	res = domain.HomeResponse{NameZh: mock.UserName}
	return res, nil
}

// todo
func (u *Usecase) Login(ctx context.Context) {}
