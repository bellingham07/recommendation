package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "dsn"
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
