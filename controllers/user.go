package controllers

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/yushiga/go_echo/db"
	"github.com/yushiga/go_echo/logger"
	"gopkg.in/go-playground/validator.v9"
)

var dbconn *gorm.DB

/*
初期化
*/
func Init() {
	dbconn = db.DbConnection()
}

/*
全件検索
*/
func FindUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.ZapLog.Info("===START FindUsers===")
		users := []db.User{}
		result := dbconn.Find(&users)
		return c.JSON(http.StatusOK, result)
	}
}

/*
指定ID検索
*/
func GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.ZapLog.Info("===START GetUsers===")
		user := db.User{}
		parseID, _ := strconv.Atoi(c.Param("id"))
		user.ID = uint(parseID)
		dbconn.First(&user)
		return c.JSON(http.StatusOK, user)
	}
}

/*
INSERT処理
*/
func AddUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		logger.ZapLog.Info("===START AddUser===")
		v := validator.New()
		user := new(db.User)
		if err = c.Bind(user); err != nil {
			logger.ZapLog.Error("Bind Error")
			return c.JSON(http.StatusBadRequest, user)
		}

		if err = v.Struct(user); err != nil {
			logger.ZapLog.Error("Validation Error")
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		dbconn.Create(user)

		return c.JSON(http.StatusOK, user)
	}
}

/*
UPDATE処理
*/
func UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		logger.ZapLog.Info("===START UpdateUser===")
		v := validator.New()
		user := new(db.User)
		if err = c.Bind(user); err != nil {
			logger.ZapLog.Error("Bind Error")
			return c.JSON(http.StatusBadRequest, user)
		}

		if err = v.Struct(user); err != nil {
			logger.ZapLog.Error("Validation Error")
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		search := SearchUser(user.ID)
		user.CreatedAt = search.CreatedAt
		dbconn.Save(&user)
		//dbconn.Model(&user).Update("user_id", "gyagya")

		return c.JSON(http.StatusOK, user)
	}
}

func SearchUser(id uint) db.User {
	logger.ZapLog.Info("===START SearchUsers===")
	user := db.User{}
	user.ID = id
	dbconn.First(&user)
	return user
}

/*
DELETE処理
type: 0=物理削除, 1=論理削除
*/
//func DeleteUser() echo.HandlerFunc {
//	return func(c echo.Context) (err error) {
//		logger.ZapLog.Info("===START DeleteUser===")
//
//		if c.Param("type") == "0" {
//			logger.ZapLog.Info("---DELETE---")
//		} else {
//			logger.ZapLog.Info("---SOFT DELETE---")
//		}
//		return c.JSON(http.StatusOK, user)
//	}
//
