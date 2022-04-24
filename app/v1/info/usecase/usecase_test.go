package usecase

import (
	"context"
	"teacher-site/domain"
	"teacher-site/mock"
	mockRepo "teacher-site/mock/info/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ctx       = context.Background()
	dbRepo    = mockRepo.NewInfoDbRepository()
	cacheRepo = mockRepo.NewInfoCacheRepository()
	usecase   = NewUsecase(dbRepo, cacheRepo, mock.Conf, mock.Log)
)

// func TestCreateInfo(t *testing.T) {
// 	var req domain.CreateInfoBulletinRequest
// 	_, err := usecase.Create(ctx, &req)
// 	assert.Nil(t, err, err)
// }

type getInfoResult struct {
	err error
}

func TestGetInfo(t *testing.T) {
	var req domain.GetInfoBulletinRequest

	testCases := []struct {
		desc         string
		lastModified string
		result       getInfoResult
	}{
		{
			desc:         "data is up to date",
			lastModified: mockRepo.CurrLastModidied,
			result:       getInfoResult{errUnnecessaryUpdate},
		},
		{
			desc:         "the data is late than current date",
			lastModified: mockRepo.LateLastModified,
			result:       getInfoResult{nil},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req = domain.GetInfoBulletinRequest{
				TeacherDomainRequest: domain.TeacherDomainRequest{TeacherDomain: mock.TeacherDomain},
				LastModified:         tC.lastModified,
			}
			_, err := usecase.Get(ctx, &req)
			assert.Equal(t, tC.result.err, err)
		})
	}
}
