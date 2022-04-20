package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"teacher-site/domain"

	log "github.com/sirupsen/logrus"
)

var errUnnecessaryUpdate = errors.New("the data is up to date")

type InfoUsecase struct {
	dbRepository    domain.InfoDbRepository
	cacheRepository domain.InfoCacheRepository
	log             *log.Logger
}

func NewInfoUsecase(dbRepo domain.InfoDbRepository, cacheRepo domain.InfoCacheRepository, logger *log.Logger) domain.InfoUsecase {
	return &InfoUsecase{
		dbRepository:    dbRepo,
		cacheRepository: cacheRepo,
		log:             logger,
	}
}

func (i *InfoUsecase) Create(ctx context.Context, req *domain.CreateInfoBulletinRequest) (domain.CreateInfoBulletinResponse, error) {
	var res domain.CreateInfoBulletinResponse
	bulletin, err := i.dbRepository.Create(ctx, req)
	if err != nil {
		return res, err
	}
	lastModified, err := i.dbRepository.GetLastModified(ctx, req.InfoID)
	if err != nil {
		// todo: get last modified error
		i.log.Error(err)
	}
	// todo: update the data in redis
	res = domain.CreateInfoBulletinResponse{
		ID:           bulletin.AutoModel.ID,
		Date:         bulletin.AutoModel.CreatedAT.Format(domain.BulletinDateFormat),
		LastModified: lastModified,
	}
	return res, nil
}

func (i *InfoUsecase) Get(ctx context.Context, req *domain.GetInfoBulletinRequest) (domain.GetInfoBulletinResponse, error) {
	var (
		res      domain.GetInfoBulletinResponse
		bulletin []domain.InfoBulletinResponse
	)
	info, err := i.dbRepository.GetByTeacherDomain(ctx, req.TeacherDomain)
	// lastModified, err := i.dbRepository.GetLastModified(ctx, req.TeacherDomain)
	if err != nil {
		return res, err
	}
	// Unnecessary to get new data if request last modified value is equal the last modified of repository value
	if req.LastModified == info.LastModified {
		i.log.Error(errUnnecessaryUpdate)
		return res, errUnnecessaryUpdate
	}
	res.SetLastModified(info.LastModified)
	res.SetID(info.AutoModel.ID)
	// Get by cache
	data, err := i.cacheRepository.Get(ctx, req)
	if err != nil {
		// todo: make error handle of "get data"(by cache)
		i.log.Error(err)
	}
	if err == nil {
		if err = json.Unmarshal([]byte(data), &bulletin); err == nil {
			// return resInfo, err
			res.SetBulletins(bulletin)
			return res, nil
		}
		// todo: make error handle of "parse json"(by cache)
		i.log.Error(err)
	}

	// Get from database
	bulletins, err := i.dbRepository.GetBulletinsByInfoId(ctx, info.AutoModel.ID)
	if err != nil {
		// todo: make error handle of "get data"(by RDBMS)
		i.log.Error(err)
		return res, err
	}
	res.SetBulletins(bulletins)
	// todo: update the data and last_modified in cache
	// if err := i.cacheRepository.Update(ctx, &resInfo); err != nil {
	// todo: error handle of updating data(by cache)
	// i.log.Error(err)
	// }
	return res, nil
}
func (i *InfoUsecase) Update(ctx context.Context, req *domain.UpdateInfoBulletinRequest) (domain.UpdateInfoBulletinResponse, error) {
	var res domain.UpdateInfoBulletinResponse
	info, err := i.dbRepository.Update(ctx, req)
	if err != nil {
		// todo: make error handle of "update the data"(by RDBMS)
		i.log.Error(err)
		return res, err
	}
	// todo: update the data and last_modified in cache
	// err = i.cacheRepository.Update(ctx, req.Content, info.LastModified)
	// if err != nil {
	// todo: error handle of updating data(by cache)
	// 	return nil, err
	// }
	res = domain.UpdateInfoBulletinResponse{
		LastModified: info.LastModified,
	}
	return res, nil
}
func (i *InfoUsecase) Delete(ctx context.Context, req *domain.DeleteInfoBulletinRequest) (domain.DeleteInfoBulletinResponse, error) {
	var res domain.DeleteInfoBulletinResponse
	info, err := i.dbRepository.Delete(ctx, req)
	if err != nil {
		// todo: make error handle of "delete the data"(by RDBMS)
		return res, err
	}
	// todo: update the data and last_modified in redis
	res = domain.DeleteInfoBulletinResponse{
		LastModified: info.LastModified,
	}
	return res, nil
}
