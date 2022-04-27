package usecase

import (
	"context"
	"teacher-site/domain"
	"teacher-site/mock"
	"teacher-site/mock/auth/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ctx       = context.Background()
	dbRepo    = repository.NewDbRepository()
	cacheRepo = repository.NewCacheRepository()
	usecase   = NewUsecase(dbRepo, cacheRepo, mock.Conf, mock.Log)
)
var (
	err error
)

func TestLogin(t *testing.T) {
	var req domain.LoginRequest

	testCases := []struct {
		desc     string
		userId   string
		password string
		result   error
	}{
		{
			desc:     "invalid user id(Not found the user)",
			userId:   mock.Unknown,
			password: mock.UserPassword,
			result:   gorm.ErrRecordNotFound,
		},
		{
			desc:     "invalid password",
			userId:   mock.UserId,
			password: mock.Unknown,
			result:   bcrypt.ErrMismatchedHashAndPassword,
		},
		{
			desc:     "normal",
			userId:   mock.UserId,
			password: mock.UserPassword,
			result:   nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req = domain.LoginRequest{
				UserId:       tC.userId,
				UserPassword: tC.password,
			}
			_, err = usecase.Login(ctx, &req)
			assert.Equal(t, tC.result, err)
		})
	}
}
