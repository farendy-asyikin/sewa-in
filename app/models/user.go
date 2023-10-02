package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	FirstName    string `gorm:"size:100;not null"`
	LastName     string `gorm:"size:100;not null"`
	IDCardNumber string `gorm:"size:100;not null"`
	Phone        string `gorm:"size:100;not null;uniqueIndex"`
	Email        string `gorm:"size:100;not null;uniqueIndex"`
	Gender       string `gorm:"size:100;not null"`
	IsActive     bool
	CreatedAt    time.Time
	UpdateAt     time.Time
	DeleteAt     gorm.DeletedAt
}
