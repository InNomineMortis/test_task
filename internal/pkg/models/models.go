package models

import (
	"time"
)

var Loc *time.Location

func init() {
	Loc, _ = time.LoadLocation("Europe/Moscow")
}

const TimeFormat = time.RFC3339 //duplicate

type Post struct {
	ID     			int  		`json:"id"     db:"id"`
	Header 			string      `json:"header" db:"header"`
	Text   			string      `json:"text"   db:"text"`
	Timestamp 		string 		`json:"timestamp" db:"timestamp"`
	Persons			Person		`json:"persons"`
	Responses 		[]Comments	`json:"responses"`
}

type address struct {
	PersonID		int			`json:"id"`
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

type Comments struct {
	PostId 			int 		`json:"post_id" db:"post_id"`
	ID 				int 		`json:"id" db:"id"`
	Text 			string		`json:"text" db:"text"`
	PersonId 		int			`json:"person.id" db:"person_id"`
}

type Person struct {
	ID 				int			`json:"id" db:"id"`
	AvatarURL 		string		`json:"avatarURL" db:"avatar_url"`
	Name 			string 		`json:"name" db:"name"`
	Surname 		string		`json:"surname" db:"surname"`
	Patronymic 		string 		`json:"patronymic" db:"patronymic"`
	Address 		address  	`json:"address" db:"address"`
}

type Error struct {
	Message string            `json:"error"`
	Params  map[string]string `json:"params"`
}

