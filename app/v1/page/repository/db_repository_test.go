package repository

import (
	"context"
	"teacher-site/config"
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
			desc:   "get 999 page is too big",
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
