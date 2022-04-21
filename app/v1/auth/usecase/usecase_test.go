package usecase

import (
	"context"
	"teacher-site/domain"
	"teacher-site/mock"
	"teacher-site/mock/auth/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	ctx       = context.Background()
	dbRepo    = repository.NewDbRepository()
	cacheRepo = repository.NewCacheRepository()
	usecase   = NewAuthUsecase(dbRepo, cacheRepo, mock.Conf, mock.Log)
)
var (
	err error
)

// todo: test the login
func TestLogin(t *testing.T) {
	var req domain.LoginRequest

	testCases := []struct {
		desc     string
		id       string
		password string
		result   error
	}{
		{
			desc:     "invalid account of id",
			id:       mock.Unknown,
			password: mock.UserPassword,
			result:   gorm.ErrRecordNotFound,
		},
		{
			desc:     "invalid account of password",
			id:       mock.UserID,
			password: mock.Unknown,
			result:   gorm.ErrRecordNotFound,
		},
		{
			desc:     "normal",
			id:       mock.UserID,
			password: mock.UserPassword,
			result:   nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req = domain.LoginRequest{
				UserID:       tC.id,
				UserPassword: tC.password,
			}
			_, err = usecase.Login(ctx, &req)
			assert.Equal(t, tC.result, err)
		})
	}
}

// todo: test invalid password
