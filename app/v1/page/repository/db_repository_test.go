package repository

import (
	"context"
	"teacher-site/config"
	"teacher-site/mock"
	"teacher-site/pkg/database"
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
var err error

func TestGetTeachers(t *testing.T) {
	var offset int
	limit := conf.Limit.TeacherListPageCount

	testCases := []struct {
		desc   string
		page   int
		result error
	}{
		{
			desc:   "Not found",
			page:   999,
			result: gorm.ErrRecordNotFound,
		},
		{
			desc:   "normal",
			page:   1,
			result: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			offset = calculateOffset(limit, tC.page)
			_, err = repo.GetTeachers(ctx, limit, offset)
			assert.Equal(t, tC.result, err)
		})
	}
}
func calculateOffset(limit, page int) int {
	return (limit * page) - limit
}
func TestCheckAuthByIdAndToken(t *testing.T) {
	realUserToken, err := mock.GetUserToken(db)
	if err != nil {
		t.Fatal(err)
	}
	testCases := []struct {
		desc   string
		userId string
		token  string
		result error
	}{
		{
			desc:   "unknown the userId",
			userId: mock.Unknown,
			token:  realUserToken,
			result: gorm.ErrRecordNotFound,
		},
		{
			desc:   "unknown the token",
			userId: mock.UserId,
			token:  mock.Unknown,
			result: gorm.ErrRecordNotFound,
		},
		{
			desc:   "normal",
			userId: mock.UserId,
			token:  realUserToken,
			result: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err = repo.CheckAuthByIdAndToken(ctx, tC.userId, tC.token)
			assert.Equal(t, tC.result, err)
		})
	}
}
