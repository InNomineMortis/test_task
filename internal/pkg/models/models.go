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
	ID     			int  		`json:"id"     db:"post_id"`
	Header 			string      `json:"header" db:"header"`
	Text   			string      `json:"text"   db:"text"`
	Timestamp 		string 		`json:"timestamp" db:"timestamp"`
	Persons			Person		`json:"person"`
	Responses 		[]Comments	`json:"responses,omitempty"`
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

type Comments struct {
	PostId 			int 		`json:"post_id" db:"post_id"`
	ID 				int 		`json:"id" db:"comment_id"`
	Text 			string		`json:"text" db:"comment_text"`
	PersonId 		int			`json:"person.id" db:"person_id"`
}

type Person struct {
	ID 				int			`json:"id" db:"person_id"`
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

type GetPost struct {
	Post_ID     	int  		`json:"id"              db:"post_id"`
	Header 			string      `json:"header"          db:"header"`
	Text   			string      `json:"text"            db:"post_text"`
	Timestamp 		string 		`json:"timestamp"       db:"timestamp"`
	Person_ID 		int			`json:"id"              db:"person_id"`
	AvatarURL 		string		`json:"avatarURL"       db:"avatar_url"`
	Name 			string 		`json:"name"            db:"name"`
	Surname 		string		`json:"surname"         db:"surname"`
	Patronymic 		string 		`json:"patronymic"      db:"patronymic"`
	Address 		address  	`json:"address"`
	PersonID		int			`json:"person_id"`
	Index 	 		int 		`json:"index"           db:"index"`
	Country  		string 		`json:"country"         db:"country"`
	Region   		string 		`json:"region"          db:"region"`
	City     		string 		`json:"city"            db:"city"`
	Street   		string 		`json:"street"          db:"street"`
	Metro    		string 		`json:"metro"           db:"metro"`
	HouseNumber 	int 		`json:"houseNumber"     db:"house_number"`
	Section 		string 		`json:"section"         db:"section"`
	Flat 			string 		`json:"flat"            db:"flat"`
	PostId 			int 		`json:"post_id"         db:"post_id"`
	ID 				int 		`json:"id"              db:"comment_id"`
	Comment_Text 	string		`json:"text"            db:"comment_text"`
	PersonId 		int			`json:"person.id"       db:"person_id"`
}

type Empty struct {

}