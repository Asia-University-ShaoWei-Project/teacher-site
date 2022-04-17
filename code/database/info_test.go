package database

import (
	"fmt"
	"teacher-site/mock"
	"teacher-site/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	oldLastModified string
	newLastModified string
)

type CreateInfo struct {
	err error
}

func TestCreateInfo(t *testing.T) {
	testCases := []struct {
		desc   string
		domain string
		req    model.ReqCreateInfo
		result CreateInfo
	}{
		{
			desc:   "unknown the domain",
			domain: mock.Unknown,
			req: model.ReqCreateInfo{
				Content: mock.FailMsg,
			},
			result: CreateInfo{err: gorm.ErrRecordNotFound},
		},
		{
			desc:   "existed domain",
			domain: mock.Domain,
			req: model.ReqCreateInfo{
				Content: mock.CreateMsg,
			},
			result: CreateInfo{err: nil},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.result.err == nil {
				oldLastModified, _ = db.GetInfoLastUpdated(ctx, tC.domain)
			}
			_, err := db.CreateInfo(ctx, tC.domain, &tC.req)
			assert.Equal(t, tC.result.err, err)
			time.Sleep(1 * time.Second)
			if tC.result.err == nil {
				newLastModified, _ = db.GetInfoLastUpdated(ctx, tC.domain)
				assert.NotEqual(t, newLastModified, oldLastModified, fmt.Sprintf("new:%s, old:%s", newLastModified, oldLastModified))
			}
		})
	}
}

//
type GetInfoResult struct {
	isEmpty bool
}

func TestGetInfo(t *testing.T) {
	testCases := []struct {
		desc   string
		domain string
		result GetInfoResult
	}{
		{
			desc:   "unknown the domain",
			domain: mock.Unknown,
			result: GetInfoResult{isEmpty: true},
		},
		{
			desc:   "existed domain",
			domain: mock.Domain,
			result: GetInfoResult{isEmpty: false},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			infoBulletin, err := db.GetInfo(ctx, tC.domain)
			fmt.Println(infoBulletin)

			assert.Nil(t, err)
			if tC.result.isEmpty {
				assert.Empty(t, infoBulletin)
			} else {
				assert.NotEmpty(t, infoBulletin)
			}
		})
	}
}

type GetInfoLastUpdatedResult struct {
	isEmpty bool
}

func TestGetInfoLastUpdated(t *testing.T) {
	testCases := []struct {
		desc   string
		domain string
		result GetInfoLastUpdatedResult
	}{
		{
			desc:   "unknown the domain",
			domain: mock.Unknown,
			result: GetInfoLastUpdatedResult{
				isEmpty: true,
			},
		},
		{
			desc:   "existed domain",
			domain: mock.Domain,
			result: GetInfoLastUpdatedResult{
				isEmpty: false,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lastModified, err := db.GetInfoLastUpdated(ctx, tC.domain)
			assert.Nil(t, err)
			if tC.result.isEmpty {
				assert.Empty(t, lastModified)
			} else {
				assert.NotEmpty(t, lastModified)
			}
		})
	}
}

type UpdateInfoLastResult struct {
	err error
}

func TestUpdateInfo(t *testing.T) {
	testCases := []struct {
		desc   string
		domain string
		req    model.ReqUpdateInfoBulletin
		result UpdateInfoLastResult
	}{
		{
			desc:   "unknown the domain",
			domain: mock.Unknown,
			req: model.ReqUpdateInfoBulletin{
				BulletinID: mock.NumKey,
				Content:    mock.FailMsg + "domain",
			},
			result: UpdateInfoLastResult{
				err: gorm.ErrRecordNotFound,
			},
		},
		{
			desc:   "existed domain",
			domain: mock.Domain,
			req: model.ReqUpdateInfoBulletin{
				BulletinID: mock.NumKey,
				Content:    mock.UpdateMsg,
			},
			result: UpdateInfoLastResult{
				err: nil,
			},
		},
		{
			desc:   "unknown the index",
			domain: mock.Unknown,
			req: model.ReqUpdateInfoBulletin{
				BulletinID: mock.UnknownNumKey,
				Content:    mock.FailMsg + "index",
			},
			result: UpdateInfoLastResult{
				err: gorm.ErrRecordNotFound,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.result.err == nil {
				oldLastModified, _ = db.GetInfoLastUpdated(ctx, tC.domain)
			}
			err := db.UpdateInfo(ctx, tC.domain, &tC.req)
			assert.Equal(t, tC.result.err, err)
			if tC.result.err == nil {
				newLastModified, _ = db.GetInfoLastUpdated(ctx, tC.domain)
				assert.NotEqual(t, newLastModified, oldLastModified, fmt.Sprintf("new:%s, old:%s", newLastModified, oldLastModified))
			}

		})
	}
}

// todo test delete
type DeleteInfoLastResult struct {
	err error
}

func TestDeleteInfo(t *testing.T) {
	testCases := []struct {
		desc   string
		domain string
		req    model.ReqDeleteInfo
		result DeleteInfoLastResult
	}{
		{
			desc:   "unknown the domain",
			domain: mock.Unknown,
			req: model.ReqDeleteInfo{
				BulletinID: mock.NumKey,
			},
			result: DeleteInfoLastResult{
				err: gorm.ErrRecordNotFound,
			},
		},
		{
			desc:   "existed domain",
			domain: mock.Domain,
			req: model.ReqDeleteInfo{
				BulletinID: mock.NumKey,
			},
			result: DeleteInfoLastResult{
				err: nil,
			},
		},
		{
			desc:   "unknown the key",
			domain: mock.Unknown,
			req: model.ReqDeleteInfo{
				BulletinID: mock.UnknownNumKey,
			},
			result: DeleteInfoLastResult{
				err: gorm.ErrRecordNotFound,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.result.err == nil {
				oldLastModified, _ = db.GetInfoLastUpdated(ctx, tC.domain)
			}
			err := db.DeleteInfo(ctx, tC.domain, &tC.req)
			info, _ := db.GetInfo(ctx, tC.domain)
			for _, v := range info {
				assert.NotEqual(t, tC.req.BulletinID, v.ID)
			}
			assert.Equal(t, tC.result.err, err)
			if tC.result.err == nil {
				newLastModified, _ = db.GetInfoLastUpdated(ctx, tC.domain)
				assert.NotEqual(t, newLastModified, oldLastModified, fmt.Sprintf("new:%s, old:%s", newLastModified, oldLastModified))
			}
			// db.GetInfo(ctx, tC.domain)

		})
	}
}
