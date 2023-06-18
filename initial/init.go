package initial

import (
	"recommendation/database"
	"recommendation/redis"
)

func Init() {
	database.InitDB()
	redis.InitRedis()
}
