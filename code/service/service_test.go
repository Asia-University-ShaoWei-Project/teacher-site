package service

import (
	"context"
	"teacher-site/cache"
	"teacher-site/database"
	"teacher-site/logsrv"
	"teacher-site/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	ctx    = context.Background()
	c      = cache.NewCache()
	logger = logsrv.NewLogrus(ctx)
	db     = database.NewSqlite("../database", logger)
	conf   = model.NewTMPConfig()
	srv    = NewService(db, c, logger, conf)
)

func TestIsExistDomain(t *testing.T) {
	var err error
	testCases := []struct {
		desc   string
		domain string
		result error
	}{
		{
			desc:   "domain = teacher_domain",
			domain: "teacher_domain",
			result: nil,
		},
		{
			desc:   "domain = unknown domain",
			domain: "other_domain",
			result: gorm.ErrRecordNotFound,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err = srv.IsExistDomain(ctx, &tC.domain)
			assert.Equal(t, tC.result, err)
		})
	}
}
