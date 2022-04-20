package domain

type AutoModel struct {
	ID uint `gorm:"primaryKey; autoIncrement"`
	Time
}
type File struct {
	Title string `json:"title"`
	Type  string `json:"type"`
	Url   string `json:"url"`
}
type TeacherDomainRequest struct {
	TeacherDomain string `json:"teacher_domain"  binding:"required"`
}
