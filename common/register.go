package common

import (
	"gorm.io/gorm"
	"recommendation/model"
)

func IsEmailExisted(db *gorm.DB, email string) bool {
	var user model.TbEshop
	db.Where("email=?", email).First(&user)
	if user.Id != "" {
		return true
	}
	return false
}
