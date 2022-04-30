package domain

import "mime/multipart"

type Slides struct {
	AutoModel AutoModel `gorm:"embedded"`
	CourseId  uint
	Chapter   string
	File      File `gorm:"embedded"`
}

func (s *Slides) SetFilename(name string) {
	s.File.Name = name
}

// request & response

type CreateSlideUriRequest struct {
	TeacherDomainRequest
	CourseId uint `uri:"courseId" binding:"required"`
}
type CreateCourseSlideRequest struct {
	TeacherDomain string
	CourseId      uint
	Chapter       string                `form:"chapter" binding:"required"`
	FileTitle     string                `form:"fileTitle" binding:"required"`
	File          *multipart.FileHeader `form:"file"`
	Filename      string
}

func (u *CreateCourseSlideRequest) SetupUri(uri *CreateSlideUriRequest) {
	u.TeacherDomain = uri.TeacherDomainRequest.TeacherDomain
	u.CourseId = uri.CourseId
}
func (u *CreateCourseSlideRequest) SetFilename(name string) {
	u.Filename = name
}

type CreateCourseSlideResponse struct {
	Id           uint   `json:"id"`
	Filename     string `json:"filename"`
	LastModified string `json:"lastModified"`
}
type CourseSlideUriRequest struct {
	TeacherDomainRequest
	CourseId uint `uri:"courseId" binding:"required"`
	SlideId  uint `uri:"slideId" binding:"required"`
}
type UpdateCourseSlideRequest struct {
	TeacherDomain string
	CourseId      uint
	SlideId       uint
	Chapter       string                `form:"chapter" binding:"required"`
	FileTitle     string                `form:"fileTitle" binding:"required"`
	File          *multipart.FileHeader `form:"file"`
	Filename      string
}

func (u *UpdateCourseSlideRequest) SetupUri(uri *CourseSlideUriRequest) {
	u.TeacherDomain = uri.TeacherDomainRequest.TeacherDomain
	u.CourseId = uri.CourseId
	u.SlideId = uri.SlideId
}
func (u *UpdateCourseSlideRequest) SetFilename(name string) {
	u.Filename = name
}

type UpdateCourseSlideResponse struct {
	Filename     string `json:"filename"`
	LastModified string `json:"lastModified"`
}

type DeleteCourseSlideRequest struct {
	CourseSlideUriRequest
}
type DeleteCourseSlideResponse struct {
	LastModified string `json:"lastModified"`
}
