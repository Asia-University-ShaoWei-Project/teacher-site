package repository

import (
	"context"
	"fmt"
	"teacher-site/config"
	"teacher-site/domain"
	"teacher-site/pkg/message"

	"github.com/go-redis/redis"
)

const (
	courseLastModifiedKey = `lastModified:course:%d`
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

func (c *CacheRepository) GetLastModifiedByCourseId(ctx context.Context, courseId uint) (string, error) {
	key := fmt.Sprintf(courseLastModifiedKey, courseId)
	data, err := c.db.Get(key).Result()
	return data, err
}
func (c *CacheRepository) UpdateLastModifiedByCourseId(ctx context.Context, courseId uint, lastModified string) error {
	key := fmt.Sprintf(courseLastModifiedKey, courseId)
	for i := 0; i < c.conf.MaxReTry; i++ {
		err := c.db.Watch(func(tx *redis.Tx) error {
			tx.Set(key, lastModified, c.conf.LastModifiedExpire)
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
	return message.ErrMaximumRetry
}
