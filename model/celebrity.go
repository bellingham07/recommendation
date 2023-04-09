package model

type TbCelebrity struct {
	Id          int
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phonenumber"`
	Email       string
	Name        string `json:"name"`
	RealName    string
	Sex         int
	Age         int
	Avatar      string
	Intro       string
	PlatformUrl string
	Platform    string
	Category    int
	CreditPoint int
	Likes       int
	Identity    int
}
