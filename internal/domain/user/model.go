package user

import (
	"time"
)

type User struct {
	ID        uint    `gorm:"primaryKey"`
	Username  string  `gorm:"uniqueIndex;not null"`
	Email     string  `gorm:"uniqueIndex;not null"`
	Password  string  `gorm:"not null"`
	Bio       *string `gorm:"type:text"`
	Image     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
