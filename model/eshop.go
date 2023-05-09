package model

type TbEshop struct {
	Id          string
	Username    string `json:"account" form:"account"`
	Password    string `json:"password" form:"password"`
	Name        string `json:"Name" form:"name"`
	Tel         string `json:"Phonenumber" form:"phonenumber"`
	Email       string `json:"Email" form:"email"`
	Seller      string `json:"Seller" form:"seller"`
	Avatar      string `json:"Avatar" form:"avatar"`
	Intro       string `json:"Intro" form:"intro"`
	PlatformUrl string `json:"PlatformUrl" form:"PlatformUrl"`
	Platform    string `json:"Platform" form:"Platform"`
	Category    int    `json:"Category" form:"Category"`
	CreditPoint int    `json:"CreditPoint,string" form:"CreditPoint"`
	Likes       int    `json:"Likes" form:"Likes"`
	Age         int    `json:"Age,string" form:"Age"`
}
