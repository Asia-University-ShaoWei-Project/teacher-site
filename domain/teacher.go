package domain

type BindDomain struct {
	Domain string `uri:"domain" binding:"required"`
}

type Teachers struct {
	Domain    string `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	NameZH    string
	NameUS    string
	Office    string
	Call      string
	Education string
	AuthID    string
	Courses   []Courses `gorm:"foreignKey:TeacherID;references:Domain"`
	Infos     []Infos   `gorm:"foreignKey:TeacherID;references:Domain"`
	Time      Time      `gorm:"embedded"`
}
