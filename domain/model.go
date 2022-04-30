package domain

type AutoModel struct {
	Id uint `gorm:"primaryKey; autoIncrement"`
	Time
}
