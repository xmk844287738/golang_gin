package router

import (
	"demo_items/gin_project/gin_vue_v2/controller"
	"github.com/gin-gonic/gin"
)

func CollecRouter(r *gin.Engine) (*gin.Engine) {

	r.POST("/api/auto/register", controller.Register) // 注册路由

	r.POST("/api/auto/login", controller.Login) // 注册路由

	return r
}