package user

import (
	"github.com/fajarsep12dev/go-api/api/db/models"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{
		DB: DB,
	}
}

func (userRepo *UserRepository) GetAllUsers() []models.User {
	var users []models.User
	userRepo.DB.Find(&users)

	return users
}