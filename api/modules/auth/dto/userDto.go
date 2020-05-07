package dto

// UserDto model
type UserDto struct {
	Email   	string  `json:"email" valid:"Email"`
	FirstName   string	`json:"firstName" valid:"Required;MaxSize(25)"`
	LastName    string	`json:"lastName" valid:"Required;MaxSize(25)"`
	Password  	string  `json:"password" valid:"Required;MaxSize(12)"`
}