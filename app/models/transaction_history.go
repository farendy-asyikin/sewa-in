package models

type TransactionHistory struct {
	ID            string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	TransactionId string `gorm:"size:100;not null"`
	IsSuccess     bool   `gorm:"not null"`
}
