package seed

import (
	"github.com/jinzhu/gorm"

	"github.com/fajarsep12dev/go-api/api/db/models"
	log "github.com/fajarsep12dev/go-api/api/utils/logger"
)

var credentials = []models.Credentials{
	{
		Email: "fajarsep12dev@gmail.com",
		UserId: 1,
	},
}

func InitSeedCredentials(db *gorm.DB) {
	var existingCredentias []models.Credentials

	db.Find(&existingCredentias)

	if len(existingCredentias) == 0 {
		for i := range credentials {
			err := db.Debug().
					Set("audited:current_user", 1).
					Model(&models.Credentials{}).
					Create(&credentials[i]).
					Error
	
			if err != nil {
				log.Fatal("cannot seed credentials table: %v", err)
			}
		}
	}	
}