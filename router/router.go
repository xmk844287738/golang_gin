package router

import (
	"demo_items/gin_project/gin_vue_v2/controller"
	"demo_items/gin_project/gin_vue_v2/middleware"
	"github.com/gin-gonic/gin"
)

func CollecRouter(r *gin.Engine) (*gin.Engine) {

	r.POST("/api/auto/register", controller.Register) // 注册路由

	r.POST("/api/auto/login", controller.Login) // 登录路由

	r.GET("/api/auto/info", middleware.AuthMiddleware(),controller.Info) // 用户信息路由

	return r
}