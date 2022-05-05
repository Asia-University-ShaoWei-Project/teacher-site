package repository

import (
	"teacher-site/domain"
	"teacher-site/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateSlide(t *testing.T) {
	var slide domain.Slides
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
			slide.CourseId = tC.courseId
			_, err = repo.CreateSlide(ctx, &slide)

			assert.Equal(t, tC.result, err)
		})
	}
}

func TestUpdateSlideById(t *testing.T) {
	var slide domain.Slides
	_newSlide := domain.Slides{CourseId: mock.NumPk}

	if _, err = repo.CreateSlide(ctx, &_newSlide); err != nil {
		t.Fatal("create tmp slide error:", err)
	}
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
			slideId:  _newSlide.AutoModel.Id,
			result:   gorm.ErrRecordNotFound,
		},
		{
			desc:     "Not found, unknown slide id",
			courseId: _newSlide.CourseId,
			slideId:  mock.UnknownNumPK,
			result:   gorm.ErrRecordNotFound,
		},
		{
			desc:     "Normal",
			courseId: _newSlide.CourseId,
			slideId:  _newSlide.AutoModel.Id,
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
			_, err := repo.UpdateSlideById(ctx, &slide)

			assert.Equal(t, tC.result, err)
		})
	}
}
func TestDeleteSlideById(t *testing.T) {
	var slide domain.Slides
	_newSlide := domain.Slides{CourseId: mock.NumPk}

	if _, err = repo.CreateSlide(ctx, &_newSlide); err != nil {
		t.Fatal("create tmp slide error:", err)
	}

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
			slideId:  _newSlide.AutoModel.Id,
			result:   gorm.ErrRecordNotFound,
		},
		{
			desc:     "Not found, unknown slide id",
			courseId: _newSlide.AutoModel.Id,
			slideId:  mock.UnknownNumPK,
			result:   gorm.ErrRecordNotFound,
		},
		{
			desc:     "Normal",
			courseId: _newSlide.CourseId,
			slideId:  _newSlide.AutoModel.Id,
			result:   nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			slide = domain.Slides{
				AutoModel: domain.AutoModel{Id: tC.slideId},
				CourseId:  tC.courseId,
			}
			_, err := repo.DeleteSlideById(ctx, &slide)
			assert.Equal(t, tC.result, err)
		})
	}
}
