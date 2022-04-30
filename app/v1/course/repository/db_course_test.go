package repository

import (
	"fmt"
	"teacher-site/domain"
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
var (
	err error
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

func TestGetContentByCourseId(t *testing.T) {
	testCases := []struct {
		desc     string
		courseId uint
		result   error
	}{
		{
			desc:     "The teacher domian is not existed",
			courseId: mock.UnknownNumPK,
			result:   gorm.ErrRecordNotFound,
		},
		{
			desc:     "Normal",
			courseId: mock.PkNum,
			result:   nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			a, err := repo.GetContentByCourseId(ctx, tC.courseId)
			fmt.Println(a)

			assert.Equal(t, tC.result, err)
		})
	}
}
func TestCheckByDomainAndCourseId(t *testing.T) {
	var course domain.Courses
	testCases := []struct {
		desc          string
		teacherDomain string
		courseId      uint
		result        error
	}{
		{
			desc:          "Not existed the course id of the teacher",
			teacherDomain: mock.TeacherDomain,
			courseId:      mock.UnknownNumPK,
			result:        gorm.ErrRecordNotFound,
		},
		{
			desc:          "Not have the course id",
			teacherDomain: mock.Unknown,
			courseId:      mock.PkNum,
			result:        gorm.ErrRecordNotFound,
		},
		{
			desc:          "Normal",
			teacherDomain: mock.TeacherDomain,
			courseId:      mock.PkNum,
			result:        nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			course = domain.Courses{
				TeacherId: tC.teacherDomain,
				AutoModel: domain.AutoModel{Id: tC.courseId},
			}
			err := repo.CheckByDomainAndCourseId(ctx, &course)
			assert.Equal(t, tC.result, err)
		})
	}
}
func TestGetLastModifiedByCourseId(t *testing.T) {
	testCases := []struct {
		desc     string
		courseId uint
		result   error
	}{
		{
			desc:     "Not found the course id",
			courseId: mock.UnknownNumPK,
			result:   gorm.ErrRecordNotFound,
		},
		{
			desc:     "Normal",
			courseId: mock.PkNum,
			result:   nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := repo.GetLastModifiedByCourseId(ctx, tC.courseId)
			assert.Equal(t, tC.result, err)
		})
	}
}
