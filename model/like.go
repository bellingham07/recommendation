package model

type TbLike struct {
	Id      string `json:"Id" form:"Id"`
	LikeId  string `json:"LikeId" form:"LikeId"`   // 点赞者id
	LikedId string `json:"LikedId" form:"LikedId"` // 被点赞者id 可以是人，也可以是商品，反正uid不重复
}
