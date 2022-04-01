package service

import (
	"context"
	"encoding/json"
	"teacher-site/model"
)

func (srv *Service) GetInit(ctx context.Context) (model.Init, error) {
	var init *model.Init
	data, err := srv.cache.GetInit(srv.domain)
	if err != nil {
		srv.log.Error(err)
		return model.Init{}, err
	}
	err = json.Unmarshal([]byte(data), init)
	if err != nil {
		srv.log.Error(err)
		init, err = srv.db.GetInit(srv.domain)
		if err != nil {
			srv.log.Error(err)
			return model.Init{}, err
		}
		err := srv.cache.SetInit(srv.domain, init)
		if err != nil {
			srv.log.Error(err)
			return model.Init{}, err
		}
	}
	return *init, nil
}

// func (srv *Service) getInitDataCache(ctx context.Context, key string) (InitData, error) {
// 	result, err := srv.cache.Get(key).Result()
// 	var data InitData
// 	if err != nil {
// 		return data, err
// 	}
// 	err = json.Unmarshal([]byte(result), &data)
// 	if err != nil {
// 		// TODO: log = get course by redis but got the error.
// 		srv.checkErr(err)
// 	}
// 	return data, nil
// }
// func (srv *Service) setInitDataCache(ctx context.Context, key string, initData InitData, expiration time.Duration) {
// 	// SELECT c.name_zh, c.name_us, c.token, c.content_updated
// 	data, err := json.Marshal(initData)
// 	if err != nil {
// 		// TODO: log = "save courses to redis was error"
// 		srv.checkErr(err)
// 	}
// 	fmt.Println("Value marshaled, start setting value!")
// 	srv.cache.Set(key, string(data), expiration)
// }

// func (srv *Service) getInitDataDB(ctx context.Context) InitData {
// 	var infos []model.Informations
// 	var coursesName []model.Courses

// infoSQL := `
// SELECT info, created_date
// FROM informations i
// INNER JOIN teachers t
// ON i.teacher_id = t.domain
// WHERE t.domain = ?
// `
// coursesSQL := `
// SELECT c.id, c.name_zh, c.name_us
// FROM courses c
// INNER JOIN teachers t
// ON c.teacher_id = t.domain
// WHERE t.domain=?`

// srv.db.Raw(infoSQL, srv.domain).Scan(&infos)
// srv.db.Raw(coursesSQL, srv.domain).Scan(&coursesName)

// 	return InitData{
// 		Courses:      coursesName,
// 		Informations: infos,
// 	}
// }
