package models

import "time"

type TransactionOngoing struct {
	ID        string    `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ItemId    string    `gorm:"size:100;not null"`
	UserId    string    `gorm:"size:100;not null"`
	Note      string    `gorm:"size:100;not null"`
	RentStart time.Time `gorm:"not null"`
	ReturnAt  time.Time `gorm:"not null"`
}
