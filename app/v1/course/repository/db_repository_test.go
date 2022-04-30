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

func TestCreateBulletin(t *testing.T) {
	var bulletin domain.BulletinBoards
	testCases := []struct {
		desc     string
		courseId uint
		result   error
	}{
		{
			desc:     "Unknown the PK",
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
			bulletin = domain.BulletinBoards{CourseId: tC.courseId, Content: mock.NewMsg()}
			err = repo.CreateBulletin(ctx, &bulletin)
			assert.Equal(t, tC.result, err)
		})
	}
}
func TestCreateSlide(t *testing.T) {
	// var slide domain.Slides
	slide := domain.Slides{Chapter: mock.NewMsg(), File: domain.File{Title: mock.NewMsg()}}
	testCases := []struct {
		desc     string
		courseId uint
		result   error
	}{
		{
			desc:     "Unknown the PK",
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
			slide.CourseId = tC.courseId
			err = repo.CreateSlide(ctx, &slide)
			fmt.Println("slide id:", slide.AutoModel.Id)

			assert.Equal(t, tC.result, err)
		})
	}
}
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

func TestUpdateBulletinById(t *testing.T) {
	var bulletin domain.BulletinBoards
	testCases := []struct {
		desc       string
		courseId   uint
		bulletinId uint
		result     error
	}{
		{
			desc:       "Both the id is not found",
			courseId:   mock.UnknownNumPK,
			bulletinId: mock.UnknownNumPK,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "Not found, unknown course id",
			courseId:   mock.UnknownNumPK,
			bulletinId: mock.PkNum,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "Not found, unknown bulletin id",
			courseId:   mock.PkNum,
			bulletinId: mock.UnknownNumPK,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "Normal",
			courseId:   mock.PkNum,
			bulletinId: mock.PkNum,
			result:     nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			bulletin = domain.BulletinBoards{
				AutoModel: domain.AutoModel{Id: tC.bulletinId},
				CourseId:  tC.courseId,
				Content:   mock.NewMsg(),
			}
			err := repo.UpdateBulletinById(ctx, &bulletin)
			assert.Equal(t, tC.result, err)
		})
	}
}
func TestUpdateSlideById(t *testing.T) {
	var slide domain.Slides
	testCases := []struct {
		desc     string
		courseId uint
		slideId  uint
		result   error
	}{
		{
			desc:     "Both the id is not found",
			courseId: mock.UnknownNumPK,
			slideId:  mock.UnknownNumPK,
			result:   gorm.ErrRecordNotFound,
		},
		{
			desc:     "Not found, unknown course id",
			courseId: mock.UnknownNumPK,
			slideId:  mock.PkNum,
			result:   gorm.ErrRecordNotFound,
		},
		{
			desc:     "Not found, unknown bulletin id",
			courseId: mock.PkNum,
			slideId:  mock.UnknownNumPK,
			result:   gorm.ErrRecordNotFound,
		},
		{
			desc:     "Normal",
			courseId: mock.PkNum,
			slideId:  mock.PkNum,
			result:   nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			slide = domain.Slides{
				AutoModel: domain.AutoModel{Id: tC.slideId},
				CourseId:  tC.courseId,
				Chapter:   mock.NewMsg(),
				File: domain.File{
					Title: mock.NewMsg(),
					Name:  mock.NewMsg(),
				},
			}
			l, err := repo.UpdateSlideById(ctx, &slide)
			fmt.Println(l)

			assert.Equal(t, tC.result, err)
		})
	}
}
