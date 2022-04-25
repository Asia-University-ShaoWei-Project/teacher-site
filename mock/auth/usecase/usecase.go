package usecase

import (
	"context"
	"teacher-site/domain"
)

type Usecase struct{}

func NewUsecase() domain.AuthUsecase {
	return &Usecase{}
}

// todo
func (u *Usecase) Login(ctx context.Context, req *domain.LoginRequest) (domain.LoginResponse, error) {
	return domain.LoginResponse{}, nil
}

// todo

func (u *Usecase) Logout(ctx context.Context, id string) error {
	return nil
}

// func (u *Usecase) Get(ctx context.Context, req *domain.GetInfoBulletinRequest) (domain.GetInfoBulletinResponse, error) {
// 	return domain.GetInfoBulletinResponse{}, nil
// }
// func (u *Usecase) Update(ctx context.Context, req *domain.UpdateInfoBulletinRequest) (domain.UpdateInfoBulletinResponse, error) {
// 	return domain.UpdateInfoBulletinResponse{}, nil
// }
// func (u *Usecase) Delete(ctx context.Context, req *domain.DeleteInfoBulletinRequest) (domain.DeleteInfoBulletinResponse, error) {
// 	return domain.DeleteInfoBulletinResponse{}, nil
// }
