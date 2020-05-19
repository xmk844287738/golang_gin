package controller

import (
	"demo_items/gin_project/gin_vue_v2/connect"
	"demo_items/gin_project/gin_vue_v2/model"
	"demo_items/gin_project/gin_vue_v2/tools"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// 定义全局数据库操作对象
//var DB *gorm.DB


//func CollectRouter(r *gin.Engine)  (*gin.Engine){
//
//	db := model.GetDB()
//
//	// 定义路由、路由处理函数
//	autoGroup := r.Group("/api/auto")
//	{
//		// 获取前台的注册响应
//		autoGroup.POST("/register", func(c *gin.Context) {
//			// 获取表单相应的参数
//			uname := c.PostForm("name")
//			upassword := c.PostForm("password")
//			telephone := c.PostForm("telephone")
//
//			// 验证表单的参数
//			if len(uname) == 0 {
//				// 调用随机生成的6个大小写字母作为用户名
//				uname = tools.RandUserName(6)
//			}
//
//			// 判断该用户名是否已经存在
//			if tools.IsUserNameExist(db, uname) {
//				c.JSON(http.StatusOK, gin.H{
//					"error": "用户名已存在",
//				})
//				return
//			}
//
//			if len(upassword) < 6 {
//				c.JSON(http.StatusOK, gin.H{
//					"error": "用户密码需6位以上",
//				})
//				return
//			}
//
//			if len(telephone) < 11 {
//				c.JSON(http.StatusOK, gin.H{
//					"error": "手机号不够11位",
//				})
//				return
//			}
//			// 对数据库增加数据
//			newUser := model.User{
//				Name:      uname,
//				Password:  upassword,
//				Telephone: telephone,
//			}
//			db.Debug().Create(&newUser)
//
//			// 对前台进行响应
//			c.JSON(http.StatusOK, gin.H{
//				"status": "创建用户成功",
//			})
//		})
//	}
//
//	return r
//}

func Register(c *gin.Context)  { // 注册路由函数

	db := connect.GetDB()

	// 获取表单相应的参数
	uname := c.PostForm("name")
	upassword := c.PostForm("password")
	telephone := c.PostForm("telephone")

	// 验证表单的参数
	if len(uname) == 0 {
		// 调用随机生成的6个大小写字母作为用户名
		uname = tools.RandUserName(6)
	}

	// 判断该用户名是否已经存在
	if tools.IsUserNameExist(db, uname) {
		c.JSON(http.StatusOK, gin.H{
			"error": "用户名已存在",
		})
		return
	}

	if len(upassword) < 6 {
		c.JSON(http.StatusOK, gin.H{
			"error": "用户密码需6位以上",
		})
		return
	}

	if len(telephone) < 11 {
		c.JSON(http.StatusOK, gin.H{
			"error": "手机号不够11位",
		})
		return
	}

	// 对用户密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(upassword),bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "密码加密错误",
		})
		return
	}

	// 对数据库增加数据,添加用户记录
	newUser := model.User{
		Name:      uname,
		Password:  string(hashedPassword),
		Telephone: telephone,
	}
	db.Debug().Create(&newUser)

	// 对前台进行响应
	c.JSON(http.StatusOK, gin.H{
		"status": "创建用户成功",
	})
}

func Login(c *gin.Context)  { // 登录路由函数

	db := connect.GetDB()

	// 获取表单相应的参数
	uname := c.PostForm("name")
	upassword := c.PostForm("password")

	// 参数验证
	// 验证用户名是否存在
	var user model.User
	db.Where("name=?", uname).Find(&user)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"error": "用户不存在",
		})
		return
	}

	// 验证用户的密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(upassword)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "用户密码错误，请检查",
		})
		return
	}

	// 发放token
	token, err := connect.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "token生产失败",})
		log.Printf("token generte faild error of: %v\n", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":gin.H{"token": token},
		"success": "登录成功",
	})
}

func Info(c *gin.Context) { // 用户信息路由函数
	// 通过上下文获取 user 字段信息
	user, _ := c.Get("user")

	// 根据业务逻辑情况返回 user 字段的对应信息
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{"user": user},
	})

}