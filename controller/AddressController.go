package controller

import (
	"github.com/gin-gonic/gin"
	"recommendation/common"
	"recommendation/model"
	"recommendation/response"
)

func Address(c *gin.Context) {
	// 获取新地址信息
	var address model.TbAddress
	err := c.ShouldBind(&address)
	if err != nil {
		panic(err)
	}

	// 获取userid
	address.User = common.GetId(c)
	// 生成全局唯一id
	address.Id = common.GenerateId()

	db := common.GetDB()
	tx := db.Debug().Save(&address)
	if tx.RowsAffected == 0 {
		panic(tx.Error)
	}
	response.Success(c, nil)
}

func GetAddrById(c *gin.Context) {
	var address []model.TbAddress

	db := common.GetDB()
	tx := db.Debug().Where("user=?", common.GetId(c)).Find(&address)
	if tx.Error != nil {
		panic(tx.Error)
	}
	response.Success(c, gin.H{"data": address})
}

func UpdateAddr(c *gin.Context) {
	var address model.TbAddress
	err := c.ShouldBind(&address)
	if err != nil {
		panic(err)
	}
	db := common.GetDB()
	tx := db.Model(&address).Where("id=?", address.Id).Updates(model.TbAddress{Name: address.Name, Phonenumber: address.Phonenumber, Detail: address.Detail, Province: address.Province, City: address.City, Area: address.Area, Town: address.Town})
	if tx.RowsAffected == 0 {
		panic(tx.Error)
	}
	response.Success(c, nil)
}

func DeleteAddr(c *gin.Context) {
	var address model.TbAddress
	address.Id = c.PostForm("id")
	db := common.GetDB()
	tx := db.Delete(&address)
	if tx.RowsAffected == 0 {
		panic(tx.Error)
	}
	response.Success(c, nil)
}
