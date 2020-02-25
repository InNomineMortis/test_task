package presenters

import (
	"github.com/jmoiron/sqlx"
	"test_task/internal/pkg/models"
)

type AddressMock struct {
}

func (AddressMock) GetAddress(_ *sqlx.DB, id int) *models.Address {
	return &models.Address{
		ID:          id,
		Index:       111111,
		Country:     "Lorem",
		Region:      "Lorem",
		City:        "Ipsum",
		Street:      "Dolores",
		Metro:       "Igma",
		HouseNumber: 10,
		Section:     "2",
		Flat:        "89",
	}
}
