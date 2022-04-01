package model

type Auths struct {
	UserID       string `gorm:"primaryKey" json:"user_id"`
	UserPassword string `json:"user_password"`
	Token        string
	Teacher      Teachers `gorm:"foreignKey:AuthID;references:UserID"`
	Time
}
type BindAuth struct {
	UserID       string `json:"user_id" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
}
type BindRegister struct {
	UserID       string `json:"user_id" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
	Domain       string `json:"domain" binding:"required"`
	Email        string `json:"email" binding:"required"`
	NameZH       string `json:"name_zh" binding:"required"`
}
