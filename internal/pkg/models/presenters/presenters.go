package presenters

import "test_task/internal/pkg/models"

var (
	PostPresenter    models.PostPresenter
	AddressPresenter models.AddressPresenter
	CommentPresenter models.CommentPresenter
	PersonPresenter  models.PersonPresenter
)

func init() {
	Setup()
}

func Setup()  {
	PostPresenter = Post{}
	AddressPresenter = Address{}
	CommentPresenter = Comment{}
	PersonPresenter = Person{}
}

func SetupMocks() {
	AddressPresenter = AddressMock{}
}