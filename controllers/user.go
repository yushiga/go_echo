package controllers

import (
	"github.com/hmhr/go_sample/logger"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"

	"github.com/hmhr/go_sample/db"
	//"github.com/hmhr/sample_echo/validate"
)

var dbconn *gorm.DB

func Init() {
	dbconn = db.DbConnection()
}

func FindUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.ZapLog.Info("===START FindUsers===")
		users := []db.User{}
		result := dbconn.Find(&users)
		return c.JSON(http.StatusOK, result)
	}
}

func GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.ZapLog.Info("===START GetUsers===")
		user := db.User{}
		user_id := c.Param("id")
		user.User_id = user_id
		result := dbconn.Where("user_id = ?", user_id).First(&user)
		return c.JSON(http.StatusOK, result)
	}
}

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

		if user.ID == 0 {
			logger.ZapLog.Info("---CREATE---")
			dbconn.Create(user)
		} else {
			logger.ZapLog.Info("---UPDATE---")
			dbconn.Save(user)
		}
		return c.JSON(http.StatusOK, user)
	}
}
