package models

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: User{}},
		{Model: Stock{}},
		{Model: RentChart{}},
		{Model: TransactionOngoing{}},
		{Model: TransactionHistory{}},
	}
}
