package middleware

import (
	"context"
	"teacher-site/cache"
	"teacher-site/database"
	"teacher-site/logsrv"
	"teacher-site/mock"
	"teacher-site/model"
	"teacher-site/service"
	"testing"
)

var (
	cacheConf = model.NewMockCacheConfig()
	_cache    = cache.NewCache(cacheConf)
	db        = database.NewSqlite("../database", logger)
	logger    = logsrv.NewLogrus(ctx)
	ctx       = context.Background()
	conf      = model.NewMockServiceConfig()
	srv       = service.NewService(db, _cache, logger, conf)
)

func TestVerifyJwtValid(t *testing.T) {
	var (
		err   error
		token string
	)
	bind := model.BindAuth{UserID: mock.UserID}
	token, _ = srv.NewJwtToken(ctx, &bind)
	tC := []struct {
		desc  string
		token string
	}{
		{
			desc:  "real token",
			token: token,
		},
		{
			desc:  "fail token",
			token: mock.Unknown,
		},
	}
	for _, v := range tC {
		t.Run(v.desc, func(t *testing.T) {
			bearerToken := "Bearer " + v.token
			err = verifyJwtValid(ctx, srv, bearerToken)
			// todo: error handle
			srv.Error(err)
		})
	}
}
