package migrations

import (
	"github.com/fajarsep12dev/go-api/core/db/models"
	log "github.com/fajarsep12dev/go-api/core/utils/logger"

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
