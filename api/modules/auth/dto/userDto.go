package dto

// UserDto model
type UserDto struct {
	Email   	string  `gorm:"not null;unique_index:idx_email"`
	FirstName   string	`gorm:"size:50;not null" json:"firstname"`
	LastName    string	`gorm:"size:50;not null" json:"lastName"`
	Password  	string  `gorm:"size:100;not null;" json:"password"`
}