package service

import (
	"teacher-site/mock"
	"teacher-site/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var mockUser = &model.BindRegister{
	UserID:       mock.UserID,
	UserPassword: mock.UserPassword,
	Domain:       mock.Domain,
	Email:        mock.Email,
	NameZH:       mock.UserName,
}

func TestLogin(t *testing.T) {
	var auth *model.BindAuth

	tC := []struct {
		desc         string
		userID       string
		userPassword string
		result       error
	}{
		{
			desc:         "Real user",
			userID:       mock.UserID,
			userPassword: mock.UserPassword,
			result:       nil,
		},
		{
			desc:         "Not found the account",
			userID:       mock.Unknown,
			userPassword: mock.UserPassword,
			result:       gorm.ErrRecordNotFound,
		},
		{
			desc:         "Fail password",
			userID:       mock.UserID,
			userPassword: mock.Unknown,
			result:       nil,
		},
	}
	for _, v := range tC {
		t.Run(v.desc, func(t *testing.T) {
			auth = &model.BindAuth{
				UserID:       v.userID,
				UserPassword: v.userPassword,
			}
			err := srv.Login(ctx, auth)
			assert.Equal(t, v.result, err)

		})
	}
}

func TestUpdateJwtToken(t *testing.T) {
	tC := struct {
		userID string
		result error
	}{
		userID: mockUser.UserID,
		result: nil,
	}

	bindAuth := &model.BindAuth{
		UserID: tC.userID,
	}
	err := srv.UpdateJwtToken(ctx, bindAuth)
	assert.Equal(t, tC.result, err)
}
func TestRegister(t *testing.T) {
	tC := []struct {
		desc   string
		data   *model.BindRegister
		result bool
	}{
		{
			desc:   "First user",
			data:   mockUser,
			result: true,
		},
		{
			desc:   "Repeat id or email",
			data:   mockUser,
			result: false,
		},
	}
	for _, v := range tC {
		t.Run(v.desc, func(t *testing.T) {
			ok := srv.Register(ctx, v.data)
			assert.Equal(t, v.result, ok, "fail to register")
		})
	}

}
