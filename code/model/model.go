package model

type AutoModel struct {
	ID uint `gorm:"primaryKey; autoIncrement" json:"id"`
	Time
}
