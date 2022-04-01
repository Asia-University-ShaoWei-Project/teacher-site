package service

import (
	"context"
	"teacher-site/message"
	"teacher-site/model"
	"time"
)

func (srv *Service) CreateInfo(ctx context.Context, reqInfo *model.BindInfo) error {
	date := time.Now().Format(model.InfoDateFormat)
	info := &model.Informations{
		TeacherID:   srv.domain,
		CreatedDate: date,
		Info:        reqInfo.Info,
	}
	err := srv.db.CreateInformation(info)
	return err
}

func (srv *Service) UpdateInfo(ctx context.Context, reqInfo *model.BindInfo) error {
	if reqInfo.Info == "" {
		return message.ErrDataEmpty
	}
	info := &model.Informations{
		AutoModel:   model.AutoModel{ID: reqInfo.ID},
		CreatedDate: reqInfo.CreateDate,
		Info:        reqInfo.Info,
	}
	return srv.db.UpdateInformation(info)
}
func (srv *Service) DeleteInfo(ctx context.Context, reqInfo *model.BindInfo) error {
	return srv.db.DeleteInformation(reqInfo.ID)
}

// func (srv *Service) CreateBulletin(ctx context.Context, bulletin model.BulletinBoard) {
// 	updateTime := string(time.Now().Unix())
// 	srv.db.Create(bulletin)
// 	course := model.Courses{AutoModel: model.AutoModel{ID: bulletin.CourseID}}
// 	srv.db.UpdateColumn("last_updated", updateTime).Find(&course)

// srv.setCourseLastUpdatedCache(ctx, key string, courseLastUpdated string)

// course, err := srv.getCourseCache(ctx, key)
// if err != nil {
// 	fmt.Println("start get course data by db")
// 	course = srv.getCourseDB(ctx, id)
// 	fmt.Println("start save course data into redis")
// 	err = srv.setCourseCache(ctx, course, key, cache.TimeForever)
// 	srv.checkErr(err)
// 	fmt.Println("store last_updated")
// 	err = srv.setCourseLastUpdatedCache(ctx, key, course.LastUpdated)
// 	srv.checkErr(err)
// }
// return course
// }

// func (srv *Service) authTeacher(ctx context.Context, token string) (*model.Teachers, error) {
// 	var teacher model.Teachers
// 	sql := `SELECT domain
// 	FROM teachers t
// 	INNER JOIN auths a
// 	ON t.auth_id = a.user_id
// 	WHERE a.token = ?`
// 	err := srv.db.Raw(sql, token).Scan(&teacher).Error
// 	return &teacher, err
// }
