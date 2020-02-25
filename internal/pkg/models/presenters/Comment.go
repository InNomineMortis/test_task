package presenters

import (
	"github.com/jmoiron/sqlx"
	"log"
	"test_task/internal/pkg/models"
)

type Comment struct {
}

func (Comment) SelectCommentsByPost(db *sqlx.DB, postID int) []models.Comment {
	var comments []models.Comment
	if err := db.Select(&comments, "SELECT * FROM comment WHERE post_id=$1", postID); err != nil {
		log.Printf("Error - select comments by post(id=%d): %s", postID, err.Error())
		return nil
	}

	for i, comment := range comments {
		p := PersonPresenter.GetPerson(db, comment.PersonID)
		if p == nil {
			log.Printf("Error - get comment's (id=%d) person: Comment should have person", comment.ID)
			return nil
		}
		comments[i].Person = p
	}

	return comments
}
