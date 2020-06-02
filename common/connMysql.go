package common

import (
	"demo_items/gin_project/gin_vue_v2/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"net/url"
)

var DB *gorm.DB

func InitDB()  (*gorm.DB) {
	//db, err := gorm.Open("mysql", "root:root@(192.168.233.1:3306)/demo1?charset=utf8mb4&parseTime=True&loc=Local")

	// 使用viper 初始化配置数据库
	driveName := viper.GetString("dataSource.driveName") // 使用的数据库驱动类型
	userName := viper.GetString("dataSource.userName")
	password := viper.GetString("dataSource.password")
	host := viper.GetString("dataSource.host")
	port := viper.GetString("dataSource.port")
	dataBaseName := viper.GetString("dataSource.dataBaseName")
	charset := viper.GetString("dataSource.charset")
	loc := viper.GetString("dataSource.loc")

	// 参数组合
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		userName,
		password,
		host,
		port,
		dataBaseName,
		charset,
		url.QueryEscape(loc))

	db, err := gorm.Open(driveName, args)
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