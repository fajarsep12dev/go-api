package dto

// UserDto model
type UserDto struct {
	Email   	string  `valid:"Email"`
	FirstName   string	`form:"firstname" valid:"Required;MaxSize(25)"`
	LastName    string	`form:"lastName" valid:"Required;MaxSize(25)"`
	Password  	string  `form:"password" valid:"Required;MaxSize(12)"`
}