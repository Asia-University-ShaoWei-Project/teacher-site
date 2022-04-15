package service

import (
	"context"
	"encoding/json"
	"teacher-site/model"
)

func (srv *Service) GetInit(ctx context.Context, init *model.Init) error {
	// get from cache
	data, err := srv.cache.GetInit(ctx, srv.domain)
	if err == nil {
		if err = json.Unmarshal([]byte(data), &init); err == nil {
			// todo: handle error
			return nil
		}
	}
	// get from database
	if err != nil {
		// todo: error handle
		if err = srv.db.GetInit(ctx, init, srv.domain); err != nil {
			// todo: error handle
			return err
		}
		b, err := json.Marshal(init)
		if err != nil {
			// todo: error handle
			return err
		}
		if err = srv.cache.SetInit(ctx, srv.domain, string(b)); err != nil {
			// todo: error handle(e.g. connection refused)
			srv.Error(err)
		}
	}
	return nil
}
