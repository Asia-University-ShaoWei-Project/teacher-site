package database

import (
	"teacher-site/config"

	"github.com/go-redis/redis"
)

func NewRedis(conf *config.Redis) *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.Database,
	})
	return redis
}
