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
	repo = NewInfoRepository(db, conf.DB)
)

type createInfoResult struct {
	err error
}

func TestCreateInfo(t *testing.T) {
	var req domain.CreateInfoBulletinRequest
	testCases := []struct {
		desc    string
		domain  string
		content string
		result  createInfoResult
	}{
		{
			desc:    "unknown the domain",
			domain:  mock.Unknown,
			content: mock.FailMsg,
			result:  createInfoResult{err: gorm.ErrRecordNotFound},
		},
		{
			desc:    "existed domain",
			domain:  mock.TeacherDomain,
			content: mock.CreateMsg,
			result:  createInfoResult{err: nil},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req = domain.CreateInfoBulletinRequest{
				TeacherDomainRequest: domain.TeacherDomainRequest{TeacherDomain: tC.domain},
				Content:              tC.content,
			}
			if tC.result.err == nil {
				oldLastModified, _ = repo.GetLastModified(ctx, req.TeacherDomain)
			}
			_, err := repo.Create(ctx, &req)
			assert.Equal(t, tC.result.err, err)
			if tC.result.err == nil {
				newLastModified, _ = repo.GetLastModified(ctx, req.TeacherDomain)
				assert.NotEqual(t, newLastModified, oldLastModified, fmt.Sprintf("new:%s, old:%s", newLastModified, oldLastModified))
			}
		})
	}
}

//
type getInfoResult struct {
	isEmpty bool
}

func TestGetInfo(t *testing.T) {
	var req domain.GetInfoBulletinRequest

	testCases := []struct {
		desc   string
		domain string
		result getInfoResult
	}{
		{
			desc:   "unknown the domain",
			domain: mock.Unknown,
			result: getInfoResult{isEmpty: true},
		},
		{
			desc:   "real domain",
			domain: mock.TeacherDomain,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req = domain.GetInfoBulletinRequest{
				TeacherDomainRequest: domain.TeacherDomainRequest{TeacherDomain: tC.domain},
			}
			infoBulletin, err := repo.Get(ctx, &req)
			assert.Nil(t, err)
			fmt.Println(infoBulletin)

			if tC.result.isEmpty {
				assert.Empty(t, infoBulletin)
			} else {
				assert.NotEmpty(t, infoBulletin)
			}
		})
	}
}

type getInfoLastUpdatedResult struct {
	isEmpty bool
}

func TestGetInfoLastUpdated(t *testing.T) {
	testCases := []struct {
		desc   string
		domain string
		result getInfoLastUpdatedResult
	}{
		{
			desc:   "unknown the domain",
			domain: mock.Unknown,
			result: getInfoLastUpdatedResult{
				isEmpty: true,
			},
		},
		{
			desc:   "existed domain",
			domain: mock.TeacherDomain,
			result: getInfoLastUpdatedResult{
				isEmpty: false,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lastModified, err := repo.GetLastModified(ctx, tC.domain)
			assert.Nil(t, err)
			if tC.result.isEmpty {
				assert.Empty(t, lastModified)
			} else {
				assert.NotEmpty(t, lastModified)
			}
		})
	}
}

type updateInfoLastResult struct {
	err error
}

func TestUpdateInfo(t *testing.T) {
	testCases := []struct {
		desc   string
		domain string
		req    domain.UpdateInfoBulletinRequest
		result updateInfoLastResult
	}{
		{
			desc:   "unknown the domain",
			domain: mock.Unknown,
			req: domain.UpdateInfoBulletinRequest{
				TeacherDomainRequest: domain.TeacherDomainRequest{TeacherDomain: mock.Unknown},
				BulletinID:           mock.NumPK,
				Content:              mock.FailMsg + "unknown the domain",
			},
			result: updateInfoLastResult{
				err: gorm.ErrRecordNotFound,
			},
		},
		{
			desc:   "existed domain",
			domain: mock.TeacherDomain,
			req: domain.UpdateInfoBulletinRequest{
				TeacherDomainRequest: domain.TeacherDomainRequest{TeacherDomain: mock.TeacherDomain},
				BulletinID:           mock.NumPK,
				Content:              mock.UpdateMsg + "existed domain",
			},
			result: updateInfoLastResult{
				err: nil,
			},
		},
		{
			desc:   "unknown the index",
			domain: mock.Unknown,
			req: domain.UpdateInfoBulletinRequest{
				TeacherDomainRequest: domain.TeacherDomainRequest{TeacherDomain: mock.TeacherDomain},
				BulletinID:           mock.UnknownNumPK,
				Content:              mock.FailMsg + ", unknown the index",
			},
			result: updateInfoLastResult{
				err: gorm.ErrRecordNotFound,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.result.err == nil {
				oldLastModified, _ = repo.GetLastModified(ctx, tC.req.TeacherDomain)

			}
			// todo: test info value
			_, err := repo.Update(ctx, &tC.req)
			assert.Equal(t, tC.result.err, err)
			if tC.result.err == nil {
				newLastModified, _ = repo.GetLastModified(ctx, tC.req.TeacherDomain)
				assert.NotEqual(t, newLastModified, oldLastModified, fmt.Sprintf("new:%s, old:%s", newLastModified, oldLastModified))
			}

		})
	}
}

type deleteInfoLastResult struct {
	err error
}

func TestDeleteInfo(t *testing.T) {
	var req domain.DeleteInfoBulletinRequest
	// generate a temporary bulletin for testcase(id)
	bulletin, _ := repo.Create(ctx, &domain.CreateInfoBulletinRequest{
		TeacherDomainRequest: domain.TeacherDomainRequest{TeacherDomain: mock.TeacherDomain},
		Content:              mock.CreateMsg,
	})
	testCases := []struct {
		desc   string
		id     uint
		result deleteInfoLastResult
	}{

		{
			desc: "existed id",
			id:   bulletin.AutoModel.ID,
			result: deleteInfoLastResult{
				err: nil,
			},
		},
		{
			desc: "unknown id",
			id:   mock.UnknownNumPK,
			result: deleteInfoLastResult{
				err: gorm.ErrRecordNotFound,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req = domain.DeleteInfoBulletinRequest{
				TeacherDomainRequest: domain.TeacherDomainRequest{TeacherDomain: mock.TeacherDomain},
				BulletinID:           tC.id,
			}
			if tC.result.err == nil {
				oldLastModified, _ = repo.GetLastModified(ctx, mock.TeacherDomain)
			}
			info, err := repo.Delete(ctx, &req)
			assert.Equal(t, tC.result.err, err)
			if tC.result.err == nil {
				err := testCheckInfoIsExistById(req.BulletinID, t)
				// the data is deleted, so that should be not found the data
				assert.Equal(t, gorm.ErrRecordNotFound, err)
				// the info.LastModified is newest, that should be not equal to old date
				assert.NotEqual(t, info.LastModified, oldLastModified, fmt.Sprintf("new:%s, old:%s", info.LastModified, oldLastModified))
			}
		})
	}
}
func testCheckInfoIsExistById(id uint, t *testing.T) error {
	info := domain.InfoBulletinBoards{AutoModel: domain.AutoModel{ID: id}}
	result := db.Where(`id=? AND deleted_at IS NULL`, id).Find(&info)
	t.Error(result.Error)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
