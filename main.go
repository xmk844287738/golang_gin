package main

import (
	"demo_items/gin_project/gin_vue_v2/common"
	"demo_items/gin_project/gin_vue_v2/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
)

// 主程序入口

func InitConfig()  {
	// 获取当前工作项目的路径
	workDir, err := os.Getwd()
	if err != nil {
		log.Printf("获取工作项目路径失败:%v\n", err)
		return
	}

	// 设置配置文件的文件类型
	viper.SetConfigType("yml")

	// 设置配置文件的文件名
	viper.SetConfigName("Application")

	// 组成配置文件的绝对路径
	viper.AddConfigPath(workDir + "/config")

	// 读取配置文件的信息
	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("读取配置文件的信息发生错误:%v\n", err)
		return
	}

}

func main() {

	InitConfig()// 读取配置文件的信息

	// 连接数据库
	db := common.InitDB()
	defer db.Close() // 延迟关闭数据库

	// 定义一个全局路由
	r := gin.Default()

	// 定义路由集合
	r = router.CollecRouter(r)

	// 监听服务
	r.Run("127.0.0.1:9000")
}
