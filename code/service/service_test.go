package service

import (
	"context"
	"teacher-site/cache"
	"teacher-site/database"
	"teacher-site/log"
	"teacher-site/message"
	"teacher-site/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	db  = database.NewSqlite("../database")
	c   = cache.NewCache()
	ctx = context.Background()
	// logger     = log.NewLog(ctx)
	logger = log.NewLogrus(ctx)
	srv    = NewService(db, c, logger)
	cfg    = &model.Config{
		JWTSecure:      []byte(`secure`),
		PasswordSecure: `secure`,
		HashCost:       10,
	}
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
			result: message.ErrQueryNotFound,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err = srv.IsExistDomain(ctx, &tC.domain)
			assert.Equal(t, tC.result, err)
		})
	}
}

// token: "b6e72454-eaf5-4477-8148-5dd601182af4",

// jwt update
// no jwt update
// info content error
