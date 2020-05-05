package models

import (
	"github.com/qor/audited"
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Email   	string  `gorm:"not null;unique_index:idx_email"`
	FirstName   string	`gorm:"size:50;not null" json:"firstname"`
	LastName    string	`gorm:"size:50;not null" json:"lastName"`
	audited.AuditedModel
}
