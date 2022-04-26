package domain

import (
	"context"
	"mime/multipart"
)

type Courses struct {
	AutoModel     AutoModel `gorm:"embedded"`
	TeacherID     string
	NameZH        string           `gorm:"not null"`
	NameUS        string           `gorm:"not null"`
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
	// Create(ctx context.Context, req *CreateInfoBulletinRequest) (InfoBulletinBoards, error)
	GetByTeacherDomain(ctx context.Context, teacherDomain string) ([]CourseResponse, error)
	// GetBulletinsByInfoId(ctx context.Context, id uint) ([]InfoBulletinResponse, error)
	// GetLastModified(ctx context.Context, id uint) (string, error)
	// Update(ctx context.Context, req *UpdateInfoBulletinRequest) (Infos, error)
	// Delete(ctx context.Context, req *DeleteInfoBulletinRequest) (Infos, error)
}
type CourseCacheRepository interface {
	// Get(ctx context.Context, req *GetInfoBulletinRequest) (string, error)
	// GetLastModified(ctx context.Context, teacherDomain string) (string, error)
	// Update(ctx context.Context, req *UpdateInfoBulletinRequest) error
	// UpdateInfoLastModified(ctx context.Context, req *UpdateInfoBulletinRequest, lastModified string) error
}

//* request & response
// course
type CreateCourseRequest struct {
	NameZH string `json:"nameZh"  binding:"required"`
	NameUS string `json:"nameUs"`
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
	NameUS string `json:"nameUs"`
}
type GetCourseContentRequest struct {
	AutoModel   uint   `uri:"courseId"`
	LastUpdated string `form:"lastModified"`
}
type GetCourseContentResponse struct {
	Bulletin    []BulletinBoards `json:"bulletins"`
	Slides      []Slides         `json:"slides"`
	Homeworks   []Homeworks      `json:"homeworks"`
	LastUpdated string           `json:"lastModified"`
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
