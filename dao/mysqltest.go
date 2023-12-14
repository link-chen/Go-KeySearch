package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db       *gorm.DB
	Mysqlerr error
)

const (
	Mysqldb = "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8"
)

func Check() {
	db, err := gorm.Open(mysql.Open(Mysqldb), &gorm.Config{})
	if db != nil {

	}
	if err != nil {

	}
}
