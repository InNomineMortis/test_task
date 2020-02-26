package models

type CommentPresenter interface {
	SelectCommentsByPost(postID int) []Comment
}

type Comment struct {
	ID       int     `json:"id" db:"comment_id"`
	Text     string  `json:"text" db:"comment_text"`
	PostId   int     `json:"-" db:"post_id"`
	PersonID int     `json:"-" db:"person_id"`
	Person   *Person `json:"person"`
}
