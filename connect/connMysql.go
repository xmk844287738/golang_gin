package connect

import (
	"demo_items/gin_project/gin_vue_v2/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB()  (*gorm.DB) {
	db, err := gorm.Open("mysql", "root:root@(192.168.233.1:3306)/demo1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("conn mysql faild of error: %#v\n", err)
	}

	db.AutoMigrate(&model.User{}) // 与数据的数据表建立对应(映射)

	DB = db

	return db
}

func GetDB() (*gorm.DB) {
	return DB
}