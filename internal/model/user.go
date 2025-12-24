package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Email    string `gorm:"size:255;unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"size:20;default:user"`
	Avatar   string `gorm:"size:500"`
}
