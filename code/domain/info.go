package domain

import (
	"context"
)

const BulletinDateFormat = "2006-01-02"

type Infos struct {
	AutoModel      AutoModel            `gorm:"embedded"`
	TeacherID      string               `gorm:"index"`
	BulletinBoards []InfoBulletinBoards `gorm:"foreignKey:InfoID;references:ID"`
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
	InfoID  uint   `uri:"info_id"`
	Content string `json:"content"`
}
type CreateInfoBulletinResponse struct {
	ID           uint   `json:"id"`
	Date         string `json:"date"`
	LastModified string `json:"last_modified"`
}

type GetInfoBulletinRequest struct {
	TeacherDomainRequest
	LastModified string `form:"last_modified"`
}
type GetInfoBulletinResponse struct {
	ID           uint                   `json:"id"`
	LastModified string                 `json:"last_modified"`
	Bulletins    []InfoBulletinResponse `json:"bulletins"`
}

func (i *GetInfoBulletinResponse) SetID(id uint) {
	i.ID = id
}
func (i *GetInfoBulletinResponse) SetLastModified(lastModified string) {
	i.LastModified = lastModified
}
func (i *GetInfoBulletinResponse) SetBulletins(bulletins []InfoBulletinResponse) {
	i.Bulletins = bulletins
}

type UpdateInfoBulletinRequest struct {
	TeacherDomainRequest
	InfoID     uint   `uri:"info_id"`
	BulletinID uint   `uri:"bulletin_id"`
	Content    string `json:"content"`
}
type UpdateInfoBulletinResponse struct {
	LastModified string `json:"last_modified"`
}
type DeleteInfoBulletinRequest struct {
	TeacherDomainRequest
	InfoID     uint `uri:"info_id" binding:"required"`
	BulletinID uint `uri:"bulletin_id" binding:"required"`
}
type DeleteInfoBulletinResponse struct {
	LastModified string `json:"last_modified"`
}
type InfoBulletinResponse struct {
	ID      uint   `json:"id"`
	Date    string `json:"date"`
	Content string `json:"content"`
}
