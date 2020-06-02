package router

import (
	"demo_items/gin_project/gin_vue_v2/controller"
	"demo_items/gin_project/gin_vue_v2/middleware"
	"github.com/gin-gonic/gin"
)

func CollecRouter(r *gin.Engine) (*gin.Engine) {

	r.Use(middleware.CORSMiddleware()) // 添加全局中间件, 处理前台访问时发生的跨域问题

	r.POST("/api/auto/register", controller.Register) // 注册路由

	r.POST("/api/auto/login", controller.Login) // 登录路由

	r.GET("/api/auto/info", middleware.AuthMiddleware(),controller.Info) // 用户信息路由

	categoryRouter := r.Group("/categories")
	categoryController := controller.NewCategoryController()  // 定义一个关于文章分类的路由管理器
	categoryRouter.POST("", categoryController.Create)  // 增加分类
	categoryRouter.DELETE("/:id", categoryController.Delete) // 删除分类
	categoryRouter.PUT("/:id", categoryController.Update) // 修改分类
	categoryRouter.GET("/:id", categoryController.Show)  // 查看分类

	return r
}