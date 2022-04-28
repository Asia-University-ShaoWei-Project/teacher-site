package usecase

import (
	"context"
	"encoding/json"
	"teacher-site/config"
	"teacher-site/domain"
	"teacher-site/pkg/message"

	log "github.com/sirupsen/logrus"
)

type Usecase struct {
	DbRepository    domain.InfoDbRepository
	CacheRepository domain.InfoCacheRepository
	conf            *config.Config
	log             *log.Logger
}

func NewUsecase(dbRepo domain.InfoDbRepository, cacheRepo domain.InfoCacheRepository, conf *config.Config, logger *log.Logger) domain.InfoUsecase {
	return &Usecase{
		DbRepository:    dbRepo,
		CacheRepository: cacheRepo,
		conf:            conf,
		log:             logger,
	}
}

func (i *Usecase) Create(ctx context.Context, req *domain.CreateInfoBulletinRequest) (domain.CreateInfoBulletinResponse, error) {
	var res domain.CreateInfoBulletinResponse
	bulletin, err := i.DbRepository.Create(ctx, req)
	if err != nil {
		return res, err
	}
	lastModified, err := i.DbRepository.GetLastModified(ctx, req.InfoId)
	if err != nil {
		// todo: get last modified error
		i.log.Error(err)
	}
	// todo: update the data in redis
	res = domain.CreateInfoBulletinResponse{
		Id:           bulletin.AutoModel.Id,
		Date:         bulletin.AutoModel.CreatedAT.Format(domain.BulletinDateFormat),
		LastModified: lastModified,
	}
	return res, nil
}

func (i *Usecase) Get(ctx context.Context, req *domain.GetInfoBulletinRequest) (domain.GetInfoBulletinResponse, error) {
	var (
		res      domain.GetInfoBulletinResponse
		bulletin []domain.InfoBulletinResponse
	)
	info, err := i.DbRepository.GetByTeacherDomain(ctx, req.TeacherDomain)
	// lastModified, err := i.DbRepository.GetLastModified(ctx, req.TeacherDomain)
	if err != nil {
		return res, err
	}
	// Unnecessary to get new data if request last modified value is equal the last modified of repository value
	if req.LastModified == info.LastModified {
		return res, message.ErrUnnecessaryUpdate
	}
	res.SetLastModified(info.LastModified)
	res.SetId(info.AutoModel.Id)
	// Get by cache
	data, err := i.CacheRepository.Get(ctx, req)
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
	bulletins, err := i.DbRepository.GetBulletinsByInfoId(ctx, info.AutoModel.Id)
	if err != nil {
		// todo: make error handle of "get data"(by RDBMS)
		i.log.Error(err)
		return res, err
	}
	res.SetBulletins(bulletins)
	// todo: update the data and last_modified in cache
	// if err := i.CacheRepository.Update(ctx, &resInfo); err != nil {
	// todo: error handle of updating data(by cache)
	// i.log.Error(err)
	// }
	return res, nil
}
func (i *Usecase) Update(ctx context.Context, req *domain.UpdateInfoBulletinRequest) (domain.UpdateInfoBulletinResponse, error) {
	var res domain.UpdateInfoBulletinResponse
	if err := i.DbRepository.CheckByDomainAndId(ctx, req.TeacherDomain, req.InfoId); err != nil {
		i.log.Error(err)
		return res, err
	}
	info, err := i.DbRepository.Update(ctx, req)
	if err != nil {
		// todo: make error handle of "update the data"(by RDBMS)
		i.log.Error(err)
		return res, err
	}
	// todo: update the data and last_modified in cache
	// err = i.CacheRepository.Update(ctx, req.Content, info.LastModified)
	// if err != nil {
	// todo: error handle of updating data(by cache)
	// 	return nil, err
	// }
	res = domain.UpdateInfoBulletinResponse{
		LastModified: info.LastModified,
	}
	return res, nil
}
func (i *Usecase) Delete(ctx context.Context, req *domain.DeleteInfoBulletinRequest) (domain.DeleteInfoBulletinResponse, error) {
	var res domain.DeleteInfoBulletinResponse
	info, err := i.DbRepository.Delete(ctx, req)
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
