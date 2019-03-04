package db

import (
	"fmt"
	"log"
	"os"

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
	logPath := config.Config.Database.LogPath
	logMode := config.Config.Database.LogMode
	db, err = gorm.Open(driver, user+":"+password+"@/"+dbname+"?charset=utf8&parseTime=True&loc=Local")

	db.AutoMigrate(&User{})

	// クエリログ
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	db.LogMode(logMode)
	db.SetLogger(log.New(file, "", 0))

	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
}

// DB接続
func DbConnection() *gorm.DB {
	return db
}
