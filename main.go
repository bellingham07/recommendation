package main

import (
	"recommendation/common"
	"recommendation/initial"
)

func main() {
	db := common.InitDB() //初始化数据库
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic("failed to close database" + err.Error())
		}
		err = sqlDB.Close()
		if err != nil {
			return
		}
	}()
	r := initial.Routers()
	panic(r.Run(":9090"))

}
