package domain

type AutoModel struct {
	Id uint `gorm:"primaryKey; autoIncrement"`
	Time
}
type File struct {
	Title string `json:"title"`
	Type  string `json:"type"`
	Url   string `json:"url"`
}
