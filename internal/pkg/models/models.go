package models

import (
	"github.com/google/uuid"
	"time"
)

var Loc *time.Location

func init() {
	Loc, _ = time.LoadLocation("Europe/Moscow")
}

const TimeFormat = time.RFC3339 //duplicate

type Post struct {
	ID     			uuid.UUID   `json:"id"     db:"id"`
	PersonID 		uuid.UUID   `json:"person_id" db:"person_id"`
	Header 			string      `json:"header" db:"header"`
	Text   			string      `json:"text"   db:"text"`
	Timestamp 		string 		`json:"timestamp" db:"timestamp"`
	Address 		address  	`json:"address" db:"address"`
	Comments 		[]comments  `json:"comments"`
}

type address struct {
	Index 	 		int 		`json:"index"`
	Country  		string 		`json:"country"`
	Region   		string 		`json:"region"`
	City     		string 		`json:"city"`
	Street   		string 		`json:"street"`
	Metro    		string 		`json:"metro"`
	HouseNumber 	int 		`json:"houseNumber"`
	Section 		string 		`json:"section"`
	Flat 			string 		`json:"flat"`
}

type comments struct {
	PostId 			uuid.UUID 	`json:"post_id" db:"post_id"`
	ID 				uuid.UUID 	`json:"id" db:"id"`
	Text 			string		`json:"text" db:"text"`
	PersonId 		uuid.UUID 	`json:"person_id" db:"person_id"`
}

type Person struct {
	ID 				uuid.UUID	`json:"id" db:"id"`
	AvatarURL 		string		`json:"avatar_url" db:"avatar_url"`
	Name 			string 		`json:"name" db:"name"`
	Surname 		string		`json:"surname" db:"surname"`
	Patronymic 		string 		`json:"patronymic" db:"patronymic"`
	Address 		address  	`json:"address" db:"address"`
}

type Error struct {
	Message string            `json:"error"`
	Params  map[string]string `json:"params"`
}

