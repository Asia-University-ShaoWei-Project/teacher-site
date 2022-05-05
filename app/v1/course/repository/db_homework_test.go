package repository

import (
	"teacher-site/domain"
	"teacher-site/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateHomework(t *testing.T) {
	var homework domain.Homeworks
	testCases := []struct {
		desc     string
		courseId uint
		result   error
	}{
		{
			desc:     "Unknown the course id",
			courseId: mock.UnknownNumPK,
			result:   gorm.ErrRecordNotFound,
		},
		{
			desc:     "Normal",
			courseId: mock.NumPk,
			result:   nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			homework.CourseId = tC.courseId
			_, err = repo.CreateHomework(ctx, &homework)

			assert.Equal(t, tC.result, err)
		})
	}
}

func TestUpdateHomeworkById(t *testing.T) {
	var homework domain.Homeworks
	_newHomework := domain.Homeworks{CourseId: mock.NumPk}

	if _, err = repo.CreateHomework(ctx, &_newHomework); err != nil {
		t.Fatal("create tmp homework error:", err)
	}
	testCases := []struct {
		desc       string
		courseId   uint
		homeworkId uint
		result     error
	}{
		{
			desc:       "Both the id is not found",
			courseId:   mock.UnknownNumPK,
			homeworkId: mock.UnknownNumPK,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "Not found, unknown course id",
			courseId:   mock.UnknownNumPK,
			homeworkId: _newHomework.AutoModel.Id,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "Not found, unknown homework id",
			courseId:   _newHomework.CourseId,
			homeworkId: mock.UnknownNumPK,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "Normal",
			courseId:   _newHomework.CourseId,
			homeworkId: _newHomework.AutoModel.Id,
			result:     nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			homework = domain.Homeworks{
				AutoModel: domain.AutoModel{Id: tC.homeworkId},
				CourseId:  tC.courseId,
				Number:    mock.NewMsg(),
				File: domain.File{
					Title: mock.NewMsg(),
					Name:  mock.NewMsg(),
				},
			}
			_, err := repo.UpdateHomeworkById(ctx, &homework)

			assert.Equal(t, tC.result, err)
		})
	}
}
func TestDeleteHomeworkById(t *testing.T) {
	var homework domain.Homeworks
	_newHomework := domain.Homeworks{CourseId: mock.NumPk}

	if _, err = repo.CreateHomework(ctx, &_newHomework); err != nil {
		t.Fatal("create tmp homework error:", err)
	}

	testCases := []struct {
		desc       string
		courseId   uint
		homeworkId uint
		result     error
	}{
		{
			desc:       "Both the id is not found",
			courseId:   mock.UnknownNumPK,
			homeworkId: mock.UnknownNumPK,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "Not found, unknown course id",
			courseId:   mock.UnknownNumPK,
			homeworkId: _newHomework.AutoModel.Id,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "Not found, unknown homework id",
			courseId:   _newHomework.AutoModel.Id,
			homeworkId: mock.UnknownNumPK,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "Normal",
			courseId:   _newHomework.CourseId,
			homeworkId: _newHomework.AutoModel.Id,
			result:     nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			homework = domain.Homeworks{
				AutoModel: domain.AutoModel{Id: tC.homeworkId},
				CourseId:  tC.courseId,
			}
			_, err := repo.DeleteHomeworkById(ctx, &homework)
			assert.Equal(t, tC.result, err)
		})
	}
}
