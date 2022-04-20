package middleware

import (
	"context"
	"teacher-site/config"
	"teacher-site/mock"
	"teacher-site/pkg/log"
	"teacher-site/pkg/util"
	"testing"
)

var (
	logger = log.NewLogrus(ctx)
	ctx    = context.Background()
	conf   = config.New()
)

func TestVerifyJwtValid(t *testing.T) {
	token, _ := util.GenerateJwt(conf.Jwt, mock.UserID)
	tC := []struct {
		desc  string
		token string
	}{
		{
			desc:  "real token",
			token: token,
		},
		{
			desc:  "fail token",
			token: mock.Unknown,
		},
	}
	for _, v := range tC {
		t.Run(v.desc, func(t *testing.T) {
			bearerToken := "Bearer " + v.token
			err := verifyJwtValid(ctx, conf.Jwt.Secure, bearerToken)
			logger.Error(err)
		})
	}
}
