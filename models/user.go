package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"uniqueIndex;not null"`
	Password    string `gorm:"not null"`
	IsActivated bool   `gorm:"default:false"`
	Roles       string `gorm:"not null;default:'user'"`
}
