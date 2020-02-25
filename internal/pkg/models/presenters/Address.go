package presenters

import (
	"github.com/jmoiron/sqlx"
	"log"
	"test_task/internal/pkg/models"
)

type Address struct {
}

func (Address) GetAddress(db *sqlx.DB, id int) *models.Address {
	var address models.Address
	if err := db.Get(&address, "SELECT * FROM address WHERE address_id=$1", id); err != nil {
		log.Printf("Error - get address (id=%d): %s", id, err.Error())
		return nil
	}

	return &address
}
