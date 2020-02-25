package presenters

import (
	"github.com/jmoiron/sqlx"
	"log"
	"test_task/internal/pkg/models"
)

type Person struct {
}

func (Person) GetPerson(db *sqlx.DB, id int) *models.Person {
	var person models.Person
	if err := db.Get(&person, "SELECT * FROM person WHERE person_id = $1", id); err != nil {
		log.Printf("Error - get person (id=%d): %s", id, err.Error())
		return nil
	}

	person.Address = AddressPresenter.GetAddress(db, person.AddressID)

	if person.Address == nil {
		log.Printf("Error - get person's (id=%d) adrress: Person should have address", id)
		return nil
	}

	return &person
}
