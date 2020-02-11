package service

import (
	"os"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = connect()
}

func password() string {
	return os.Getenv("PASSWORD")
}

func connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:"+password()+"@/db?charset=utf8mb4")
	if err != nil {
		panic("数据库连接失败")
	}
	return db
}
