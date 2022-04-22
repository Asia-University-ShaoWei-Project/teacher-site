package middleware

import (
	"context"
	"fmt"
	"teacher-site/config"
	"teacher-site/mock"
	"teacher-site/pkg/util"
	"testing"
)

var (
	ctx  = context.Background()
	conf = config.New()
)

func TestVerifyJwtValid(t *testing.T) {
	token, _ := util.GenerateJwt(conf.Jwt, mock.UserID)
	testCases := []struct {
		desc        string
		bearerToken string
	}{
		{
			desc:        "fail token",
			bearerToken: mock.EmptyStr,
			// result:      errors.New("token contains an invalid number of segments"),
		},
		{
			desc:        "real token",
			bearerToken: token,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := verifyJwtValid(ctx, tC.bearerToken, conf.Jwt.Secret)
			// todo: error handle
			// assert.Equal(t, tC.result.Error(), err.Error())
			fmt.Println(err)

		})
	}
}
