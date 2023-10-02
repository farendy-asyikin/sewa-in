package models

type Stock struct {
	ID                string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ItemName          string `gorm:"size:100;not null"`
	ItemType          string `gorm:"size:100;not null"`
	ItemBrand         string `gorm:"size:100;not null"`
	ItemSpesification string `gorm:"size:100;not null"`
}
