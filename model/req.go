package model

type ReqListUsers struct {
	Field  string `form:"field" binding:"omitempty,oneof=username nickname"`
	Key    string `form:"key"`
	Offset int64  `form:"offset" binding:"omitempty,min=0"`
	Limit  int64  `form:"limit,default=20" binding:"omitempty,min=1,max=100"`
	Sort   string `form:"sort,default=id" binding:"omitempty,oneof=id createAt lastLoginAt"`
	Order  string `form:"order,default=asc" binding:"omitempty,oneof=asc desc"`
}
