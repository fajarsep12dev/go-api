package auth

import (
	"github.com/fajarsep12dev/go-api/api/modules/auth/dto"
	"github.com/fajarsep12dev/go-api/api/db/models"

)

func ToUser(userDto dto.UserDto) models.User {
	return models.User{
		Email: userDto.Email,
		FirstName: userDto.FirstName,
		LastName: userDto.LastName,
		Password: userDto.Password,
	}
}

func ToUserDTO(user models.User) dto.UserDto {
	return dto.UserDto{
		Email: user.Email,
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