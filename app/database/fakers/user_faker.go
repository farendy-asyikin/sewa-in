package fakers

import (
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sewa-in/app/models"
	"time"
)

func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		ID:           uuid.New().String(),
		FirstName:    faker.FirstName(),
		LastName:     faker.LastName(),
		IDCardNumber: "01234554321",
		Phone:        faker.Phonenumber(),
		Email:        faker.Email(),
		Gender:       faker.Gender(),
		IsActive:     true,
		CreatedAt:    time.Time{},
		UpdateAt:     time.Time{},
		DeleteAt:     time.Time{},
	}
}
