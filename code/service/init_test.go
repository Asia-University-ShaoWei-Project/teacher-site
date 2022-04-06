package service

import (
	"testing"
)

func TestGetInit(t *testing.T) {
	testCases := []struct {
		desc   string
		domain string
	}{
		{
			desc:   `Use "teacher-1" domain to get option`,
			domain: "teacher_domain",
		},
		{
			desc:   `Use "unknown" domain to get option`,
			domain: "unknown",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			srv.SetDomain(ctx, &tC.domain)
			//TODO: When deleted key in redis, can you get data from db?
			init, _ := srv.GetInit(ctx)
			srv.Info(init)
		})
	}
}
