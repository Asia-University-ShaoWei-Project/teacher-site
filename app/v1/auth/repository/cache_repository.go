package repository

import (
	"teacher-site/config"
	"teacher-site/domain"

	"github.com/go-redis/redis"
)

// hash keys & fields
const (
	keyInfo       = `info:%s`
	fContent      = `content`
	fLastModified = `last_modified`
)

type cacheRepository struct {
	db   *redis.Client
	conf *config.Redis
}

func NewCacheRepository(db *redis.Client, conf *config.Redis) domain.AuthCacheRepository {
	return &cacheRepository{
		db:   db,
		conf: conf,
	}
}
