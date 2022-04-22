package usecase

import (
	"context"
	"teacher-site/config"
	"teacher-site/domain"
	"teacher-site/pkg/util"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	dbRepository    domain.AuthDbRepository
	cacheRepository domain.AuthCacheRepository
	conf            *config.Config
	log             *logrus.Logger
}

func NewUsecase(dbRepo domain.AuthDbRepository, cacheRepo domain.AuthCacheRepository, conf *config.Config, logger *logrus.Logger) domain.AuthUsecase {
	return &AuthUsecase{
		dbRepository:    dbRepo,
		cacheRepository: cacheRepo,
		conf:            conf,
		log:             logger,
	}
}
func (auth *AuthUsecase) Login(ctx context.Context, req *domain.LoginRequest) (string, error) {
	var token string

	// todo: check logged in
	account, err := auth.dbRepository.GetAccountByUserId(ctx, req.UserID)
	if err != nil {
		auth.log.Error(err)
		return token, err
	}

	// compare password
	saltPassword := []byte(req.UserPassword + account.Salt)
	if err = bcrypt.CompareHashAndPassword([]byte(account.UserPassword), saltPassword); err != nil {
		auth.log.Error(err)
		return token, err
	}

	// generate new token for the header of client(authorization)
	token, err = util.GenerateJwt(auth.conf.Jwt, account.UserID)
	if err != nil {
		// todo: try again
		auth.log.Error(err)
	}
	if err = auth.dbRepository.UpdateTokenByUserId(ctx, account.UserID, token); err != nil {
		// todo: error handle of update token
		auth.log.Error(err)
	}
	return token, nil
}

func (auth *AuthUsecase) Logout(ctx context.Context, id string) error {
	return auth.dbRepository.DeleteToken(ctx, id)
}

// func (auth *AuthUsecase) Create(ctx context.Context, req *domain.ReqCreateInfo) (domain.InfoBulletinBoards, error) {
// 	// todo: use TX to read and write the last_updated
// 	return auth.dbRepository.Create(ctx, req)
// }
// func (auth *AuthUsecase) Get(ctx context.Context, req *domain.ReqGetInfo) (domain.ResGetInfo, error) {
// 	var (
// 		resInfo  domain.ResGetInfo
// 		bulletin []domain.GetInfoBulletin
// 	)

// 	lastModified, err := auth.cacheRepository.GetLastModified(ctx, req)
// 	// Not need to get new data if request last modified value is equal the last modified of repository value
// 	if err == nil && (req.LastModified == lastModified) {
// 		auth.log.Info("Not need to update")
// 		return resInfo, err
// 	}
// 	resInfo.SetLastModified(lastModified)
// 	// Get by cache
// 	data, err := auth.cacheRepository.Get(ctx, req)
// 	if err == nil {
// 		// return resInfo, err
// 		if err = json.Unmarshal([]byte(data), &bulletin); err == nil {
// 			resInfo.SetBulletins(bulletin)
// 			return resInfo, nil
// 		}
// 		// todo: parse json error handle
// 		auth.log.Error(err)
// 	}
// 	// todo: cache error handle
// 	auth.log.Error(err)
// 	// Get from database
// 	bulletins, err := auth.dbRepository.Get(ctx, req)
// 	if err != nil {
// 		auth.log.Error(err)
// 		return resInfo, err
// 	}
// 	resInfo.SetBulletins(bulletins)
// 	// todo: Update the info in cache & cache update error handle
// 	// if err := auth.cacheRepository.Update(ctx, &resInfo); err != nil {
// 	// auth.log.Error(err)
// 	// }
// 	return resInfo, nil
// }
// func (auth *AuthUsecase) Update(ctx context.Context, req *domain.ReqUpdateInfoBulletin) (domain.ResUpdateInfo, error) {
// 	var res domain.ResUpdateInfo
// 	bulletin, err := auth.dbRepository.Update(ctx, req)
// 	if err != nil {
// 		auth.log.Error(err)
// 		return res, err
// 	}
// 	// todo: update the data in redis
// 	// err = auth.cacheRepository.Update(ctx, req*domain.ResGetInfo)
// 	res = domain.ResUpdateInfo{
// 		ID:   bulletin.AutoModel.ID,
// 		Date: bulletin.AutoModel.CreatedAT.Format(domain.BulletinDateFormat),
// 	}
// 	return res, nil
// }
// func (auth *AuthUsecase) Delete(ctx context.Context, req *domain.ReqDeleteInfo) error {
// 	// todo: update the data in redis
// 	return auth.dbRepository.Delete(ctx, req)
// }
