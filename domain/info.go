package domain

import (
	"context"
)

const BulletinDateFormat = "2006-01-02"

type Infos struct {
	AutoModel      AutoModel            `gorm:"embedded"`
	TeacherId      string               `gorm:"index"`
	BulletinBoards []InfoBulletinBoards `gorm:"foreignKey:InfoId;references:Id"`
	LastModified   string
}
type InfoBulletinBoards struct {
	AutoModel AutoModel `gorm:"embedded"`
	InfoId    uint
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
	CheckByDomainAndId(ctx context.Context, teacherDomain string, infoId uint) error
	GetBulletinsByInfoId(ctx context.Context, id uint) ([]InfoBulletinResponse, error)
	GetLastModified(ctx context.Context, id uint) (string, error)
	Update(ctx context.Context, req *UpdateInfoBulletinRequest) (Infos, error)
	Delete(ctx context.Context, req *DeleteInfoBulletinRequest) (Infos, error)
}
type InfoCacheRepository interface {
	GetLastModifiedByTeacherDomain(ctx context.Context, teacherDomain string) (string, error)
	UpdateLastModifiedByTeacherDomain(ctx context.Context, teacherDomain, lastModified string) error
}

//* request & response

type CreateInfoBulletinRequest struct {
	TeacherDomainRequest
	InfoId  uint   `uri:"infoId"`
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

func (i *GetInfoBulletinResponse) SetId(id uint) {
	i.Id = id
}
func (i *GetInfoBulletinResponse) SetLastModified(lastModified string) {
	i.LastModified = lastModified
}
func (i *GetInfoBulletinResponse) SetBulletins(bulletins []InfoBulletinResponse) {
	i.Bulletins = bulletins
}

type UpdateInfoBulletinUriRequest struct {
	TeacherDomainRequest
	InfoId     uint `uri:"infoId" binding:"required"`
	BulletinId uint `uri:"bulletinId" binding:"required"`
}
type UpdateInfoBulletinRequest struct {
	TeacherDomain string
	InfoId        uint
	BulletinId    uint
	Content       string `json:"content" binding:"required"`
}

func (u *UpdateInfoBulletinRequest) SetupUri(req *UpdateInfoBulletinUriRequest) {
	u.TeacherDomain = req.TeacherDomain
	u.InfoId = req.InfoId
	u.BulletinId = req.BulletinId
}

type UpdateInfoBulletinResponse struct {
	LastModified string `json:"lastModified"`
}
type DeleteInfoBulletinRequest struct {
	TeacherDomainRequest
	InfoId     uint `uri:"infoId" binding:"required"`
	BulletinId uint `uri:"bulletinId" binding:"required"`
}
type DeleteInfoBulletinResponse struct {
	LastModified string `json:"lastModified"`
}
type InfoBulletinResponse struct {
	Id      uint   `json:"id"`
	Date    string `json:"date"`
	Content string `json:"content"`
}
