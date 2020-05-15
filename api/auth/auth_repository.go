package auth

import (
	"github.com/fajarsep12dev/go-api/core/db/models"

	"github.com/jinzhu/gorm"
)

// AuthRepository Interface
type AuthRepository struct {
	DB *gorm.DB
}

// ProvideAuthRepository constructor of AuthRepository
func ProvideAuthRepository(DB *gorm.DB) AuthRepository {
	return AuthRepository{
		DB: DB,
	}
}

// GetAll get all data users
func (authRepo *AuthRepository) GetAll() []models.User {
	var users []models.User
	authRepo.DB.Find(&users)

	return users
}

// Save data users
func (authRepo *AuthRepository) Save(user models.User) models.User {
	authRepo.DB.Save(&user)

	return user
}

// Delete data users
func (authRepo *AuthRepository) Delete(user models.User) {
	authRepo.DB.Delete(&user)
}