package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/audited"

	"github.com/fajarsep12dev/go-api/core/utils/app"
)

// Credentials Model
type Credentials struct {
	gorm.Model
	Email   		string  `gorm:"not null;unique_index:idx_email"`
	EmailConfirmed 	bool	`gorm:"default:false"`
	IsActive 		bool	`gorm:"default:false"`
	Password  		string  `gorm:"size:100;not null;" json:"password"`
	UserId			uint	`json:"user_id"`
	audited.AuditedModel
}

// BeforeSave run for hashing password while save
func (c *Credentials) BeforeSave() error {
	hashedPassword, err := app.Hash(c.Password)
	if err != nil {
		return err
	}
	c.Password = string(hashedPassword)

	return nil
}