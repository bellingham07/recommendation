package common

import (
	"context"
	"gorm.io/gorm"
	"recommendation/model"
	"recommendation/redis"
)

func IsEmailExisted(db *gorm.DB, email string) bool {
	var user model.TbEshop
	db.Where("email=?", email).First(&user)
	if user.Id != "" {
		return true
	}
	return false
}

func IsEmailCodeValid(code string, email string) bool {
	rds := redis.GetRedis()
	cacheCode := rds.Get(context.Background(), email)
	if cacheCode.Val() == code {
		return true
	}
	return false
}
