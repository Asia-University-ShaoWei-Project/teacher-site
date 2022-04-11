package service

import (
	"context"
	"encoding/json"
	"teacher-site/model"
	"time"
)

type NeedUpdate bool

func (srv *Service) GetCourse(ctx context.Context, courseBind *model.BindCourse, course *model.Courses) error {
	if needUpdate := srv.courseNeedUpdate(ctx, courseBind); needUpdate {
		data, err := srv.cache.GetCourseContent(ctx, srv.domain, courseBind.ID)
		if err == nil {
			if err = json.Unmarshal([]byte(data), &course); err != nil {
				return err
			} else {
				return nil
			}
		}
		_course, err := srv.db.GetCourseWithContent(ctx, course.ID)
		*course = _course
		if err != nil {
			return err
		}
		if err = srv.cache.SetCourseContent(ctx, srv.domain, course.ID, course); err != nil {
			return err
		}
		updatedTime := time.Now().Unix()
		if err = srv.cache.SetCourseLastUpdated(ctx, srv.domain, course.ID, updatedTime); err != nil {
			return err
		}
	}
	return nil
}

func (srv *Service) courseNeedUpdate(ctx context.Context, courseBind *model.BindCourse) NeedUpdate {
	lastUpdated, err := srv.cache.GetCourseLastUpdated(ctx, srv.domain, courseBind.ID)
	if err != nil {
		// todo error handle
		return false
	}
	if courseBind.LastUpdated == lastUpdated {
		return false
	}
	return true
}
