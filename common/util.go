package common

import "recommendation/model"

func IsGoodExist(name string) bool {
	db := GetDB()
	var good model.TbGood
	db.Where("name=?", name).First(&good)
	if good.Id != "" {
		return true
	}
	return false
}
