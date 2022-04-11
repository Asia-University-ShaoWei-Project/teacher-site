package service

import (
	"context"
	"teacher-site/cache"
	"teacher-site/database"
	"teacher-site/logsrv"
	"teacher-site/model"
)

var (
	ctx       = context.Background()
	cacheConf = model.NewMockCacheConfig()
	c         = cache.NewCache(cacheConf)
	logger    = logsrv.NewLogrus(ctx)
	db        = database.NewSqlite("../database", logger)
	conf      = model.NewMockServiceConfig()
	srv       = NewService(db, c, logger, conf)
)
