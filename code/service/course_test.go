package service

import (
	"teacher-site/mock"
	"teacher-site/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCourseNeedUpdate(t *testing.T) {
	var courseBind model.BindCourse
	realLastUpdated, _ := srv.db.GetCourseLastUpdated(ctx, mock.CourseID)
	tC := []struct {
		desc        string
		courseID    uint
		lastUpdated string
		result      NeedUpdate
	}{
		{
			desc:        "equal last_update",
			courseID:    mock.CourseID,
			lastUpdated: realLastUpdated,
			result:      false,
		},
		{
			desc:        "before the last_update",
			courseID:    mock.CourseID,
			lastUpdated: "0",
			result:      true,
		},
		{
			desc:        "fatal the course id",
			courseID:    mock.CourseID,
			lastUpdated: "0",
			result:      true,
		},
	}
	for _, v := range tC {
		t.Run(v.desc, func(t *testing.T) {
			courseBind.ID = v.courseID
			courseBind.LastUpdated = v.lastUpdated
			needUpdate := srv.courseNeedUpdate(ctx, &courseBind)
			assert.Equal(t, v.result, needUpdate)
		})
	}
}
