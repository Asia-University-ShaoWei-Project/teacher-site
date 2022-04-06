package service

import (
	"teacher-site/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const password = "password"

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
			userID:       "teacher_id",
			userPassword: password,
			result:       nil,
		},
		{
			desc:         "Not found the account",
			userID:       "unknown_id",
			userPassword: password,
			result:       gorm.ErrRecordNotFound,
		},
		{
			desc:         "Fail password",
			userID:       "teacher_id",
			userPassword: password + "additional",
			result:       errAuthNotMatch,
		},
	}
	for _, v := range tC {
		t.Run(v.desc, func(t *testing.T) {
			auth = &model.BindAuth{
				UserID:       v.userID,
				UserPassword: v.userPassword,
			}
			_, err := srv.LoginAndGetNewToken(ctx, auth)
			assert.Equal(t, v.result, err)
		})
	}
}

func TestRegister(t *testing.T) {
	tC := struct {
		desc   string
		data   *model.BindRegister
		result bool
	}{
		desc: "e",
		data: &model.BindRegister{
			UserID:       "user-4",
			UserPassword: "password",
			Domain:       "teacher-domain",
			Email:        "xx@gmail.com",
			NameZH:       "My-name",
		},
		result: true,
	}
	t.Run(tC.desc, func(t *testing.T) {
		ok := srv.Register(ctx, tC.data)
		if ok != tC.result {
			t.Fatal("fail to register")
		}
	})
}
