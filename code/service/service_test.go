package service

import (
	"context"
	"teacher-site/cache"
	"teacher-site/database"
	"teacher-site/logsrv"
	"teacher-site/mock"
	"teacher-site/model"
)

var (
	ctx       = context.Background()
	cacheConf = model.NewMockCacheConfig()
	c         = cache.NewCache(cacheConf)
	logger    = logsrv.NewLogrus(ctx)
	db        = database.NewSqlite("../database", logger)
	conf      = model.NewMockServiceConfig()
	srv       = Service{
		domain: mock.Domain,
		db:     db,
		cache:  c,
		log:    logger,
		conf:   conf,
	}
)
