package repository

import (
	"teacher-site/config"
	"teacher-site/domain"

	"github.com/go-redis/redis"
)

type CacheRepository struct {
	db   *redis.Client
	conf *config.Redis
}

func NewCacheRepository(db *redis.Client, conf *config.Redis) domain.CourseCacheRepository {
	return &CacheRepository{
		db:   db,
		conf: conf,
	}
}
