package seed

import (
	"github.com/jinzhu/gorm"
)

// Initialize seeding some table
func Initialize(db *gorm.DB) {
	InitSeedUsers(db)
	InitSeedCredentials(db)
}