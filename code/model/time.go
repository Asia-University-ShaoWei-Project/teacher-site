package model

import (
	"time"

	"gorm.io/gorm"
)

type Time struct {
	CreatedAT time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAT time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAT gorm.DeletedAt `json:"deleted_at"`
}
