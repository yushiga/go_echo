package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/yushiga/go_echo/db"
	"github.com/yushiga/go_echo/form"
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
サンプルデータ作成
*/
func CreateSampleData() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		logger.ZapLog.Info("===START CreateSampleData===")
		user := new(db.User)
		user.UserID = "sample"
		user.Name = "sample"
		user.Remark = "sample"
		user.CreatedAt = time.Now()
		dbconn.Create(&user)

		return c.JSON(http.StatusOK, user)
	}
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
		user := new(db.User)
		if err = c.Bind(user); err != nil {
			logger.ZapLog.Error("Bind Error")
			return c.JSON(http.StatusBadRequest, user)
		}

		v := validator.New()
		if err = v.Struct(user); err != nil {
			logger.ZapLog.Error("Validation Error")
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		dbconn.Create(&user)

		return c.JSON(http.StatusOK, user)
	}
}

/*
UPDATE処理
*/
func UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		logger.ZapLog.Info("===START UpdateUser===")
		user := new(form.User)
		if err = c.Bind(user); err != nil {
			logger.ZapLog.Error("Bind Error")
			return c.JSON(http.StatusBadRequest, user)
		}

		v := validator.New()
		if err = v.Struct(user); err != nil {
			logger.ZapLog.Error("Validation Error")
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		dbconn.Model(&user).Updates(&user)

		return c.JSON(http.StatusOK, user)
	}
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
