package common

import (
	"gorm.io/gorm"
	"recommendation/model"
)

func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.TbEshop
	db.Where("tel=?", telephone).First(&user)
	if user.Id != "" {
		return true
	}
	return false
}
