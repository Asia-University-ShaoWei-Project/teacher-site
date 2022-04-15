package model

const InfoDateFormat = "2006-01-02"

type BindInfoID struct {
	ID uint `uri:"id"`
}
type BindInfo struct {
	ID         uint   `uri:"id"`
	CreateDate string `json:"create_date" binding:"required"`
	Info       string `json:"info" binding:"required"`
}
type Informations struct {
	AutoModel
	TeacherID   string `json:"teacher_id"`
	CreatedDate string `json:"created_date"`
	Info        string `json:"info"`
}
