package entity

import (
	"time"

	"gorm.io/gorm"
)

type AuthUser struct {
	UserId        string         `gorm:"column:userId"`
	Email           string         `gorm:"column:email"`
	IsEmailVerified bool           `gorm:"column:isEmailVerified"`
	PasswordHash    string         `gorm:"column:passwordHash"`
	Status          int            `gorm:"column:status"`
	CreatedAt       time.Time      `gorm:"column:createdAt;autoCreateTime:nano"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt;autoUpdateTime:nano"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deletedAt"`
}

func (AuthUser) TableName() string {
	return "authUsers"
}
