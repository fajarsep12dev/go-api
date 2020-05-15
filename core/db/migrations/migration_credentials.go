package migrations

import (
	"github.com/jinzhu/gorm"

	"github.com/fajarsep12dev/go-api/core/db/models"
)

func MigrateCredential(db *gorm.DB) {	
	db.Model(&models.Credentials{}).AddForeignKey("user_id", "cafe_user(id)", "RESTRICT", "RESTRICT")
}