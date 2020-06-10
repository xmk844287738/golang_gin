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

	// 文章分类路由组
	categoryRouter := r.Group("/category")
	categoryController := controller.NewCategoryController() // 定义一个关于文章分类的路由管理器
	{
		categoryRouter.POST("", categoryController.Create)       // 增加分类
		categoryRouter.DELETE("/:id", categoryController.Delete) // 删除分类
		categoryRouter.PUT("/:id", categoryController.Update)    // 修改分类
		categoryRouter.GET("/:id", categoryController.Show)      // 查看对应的id分类
	}
	r.GET("/categories/showAllCategories", categoryController.ShowAllCategories)      // 查看所有分类

	// 文章路由组
	PostRouter := r.Group("/post")
	PostRouter.Use(middleware.AuthMiddleware())	// 认证后用户登录进行文章增删改查操作
	postController := controller.NewPostcontroller() // 定义一个关于文章的路由管理器
	{
		PostRouter.POST("", postController.Create) // 增加文章
		PostRouter.DELETE("/:postId", postController.Delete) // 删除文章
		PostRouter.PUT("/:postId", postController.Update) // 修改文章
		//PostRouter.GET("/:postId", postController.Show)	// 查询一篇文章
	}
	// 游客身份浏览,不用添加middleware.AuthMiddleware() 中间件认证
	r.GET("/anonymousUser/post/:postId", postController.Show)	// 查询一篇文章
	r.GET("/anonymousUser/posts/*postNum", postController.MultiplyShows) // 查询多篇文章

	return r
}