package seed

import (
	"github.com/fajarsep12dev/go-api/api/db/models"
	log "github.com/fajarsep12dev/go-api/api/utils/logger"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Email: "fajarsep12dev@gmail.com",
		FirstName: "Fajar",
		LastName: "Septiawan",
	},
}

// Initialize seeding some table
func Initialize(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.User{}).Error

	if err != nil {
		log.Fatal("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatal("cannot migrate table: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal("cannot seed users table: %v", err)
		}
	}
}