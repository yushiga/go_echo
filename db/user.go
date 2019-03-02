package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	UserID string `gorm:"unique;not null" validate:"required"`
	Name   string
	Remark string
}
