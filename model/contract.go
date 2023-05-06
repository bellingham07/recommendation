package model

import "time"

type TbContract struct {
	Id         string    `form:"Id"`
	Eshop      string    `form:"Eshop"`     //电商id
	Celebrity  string    `form:"celebrity"` //网红id
	Good       string    //商品id
	CreateTime time.Time //创建时间
	CreateBy   string    //谁发起的 商家或者网红id
	StartTime  time.Time
	EndTime    time.Time
	Status     string //状态 1，申请合作 2，接受合作，合约开始生效 3，合约失效 4，合约取消
}
