package database

import (
	"teacher-site/mock"
	"teacher-site/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetCourseWithContent(t *testing.T) {
	tC := struct {
		courseID uint
		result   error
	}{
		courseID: 1,
		result:   nil,
	}
	_, err := db.GetCourseContent(ctx, tC.courseID)
	assert.Equal(t, tC.result, err)
}

func TestGetInit(t *testing.T) {
	tC := struct {
		domain string
		result error
	}{
		domain: mock.Domain,
		result: nil,
	}

	init := model.Init{}
	err := db.GetInit(ctx, &init, tC.domain)
	assert.Equal(t, tC.result, err)
}

func TestDomainIsExist(t *testing.T) {
	tC := []struct {
		desc   string
		domain string
		result error
	}{
		{
			desc:   "Real domain",
			domain: mock.Domain,
			result: nil,
		},
		{
			desc:   "Not existed of domain",
			domain: mock.Unknown,
			result: gorm.ErrRecordNotFound,
		},
	}
	for _, v := range tC {
		t.Run(v.desc, func(t *testing.T) {
			err := db.DomainIsExist(ctx, v.domain)
			assert.Equal(t, v.result, err)
		})
	}
}
func TestUserIsExist(t *testing.T) {
	tC := []struct {
		desc   string
		userID string
		result error
	}{
		{
			desc:   "Real user",
			userID: mock.UserID,
			result: nil,
		},
		{
			desc:   "Not existed of user",
			userID: mock.Unknown,
			result: gorm.ErrRecordNotFound,
		},
	}
	for _, v := range tC {
		t.Run(v.desc, func(t *testing.T) {
			err := db.UserIsExist(ctx, v.userID)
			assert.Equal(t, v.result, err)
		})
	}
}

func TestGetAuth(t *testing.T) {
	var auth *model.Auths
	tC := []struct {
		desc   string
		userID string
		result error
	}{
		{
			desc:   "Real user",
			userID: mock.UserID,
			result: nil,
		},
		{
			desc:   "Not existed of user",
			userID: mock.Unknown,
			result: gorm.ErrRecordNotFound,
		},
	}
	for _, v := range tC {
		t.Run(v.desc, func(t *testing.T) {
			auth = &model.Auths{
				UserID: v.userID,
			}
			err := db.GetAuth(ctx, auth)
			assert.Equal(t, v.result, err)
		})
	}
}
