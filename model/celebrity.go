package model

type TbCelebrity struct {
	Id          string
	Username    string `json:"username" form:"account"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phonenumber" form:"phonenumber"`
	Email       string
	Name        string `json:"name" form:"name"`
	RealName    string
	Sex         string `form:"Sex"`
	Age         string `form:"Age"`
	Avatar      string
	Intro       string `form:"Intro"`
	PlatformUrl string
	Platform    string
	Category    int
	CreditPoint int
	Likes       int
	Identity    int
}
