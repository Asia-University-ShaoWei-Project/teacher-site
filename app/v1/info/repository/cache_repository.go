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
	infoLastModifiedKey = `lastModified:info:%s`
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

// func (c *CacheRepository) Get(ctx context.Context, req *domain.GetInfoBulletinRequest) (string, error) {
// 	k := fmt.Sprintf(keyInfo, req.TeacherDomain)
// 	return c.db.HGet(k, fContent).Result()
// }

func (c *CacheRepository) GetLastModifiedByTeacherDomain(ctx context.Context, teacherDomain string) (string, error) {
	key := fmt.Sprintf(infoLastModifiedKey, teacherDomain)
	data, err := c.db.Get(key).Result()
	return data, err
}

// func (c *CacheRepository) Update(ctx context.Context, req *domain.UpdateInfoBulletinRequest) error {
// 	k := fmt.Sprintf(keyInfo, req.TeacherDomain)
// 	// todo: how to append the value into the hash
// 	for i := 0; i < c.conf.MaxReTry; i++ {
// 		err := c.db.Watch(func(tx *redis.Tx) error {
// 			tx.HSet(key, field, value)
// 			return nil
// 		}, key)
// 		if err == nil {
// 			return nil
// 		}
// 		if err == redis.TxFailedErr {
// 			continue
// 		}
// 		return err
// 	}
// 	return errMaximumRetry

// 	return nil
// }
func (c *CacheRepository) UpdateLastModifiedByTeacherDomain(ctx context.Context, teacherDomain, lastModified string) error {
	// key := fmt.Sprintf(keyInfo, req.TeacherDomain)
	key := fmt.Sprintf(infoLastModifiedKey, teacherDomain)
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
