package repository

import (
	"teacher-site/domain"
	"teacher-site/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateBulletin(t *testing.T) {
	var bulletin domain.BulletinBoards
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
			courseId: mock.PkNum,
			result:   nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			bulletin.CourseId = tC.courseId
			_, err = repo.CreateBulletin(ctx, &bulletin)
			assert.Equal(t, tC.result, err)
		})
	}
}

func TestUpdateBulletinById(t *testing.T) {
	var bulletin domain.BulletinBoards
	_newBulletin := domain.BulletinBoards{CourseId: mock.PkNum}

	if _, err = repo.CreateBulletin(ctx, &_newBulletin); err != nil {
		t.Fatal("create tmp bulletin error:", err)
	}

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
			bulletinId: _newBulletin.AutoModel.Id,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "Not found, unknown bulletin id",
			courseId:   _newBulletin.CourseId,
			bulletinId: mock.UnknownNumPK,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "Normal",
			courseId:   _newBulletin.CourseId,
			bulletinId: _newBulletin.AutoModel.Id,
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
			_, err := repo.UpdateBulletinById(ctx, &bulletin)
			assert.Equal(t, tC.result, err)
		})
	}
}

func TestDeleteBulletinById(t *testing.T) {
	var bulletin domain.BulletinBoards
	_newBulletin := domain.BulletinBoards{CourseId: mock.PkNum}

	if _, err = repo.CreateBulletin(ctx, &_newBulletin); err != nil {
		t.Fatal("create tmp bulletin error:", err)
	}

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
			bulletinId: _newBulletin.AutoModel.Id,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "Not found, unknown bulletin id",
			courseId:   _newBulletin.AutoModel.Id,
			bulletinId: mock.UnknownNumPK,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "Normal",
			courseId:   _newBulletin.CourseId,
			bulletinId: _newBulletin.AutoModel.Id,
			result:     nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			bulletin = domain.BulletinBoards{
				AutoModel: domain.AutoModel{Id: tC.bulletinId},
				CourseId:  tC.courseId,
			}
			_, err := repo.DeleteBulletinById(ctx, &bulletin)
			assert.Equal(t, tC.result, err)
		})
	}
}
