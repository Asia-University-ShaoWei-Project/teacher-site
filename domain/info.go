package domain

import (
	"context"
)

const BulletinDateFormat = "2006-01-02"

type Infos struct {
	AutoModel      AutoModel            `gorm:"embedded"`
	TeacherID      string               `gorm:"index"`
	BulletinBoards []InfoBulletinBoards `gorm:"foreignKey:InfoID;references:Id"`
	LastModified   string
}

type InfoBulletinBoards struct {
	AutoModel AutoModel `gorm:"embedded"`
	InfoID    uint
	Content   string
}

// usecase & repository
type InfoUsecase interface {
	Create(ctx context.Context, req *CreateInfoBulletinRequest) (CreateInfoBulletinResponse, error)
	Get(ctx context.Context, req *GetInfoBulletinRequest) (GetInfoBulletinResponse, error)
	Update(ctx context.Context, req *UpdateInfoBulletinRequest) (UpdateInfoBulletinResponse, error)
	Delete(ctx context.Context, req *DeleteInfoBulletinRequest) (DeleteInfoBulletinResponse, error)
}
type InfoDbRepository interface {
	Create(ctx context.Context, req *CreateInfoBulletinRequest) (InfoBulletinBoards, error)
	GetByTeacherDomain(ctx context.Context, teacherDomain string) (Infos, error)
	GetBulletinsByInfoId(ctx context.Context, id uint) ([]InfoBulletinResponse, error)
	GetLastModified(ctx context.Context, id uint) (string, error)
	Update(ctx context.Context, req *UpdateInfoBulletinRequest) (Infos, error)
	Delete(ctx context.Context, req *DeleteInfoBulletinRequest) (Infos, error)
}
type InfoCacheRepository interface {
	Get(ctx context.Context, req *GetInfoBulletinRequest) (string, error)
	GetLastModified(ctx context.Context, teacherDomain string) (string, error)
	Update(ctx context.Context, req *UpdateInfoBulletinRequest) error
	UpdateInfoLastModified(ctx context.Context, req *UpdateInfoBulletinRequest, lastModified string) error
}

//* request & response

// todo: binding:"required"
type CreateInfoBulletinRequest struct {
	TeacherDomainRequest
	InfoID  uint   `uri:"infoId"`
	Content string `json:"content"`
}
type CreateInfoBulletinResponse struct {
	Id           uint   `json:"id"`
	Date         string `json:"date"`
	LastModified string `json:"lastModified"`
}

type GetInfoBulletinRequest struct {
	TeacherDomainRequest
	LastModified string `form:"lastModified"`
}
type GetInfoBulletinResponse struct {
	Id           uint                   `json:"id"`
	LastModified string                 `json:"lastModified"`
	Bulletins    []InfoBulletinResponse `json:"bulletins"`
}

func (i *GetInfoBulletinResponse) SetID(id uint) {
	i.Id = id
}
func (i *GetInfoBulletinResponse) SetLastModified(lastModified string) {
	i.LastModified = lastModified
}
func (i *GetInfoBulletinResponse) SetBulletins(bulletins []InfoBulletinResponse) {
	i.Bulletins = bulletins
}

type UpdateInfoBulletinRequest struct {
	TeacherDomainRequest
	InfoID     uint   `uri:"infoId"`
	BulletinID uint   `uri:"bulletinId"`
	Content    string `json:"content"`
}
type UpdateInfoBulletinResponse struct {
	LastModified string `json:"lastModified"`
}
type DeleteInfoBulletinRequest struct {
	TeacherDomainRequest
	InfoID     uint `uri:"infoId" binding:"required"`
	BulletinID uint `uri:"bulletinId" binding:"required"`
}
type DeleteInfoBulletinResponse struct {
	LastModified string `json:"lastModified"`
}
type InfoBulletinResponse struct {
	Id      uint   `json:"id"`
	Date    string `json:"date"`
	Content string `json:"content"`
}
