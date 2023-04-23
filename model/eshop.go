package model

type TbEshop struct {
	Id          string
	Password    string `json:"Password"`
	Username    string `json:"Username"`
	Name        string `json:"Name"`
	Tel         string `json:"Phonenumber"`
	Email       string `json:"Email"`
	Seller      string `json:"Seller"`
	Avatar      string `json:"Avatar"`
	Intro       string `json:"Intro"`
	PlatformUrl string `json:"PlatformUrl"`
	Platform    string `json:"Platform"`
	Category    int    `json:"Category"`
	CreditPoint int    `json:"CreditPoint,string"`
	Likes       int    `json:"Likes"`
	Age         int    `json:"Age,string"`
}
