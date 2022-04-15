package service

import (
	"fmt"
	"teacher-site/mock"
	"teacher-site/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

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
func TestNewJwtToken(t *testing.T) {
	tC := struct {
		userID string
		result error
	}{
		userID: mock.UserID,
		result: nil,
	}
	bindAuth := &model.BindAuth{
		UserID: tC.userID,
	}
	token, err := srv.NewJwtToken(ctx, bindAuth)
	assert.Equal(t, tC.result, err)
	fmt.Println(token)

}
func TestUpdateJwtToken(t *testing.T) {
	tC := struct {
		userID string
		result error
	}{
		userID: mock.UserID,
		result: nil,
	}

	bindAuth := &model.BindAuth{
		UserID: tC.userID,
	}
	token, err := srv.NewJwtToken(ctx, bindAuth)
	assert.NotNil(t, err, err)
	err = srv.UpdateJwtToken(ctx, token, bindAuth)
	// err = srv.
	assert.Equal(t, tC.result, err)
}

// done 04/13
func TestRegister(t *testing.T) {
	mockBasicData := &model.BindRegister{
		UserPassword: mock.UserPassword,
		NameZH:       mock.UserName,
	}

	tC := []struct {
		desc   string
		userID string
		domain string
		result error
	}{
		{
			desc:   "User is existed",
			userID: mock.UserID,
			domain: mock.Unknown,
			result: errIsExisted,
		},
		{
			desc:   "Domain is existed",
			userID: mock.Unknown,
			domain: mock.Domain,
			result: errIsExisted,
		},
		{
			desc:   "Success register",
			userID: mock.Unknown,
			domain: mock.Unknown,
			result: nil,
		},
	}
	for _, v := range tC {
		t.Run(v.desc, func(t *testing.T) {
			mockBasicData.UserID = v.userID
			mockBasicData.Domain = v.domain
			err := srv.Register(ctx, mockBasicData)
			assert.Equal(t, v.result, err, err)
		})
	}
}
