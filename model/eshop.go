package model

type TbEshop struct {
	Id          int
	Password    string `json:"password"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Tel         string `json:"phonenumber"`
	Email       string
	Seller      string
	Avatar      string
	Intro       string
	PlatformUrl string
	Platform    string
	Category    int
	CreditPoint int
	Likes       int
	Age         int
}
