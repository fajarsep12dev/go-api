package seed

import (
	"github.com/fajarsep12dev/go-api/api/db/models"
	log "github.com/fajarsep12dev/go-api/api/utils/logger"

	"github.com/jinzhu/gorm"
)

// Initialize seeding some table
func Initialize(db *gorm.DB) {
	// Migrate table from models
	err := db.AutoMigrate(
			&models.User{},
			&models.Credentials{},
		).Error

	// Initialize FK table
	db.Model(&models.Credentials{}).AddForeignKey("user_id", "cafe_user(id)", "RESTRICT", "RESTRICT")

	if err != nil {
		log.Fatal("cannot migrate, table: %v", err)
	}

	InitSeedUsers(db)
	InitSeedCredentials(db)
}