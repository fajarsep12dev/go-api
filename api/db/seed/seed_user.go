package seed

import (
	"github.com/jinzhu/gorm"

	"github.com/fajarsep12dev/go-api/api/db/models"
	log "github.com/fajarsep12dev/go-api/api/utils/logger"
)

var users = []models.User{
	{
		FirstName: "Fajar",
		LastName: "Septiawan",
	},
}

func InitSeedUsers(db *gorm.DB) {
	var existingUser []models.User
	db.Find(&existingUser)

	if len(existingUser) == 0 {
		for i := range users {
			err := db.Debug().
					Set("audited:current_user", 1).
					Model(&models.User{}).
					Create(&users[i]).
					Error
					
			if err != nil {
				log.Fatal("cannot seed users table: %v", err)
			}
		}
	}
}