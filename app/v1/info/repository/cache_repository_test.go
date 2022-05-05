package repository

import (
	"teacher-site/mock"
	"teacher-site/pkg/database"
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var (
	cache     = database.NewRedis(conf.Redis)
	cacheRepo = NewCacheRepository(cache, conf.Redis)
)

func TestGetLastModifiedByTeacherDomain(t *testing.T) {
	testCases := []struct {
		desc          string
		teacherDomain string
		result        error
	}{
		{
			desc:          "Fail to get last modified",
			teacherDomain: mock.Unknown,
			result:        redis.Nil,
		},
		{
			desc:          "normal",
			teacherDomain: "rikki",
			// teacherDomain: mock.TeacherDomain,
			result: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := cacheRepo.GetLastModifiedByTeacherDomain(ctx, tC.teacherDomain)
			assert.Equal(t, tC.result, err)
		})
	}
}
