package dto

import "UserServiceQF/models"

type UserDto struct {
	Name     string `json:"name" validate:"required"`
	Contact  string `json:"contact" validate:"required, gt=0 lt=11"`
	Role     string `json:"role"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Image    string `json:"image"`
}

func UsrToDto(usr models.Users) *UserDto {
	return &UserDto{
		Name:    usr.Name,
		Contact: usr.Contact,
		Email:   usr.Email,
		Image:   usr.Image,
	}
}
