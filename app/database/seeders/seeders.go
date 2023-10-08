package seeders

import (
	"gorm.io/gorm"
	"sewa-in/app/database/fakers"
)

type Seeders struct {
	Seeders interface{}
}

func RegisterSeeders(db *gorm.DB) []Seeders {
	return []Seeders{
		{Seeders: fakers.UserFaker(db)},
		{Seeders: fakers.StockFaker(db)},
	}
}
func DBSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeeders(db) {
		err := db.Debug().Create(seeder.Seeders).Error
		if err != nil {
			return err
		}
	}

	return nil
}
