package repository

import (
	"context"
	"fmt"
	"teacher-site/config"
	"teacher-site/domain"
	"teacher-site/mock"
	"teacher-site/pkg/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// todo: remove the unknown teacher domain of testcase

var (
	oldLastModified, newLastModified string
)
var (
	ctx  = context.Background()
	conf = config.New()
	db   = database.NewDB("../../../../pkg/database", conf.DB)
	repo = NewDbRepository(db, conf.DB)
)

func TestCreateInfo(t *testing.T) {
	var req domain.CreateInfoBulletinRequest
	testCases := []struct {
		desc   string
		infoID uint
		result error
	}{
		{
			desc:   "fail info id",
			infoID: mock.UnknownNumPK,
			result: gorm.ErrRecordNotFound,
		},
		{
			desc:   "existed info id",
			infoID: mock.NumPK,
			result: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req = domain.CreateInfoBulletinRequest{
				TeacherDomainRequest: domain.TeacherDomainRequest{TeacherDomain: mock.TeacherDomain},
				InfoID:               tC.infoID,
				Content:              mock.NewMsg(),
			}
			if tC.result == nil {
				oldLastModified, _ = repo.GetLastModified(ctx, req.InfoID)
			}
			_, err := repo.Create(ctx, &req)
			assert.Equal(t, tC.result, err)
			if tC.result == nil {
				newLastModified, _ = repo.GetLastModified(ctx, req.InfoID)
				assert.NotEqual(t, newLastModified, oldLastModified, fmt.Sprintf("new:%s, old:%s", newLastModified, oldLastModified))
			}
		})
	}
}
func TestGetByTeacherDomain(t *testing.T) {
	testCases := []struct {
		desc   string
		domain string
		result error
	}{
		{
			desc:   "fail teacher domain",
			domain: mock.Unknown,
			result: gorm.ErrRecordNotFound,
		},
		{
			desc:   "normal",
			domain: mock.TeacherDomain,
			result: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := repo.GetByTeacherDomain(ctx, tC.domain)
			assert.Equal(t, tC.result, err)
		})
	}
}
func TestGetBulletinsByInfoId(t *testing.T) {
	testCases := []struct {
		desc   string
		infoID uint
		result error
	}{
		{
			desc:   "fail info id",
			infoID: mock.UnknownNumPK,
			result: gorm.ErrRecordNotFound,
		},
		{
			desc:   "existed info id",
			infoID: mock.NumPK,
			result: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := repo.GetBulletinsByInfoId(ctx, tC.infoID)
			assert.Equal(t, tC.result, err)
		})
	}
}

func TestGetInfoLastUpdated(t *testing.T) {
	testCases := []struct {
		desc   string
		infoID uint
		result error
	}{
		{
			desc:   "fail info id",
			infoID: mock.UnknownNumPK,
			result: gorm.ErrRecordNotFound,
		},
		{
			desc:   "existed info id",
			infoID: mock.NumPK,
			result: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := repo.GetLastModified(ctx, tC.infoID)
			assert.Equal(t, tC.result, err)
		})
	}
}

func TestUpdateInfo(t *testing.T) {
	var req domain.UpdateInfoBulletinRequest
	testCases := []struct {
		desc       string
		infoID     uint
		bulletinID uint
		result     error
	}{

		{
			desc:       "fail info id",
			infoID:     mock.UnknownNumPK,
			bulletinID: mock.NumPK,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "fail bulletin id",
			infoID:     mock.NumPK,
			bulletinID: mock.UnknownNumPK,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "existed info, bulletin id",
			infoID:     mock.NumPK,
			bulletinID: mock.NumPK,
			result:     nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.result == nil {
				oldLastModified, _ = repo.GetLastModified(ctx, tC.infoID)
			}
			req = domain.UpdateInfoBulletinRequest{
				TeacherDomainRequest: domain.TeacherDomainRequest{TeacherDomain: mock.TeacherDomain},
				InfoID:               tC.infoID,
				BulletinID:           tC.bulletinID,
				Content:              mock.NewMsg(),
			}
			_, err := repo.Update(ctx, &req)
			assert.Equal(t, tC.result, err)
			if tC.result == nil {
				newLastModified, _ = repo.GetLastModified(ctx, tC.infoID)
				assert.NotEqual(t, newLastModified, oldLastModified, fmt.Sprintf("new:%s, old:%s", newLastModified, oldLastModified))
			}
		})
	}
}

func TestDeleteInfo(t *testing.T) {
	var req domain.DeleteInfoBulletinRequest
	// generate a temporary bulletin for testcase(id)
	bulletin, _ := repo.Create(ctx, &domain.CreateInfoBulletinRequest{
		TeacherDomainRequest: domain.TeacherDomainRequest{TeacherDomain: mock.TeacherDomain},
		InfoID:               mock.NumPK,
		Content:              mock.NewMsg(),
	})
	testCases := []struct {
		desc       string
		infoID     uint
		bulletinID uint
		result     error
	}{

		{
			desc:       "fail info id",
			infoID:     mock.UnknownNumPK,
			bulletinID: bulletin.AutoModel.ID,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "fail bulletin id",
			infoID:     mock.NumPK,
			bulletinID: mock.UnknownNumPK,
			result:     gorm.ErrRecordNotFound,
		},
		{
			desc:       "existed info, bulletin id",
			infoID:     mock.NumPK,
			bulletinID: bulletin.AutoModel.ID,
			result:     nil,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req = domain.DeleteInfoBulletinRequest{
				TeacherDomainRequest: domain.TeacherDomainRequest{TeacherDomain: mock.TeacherDomain},
				InfoID:               tC.infoID,
				BulletinID:           tC.bulletinID,
			}
			if tC.result == nil {
				oldLastModified, _ = repo.GetLastModified(ctx, tC.infoID)
				fmt.Println(`before time:`, oldLastModified)

			}
			info, err := repo.Delete(ctx, &req)
			assert.Equal(t, tC.result, err)
			if tC.result == nil {
				err := testCheckBulletinIsExistById(req.BulletinID, t)
				// the data is deleted, so that should be not found the data
				assert.Equal(t, gorm.ErrRecordNotFound, err)
				// the info.LastModified is newest, that should be not equal to old date
				newLastModified, _ = repo.GetLastModified(ctx, tC.infoID)
				// todo: test error! issue: new last-modified == old last-modified
				assert.NotEqual(t, newLastModified, oldLastModified, fmt.Sprintf("new:%s, old:%s", info.LastModified, oldLastModified))
			}
		})
	}
}
func testCheckBulletinIsExistById(id uint, t *testing.T) error {
	info := domain.InfoBulletinBoards{AutoModel: domain.AutoModel{ID: id}}
	result := db.Where(`id=? AND deleted_at IS NULL`, id).Find(&info)
	t.Error(result.Error)
	return checkErrAndExist(result)
}
