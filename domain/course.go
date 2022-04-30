package domain

import (
	"context"
)

type Courses struct {
	AutoModel     AutoModel `gorm:"embedded"`
	TeacherId     string
	NameZh        string `gorm:"not null"`
	NameUs        string
	BulletinBoard []BulletinBoards `gorm:"foreignKey:CourseId;references:Id"`
	Slide         []Slides         `gorm:"foreignKey:CourseId;references:Id"`
	Homework      []Homeworks      `gorm:"foreignKey:CourseId;references:Id"`
	LastModified  string
}
type File struct {
	Name  string
	Title string
}

// * Usecase, repository

type CourseUsecase interface {
	Create(ctx context.Context, req *CreateCourseRequest) (CreateCourseResponse, error)
	Get(ctx context.Context, req *GetCourseRequest) (GetCourseResponse, error)
	GetContent(ctx context.Context, req *GetCourseContentRequest) (GetCourseContentResponse, error)
	// Update()
	// Delete()

	CreateBulletin(ctx context.Context, req *CreateCourseBulletinRequest) (CreateCourseBulletinResponse, error)
	UpdateBulletin(ctx context.Context, req *UpdateCourseBulletinRequest) (UpdateCourseBulletinResponse, error)
	DeleteBulletin(ctx context.Context, req *DeleteCourseBulletinRequest) (DeleteCourseBulletinResponse, error)

	CreateSlide(ctx context.Context, req *CreateCourseSlideRequest) (CreateCourseSlideResponse, error)
	UpdateSlide(ctx context.Context, req *UpdateCourseSlideRequest) (UpdateCourseSlideResponse, error)
	DeleteSlide(ctx context.Context, req *DeleteCourseSlideRequest) (DeleteCourseSlideResponse, error)

	CreateHomework(ctx context.Context, req *CreateCourseHomeworkRequest) (CreateCourseHomeworkResponse, error)
	UpdateHomework(ctx context.Context, req *UpdateCourseHomeworkRequest) (UpdateCourseHomeworkResponse, error)
	DeleteHomework(ctx context.Context, req *DeleteCourseHomeworkRequest) (DeleteCourseHomeworkResponse, error)
}
type CourseDbRepository interface {
	CreateBulletin(ctx context.Context, bulletin *BulletinBoards) (string, error)
	CreateSlide(ctx context.Context, slide *Slides) (string, error)
	CreateHomework(ctx context.Context, homework *Homeworks) (string, error)

	GetByTeacherDomain(ctx context.Context, teacherDomain string) ([]CourseResponse, error)
	GetContentByCourseId(ctx context.Context, courseId uint) (GetCourseContentResponse, error)
	GetLastModifiedByCourseId(ctx context.Context, courseId uint) (string, error)

	// Update(ctx context.Context,  *) (, error)
	UpdateBulletinById(ctx context.Context, bulletin *BulletinBoards) (string, error)
	UpdateSlideById(ctx context.Context, slide *Slides) (string, error)
	UpdateHomeworkById(ctx context.Context, homework *Homeworks) (string, error)

	// Delete(ctx context.Context,  *) (, error)
	DeleteBulletinById(ctx context.Context, bulletin *BulletinBoards) (string, error)
	DeleteSlideById(ctx context.Context, slide *Slides) (string, error)
	DeleteHomeworkById(ctx context.Context, homework *Homeworks) (string, error)

	CheckByDomainAndCourseId(ctx context.Context, course *Courses) error
}

// todo: implement interface of the cache
type CourseCacheRepository interface {
}

//* request & response
type CreateCourseRequest struct {
	NameZh string `json:"nameZh"  binding:"required"`
	NameUs string `json:"nameUs"`
}
type CreateCourseResponse struct {
	Id           uint   `json:"id"`
	LastModified string `json:"lastModified"`
}

type GetCourseRequest struct {
	TeacherDomainRequest
}
type GetCourseResponse struct {
	Courses []CourseResponse `json:"courses"`
}
type CourseResponse struct {
	Id     uint   `json:"id"`
	NameZh string `json:"nameZh"`
	NameUs string `json:"nameUs"`
}

type GetCourseContentRequest struct {
	Id           uint   `uri:"courseId"`
	LastModified string `form:"lastModified"`
}

type GetCourseContentResponse struct {
	Id            uint                     `json:"id"`
	BulletinBoard []CourseBulletinResponse `json:"bulletins" gorm:"foreignKey:Id;references:Id"`
	Slide         []CourseSlideResponse    `json:"slides" gorm:"foreignKey:Id;references:Id"`
	Homework      []CourseHomeworkResponse `json:"homeworks" gorm:"foreignKey:Id;references:Id"`
	LastModified  string                   `json:"lastModified"`
}
type CourseBulletinResponse struct {
	Id      uint   `json:"id"`
	Date    string `json:"date"`
	Content string `json:"content"`
}
type CourseSlideResponse struct {
	Id      uint   `json:"id"`
	Chapter string `json:"chapter"`
	Title   string `json:"fileTitle"`
	Name    string `json:"filename"`
}
type CourseHomeworkResponse struct {
	Id     uint   `json:"id"`
	Number string `json:"number"`
	Title  string `json:"fileTitle"`
	Name   string `json:"filename"`
}
