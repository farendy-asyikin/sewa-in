package fakers

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sewa-in/app/models"
	"time"
)

func StockFaker(db *gorm.DB) *models.Stock {
	return &models.Stock{
		ID:                uuid.New().String(),
		ItemName:          "Supra",
		ItemType:          "Motor Cycle",
		ItemBrand:         "Honda",
		ItemSpesification: "98cc, SOHC 2 Valve",
		IsActive:          true,
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		DeletedAt:         time.Time{},
	}
}
