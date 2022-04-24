package usecase

import (
	"context"
	"teacher-site/domain"
	"teacher-site/mock"
	"teacher-site/mock/page/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	ctx       = context.Background()
	dbRepo    = repository.NewDbRepository()
	cacheRepo = repository.NewCacheRepository()
	usecase   = NewUsecase(dbRepo, cacheRepo, mock.Conf, mock.Log)
)
var err error

// todo TestTeacherList
func TestHome(t *testing.T) {
	var req domain.HomeRequest

	testCases := []struct {
		desc          string
		teacherDomain string
		result        error
	}{
		{
			desc:          "unknown teacher domain",
			teacherDomain: mock.Unknown,
			result:        gorm.ErrRecordNotFound,
		},
		{
			desc:          "normal",
			teacherDomain: mock.TeacherDomain,
			result:        nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req = domain.HomeRequest{Domain: tC.teacherDomain}
			_, err = usecase.Home(ctx, &req)
			assert.Equal(t, tC.result, err)
		})
	}
}
