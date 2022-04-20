package usecase

import (
	"context"
	"encoding/json"
	"teacher-site/domain"

	log "github.com/sirupsen/logrus"
)

type AuthUsecase struct {
	dbRepository    domain.AuthDbRepository
	cacheRepository domain.AuthCacheRepository
	log             *log.Logger
}

func NewAuthUsecase(dbRepo domain.AuthDbRepository, cacheRepo domain.AuthCacheRepository, logger *log.Logger) domain.AuthUsecase {
	return &AuthUsecase{
		dbRepository:    dbRepo,
		cacheRepository: cacheRepo,
		log:             logger,
	}
}

func (auth *AuthUsecase) Login(ctx context.Context, req *domain.ReqCreateInfo) (domain.InfoBulletinBoards, error) {
	// todo: use TX to read and write the last_updated
	return auth.dbRepository.Create(ctx, req)
}

func (auth *AuthUsecase) Create(ctx context.Context, req *domain.ReqCreateInfo) (domain.InfoBulletinBoards, error) {
	// todo: use TX to read and write the last_updated
	return auth.dbRepository.Create(ctx, req)
}
func (auth *AuthUsecase) Get(ctx context.Context, req *domain.ReqGetInfo) (domain.ResGetInfo, error) {
	var (
		resInfo  domain.ResGetInfo
		bulletin []domain.GetInfoBulletin
	)

	lastModified, err := auth.cacheRepository.GetLastModified(ctx, req)
	// Not need to get new data if request last modified value is equal the last modified of repository value
	if err == nil && (req.LastModified == lastModified) {
		auth.log.Info("Not need to update")
		return resInfo, err
	}
	resInfo.SetLastModified(lastModified)
	// Get by cache
	data, err := auth.cacheRepository.Get(ctx, req)
	if err == nil {
		// return resInfo, err
		if err = json.Unmarshal([]byte(data), &bulletin); err == nil {
			resInfo.SetBulletins(bulletin)
			return resInfo, nil
		}
		// todo: parse json error handle
		auth.log.Error(err)
	}
	// todo: cache error handle
	auth.log.Error(err)
	// Get from database
	bulletins, err := auth.dbRepository.Get(ctx, req)
	if err != nil {
		auth.log.Error(err)
		return resInfo, err
	}
	resInfo.SetBulletins(bulletins)
	// todo: Update the info in cache & cache update error handle
	// if err := auth.cacheRepository.Update(ctx, &resInfo); err != nil {
	// auth.log.Error(err)
	// }
	return resInfo, nil
}
func (auth *AuthUsecase) Update(ctx context.Context, req *domain.ReqUpdateInfoBulletin) (domain.ResUpdateInfo, error) {
	var res domain.ResUpdateInfo
	bulletin, err := auth.dbRepository.Update(ctx, req)
	if err != nil {
		auth.log.Error(err)
		return res, err
	}
	// todo: update the data in redis
	// err = auth.cacheRepository.Update(ctx, req*domain.ResGetInfo)
	res = domain.ResUpdateInfo{
		ID:   bulletin.AutoModel.ID,
		Date: bulletin.AutoModel.CreatedAT.Format(domain.BulletinDateFormat),
	}
	return res, nil
}
func (auth *AuthUsecase) Delete(ctx context.Context, req *domain.ReqDeleteInfo) error {
	// todo: update the data in redis
	return auth.dbRepository.Delete(ctx, req)
}
