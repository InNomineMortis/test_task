package presenters

import (
	"log"
	"test_task/internal/pkg/database"
	"test_task/internal/pkg/models"
)

type Person struct {
}

func (Person) GetPerson(id int) *models.Person {
	var person models.Person
	if err := database.DB.Get(&person, "SELECT * FROM person WHERE person_id = $1", id); err != nil {
		log.Printf("Error - get person (id=%d): %s", id, err.Error())
		return nil
	}

	person.Address = AddressPresenter.GetAddress(person.AddressID)

	if person.Address == nil {
		log.Printf("Error - get person's (id=%d) adrress: Person should have address", id)
		return nil
	}

	return &person
}
