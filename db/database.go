package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/yushiga/go_echo/config"
)

var db *gorm.DB
var err error

// DB接続情報
func Init() {
	driver := config.Config.Database.Driver
	user := config.Config.Database.User
	password := config.Config.Database.Password
	dbname := config.Config.Database.Name
	db, err = gorm.Open(driver, user+":"+password+"@/"+dbname+"?charset=utf8&parseTime=True&loc=Local")

	db.AutoMigrate(&User{})

	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
}

// DB接続
func DbConnection() *gorm.DB {
	return db
}
