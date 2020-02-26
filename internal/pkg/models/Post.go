package models

type PostPresenter interface {
	SelectPosts() []Post
}

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
