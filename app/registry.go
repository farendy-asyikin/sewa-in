package app

import "sewa-in/app/models"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Stock{}},
		{Model: models.RentChart{}},
		{Model: models.TransactionOngoing{}},
		{Model: models.TransactionHistory{}},
	}
}
