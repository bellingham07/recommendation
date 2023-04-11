package dto

import "recommendation/model"

type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ToUserDto(user model.TbEshop) UserDto {
	return UserDto{
		Username: user.Name,
		Password: user.Password,
	}
}

func ToCeleDto(user model.TbCelebrity) UserDto {
	return UserDto{
		Username: user.Name,
		Password: user.Password,
	}
}
