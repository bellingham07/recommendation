package main

import (
	"recommendation/common"
	"recommendation/initial"
)

func main() {
	//InitConfig()          //加在配置类
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
	//r := gin.Default()
	//r = myRouters.CollectRoute(r)          //配置路由
	r := initial.Routers()
	//port := viper.GetString("server.port") //获取配置类所设置的端口号
	//if port != "" {
	panic(r.Run(":9090"))
	//}
	//panic(r.Run())
}

//func InitConfig() {
//	//初始化配置类
//	workDir, _ := os.Getwd()
//	fmt.Println("曹金波", workDir)
//	viper.SetConfigName("application")
//	viper.SetConfigType("yml")
//	viper.AddConfigPath("./config")
//	err := viper.ReadInConfig()
//	if err != nil {
//		panic(err)
//	}
//}
