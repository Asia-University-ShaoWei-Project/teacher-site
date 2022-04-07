package service

import (
	"context"
	"encoding/json"
	"errors"
	"teacher-site/model"

	"github.com/go-redis/redis"
)

func (srv *Service) GetInit(ctx context.Context) (*model.Init, error) {
	var (
		init *model.Init
		err  error
	)
	data, err := srv.cache.GetInit(ctx, srv.domain)
	if errors.Is(err, redis.Nil) {
		init, err = srv.db.GetInit(ctx, srv.domain)
		if err != nil {
			srv.log.Error(err)
			return init, err
		}
		b, err := json.Marshal(init)
		if err != nil {
			srv.log.Error(err)
			return init, err
		}
		if err = srv.cache.SetInit(ctx, srv.domain, string(b)); err != nil {
			srv.log.Error(err)
			return init, err
		}
	}
	if err = json.Unmarshal([]byte(data), &init); err != nil {
		return init, err
	}
	return init, nil
}
