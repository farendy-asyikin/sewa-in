package models

import "time"

type RentChart struct {
	ID        string    `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ItemId    string    `gorm:"size:36;not null"`
	Duration  time.Time `gorm:"not mull"`
	Deadline  time.Time `gorm:"not mull"`
	UserId    string    `gorm:"size:36;not null"`
	Price     int16     `gorm:"not mull"`
	Subtotal  int16     `gorm:"not mull"`
	Note      string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"not mull"`
	UpdatedAt time.Time `gorm:"not mull"`
}
