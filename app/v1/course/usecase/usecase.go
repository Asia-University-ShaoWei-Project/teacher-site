package usecase

import (
	"context"
	"fmt"
	"os"
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
		log.Error(err)
		return res, err
	}
	res = domain.GetCourseResponse{
		Courses: courses,
	}
	return res, nil
}

func (u *Usecase) GetContent(ctx context.Context, req *domain.GetCourseContentRequest) (domain.GetCourseContentResponse, error) {
	var res domain.GetCourseContentResponse

	lastModified, err := u.DbRepository.GetLastModifiedByCourseId(ctx, req.Id)
	if err != nil {
		return res, err
	}
	// Unnecessary to get new data if request last modified value is equal the last modified of repository value
	if req.LastModified == lastModified {
		return res, message.ErrUnnecessaryUpdate
	}

	res, err = u.DbRepository.GetContentByCourseId(ctx, req.Id)
	if err != nil {
		log.Error(err)
		return res, err
	}
	return res, nil
}

// todo
// Update()
// Delete()

// todo
func (u *Usecase) CreateBulletin(ctx context.Context, req *domain.CreateCourseBulletinRequest) (domain.CreateCourseBulletinResponse, error) {
	var res domain.CreateCourseBulletinResponse
	if err := u.checkByDomainAndCourseId(ctx, req.TeacherDomain, req.CourseId); err != nil {
		return res, err
	}
	bulletin := domain.BulletinBoards{
		CourseId: req.CourseId,
		Content:  req.Content,
	}
	err := u.DbRepository.CreateBulletin(ctx, &bulletin)
	if err != nil {
		return res, err
	}
	lastModified, err := u.DbRepository.GetLastModifiedByCourseId(ctx, req.CourseId)
	if err != nil {
		// todo: get last modified error
		u.log.Error(err)
	}
	// todo: update the data in redis
	res = domain.CreateCourseBulletinResponse{
		Id:           bulletin.AutoModel.Id,
		Date:         bulletin.AutoModel.CreatedAT.Format(domain.BulletinDateFormat),
		LastModified: lastModified,
	}
	return res, nil
}

func (u *Usecase) UpdateBulletin(ctx context.Context, req *domain.UpdateCourseBulletinRequest) (domain.UpdateCourseBulletinResponse, error) {
	var res domain.UpdateCourseBulletinResponse
	if err := u.checkByDomainAndCourseId(ctx, req.TeacherDomain, req.CourseId); err != nil {
		return res, err
	}
	bulletin := domain.BulletinBoards{
		CourseId: req.CourseId,
		Content:  req.Content,
	}

	if err := u.DbRepository.UpdateBulletinById(ctx, &bulletin); err != nil {
		return res, err
	}
	lastModified, err := u.DbRepository.GetLastModifiedByCourseId(ctx, req.CourseId)
	if err != nil {
		// todo: error at get last modified
		u.log.Error(err)
	}
	// todo: cache update
	res = domain.UpdateCourseBulletinResponse{
		LastModified: lastModified,
	}
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
	if err := u.checkByDomainAndCourseId(ctx, req.TeacherDomain, req.CourseId); err != nil {
		u.log.Error(err)
		return res, err
	}
	slide := domain.Slides{
		CourseId: req.CourseId,
		Chapter:  req.Chapter,
		File: domain.File{
			Title: req.FileTitle,
		},
	}
	if req.Filename != "" {
		slide.SetFileName(req.Filename)
	}

	if err := u.DbRepository.CreateSlide(ctx, &slide); err != nil {
		u.log.Error(err)
		return res, err
	}
	lastModified, err := u.DbRepository.GetLastModifiedByCourseId(ctx, req.CourseId)
	if err != nil {
		// todo: error at get last modified
		u.log.Error(err)
	}
	// todo: create data from the cache
	res = domain.CreateCourseSlideResponse{
		Id:           slide.AutoModel.Id,
		Filename:     slide.File.Name,
		LastModified: lastModified,
	}
	return res, nil

}

// todo
func (u *Usecase) UpdateSlide(ctx context.Context, req *domain.UpdateCourseSlideRequest) (domain.UpdateCourseSlideResponse, error) {
	var res domain.UpdateCourseSlideResponse
	if err := u.checkByDomainAndCourseId(ctx, req.TeacherDomain, req.CourseId); err != nil {
		u.log.Error(err)
		return res, err
	}
	slide := domain.Slides{
		AutoModel: domain.AutoModel{Id: req.SlideId},
		CourseId:  req.CourseId,
		Chapter:   req.Chapter,
		File: domain.File{
			Title: req.FileTitle,
		},
	}
	if req.Filename != "" {
		slide.SetFileName(req.Filename)
	}
	lastModified, err := u.DbRepository.UpdateSlideById(ctx, &slide)
	if err != nil {
		u.log.Error(err)
		return res, err
	}
	// todo: cache update
	res = domain.UpdateCourseSlideResponse{
		Filename:     slide.File.Name,
		LastModified: lastModified,
	}
	return res, nil
}

// todo
func (u *Usecase) DeleteSlide(ctx context.Context, req *domain.DeleteCourseSlideRequest) (domain.DeleteCourseSlideResponse, error) {
	var res domain.DeleteCourseSlideResponse
	if err := u.checkByDomainAndCourseId(ctx, req.TeacherDomain, req.CourseId); err != nil {
		u.log.Error(err)
		return res, err
	}
	slide := domain.Slides{
		AutoModel: domain.AutoModel{Id: req.SlideId},
		CourseId:  req.CourseId,
	}
	lastModified, err := u.DbRepository.DeleteSlideById(ctx, &slide)
	if err != nil {
		u.log.Error(err)
		return res, err
	}
	// todo: delete the file
	fmt.Println("slide filename:", slide.File.Name)
	path := fmt.Sprintf(u.conf.Server.SlidePathFormat, req.TeacherDomain, slide.File.Name)
	if slide.File.Name != "" {
		u.log.Info("file name is not empty")
		if err := os.Remove(path); err != nil {
			u.log.Info("remove the file, the path:", path)
			// todo: put the task to queue(remove again)
			u.log.Error(err)
		}
	}

	res = domain.DeleteCourseSlideResponse{
		LastModified: lastModified,
	}
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
func (u *Usecase) checkByDomainAndCourseId(ctx context.Context, teacherDomain string, courseId uint) error {
	var err error
	course := domain.Courses{
		TeacherId: teacherDomain,
		AutoModel: domain.AutoModel{Id: courseId},
	}
	err = u.DbRepository.CheckByDomainAndCourseId(ctx, &course)
	return err
}
