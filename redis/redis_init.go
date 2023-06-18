package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var client *redis.Client

func InitRedis() {

	config := viper.New()
	config.SetConfigName("application")
	config.AddConfigPath("./config")
	config.SetConfigType("yaml")
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	client = redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.addr"),
		Password: config.GetString("redis.password"),
		DB:       config.GetInt("redis.db"),
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	//ctx.Value(config.GetString("redis.password"))
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		panic("连接redis失败：" + err.Error())
	}

}

func GetRedis() *redis.Client {
	return client
}
