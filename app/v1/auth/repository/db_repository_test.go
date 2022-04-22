package repository

import (
	"context"
	"teacher-site/config"
	"teacher-site/mock"
	"teacher-site/pkg/database"
	"teacher-site/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	ctx  = context.Background()
	conf = config.New()
	db   = database.NewDB("../../../../pkg/database", conf.DB)
	repo = NewDbRepository(db, conf.DB)
)
var (
	err error
)

func TestGetAccountByUserId(t *testing.T) {
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
			desc:     "normal",
			id:       mock.UserID,
			password: mock.UserPassword,
			result:   nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err = repo.GetAccountByUserId(ctx, tC.id)
			assert.Equal(t, tC.result, err)
		})
	}
}
func TestUpdateTokenByUserId(t *testing.T) {
	token, _ := util.GenerateJwt(conf.Jwt, mock.UserID)
	testCases := []struct {
		desc   string
		id     string
		result error
	}{
		{
			desc:   "invalid account of id",
			id:     mock.Unknown,
			result: gorm.ErrRecordNotFound,
		},
		{
			desc:   "normal",
			id:     mock.UserID,
			result: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err = repo.UpdateTokenByUserId(ctx, tC.id, token)
			assert.Equal(t, tC.result, err)
		})
	}
}

func TestDeleteToken(t *testing.T) {
	// token, _ := util.GenerateJwt(conf.Jwt, mock.UserID)
	testCases := []struct {
		desc   string
		id     string
		result error
	}{
		{
			desc:   "invalid account of id",
			id:     mock.Unknown,
			result: gorm.ErrRecordNotFound,
		},
		{
			desc:   "normal",
			id:     mock.UserID,
			result: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err = repo.DeleteToken(ctx, tC.id)
			assert.Equal(t, tC.result, err)
		})
	}
}
