package service

import (
	"teacher-site/mock"
	"teacher-site/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInit(t *testing.T) {
	tC := []struct {
		desc   string
		domain string
		result error
	}{
		{
			desc:   `Real domain`,
			domain: mock.Domain,
			result: nil,
		},
		{
			desc:   `Unknown domain`,
			domain: mock.Unknown,
			// todo
			result: nil,
		},
	}
	for _, v := range tC {
		t.Run(v.desc, func(t *testing.T) {
			srv.SetDomain(ctx, v.domain)
			//TODO: When deleted key in redis, can you get data from db?
			err := srv.GetInit(ctx, &model.Init{})
			assert.Equal(t, v.result, err, err)
		})
	}
}
