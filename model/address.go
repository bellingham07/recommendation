package model

type TbAddress struct {
	Id          string `form:"id"`
	User        string // 用户id
	Name        string `form:"name"` // 收货人or发货人名字
	Phonenumber string `form:"tel"`  // 手机号
	Province    string `form:"province"`
	City        string `form:"city"`
	Area        string `form:"area"`
	Town        string `form:"town"`
	Detail      string `form:"detail"` // 具体地址
	Type        string // 地址类型
	Post        string // 邮编
}
