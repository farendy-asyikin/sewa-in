package models

type Review struct {
	ID   string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User User
}
