package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"recommendation/common"
	"recommendation/model"
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
	fmt.Println(address.User)
}
