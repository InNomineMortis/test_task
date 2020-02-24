package models

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type Comment struct {
	ID       int     `json:"id" db:"comment_id"`
	Text     string  `json:"text" db:"comment_text"`
	PostId   int     `json:"-" db:"post_id"`
	PersonID int     `json:"-" db:"person_id"`
	Person   *Person `json:"person"`
}

func SelectCommentsByPost(db *sqlx.DB, postID int) []Comment {
	var comments []Comment
	if err := db.Select(&comments, "SELECT * FROM comment WHERE post_id=$1", postID); err != nil {
		log.Printf("Error - select comments by post(id=%d): %s", postID, err.Error())
		return nil
	}

	for i, comment := range comments {
		p := GetPerson(db, comment.PersonID)
		if p == nil {
			log.Printf("Error - get comment's (id=%d) person: Comment should have person", comment.ID)
			return nil
		}
		comments[i].Person = p
	}

	return comments
}
