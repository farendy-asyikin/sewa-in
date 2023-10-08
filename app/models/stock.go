package models

import "time"

type Stock struct {
	ID                string    `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ItemName          string    `gorm:"size:100;not null"`
	ItemType          string    `gorm:"size:100;not null"`
	ItemBrand         string    `gorm:"size:100;not null"`
	ItemSpesification string    `gorm:"size:100;not null"`
	IsActive          bool      `gorm:"not mull;default:true"`
	CreatedAt         time.Time `gorm:"not mull"`
	UpdatedAt         time.Time `gorm:"not mull"`
	DeletedAt         time.Time `gorm:"not null"`
}
