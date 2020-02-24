package models

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type Person struct {
	ID         int     `json:"id" db:"person_id"`
	AvatarURL  string  `json:"avatarURL" db:"avatar_url"`
	Name       string  `json:"name" db:"name"`
	Surname    string  `json:"surname" db:"surname"`
	Patronymic string  `json:"patronymic" db:"patronymic"`
	AddressID  int     `json:"-" db:"address_id"`
	Address    *Address `json:"address" db:"address"`
}

func GetPerson(db *sqlx.DB, id int) *Person {
	var person Person
	if err := db.Get(&person, "SELECT * FROM person WHERE person_id = $1", id); err != nil {
		log.Printf("Error - get person (id=%d): %s", id, err.Error())
		return nil
	}

	person.Address = GetAddress(db, person.AddressID)

	if person.Address == nil {
		log.Printf("Error - get person's (id=%d) adrress: Person should have address", id)
		return nil
	}

	return &person
}
