package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func InitDB() {
	dbConfig() //初始化数据库
}

func dbConfig() *gorm.DB {
	config := viper.New()
	config.SetConfigName("application")
	config.AddConfigPath("./config")
	config.SetConfigType("yaml")
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.GetString("mysql.username"), config.GetString("mysql.password"), config.GetString("mysql.host"), config.GetString("mysql.port"), config.GetString("mysql.DB"))
	fmt.Println("dsn", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn, // Data Source Name，参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})

	if err != nil {
		panic(err)
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
