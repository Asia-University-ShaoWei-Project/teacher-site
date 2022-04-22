package domain

import "context"

type Auths struct {
	UserID       string `gorm:"primaryKey"`
	UserPassword string
	Salt         string
	Token        string
	Teacher      Teachers `gorm:"foreignKey:AuthID;references:UserID"`
	Time
}

//* usecase & repository
type AuthUsecase interface {
	Login(ctx context.Context, req *LoginRequest) (string, error)
	Logout(ctx context.Context, id string) error
	// Create(ctx context.Context, req *ReqCreateAuth) (InfoBulletinBoards, error)
	// Get(ctx context.Context, req *ReqGetAuth) (ResGetInfo, error)
	// Update(ctx context.Context, req *ReqUpdateInfoBulletin) (ResUpdateInfo, error)
	// Delete(ctx context.Context, req *ReqDeleteInfo) error
}
type AuthDbRepository interface {
	// Login(ctx context.Context, id, password string) error
	GetAccountByUserId(ctx context.Context, id string) (Auths, error)
	UpdateTokenByUserId(ctx context.Context, id, token string) error
	DeleteToken(ctx context.Context, id string) error
	// DeleteToken(ctx context.Context) error
	// Create(ctx context.Context, req *ReqCreateAuth) (InfoBulletinBoards, error)
	// Get(ctx context.Context, req *ReqGetAuth) ([]GetInfoBulletin, error)
	// Update(ctx context.Context, req *ReqUpdateInfoBulletin) (InfoBulletinBoards, error)
	// Delete(ctx context.Context, req *ReqDeleteInfo) error
}
type AuthCacheRepository interface {
	// Get(ctx context.Context, req *ReqGetAuth) (string, error)
	// GetLastModified(ctx context.Context, req *ReqGetInfo) (string, error)
	// Update(ctx context.Context, req *ResGetInfo) error
	// UpdateInfoLastModified(ctx context.Context, req *ResGetInfo) error
}

//* request & response
type LoginRequest struct {
	UserID       string `json:"id" binding:"required"`
	UserPassword string `json:"password" binding:"required"`
}

// type RegisterRequest struct {
// 	UserID       string `json:"id" binding:"required"`
// 	UserPassword string `json:"password" binding:"required"`
// 	Domain       string `json:"domain" binding:"required"`
// 	Email        string `json:"email" binding:"required"`
// 	NameZH       string `json:"name_zh" binding:"required"`
// }
