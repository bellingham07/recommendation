package model

import "time"

type TbOrder struct {
	Id              string
	Eshop           string //商家id
	Celebrity       string //网红id
	Good            string //商品id
	ConsignAddress  string //发货地址
	ReceiveAddress  string //收货地址
	Remark          string //订单备注
	Type            string //支付方式
	CreateTime      time.Time
	PayTime         time.Time
	ConsignmentTime time.Time // 发货时间
	DoneTime        time.Time //完成时间（收货时间
	Status          int       //状态 0加入购物车 1下单为支付 2支付未发货 3发货 4收货 6取消订单
	PreStatus       int
	Consignee       string //收货人
	Consignor       string //发货人
}
