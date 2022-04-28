package repository

import (
	"context"
	"teacher-site/domain"
	"teacher-site/mock"
	"teacher-site/pkg/database"

	"gorm.io/gorm"
)

var (
	db = database.NewDB("../../../../pkg/database", mock.Conf.DB)
)

// rdbms
type DbRepository struct{}

func NewDbRepository() domain.AuthDbRepository {
	return &DbRepository{}
}

func (i *DbRepository) GetAccountByUserId(ctx context.Context, id string) (domain.Auths, error) {
	if id == mock.Unknown {
		return domain.Auths{}, gorm.ErrRecordNotFound
	}
	return testGetAccount(id)
}

//todo
func (i *DbRepository) GetTeacherDomainByUserId(ctx context.Context, id string) (domain.Teachers, error) {
	return domain.Teachers{}, nil
}

func (i *DbRepository) UpdateTokenByUserId(ctx context.Context, id, token string) error {
	if id == mock.Unknown {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// todo
func (i *DbRepository) DeleteTokenById(ctx context.Context, id string) error {
	return nil
}

func testGetAccount(id string) (domain.Auths, error) {
	auth := domain.Auths{UserId: id}
	result := db.Find(&auth)
	return auth, result.Error
}

// func (i *DbRepository) GetByTeacherDomain(ctx context.Context, teacherDomain string) (domain.Infos, error) {
// 	return domain.Infos{
// 		LastModified: CurrLastModidied,
// 	}, nil
// }
// func (i *DbRepository) GetBulletinsByInfoId(ctx context.Context, id uint) ([]domain.InfoBulletinResponse, error) {
// 	return []domain.InfoBulletinResponse{}, nil
// }
// func (i *DbRepository) GetLastModified(ctx context.Context, id uint) (string, error) {
// 	return CurrLastModidied, nil
// }
// func (i *DbRepository) Update(ctx context.Context, req *domain.UpdateInfoBulletinRequest) (domain.Infos, error) {
// 	return domain.Infos{}, nil
// }
// func (i *DbRepository) Delete(ctx context.Context, req *domain.DeleteInfoBulletinRequest) (domain.Infos, error) {
// 	return domain.Infos{}, nil
// }

// // cache
type CacheRepository struct{}

func NewCacheRepository() domain.AuthCacheRepository {
	return &CacheRepository{}
}

// func (i *InfoCacheRepository) Get(ctx context.Context, req *domain.GetInfoBulletinRequest) (string, error) {
// 	return "", nil
// }
// func (i *InfoCacheRepository) GetLastModified(ctx context.Context, teacherDomain string) (string, error) {
// 	return CurrLastModidied, nil

// }
// func (i *InfoCacheRepository) Update(ctx context.Context, req *domain.UpdateInfoBulletinRequest) error {
// 	return nil

// }
// func (i *InfoCacheRepository) UpdateInfoLastModified(ctx context.Context, req *domain.UpdateInfoBulletinRequest, lastModified string) error {
// 	return nil
// }
