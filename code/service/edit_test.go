package service

import (
	"teacher-site/message"
	"teacher-site/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInfo(t *testing.T) {
	info := &model.BindInfo{
		Info: "testing case",
	}
	if err := srv.CreateInfo(ctx, info); err != nil {
		srv.Debug(err)
	}
}

func TestUpdateInfo(t *testing.T) {
	tC := []struct {
		desc   string
		info   *model.BindInfo
		result error
	}{
		{
			desc: "normal data for update",
			info: &model.BindInfo{
				ID:   1,
				Info: "testing normal data",
			},
			result: nil,
		},
		{
			desc: "id not exist",
			info: &model.BindInfo{
				ID:   999,
				Info: "testing case2",
			},
			result: message.ErrQueryNotFound,
		},
		{
			desc: "empty info",
			info: &model.BindInfo{
				ID:   2,
				Info: "",
			},
			result: message.ErrDataEmpty,
		},
	}
	for _, v := range tC {
		t.Run(v.desc, func(t *testing.T) {
			err := srv.UpdateInfo(ctx, v.info)
			assert.Equal(t, v.result, err)
		})
	}
}
