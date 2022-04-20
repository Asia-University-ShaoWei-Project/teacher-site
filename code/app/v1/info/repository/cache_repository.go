package repository

import (
	"context"
	"fmt"
	"teacher-site/config"
	"teacher-site/domain"
	"teacher-site/pkg/database"

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

func NewCacheRepository(db *redis.Client, conf *config.Redis) domain.InfoCacheRepository {
	return &cacheRepository{
		db:   db,
		conf: conf,
	}
}
func (c *cacheRepository) Get(ctx context.Context, req *domain.GetInfoBulletinRequest) (string, error) {
	k := fmt.Sprintf(keyInfo, req.TeacherDomain)
	return c.db.HGet(k, fContent).Result()
}

func (c *cacheRepository) GetLastModified(ctx context.Context, teacherDomain string) (string, error) {
	k := fmt.Sprintf(keyInfo, teacherDomain)
	return c.db.HGet(k, fLastModified).Result()
}

func (c *cacheRepository) Update(ctx context.Context, req *domain.UpdateInfoBulletinRequest) error {
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
func (c *cacheRepository) UpdateInfoLastModified(ctx context.Context, req *domain.UpdateInfoBulletinRequest, lastModified string) error {
	key := fmt.Sprintf(keyInfo, req.TeacherDomain)
	// todo: how to append the value into the hash
	for i := 0; i < c.conf.MaxReTry; i++ {
		err := c.db.Watch(func(tx *redis.Tx) error {
			tx.HSet(key, fLastModified, lastModified)
			return nil
		}, key)
		if err == nil {
			return nil
		}
		if err == redis.TxFailedErr {
			continue
		}
		return err
	}
	return database.ErrMaximumRetry
}
