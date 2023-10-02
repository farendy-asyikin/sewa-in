package models

import "time"

type TransactionOngoing struct {
	ID         string    `gorm:"size:36;not null;uniqueIndex;primary_key"`
	item_id    string    `gorm:"size:100;not null"`
	user_id    string    `gorm:"size:100;not null"`
	note       string    `gorm:"size:100;not null"`
	rent_start time.Time `gorm:"not null"`
	return_at  time.Time `gorm:"not null"`
}
