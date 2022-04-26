package repository

import (
	"teacher-site/mock"
	"teacher-site/pkg/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	ctx  = mock.Ctx
	db   = database.NewDB("../../../../pkg/database", mock.Conf.DB)
	repo = NewDbRepository(db, mock.Conf.DB)
)

func TestGetByTeacherDomain(t *testing.T) {
	testCases := []struct {
		desc          string
		teacherDomain string
		result        error
	}{
		{
			desc:          "The teacher domian is not existed",
			teacherDomain: mock.Unknown,
			result:        gorm.ErrRecordNotFound,
		},
		{
			desc:          "Normal",
			teacherDomain: mock.TeacherDomain,
			result:        nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := repo.GetByTeacherDomain(ctx, tC.teacherDomain)
			assert.Equal(t, tC.result, err)
		})
	}
}
