package database

import (
	"errors"
	"teacher-site/config"

	"github.com/go-redis/redis"
)

var (
	ErrMaximumRetry = errors.New("reached maximum number of retries")
)

func NewRedis(conf *config.Redis) *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.Database,
	})
	return redis
}
