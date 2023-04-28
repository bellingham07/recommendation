package model

import "time"

type TbGood struct {
	Id             string
	Name           string  `json:"Name" form:"Name"`
	Img            string  `json:"Img" form:"Img,*multipart.FileHeader"`
	Category       string  `json:"Category" form:"Category"`  //商品分类id
	Brand          int     `json:"Brand,string" form:"Brand"` //品牌
	MarketPrice    float64 `json:"MarketPrice,string" form:"MarketPrice"`
	CelebrityPrice float32 `json:"CelebrityPrice,string" form:"CelebrityPrice"`
	GoodUrl        string  `json:"GoodUrl" form:"GoodUrl"`
	Status         int     // 状态：0禁用，1启用
	CreateTime     time.Time
	UpdateTime     time.Time
	Intro          string `json:"Intro" form:"Intro"`
	Eshop          string
}
