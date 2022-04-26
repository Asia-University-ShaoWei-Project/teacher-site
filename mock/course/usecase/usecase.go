package usecase

import (
	"context"
	"teacher-site/domain"
)

type Usecase struct{}

func NewUsecase() domain.CourseUsecase {
	return &Usecase{}
}

// course
func (u *Usecase) Create(ctx context.Context, req *domain.CreateCourseRequest) (domain.CreateCourseResponse, error) {
	return domain.CreateCourseResponse{}, nil
}
func (u *Usecase) Get(ctx context.Context, req *domain.GetCourseRequest) (domain.GetCourseResponse, error) {
	return domain.GetCourseResponse{}, nil
}
func (u *Usecase) GetContent(ctx context.Context, req *domain.GetCourseContentRequest) (domain.GetCourseContentResponse, error) {
	return domain.GetCourseContentResponse{}, nil
}

// bulletin
func (u *Usecase) CreateBulletin(ctx context.Context, req *domain.CreateCourseBulletinRequest) (domain.CreateCourseBulletinResponse, error) {
	return domain.CreateCourseBulletinResponse{}, nil
}
func (u *Usecase) UpdateBulletin(ctx context.Context, req *domain.UpdateCourseBulletinRequest) (domain.UpdateCourseBulletinResponse, error) {
	return domain.UpdateCourseBulletinResponse{}, nil
}
func (u *Usecase) DeleteBulletin(ctx context.Context, req *domain.DeleteCourseBulletinRequest) (domain.DeleteCourseBulletinResponse, error) {
	return domain.DeleteCourseBulletinResponse{}, nil
}

// slide
func (u *Usecase) CreateSlide(ctx context.Context, req *domain.CreateCourseSlideRequest) (domain.CreateCourseSlideResponse, error) {
	return domain.CreateCourseSlideResponse{}, nil
}
func (u *Usecase) UpdateSlide(ctx context.Context, req *domain.UpdateCourseSlideRequest) (domain.UpdateCourseSlideResponse, error) {
	return domain.UpdateCourseSlideResponse{}, nil
}
func (u *Usecase) DeleteSlide(ctx context.Context, req *domain.DeleteCourseSlideRequest) (domain.DeleteCourseSlideResponse, error) {
	return domain.DeleteCourseSlideResponse{}, nil
}

// homework
func (u *Usecase) CreateHomework(ctx context.Context, req *domain.CreateCourseHomeworkRequest) (domain.CreateCourseHomeworkResponse, error) {
	return domain.CreateCourseHomeworkResponse{}, nil
}
func (u *Usecase) UpdateHomework(ctx context.Context, req *domain.UpdateCourseHomeworkRequest) (domain.UpdateCourseHomeworkResponse, error) {
	return domain.UpdateCourseHomeworkResponse{}, nil
}
func (u *Usecase) DeleteHomework(ctx context.Context, req *domain.DeleteCourseHomeworkRequest) (domain.DeleteCourseHomeworkResponse, error) {
	return domain.DeleteCourseHomeworkResponse{}, nil
}
