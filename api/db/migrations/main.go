package migrations

import (
	"github.com/fajarsep12dev/go-api/api/db/models"
	log "github.com/fajarsep12dev/go-api/api/utils/logger"

	"github.com/jinzhu/gorm"
)

func Initialize(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Credentials{},
	).Error

	if err != nil {
		log.Fatal("cannot migrate, table: %v", err)
	}

	MigrateCredential(db)
}
