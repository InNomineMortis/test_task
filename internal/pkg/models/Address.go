package models

type AddressPresenter interface {
	GetAddress(id int) *Address
}

type Address struct {
	ID          int    `json:"id" db:"address_id"`
	Index       int    `json:"index" db:"index"`
	Country     string `json:"country" db:"country"`
	Region      string `json:"region" db:"region"`
	City        string `json:"city" db:"city"`
	Street      string `json:"street" db:"street"`
	Metro       string `json:"metro" db:"metro"`
	HouseNumber int    `json:"houseNumber" db:"house_number"`
	Section     string `json:"section" db:"section"`
	Flat        string `json:"flat" db:"flat"`
}
