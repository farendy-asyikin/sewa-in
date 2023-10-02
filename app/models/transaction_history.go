package models

type TransactionHistory struct {
	ID              string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	transaction_id  string `gorm:"size:100;not null"`
	review_by_user  string `gorm:"size:100;not null"`
	review_by_owner string `gorm:"size:100;not null"`
	is_success      bool   `gorm:"not null"`
}
