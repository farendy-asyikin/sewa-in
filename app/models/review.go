package models

import "time"

type Review struct {
	ID         string    `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User       string    `gorm:"size:36;not null"`
	created_at time.Time `gorm:"not mull"`
	updated_at time.Time `gorm:"not mull"`
}
