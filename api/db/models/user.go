package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/audited"
)

// User model
type User struct {
	gorm.Model
	FirstName   string	`gorm:"size:50;not null" json:"firstname"`
	LastName    string	`gorm:"size:50;not null" json:"lastName"`
	Credentials	Credentials `gorm:"foreignkey:UserId"`
	audited.AuditedModel
}
