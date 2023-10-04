package models

import (
	"time"
)

type RentChart struct {
	ID         string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	item_id    Stock
	duration   time.Time `gorm:"not mull"`
	deadline   time.Time `gorm:"not mull"`
	user_id    User
	price      int16     `gorm:"not mull"`
	subtotal   int16     `gorm:"not mull"`
	is_active  bool      `gorm:"not mull"`
	note       string    `gorm:"size:255;not null"`
	created_at time.Time `gorm:"not mull"`
	updated_at time.Time `gorm:"not mull"`
}
