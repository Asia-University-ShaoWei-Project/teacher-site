package repository

import (
	"context"
	"teacher-site/domain"
)

const (
	CurrLastModidied = "1"
	LateLastModified = "0"
)

// rdbms
type InfoDbRepository struct{}

func NewInfoDbRepository() domain.InfoDbRepository {
	return &InfoDbRepository{}
}
func (i *InfoDbRepository) Create(ctx context.Context, req *domain.CreateInfoBulletinRequest) (domain.InfoBulletinBoards, error) {

	return domain.InfoBulletinBoards{}, nil
}

func (i *InfoDbRepository) GetByTeacherDomain(ctx context.Context, teacherDomain string) (domain.Infos, error) {
	return domain.Infos{
		LastModified: CurrLastModidied,
	}, nil
}
func (i *InfoDbRepository) GetBulletinsByInfoId(ctx context.Context, id uint) ([]domain.InfoBulletinResponse, error) {
	return []domain.InfoBulletinResponse{}, nil
}
func (i *InfoDbRepository) GetLastModified(ctx context.Context, id uint) (string, error) {
	return CurrLastModidied, nil
}
func (i *InfoDbRepository) Update(ctx context.Context, req *domain.UpdateInfoBulletinRequest) (domain.Infos, error) {
	return domain.Infos{}, nil
}
func (i *InfoDbRepository) Delete(ctx context.Context, req *domain.DeleteInfoBulletinRequest) (domain.Infos, error) {
	return domain.Infos{}, nil
}
func (i *InfoDbRepository) CheckByDomainAndId(ctx context.Context, teacherDomain string, infoId uint) error {
	return nil
}

// cache
type InfoCacheRepository struct{}

func NewInfoCacheRepository() domain.InfoCacheRepository {
	return &InfoCacheRepository{}
}
func (i *InfoCacheRepository) UpdateInfoLastModified(ctx context.Context, req *domain.UpdateInfoBulletinRequest, lastModified string) error {
	return nil
}
func (i *InfoCacheRepository) GetLastModifiedByTeacherDomain(ctx context.Context, teacherDomain string) (string, error) {
	return "", nil
}
func (i *InfoCacheRepository) UpdateLastModifiedByTeacherDomain(ctx context.Context, teacherDomain, lastModified string) error {
	return nil
}
