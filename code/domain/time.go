package domain

import (
	"time"

	"gorm.io/gorm"
)

type Time struct {
	CreatedAT time.Time `gorm:"autoCreateTime"`
	UpdatedAT time.Time `gorm:"autoUpdateTime:milli"`
	DeletedAT gorm.DeletedAt
}
