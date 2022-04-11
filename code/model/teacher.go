package model

type BindDomain struct {
	Domain string `uri:"domain" binding:"required"`
}

type Teachers struct {
	Domain       string `gorm:"primaryKey"`
	Email        string `gorm:"unique" json:"email"`
	NameZH       string `json:"name_zh"`
	NameUS       string `json:"name_us"`
	Office       string `json:"office"`
	Call         string `json:"call"`
	Education    string `json:"education"`
	AuthID       string
	Courses      []Courses      `gorm:"foreignKey:TeacherID;references:Domain" json:"courses"`
	Informations []Informations `gorm:"foreignKey:TeacherID;references:Domain" json:"informations"`
	Time
}
