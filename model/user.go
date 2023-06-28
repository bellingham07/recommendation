package model

type User struct {
	Id   string `form:"id,default=10"`
	Name string `form:"name"`
}
