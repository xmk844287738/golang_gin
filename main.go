package main

import (
	"demo_items/gin_project/gin_vue_v2/model"
	"demo_items/gin_project/gin_vue_v2/router"
	"github.com/gin-gonic/gin"
)

// 主程序入口

func main() {
	// 连接数据库
	db := model.InitDB()
	defer db.Close() // 延迟关闭数据库

	// 定义一个全局路由
	r := gin.Default()

	// 定义路由集合
	r = router.CollecRouter(r)

	// 监听服务
	r.Run("127.0.0.1:9000")
}
