package service

import (
	"context"
	"encoding/json"
	"errors"
	"teacher-site/model"
	"time"
)

type NeedUpdate bool

var ErrNotUpdate = errors.New("not need to update")

func (srv *Service) GetCourse(ctx context.Context, courseBind *model.BindCourse) (model.Courses, error) {
	// get from cache
	var course model.Courses
	needUpdate := srv.courseNeedUpdate(ctx, courseBind)
	if !needUpdate {
		return course, ErrNotUpdate
	}
	data, err := srv.cache.GetCourseContent(ctx, srv.domain, courseBind.ID)
	if err != nil {
		return course, err
	}
	if err = json.Unmarshal([]byte(data), &course); err == nil {
		// get data from cache is completed.
		return course, nil
	}
	// todo: error handle - cache
	srv.Error(err)
	// get from database
	// todo: use TX to read and write the last_updated
	course, err = srv.db.GetCourseContent(ctx, course.ID)
	if err != nil {
		return course, err
	}
	if err = srv.cache.SetCourseContent(ctx, srv.domain, &course); err != nil {
		// todo: error handle
		srv.Error(err)
	}
	updatedTime := time.Now().Unix()
	if err = srv.cache.SetCourseLastUpdated(ctx, srv.domain, course.ID, updatedTime); err != nil {
		// todo: error handle
		srv.Error(err)
	}
	return nil
}

func (srv *Service) courseNeedUpdate(ctx context.Context, courseBind *model.BindCourse) NeedUpdate {
	lastUpdated, err := srv.cache.GetCourseLastUpdated(ctx, srv.domain, courseBind.ID)
	if err != nil {
		// todo error handle
		srv.Error(err)
		// from database
		lastUpdated, err = srv.db.GetCourseLastUpdated(ctx, courseBind.ID)
		if err != nil {
			// todo: error handle
		}
	}
	// compare time. if request time equal the origin data time that should not be update
	if courseBind.LastUpdated == lastUpdated {
		return false
	}
	return true
}
