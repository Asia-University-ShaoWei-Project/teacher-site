package usecase

import (
	"context"
	"teacher-site/config"
	"teacher-site/domain"
	"teacher-site/pkg/message"

	log "github.com/sirupsen/logrus"
)

type Usecase struct {
	DbRepository    domain.CourseDbRepository
	CacheRepository domain.CourseCacheRepository
	conf            *config.Config
	log             *log.Logger
}

func NewUsecase(dbRepo domain.CourseDbRepository, cacheRepo domain.CourseCacheRepository, conf *config.Config, logger *log.Logger) domain.CourseUsecase {
	return &Usecase{
		DbRepository:    dbRepo,
		CacheRepository: cacheRepo,
		conf:            conf,
		log:             logger,
	}
}

// todo
func (u *Usecase) Create(ctx context.Context, req *domain.CreateCourseRequest) (domain.CreateCourseResponse, error) {
	var res domain.CreateCourseResponse
	// if err := u.DbRepository; err != nil {
	// 	return res, err
	// }
	return res, nil
}

func (u *Usecase) Get(ctx context.Context, req *domain.GetCourseRequest) (domain.GetCourseResponse, error) {
	var res domain.GetCourseResponse
	courses, err := u.DbRepository.GetByTeacherDomain(ctx, req.TeacherDomain)
	if err != nil {
		return res, err
	}
	res = domain.GetCourseResponse{
		Courses: courses,
	}
	return res, nil
}

func (u *Usecase) GetContent(ctx context.Context, req *domain.GetCourseContentRequest) (domain.GetCourseContentResponse, error) {
	var res domain.GetCourseContentResponse

	course, err := u.DbRepository.GetLastModifiedByCourseId(ctx, req.Id)
	if err != nil {
		return res, err
	}
	// Unnecessary to get new data if request last modified value is equal the last modified of repository value
	if req.LastModified == course.LastModified {
		return res, message.ErrUnnecessaryUpdate
	}

	res, err = u.DbRepository.GetContentByCourseId(ctx, req.Id)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Update()
// Delete()

// todo
func (u *Usecase) CreateBulletin(ctx context.Context, req *domain.CreateCourseBulletinRequest) (domain.CreateCourseBulletinResponse, error) {
	var res domain.CreateCourseBulletinResponse
	// ,err := u.DbRepository
	// if err != nil {
	// 	return res, err
	// }
	// res = domain.
	return res, nil
}

// todo
func (u *Usecase) UpdateBulletin(ctx context.Context, req *domain.UpdateCourseBulletinRequest) (domain.UpdateCourseBulletinResponse, error) {
	var res domain.UpdateCourseBulletinResponse
	// ,err := u.DbRepository
	// if err != nil {
	// 	return res, err
	// }
	// res = domain.
	return res, nil
}

// todo
func (u *Usecase) DeleteBulletin(ctx context.Context, req *domain.DeleteCourseBulletinRequest) (domain.DeleteCourseBulletinResponse, error) {
	var res domain.DeleteCourseBulletinResponse
	// ,err := u.DbRepository
	// if err != nil {
	// 	return res, err
	// }
	// res = domain.
	return res, nil
}

// todo
func (u *Usecase) CreateSlide(ctx context.Context, req *domain.CreateCourseSlideRequest) (domain.CreateCourseSlideResponse, error) {
	var res domain.CreateCourseSlideResponse
	// ,err := u.DbRepository
	// if err != nil {
	// 	return res, err
	// }
	// res = domain.
	return res, nil
}

// todo
func (u *Usecase) UpdateSlide(ctx context.Context, req *domain.UpdateCourseSlideRequest) (domain.UpdateCourseSlideResponse, error) {
	var res domain.UpdateCourseSlideResponse
	// ,err := u.DbRepository
	// if err != nil {
	// 	return res, err
	// }
	// res = domain.
	return res, nil
}

// todo
func (u *Usecase) DeleteSlide(ctx context.Context, req *domain.DeleteCourseSlideRequest) (domain.DeleteCourseSlideResponse, error) {
	var res domain.DeleteCourseSlideResponse
	// ,err := u.DbRepository
	// if err != nil {
	// 	return res, err
	// }
	// res = domain.
	return res, nil
}

// todo
func (u *Usecase) CreateHomework(ctx context.Context, req *domain.CreateCourseHomeworkRequest) (domain.CreateCourseHomeworkResponse, error) {
	var res domain.CreateCourseHomeworkResponse
	// ,err := u.DbRepository
	// if err != nil {
	// 	return res, err
	// }
	// res = domain.
	return res, nil
}

// todo
func (u *Usecase) UpdateHomework(ctx context.Context, req *domain.UpdateCourseHomeworkRequest) (domain.UpdateCourseHomeworkResponse, error) {
	var res domain.UpdateCourseHomeworkResponse
	// ,err := u.DbRepository
	// if err != nil {
	// 	return res, err
	// }
	// res = domain.
	return res, nil
}

// todo
func (u *Usecase) DeleteHomework(ctx context.Context, req *domain.DeleteCourseHomeworkRequest) (domain.DeleteCourseHomeworkResponse, error) {
	var res domain.DeleteCourseHomeworkResponse
	// ,err := u.DbRepository
	// if err != nil {
	// 	return res, err
	// }
	// res = domain.
	return res, nil
}
