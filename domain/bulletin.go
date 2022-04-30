package domain

type BulletinBoards struct {
	AutoModel AutoModel `gorm:"embedded"`
	CourseId  uint
	Content   string
}

// request & response

type CreateCourseBulletinRequest struct {
	TeacherDomainRequest
	CourseId uint   `uri:"courseId"`
	Content  string `json:"content"`
}
type CreateCourseBulletinResponse struct {
	Id           uint   `json:"id"`
	Date         string `json:"date"`
	LastModified string `json:"lastModified"`
}
type CourseBulletinUriRequest struct {
	TeacherDomainRequest
	CourseId   uint `uri:"courseId" binding:"required"`
	BulletinId uint `uri:"bulletinId" binding:"required"`
}
type UpdateCourseBulletinRequest struct {
	TeacherDomain string
	CourseId      uint
	BulletinId    uint
	Content       string `json:"content" binding:"required"`
}

func (u *UpdateCourseBulletinRequest) SetupUri(uri CourseBulletinUriRequest) {
	u.TeacherDomain = uri.TeacherDomain
	u.CourseId = uri.CourseId
	u.BulletinId = uri.BulletinId
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
