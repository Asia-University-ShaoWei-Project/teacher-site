package model

type Courses struct {
	AutoModel
	TeacherID     string          `json:"teacher_id"`
	NameZH        string          `gorm:"not null" json:"name_zh"`
	NameUS        string          `gorm:"not null" json:"name_us"`
	BulletinBoard []BulletinBoard `gorm:"foreignKey:CourseID;references:ID" json:"bulletin_board"`
	Slide         []Slides        `gorm:"foreignKey:CourseID;references:ID" json:"slide"`
	Homework      []Homeworks     `gorm:"foreignKey:CourseID;references:ID" json:"homework"`
	LastUpdated   string          `json:"last_updated"`
}
type BindCourse struct {
	ID          uint   `uri:"course_id" binding:"required"`
	LastUpdated string `uri:"last_updated" binding:"required"`
}
type BulletinBoard struct {
	AutoModel
	CourseID    uint   `json:"course_id"`
	CreatedDate string `json:"created_date"`
	Info        string `json:"info"`
}

// todo: add link field
type Slides struct {
	AutoModel
	CourseID uint   `json:"course_id"`
	Chapter  string `json:"chapter"`
	File     File   `gorm:"embedded" json:"file"`
}

// todo: add link field
type Homeworks struct {
	AutoModel
	CourseID uint   `json:"course_id"`
	Number   string `json:"number"`
	File     File   `gorm:"embedded" json:"file"`
}

type File struct {
	Title string `json:"title"`
	Type  string `json:"type"`
}
