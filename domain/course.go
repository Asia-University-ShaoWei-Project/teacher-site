package domain

type Courses struct {
	AutoModel     AutoModel        `gorm:"embedded"`
	TeacherID     string           `json:"teacher_id"`
	NameZH        string           `gorm:"not null"`
	NameUS        string           `gorm:"not null"`
	BulletinBoard []BulletinBoards `gorm:"foreignKey:CourseID;references:ID"`
	Slide         []Slides         `gorm:"foreignKey:CourseID;references:ID"`
	Homework      []Homeworks      `gorm:"foreignKey:CourseID;references:ID"`
	LastModified  string
}

type BulletinBoards struct {
	AutoModel AutoModel `gorm:"embedded"`
	CourseID  uint      `json:"course_id"`
	Content   string    `json:"content"`
}

// todo: add link field
type Slides struct {
	AutoModel AutoModel `gorm:"embedded"`
	CourseID  uint      `json:"course_id"`
	Chapter   string    `json:"chapter"`
	File      File      `gorm:"embedded" json:"file"`
}

// todo: add link field
type Homeworks struct {
	AutoModel AutoModel `gorm:"embedded"`
	CourseID  uint      `json:"course_id"`
	Number    string    `json:"number"`
	File      File      `gorm:"embedded" json:"file"`
}

// request & response
type ReqCreateCourse struct {
	NameZH string `json:"name_zh"`
	NameUS string `json:"name_us"`
}
type ReqGetCourse struct {
	ID          uint   `uri:"course_id" binding:"required"`
	LastUpdated string `uri:"last_updated" binding:"required"`
}
type ResCreateCourse struct {
	ID uint `json:"id"`
}
