package repository

import (
	"context"
	"teacher-site/domain"
)

var (
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

func (i *InfoCacheRepository) Get(ctx context.Context, req *domain.GetInfoBulletinRequest) (string, error) {
	return "", nil
}
func (i *InfoCacheRepository) GetLastModified(ctx context.Context, teacherDomain string) (string, error) {
	return CurrLastModidied, nil

}
func (i *InfoCacheRepository) Update(ctx context.Context, req *domain.UpdateInfoBulletinRequest) error {
	return nil

}
func (i *InfoCacheRepository) UpdateInfoLastModified(ctx context.Context, req *domain.UpdateInfoBulletinRequest, lastModified string) error {
	return nil
}
