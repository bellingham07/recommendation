package dto

import "recommendation/model"

type UserDto struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
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
