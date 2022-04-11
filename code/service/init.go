package service

import (
	"context"
	"encoding/json"
	"errors"
	"teacher-site/model"

	"github.com/go-redis/redis"
)

func (srv *Service) GetInit(ctx context.Context, init *model.Init) error {
	var err error

	data, err := srv.cache.GetInit(ctx, srv.domain)
	if errors.Is(err, redis.Nil) {
		if err = srv.db.GetInit(ctx, init, srv.domain); err != nil {
			return err
		}
		b, err := json.Marshal(init)
		if err != nil {
			return err
		}
		if err = srv.cache.SetInit(ctx, srv.domain, string(b)); err != nil {
			return err
		}
	}
	if err = json.Unmarshal([]byte(data), &init); err != nil {
		return err
	}
	return nil
}
