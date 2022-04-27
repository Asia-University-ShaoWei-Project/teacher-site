package domain

type Teachers struct {
	Domain    string `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	NameZh    string
	NameUs    string
	Office    string
	Call      string
	Education string
	// todo renames the variable & add foreign key to school table
	Location string
	AuthId   string
	Courses  []Courses `gorm:"foreignKey:TeacherId;references:Domain"`
	Infos    []Infos   `gorm:"foreignKey:TeacherId;references:Domain"`
	Time     Time      `gorm:"embedded"`
}
type TeacherDomainRequest struct {
	TeacherDomain string `uri:"teacherDomain"  binding:"required"`
}
