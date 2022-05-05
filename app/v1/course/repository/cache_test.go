package repository

import (
	"teacher-site/mock"
	"teacher-site/pkg/database"
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

const (
	lastModified = "0"
)

var (
	conf      = mock.Conf
	cache     = database.NewRedis(conf.Redis)
	cacheRepo = NewCacheRepository(cache, conf.Redis)
)

func TestGetLastModifiedByTeacherDomain(t *testing.T) {
	if err := cacheRepo.UpdateLastModifiedByCourseId(ctx, mock.NumPk, lastModified); err != nil {
		t.Fatal(err)
	}
	testCases := []struct {
		desc     string
		courseId uint
		result   error
	}{
		{
			desc:     "Fail to get last modified",
			courseId: mock.UnknownNumPK,
			result:   redis.Nil,
		},
		{
			desc:     "normal",
			courseId: mock.NumPk,
			result:   nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := cacheRepo.GetLastModifiedByCourseId(ctx, tC.courseId)
			assert.Equal(t, tC.result, err)
		})
	}
}
