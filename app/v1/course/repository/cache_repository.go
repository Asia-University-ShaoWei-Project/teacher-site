package repository

import (
	"context"
	"fmt"
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

type CacheRepository struct {
	db   *redis.Client
	conf *config.Redis
}

func NewCacheRepository(db *redis.Client, conf *config.Redis) domain.InfoCacheRepository {
	return &CacheRepository{
		db:   db,
		conf: conf,
	}
}

func (c *CacheRepository) Get(ctx context.Context, req *domain.ReqGetInfo) (string, error) {
	k := fmt.Sprintf(keyInfo, req.TeacherDomain)
	return c.db.HGet(k, fContent).Result()
}
func (c *CacheRepository) GetLastModified(ctx context.Context, req *domain.ReqGetInfo) (string, error) {
	k := fmt.Sprintf(keyInfo, req.TeacherDomain)
	return c.db.HGet(k, fLastModified).Result()
}

func (c *CacheRepository) Update(ctx context.Context, req *domain.ResGetInfo) error {
	// k := fmt.Sprintf(keyInfo, req.TeacherDomain)
	// todo: how to append the value into the hash
	// for i := 0; i < c.conf.MaxReTry; i++ {
	// 	err := c.db.Watch(func(tx *redis.Tx) error {
	// 		tx.HSet(key, field, value)
	// 		return nil
	// 	}, key)
	// 	if err == nil {
	// 		return nil
	// 	}
	// 	if err == redis.TxFailedErr {
	// 		continue
	// 	}
	// 	return err
	// }
	// return errMaximumRetry

	return nil
}
func (c *CacheRepository) UpdateInfoLastModified(ctx context.Context, req *domain.ResGetInfo) error {
	// k := fmt.Sprintf(keyInfo, req.TeacherDomain)
	// todo: how to append the value into the hash
	// for i := 0; i < c.conf.MaxReTry; i++ {
	// 	err := c.db.Watch(func(tx *redis.Tx) error {
	// 		tx.HSet(key, field, value)
	// 		return nil
	// 	}, key)
	// 	if err == nil {
	// 		return nil
	// 	}
	// 	if err == redis.TxFailedErr {
	// 		continue
	// 	}
	// 	return err
	// }
	// return errMaximumRetry

	return nil
}
