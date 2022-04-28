package domain

import (
	"context"
	"mime/multipart"
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

type BulletinBoards struct {
	AutoModel AutoModel `gorm:"embedded"`
	CourseId  uint
	Content   string
}

// todo: add link field
type Slides struct {
	AutoModel AutoModel `gorm:"embedded"`
	CourseId  uint
	Chapter   string
	File      File `gorm:"embedded"`
}

// todo: add link field
type Homeworks struct {
	AutoModel AutoModel `gorm:"embedded"`
	CourseId  uint
	// todo: number(str) -> int
	Number string
	File   File `gorm:"embedded"`
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
	GetByTeacherDomain(ctx context.Context, teacherDomain string) ([]CourseResponse, error)
	GetContentByCourseId(ctx context.Context, courseId uint) (GetCourseContentResponse, error)
	GetLastModifiedByCourseId(ctx context.Context, courseId uint) (Courses, error)
}
type CourseCacheRepository interface {
}

//* request & response
// course
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
	Type    string `json:"fileType"`
	Url     string `json:"fileUrl"`
}
type CourseHomeworkResponse struct {
	Id     uint   `json:"id"`
	Number string `json:"number"`
	Title  string `json:"fileTitle"`
	Type   string `json:"fileType"`
	Url    string `json:"fileUrl"`
}

// type UpdateCourseRequest struct{}
// type UpdateCourseResponse struct{}

// type DeleteCourseRequest struct{}
// type DeleteCourseResponse struct{}

// bulletin
type CreateCourseBulletinRequest struct {
	TeacherDomainRequest
	CourseId uint   `uri:"courseId"`
	Content  string `json:"content"`
}
type CreateCourseBulletinResponse struct {
	Id           uint   `json:"bulletinId"`
	Date         string `json:"date"`
	LastModified string `json:"lastModified"`
}
type UpdateCourseBulletinRequest struct {
	TeacherDomainRequest
	CourseId   uint   `uri:"courseId"`
	BulletinId uint   `uri:"bulletinId"`
	Content    string `uri:"content"`
}
type UpdateCourseBulletinResponse struct {
	LastModified string `json:"lastModified"`
}

type DeleteCourseBulletinRequest struct {
	TeacherDomainRequest
	CourseId   uint `uri:"courseId" binding:"required"`
	BulletinId uint `uri:"bulletinId" binding:"required"`
}
type DeleteCourseBulletinResponse struct {
	LastModified string `json:"lastModified"`
}

// slide
type CreateCourseSlideRequest struct {
	TeacherDomainRequest
	CourseId  uint                  `uri:"courseId"`
	Chapter   string                `json:"chapter"`
	FileTitle string                `json:"fileTitle"`
	File      *multipart.FileHeader `json:"file"`
}
type CreateCourseSlideResponse struct {
	Id           uint   `json:"id"`
	LastModified string `json:"lastModified"`
}

type UpdateCourseSlideRequest struct {
	TeacherDomainRequest
	CourseId  uint                  `uri:"courseId"`
	SlideId   uint                  `uri:"slideId"`
	Chapter   string                `json:"chapter"`
	FileTitle string                `json:"fileTitle"`
	File      *multipart.FileHeader `json:"file"`
}
type UpdateCourseSlideResponse struct {
	LastModified string `json:"lastModified"`
}

type DeleteCourseSlideRequest struct {
	TeacherDomainRequest
	CourseId uint `uri:"courseId" binding:"required"`
	SlideId  uint `uri:"slideId" binding:"required"`
}
type DeleteCourseSlideResponse struct {
	LastModified string `json:"lastModified"`
}

// homework
type CreateCourseHomeworkRequest struct {
	TeacherDomainRequest
	CourseId  uint                  `uri:"courseId"`
	Number    string                `json:"number"`
	FileTitle string                `json:"fileTitle"`
	File      *multipart.FileHeader `json:"file"`
}
type CreateCourseHomeworkResponse struct {
	Id           uint   `json:"id"`
	LastModified string `json:"lastModified"`
}

type UpdateCourseHomeworkRequest struct {
	TeacherDomainRequest
	CourseId   uint                  `uri:"courseId"`
	HomeworkId uint                  `uri:"homeworkId"`
	Number     string                `json:"number"`
	FileTitle  string                `json:"fileTitle"`
	File       *multipart.FileHeader `json:"file"`
}
type UpdateCourseHomeworkResponse struct {
	LastModified string `json:"lastModified"`
}

type DeleteCourseHomeworkRequest struct {
	TeacherDomainRequest
	CourseId   uint `uri:"courseId" binding:"required"`
	HomeworkId uint `uri:"homeworkId" binding:"required"`
}
type DeleteCourseHomeworkResponse struct {
	LastModified string `json:"lastModified"`
}
