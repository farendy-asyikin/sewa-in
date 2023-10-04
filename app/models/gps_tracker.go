package models

import (
	"time"
)

type GpsTracker struct {
	ID        string    `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Latitude  string    `gorm:"size:20;not null"`
	Longitude string    `gorm:"size:20;not null"`
	Altitide  string    `gorm:"size:20;not null"`
	Speed     float32   `gorm:"not null"`
	Satelites int16     `gorm:"not null"`
	Time      time.Time `gorm:"not null"`
}
