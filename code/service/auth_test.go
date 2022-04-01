package service

import (
	"teacher-site/message"
	"teacher-site/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// func (srv *Service) Login(ctx context.Context, auth *model.Auths) error {

	tC := []struct {
		desc   string
		auth   *model.BindAuth
		result error
	}{
		{
			desc: "Real user",
			auth: &model.BindAuth{
				UserID:       "teacher_id",
				UserPassword: "password",
			},
			result: nil,
		},
		{
			desc: "Empty fields",
			auth: &model.BindAuth{
				UserID:       "",
				UserPassword: "password",
			},
			result: message.ErrDataEmpty,
		},
		{
			desc: "Not found the account",
			auth: &model.BindAuth{
				UserID:       "unknown_id",
				UserPassword: "password",
			},
			result: message.ErrQueryNotFound,
		},
	}
	for _, v := range tC {
		t.Run(v.desc, func(t *testing.T) {
			_, err := srv.Login(ctx, v.auth, cfg)
			assert.Equal(t, v.result, err)
		})
	}
}
