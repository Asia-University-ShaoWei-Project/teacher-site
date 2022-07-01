package domain

import "context"

type Auths struct {
	UserId       string `gorm:"primaryKey"`
	UserPassword string
	Salt         []byte
	Token        string
	Teacher      Teachers `gorm:"foreignKey:AuthId;references:UserId"`
	Time
}
type JwtInfoRequest struct {
	UserId string
	Domain string
}

//* usecase & repository
type AuthUsecase interface {
	Login(ctx context.Context, req *LoginRequest) (LoginResponse, error)
	Logout(ctx context.Context, id string) error
	Register(ctx context.Context, req *RegisterRequest) error
}
type AuthDbRepository interface {
	CreateTeacher(ctx context.Context, auth *Auths) error
	GetAccountByUserId(ctx context.Context, id string) (Auths, error)
	GetTeacherDomainByUserId(ctx context.Context, id string) (Teachers, error)
	UpdateTokenByUserId(ctx context.Context, id, token string) error
	DeleteTokenById(ctx context.Context, id string) error
	CheckUserExistByUserIdAndDomain(ctx context.Context, userId, teacherDomain string) error
}
type AuthCacheRepository interface {
}

//* request & response
type LoginRequest struct {
	UserId       string `json:"id" binding:"required"`
	UserPassword string `json:"password" binding:"required"`
}
type LoginResponse struct {
	Token  string
	Domain string
}

type RegisterRequest struct {
	UserId       string `json:"id" binding:"required"`
	UserPassword string `json:"password" binding:"required"`
	Domain       string `json:"domain" binding:"required"`
	Email        string `json:"email" binding:"required"`
	NameZh       string `json:"nameZh" binding:"required"`
}
