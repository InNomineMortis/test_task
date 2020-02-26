package presenters

import (
	"log"
	"test_task/internal/pkg/database"
	"test_task/internal/pkg/models"
)


type Post struct {
}

func (Post) SelectPosts() []models.Post {
	var posts []models.Post
	if err := database.DB.Select(&posts, "SELECT * FROM post"); err != nil {
		log.Printf("Error - select posts: %s", err.Error())
		return nil
	}

	for i, post := range posts {
		p := PersonPresenter.GetPerson(post.PersonID)
		if p == nil {
			log.Printf("Error - get post's (id=%d) person: Post should have person", post.ID)
			return nil
		}
		posts[i].Person = p
		posts[i].Responses = CommentPresenter.SelectCommentsByPost(post.ID)
		posts[i].Address = AddressPresenter.GetAddress(post.AddressID)
	}

	return posts
}