package cache

import (
	"fmt"
	"teacher-site/model"

	"github.com/go-redis/redis"
)

type Cache struct {
	db *redis.Client
}
type Cacheer interface {
	GetInit(domain string) (string, error)
	GetCourseWithContent(domain string, courseID uint) (string, error)
	GetCourseLastUpdated(domain string, courseID uint) (string, error)
	SetInit(domain string, value string) error
	SetToken(domain, token string) error
	SetCourse(domain string, courseID uint, value *model.Courses) error
	SetCourseLastUpdated(domain string, courseID uint, updatedTime int64) error
}

func NewCache() Cacheer {
	addr := "localhost:6379"
	password := ""
	database := 0
	redis := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       database,
	})
	return &Cache{
		db: redis,
	}
}

const (
	initKey              = "init_data"
	tokenKey             = "token"
	courseKey            = "course"
	courseLastUpdatedKey = "course:last_updated"
)

func (c *Cache) GetInit(domain string) (string, error) {
	return c.db.HGet(initKey, domain).Result()
}
func (c *Cache) SetInit(domain string, value string) error {
	// todo: expire?
	return c.db.HSet(initKey, domain, value).Err()
}
func (c *Cache) GetCourseWithContent(domain string, courseID uint) (string, error) {
	field := fmt.Sprintf("%s:%d", domain, courseID)
	return c.db.HGet(courseKey, field).Result()
}
func (c *Cache) SetCourse(domain string, courseID uint, value *model.Courses) error {
	field := fmt.Sprintf("%s:%d", domain, courseID)
	return c.db.HSet(initKey, field, value).Err()
}
func (c *Cache) GetCourseLastUpdated(domain string, courseID uint) (string, error) {
	field := fmt.Sprintf("%s:%d", domain, courseID)
	return c.db.HGet(courseLastUpdatedKey, field).Result()
}
func (c *Cache) SetCourseLastUpdated(domain string, courseID uint, updatedTime int64) error {
	field := fmt.Sprintf("%s:%d", domain, courseID)
	return c.db.HSet(courseLastUpdatedKey, field, updatedTime).Err()
}
func (c *Cache) SetToken(domain, token string) error {
	return c.db.HSetNX(tokenKey, token, domain).Err()
}
