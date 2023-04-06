package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"recommendation/common"
	"recommendation/routes"
)

func main() {
	InitConfig()          //加在配置类
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
	r := gin.Default()
	r = routes.CollectRoute(r)             //配置路由
	port := viper.GetString("server.port") //获取配置类所设置的端口号
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	//初始化配置类
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {

	}
}
