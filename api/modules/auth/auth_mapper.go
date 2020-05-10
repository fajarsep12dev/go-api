package auth

import (
	"github.com/fajarsep12dev/go-api/api/db/models"
	"github.com/fajarsep12dev/go-api/api/modules/auth/dto"
)

func ToUser(userDto dto.UserDto) models.User {
	return models.User{
		FirstName: userDto.FirstName,
		LastName: userDto.LastName,
	}
}

func ToUserDTO(user models.User) dto.UserDto {
	return dto.UserDto{
		FirstName: user.FirstName,
		LastName: user.LastName,
	}
}

func ToUserDTOs(users []models.User) []dto.UserDto {
	userdtos := make([]dto.UserDto, len(users))

	for i, itm := range users {
		userdtos[i] = ToUserDTO(itm)
	}

	return userdtos
}