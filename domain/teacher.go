package domain

type Teachers struct {
	Domain    string `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	NameZH    string
	NameUS    string
	Office    string
	Call      string
	Education string
	// todo rename & add foreign key
	Location string
	AuthID   string
	Courses  []Courses `gorm:"foreignKey:TeacherID;references:Domain"`
	Infos    []Infos   `gorm:"foreignKey:TeacherID;references:Domain"`
	Time     Time      `gorm:"embedded"`
}
type TeacherDomainRequest struct {
	TeacherDomain string `uri:"teacherDomain"  binding:"required"`
}
