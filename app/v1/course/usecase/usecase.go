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

// ===== CREATE =====
// todo
func (u *Usecase) Create(ctx context.Context, req *domain.CreateCourseRequest) (domain.CreateCourseResponse, error) {
	var res domain.CreateCourseResponse
	// if err := u.DbRepository; err != nil {
	// 	return res, err
	// }
	return res, nil
}

func (u *Usecase) CreateBulletin(ctx context.Context, req *domain.CreateCourseBulletinRequest) (domain.CreateCourseBulletinResponse, error) {
	var res domain.CreateCourseBulletinResponse
	if err := u.checkByDomainAndCourseId(ctx, req.TeacherDomain, req.CourseId); err != nil {
		return res, err
	}
	bulletin := domain.BulletinBoards{
		CourseId: req.CourseId,
		Content:  req.Content,
	}
	lastModified, err := u.DbRepository.CreateBulletin(ctx, &bulletin)
	if err != nil {
		u.log.Error(err)
		return res, err
	}
	// todo: update the data in redis
	res = domain.CreateCourseBulletinResponse{
		Id:           bulletin.AutoModel.Id,
		Date:         bulletin.AutoModel.CreatedAT.Format(domain.BulletinDateFormat),
		LastModified: lastModified,
	}
	return res, nil
}

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
		slide.SetFilename(req.Filename)
	}
	lastModified, err := u.DbRepository.CreateSlide(ctx, &slide)
	if err != nil {
		u.log.Error(err)
		return res, err
	}
	// todo: create data from the cache
	res = domain.CreateCourseSlideResponse{
		Id:           slide.AutoModel.Id,
		Filename:     slide.File.Name,
		LastModified: lastModified,
	}
	return res, nil

}

func (u *Usecase) CreateHomework(ctx context.Context, req *domain.CreateCourseHomeworkRequest) (domain.CreateCourseHomeworkResponse, error) {
	var res domain.CreateCourseHomeworkResponse
	if err := u.checkByDomainAndCourseId(ctx, req.TeacherDomain, req.CourseId); err != nil {
		u.log.Error(err)
		return res, err
	}
	homework := domain.Homeworks{
		CourseId: req.CourseId,
		Number:   req.Number,
		File: domain.File{
			Title: req.FileTitle,
		},
	}
	if req.Filename != "" {
		homework.SetFilename(req.Filename)
	}
	lastModified, err := u.DbRepository.CreateHomework(ctx, &homework)
	if err != nil {
		u.log.Error(err)
		return res, err
	}
	// todo: create data from the cache
	res = domain.CreateCourseHomeworkResponse{
		Id:           homework.AutoModel.Id,
		Filename:     homework.File.Name,
		LastModified: lastModified,
	}
	return res, nil

}

// ===== GET =====
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
		u.log.Error(err)
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

// ===== UPDATE =====

func (u *Usecase) UpdateBulletin(ctx context.Context, req *domain.UpdateCourseBulletinRequest) (domain.UpdateCourseBulletinResponse, error) {
	var res domain.UpdateCourseBulletinResponse
	if err := u.checkByDomainAndCourseId(ctx, req.TeacherDomain, req.CourseId); err != nil {
		return res, err
	}
	bulletin := domain.BulletinBoards{
		AutoModel: domain.AutoModel{Id: req.BulletinId},
		CourseId:  req.CourseId,
		Content:   req.Content,
	}
	lastModified, err := u.DbRepository.UpdateBulletinById(ctx, &bulletin)
	if err != nil {
		u.log.Error(err)
		return res, err
	}
	// todo: cache update
	res = domain.UpdateCourseBulletinResponse{
		LastModified: lastModified,
	}
	return res, nil
}

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
		slide.SetFilename(req.Filename)
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

func (u *Usecase) UpdateHomework(ctx context.Context, req *domain.UpdateCourseHomeworkRequest) (domain.UpdateCourseHomeworkResponse, error) {
	var res domain.UpdateCourseHomeworkResponse
	if err := u.checkByDomainAndCourseId(ctx, req.TeacherDomain, req.CourseId); err != nil {
		u.log.Error(err)
		return res, err
	}
	homework := domain.Homeworks{
		AutoModel: domain.AutoModel{Id: req.HomeworkId},
		CourseId:  req.CourseId,
		Number:    req.Number,
		File: domain.File{
			Title: req.FileTitle,
		},
	}
	if req.Filename != "" {
		homework.SetFilename(req.Filename)
	}
	lastModified, err := u.DbRepository.UpdateHomeworkById(ctx, &homework)
	if err != nil {
		u.log.Error(err)
		return res, err
	}
	// todo: cache update
	res = domain.UpdateCourseHomeworkResponse{
		Filename:     homework.File.Name,
		LastModified: lastModified,
	}
	return res, nil
}

// ===== DELETE =====
func (u *Usecase) DeleteBulletin(ctx context.Context, req *domain.DeleteCourseBulletinRequest) (domain.DeleteCourseBulletinResponse, error) {
	var res domain.DeleteCourseBulletinResponse
	if err := u.checkByDomainAndCourseId(ctx, req.TeacherDomain, req.CourseId); err != nil {
		return res, err
	}
	bulletin := domain.BulletinBoards{
		AutoModel: domain.AutoModel{Id: req.BulletinId},
		CourseId:  req.CourseId,
	}
	lastModified, err := u.DbRepository.DeleteBulletinById(ctx, &bulletin)
	if err != nil {
		u.log.Error(err)
		return res, err
	}
	res = domain.DeleteCourseBulletinResponse{
		LastModified: lastModified,
	}
	return res, nil
}

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
	path := fmt.Sprintf(u.conf.Server.SlidePathFormat, req.TeacherDomain, slide.File.Name)
	if slide.File.Name != "" {
		u.log.Info("file name is not empty")
		// todo: put the task to queue(remove again)
		if err := os.Remove(path); err != nil {
			u.log.Error("remove the file, the path:%s, error got:%v", path, err)
		}
	}

	res = domain.DeleteCourseSlideResponse{
		LastModified: lastModified,
	}
	return res, nil
}

func (u *Usecase) DeleteHomework(ctx context.Context, req *domain.DeleteCourseHomeworkRequest) (domain.DeleteCourseHomeworkResponse, error) {
	var res domain.DeleteCourseHomeworkResponse
	if err := u.checkByDomainAndCourseId(ctx, req.TeacherDomain, req.CourseId); err != nil {
		u.log.Error(err)
		return res, err
	}
	homework := domain.Homeworks{
		AutoModel: domain.AutoModel{Id: req.HomeworkId},
		CourseId:  req.CourseId,
	}
	lastModified, err := u.DbRepository.DeleteHomeworkById(ctx, &homework)
	if err != nil {
		u.log.Error(err)
		return res, err
	}
	path := fmt.Sprintf(u.conf.Server.SlidePathFormat, req.TeacherDomain, homework.File.Name)
	if homework.File.Name != "" {
		// todo: put the task to queue(remove again)
		if err := os.Remove(path); err != nil {
			u.log.Error("remove the file, the path:%s, error got:%v", path, err)
		}
	}

	res = domain.DeleteCourseHomeworkResponse{
		LastModified: lastModified,
	}
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
