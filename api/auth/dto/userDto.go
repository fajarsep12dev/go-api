package dto

// UserDto model
type UserDto struct {
	FirstName   string	`json:"firstName" valid:"Required;MaxSize(25)"`
	LastName    string	`json:"lastName" valid:"Required;MaxSize(25)"`
}