package models

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type Post struct {
	ID        int       `json:"id"     db:"post_id"`
	Header    string    `json:"header" db:"header"`
	Text      string    `json:"text"   db:"text"`
	Timestamp string    `json:"timestamp" db:"timestamp"`
	PersonID  int       `json:"-" db:"person_id"`
	AddressID int       `json:"-" db:"address_id"`
	Address   *Address  `json:"address,omitempty"`
	Person    *Person   `json:"person"`
	Responses []Comment `json:"responses,omitempty"`
}

func SelectPosts(db *sqlx.DB) []Post {
	var posts []Post
	if err := db.Select(&posts, "SELECT * FROM post"); err != nil {
		log.Printf("Error - select posts: %s", err.Error())
		return nil
	}

	for i, post := range posts {
		p := GetPerson(db, post.PersonID)
		if p == nil {
			log.Printf("Error - get post's (id=%d) person: Post should have person", post.ID)
			return nil
		}
		posts[i].Person = p
		posts[i].Responses = SelectCommentsByPost(db, post.ID)
		posts[i].Address = GetAddress(db, post.AddressID)
	}

	return posts
}
