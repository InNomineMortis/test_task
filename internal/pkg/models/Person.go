package models

type PersonPresenter interface {
	GetPerson(id int) *Person
}

type Person struct {
	ID         int      `json:"id" db:"person_id"`
	AvatarURL  string   `json:"avatarURL" db:"avatar_url"`
	Name       string   `json:"name" db:"name"`
	Surname    string   `json:"surname" db:"surname"`
	Patronymic string   `json:"patronymic" db:"patronymic"`
	AddressID  int      `json:"-" db:"address_id"`
	Address    *Address `json:"address" db:"address"`
}
