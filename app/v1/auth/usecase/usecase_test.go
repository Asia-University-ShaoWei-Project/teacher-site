package usecase

import (
	"context"
	"teacher-site/domain"
	"teacher-site/mock"
	"teacher-site/mock/auth/repository"
	"teacher-site/pkg/message"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	staticPath = "../../../../static"
)

var (
	ctx       = context.Background()
	dbRepo    = repository.NewDbRepository()
	cacheRepo = repository.NewCacheRepository()
	conf      = mock.Conf
	usecase   = NewUsecase(dbRepo, cacheRepo, conf, mock.Log)
)
var (
	err error
)

// todo
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

func TestRegister(t *testing.T) {
	conf.Server.Path.StaticRelative = staticPath
	req := domain.RegisterRequest{
		UserPassword: mock.WordStr,
	}

	testCases := []struct {
		desc          string
		userId        string
		teacherDomain string
		result        error
	}{
		{
			desc:          "user id is existed",
			userId:        mock.UserId,
			teacherDomain: mock.Unknown,
			result:        message.ErrExistUserId,
		},
		{
			desc:          "teacher domain id is existed",
			userId:        mock.Unknown,
			teacherDomain: mock.TeacherDomain,
			result:        message.ErrExistTeacherDomain,
		},
		{
			desc:          "normal",
			userId:        mock.Unknown,
			teacherDomain: mock.Unknown,
			result:        nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req.UserId = tC.userId
			req.Domain = tC.teacherDomain
			err = usecase.Register(ctx, &req)
			assert.Equal(t, tC.result, err)
		})
	}
}
