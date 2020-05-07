package models

import (
	"github.com/fajarsep12dev/go-api/api/utils/app"

	"github.com/qor/audited"
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Email   	string  `gorm:"not null;unique_index:idx_email"`
	FirstName   string	`gorm:"size:50;not null" json:"firstname"`
	LastName    string	`gorm:"size:50;not null" json:"lastName"`
	Password  	string  `gorm:"size:100;not null;" json:"password"`
	audited.AuditedModel
}

// BeforeSave run for hashing password while save
func (u *User) BeforeSave() error {
	hashedPassword, err := app.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}