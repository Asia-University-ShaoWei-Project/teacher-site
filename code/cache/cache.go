package cache

// https://redis.uptrace.dev/guide/server.html
import (
	"context"
	"errors"
	"fmt"
	"teacher-site/model"

	"github.com/go-redis/redis"
)

var (
	errMaximumRetry = errors.New("reached maximum number of retries")
)

type Cache struct {
	db   *redis.Client
	conf *model.CacheConfig
}
type Cacheer interface {
	GetInit(ctx context.Context, domain string) (string, error)
	GetCourseContent(ctx context.Context, domain string, courseID uint) (string, error)
	GetCourseLastUpdated(ctx context.Context, domain string, courseID uint) (string, error)

	SetInit(ctx context.Context, domain string, value string) error
	SetCourseContent(ctx context.Context, domain string, courseID uint, value *model.Courses) error
	SetCourseLastUpdated(ctx context.Context, domain string, courseID uint, updatedTime int64) error
	SetTokenWithDomain(ctx context.Context, domain, token string) error
}

func NewCache(config *model.CacheConfig) Cacheer {
	addr := ":6379"
	password := ""
	database := 0
	redis := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       database,
	})
	return &Cache{
		db:   redis,
		conf: config,
	}
}

// keys
const (
	kInit         = "init_data:%s"
	kCourse       = "course:%s:%d"
	kTeacherToken = "teacher_token"
)

// hash fields
const (
	fCourseContent     = `content`
	fCourseLastUpdated = `last_updated`
)

func (c *Cache) GetInit(ctx context.Context, domain string) (string, error) {
	k := bindKey(kInit, domain)
	return c.db.HGet(k, domain).Result()
}
func (c *Cache) SetInit(ctx context.Context, domain string, value string) error {
	k := bindKey(kInit, domain)
	return c.txHashSet(ctx, k, domain, value)
}

func (c *Cache) GetCourseContent(ctx context.Context, domain string, courseID uint) (string, error) {
	k := bindKey(kCourse, domain, courseID)

	return c.db.HGet(k, fCourseContent).Result()
}
func (c *Cache) SetCourseContent(ctx context.Context, domain string, courseID uint, value *model.Courses) error {
	k := bindKey(kCourse, domain, courseID)

	return c.txHashSet(ctx, k, fCourseContent, value)
}

func (c *Cache) GetCourseLastUpdated(ctx context.Context, domain string, courseID uint) (string, error) {
	k := bindKey(kCourse, domain, courseID)

	return c.db.HGet(k, fCourseLastUpdated).Result()
}
func (c *Cache) SetCourseLastUpdated(ctx context.Context, domain string, courseID uint, updatedTime int64) error {
	k := bindKey(kCourse, domain, courseID)

	return c.txHashSet(ctx, k, fCourseLastUpdated, updatedTime)
}

func (c *Cache) SetTokenWithDomain(ctx context.Context, domain, token string) error {
	return c.txHashSet(ctx, kTeacherToken, token, domain)
}

func (c *Cache) txHashSet(ctx context.Context, key, field string, value interface{}) error {
	for i := 0; i < c.conf.MaxReTry; i++ {
		err := c.db.Watch(func(tx *redis.Tx) error {
			tx.HSet(key, field, value)
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
	return errMaximumRetry
}
func bindKey(format string, value ...interface{}) string {
	return fmt.Sprintf(format, value...)
}
