package service

import (
	"context"
	"encoding/json"
	"teacher-site/model"
	"time"
)

type NeedUpdate bool

func (srv *Service) GetCourse(ctx context.Context, courseBind *model.BindCourse) (model.Courses, error) {
	var course *model.Courses
	needUpdate := srv.courseNeedUpdate(ctx, courseBind)
	if needUpdate {
		data, err := srv.cache.GetCourseWithContent(srv.domain, courseBind.ID)
		if err != nil {
			srv.log.Error(err)
			course := srv.db.GetCourseWithContent(ctx, course.ID)
			err = srv.cache.SetCourse(srv.domain, course.ID, course)
			if err != nil {
				srv.log.Error(err)
				return model.Courses{}, err
			}
			updatedTime := time.Now().Unix()
			err = srv.cache.SetCourseLastUpdated(srv.domain, course.ID, updatedTime)
			if err != nil {
				srv.log.Error(err)
				return model.Courses{}, err
			}
			return model.Courses{}, err
		}
		err = json.Unmarshal([]byte(data), course)
		if err != nil {
			return model.Courses{}, err
		}
	}
	return *course, nil
}

func (srv *Service) courseNeedUpdate(ctx context.Context, courseBind *model.BindCourse) NeedUpdate {
	lastUpdated, err := srv.cache.GetCourseLastUpdated(srv.domain, courseBind.ID)
	if err != nil {
		srv.log.Error(err)
		return false
	}
	if courseBind.LastUpdated == lastUpdated {
		return false
	}
	return true
}
