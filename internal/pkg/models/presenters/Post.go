package presenters

import (
	"github.com/jmoiron/sqlx"
	"log"
	"test_task/internal/pkg/models"
)


type Post struct {
}

func (Post) SelectPosts(db *sqlx.DB) []models.Post {
	var posts []models.Post
	if err := db.Select(&posts, "SELECT * FROM post"); err != nil {
		log.Printf("Error - select posts: %s", err.Error())
		return nil
	}

	for i, post := range posts {
		p := PersonPresenter.GetPerson(db, post.PersonID)
		if p == nil {
			log.Printf("Error - get post's (id=%d) person: Post should have person", post.ID)
			return nil
		}
		posts[i].Person = p
		posts[i].Responses = CommentPresenter.SelectCommentsByPost(db, post.ID)
		posts[i].Address = AddressPresenter.GetAddress(db, post.AddressID)
	}

	return posts
}