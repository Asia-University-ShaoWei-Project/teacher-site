package domain

import "mime/multipart"

type Homeworks struct {
	AutoModel AutoModel `gorm:"embedded"`
	CourseId  uint
	// todo: number(str) -> int
	Number string
	File   File `gorm:"embedded"`
}

func (s *Homeworks) SetFilename(name string) {
	s.File.Name = name
}

// request & response

type CreateHomeworkUriRequest struct {
	TeacherDomainRequest
	CourseId uint `uri:"courseId" binding:"required"`
}
type CreateCourseHomeworkRequest struct {
	TeacherDomain string
	CourseId      uint
	Number        string                `form:"number" binding:"required"`
	FileTitle     string                `form:"fileTitle" binding:"required"`
	File          *multipart.FileHeader `form:"file"`
	Filename      string
}

func (c *CreateCourseHomeworkRequest) SetupUri(uri *CreateHomeworkUriRequest) {
	c.TeacherDomain = uri.TeacherDomainRequest.TeacherDomain
	c.CourseId = uri.CourseId
}
func (s *CreateCourseHomeworkRequest) SetFilename(name string) {
	s.Filename = name
}

type CreateCourseHomeworkResponse struct {
	Id           uint   `json:"id"`
	Filename     string `json:"filename"`
	LastModified string `json:"lastModified"`
}
type CourseHomeworkUriRequest struct {
	TeacherDomainRequest
	CourseId   uint `uri:"courseId" binding:"required"`
	HomeworkId uint `uri:"homeworkId" binding:"required"`
}

type UpdateCourseHomeworkRequest struct {
	TeacherDomain string
	CourseId      uint
	HomeworkId    uint
	Number        string                `form:"number" binding:"required"`
	FileTitle     string                `form:"fileTitle" binding:"required"`
	File          *multipart.FileHeader `form:"file"`
	Filename      string
}

func (c *UpdateCourseHomeworkRequest) SetupUri(uri *CourseHomeworkUriRequest) {
	c.TeacherDomain = uri.TeacherDomain
	c.CourseId = uri.CourseId
	c.HomeworkId = uri.HomeworkId
}
func (u *UpdateCourseHomeworkRequest) SetFilename(name string) {
	u.Filename = name
}

type UpdateCourseHomeworkResponse struct {
	Filename     string `json:"filename"`
	LastModified string `json:"lastModified"`
}

type DeleteCourseHomeworkRequest struct {
	CourseHomeworkUriRequest
}
type DeleteCourseHomeworkResponse struct {
	LastModified string `json:"lastModified"`
}
