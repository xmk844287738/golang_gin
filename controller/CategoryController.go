package controller

import (
	"demo_items/gin_project/gin_vue_v2/common"
	"demo_items/gin_project/gin_vue_v2/model"
	"demo_items/gin_project/gin_vue_v2/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

// 定义一个接口方便使用编辑器完成增删改查的代码 快捷键：1.Alt + Insert 2.ctrl + I
type ICategoryController interface {
	ResetController
}

type CategoryController struct {  // 定义一个关于文章分类的路由管理器
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB() // 连接数据库

	// 迁移 Category 数据模型; 与数据库建立模型映射
	db.AutoMigrate(&model.Category{})

	return CategoryController{DB:db,}

}

func (c CategoryController) Create(ctx *gin.Context) { //　增加分类

	// 获取body参数　绑定数据
	var createCategory = model.Category{}
	ctx.Bind(&createCategory)
	// 校验数据
	if createCategory.Name == ""{
		response.Fail(ctx, nil, "分类名称必填！")
		return
	}

	// 数据存入数据库
	if err := c.DB.Debug().Create(&createCategory).Error; err != nil { // 验证数据中是否已经存在此分类名
		response.Fail(ctx, nil, fmt.Sprint(err))
		return
	}

	//回应前台
	response.Success(ctx, gin.H{"createCategory": createCategory}, "分类创建成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {	//　删除分类
	// 获取path的数据
	requestCategoryId := ctx.Params.ByName("id")

	var requestCategory  = model.Category{}
	// 查询数据库
	c.DB.Debug().Where("id=?", requestCategoryId).First(&requestCategory)
	if requestCategory.ID == 0 {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	if err := c.DB.Delete(requestCategory).Error; err != nil {
		response.Fail(ctx, nil, "分类删除失败!")
		return
	}
	response.Success(ctx, nil, "分类删除成功!")

}

func (c CategoryController) Update(ctx *gin.Context) { //修改分类
	// 获取path的数据
	updateCategoryId := ctx.Params.ByName("id")
	var Category  = model.Category{}

	// 查询数据库
	if err := c.DB.Debug().Where("id=?", updateCategoryId).First(&Category).Error; err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	// 获取body参数　绑定数据
	var updateCategory = model.Category{}
	ctx.Bind(&updateCategory)
	// 校验数据
	if updateCategory.Name == ""{
		response.Fail(ctx, nil, "分类名称必填！")
		return
	}
	CategoryId, _ := strconv.Atoi(updateCategoryId)
	updateCategory.ID = CategoryId

	c.DB.Debug().Model(&updateCategory).Where("id=?", updateCategory.ID).Update("name", updateCategory.Name)
	response.Success(ctx, gin.H{"updateCategory": updateCategory,}, "success")

}

func (c CategoryController) Show(ctx *gin.Context) { //　查询分类
	// 获取path的数据
	requestCategoryId := ctx.Param("id")

	var requestCategory  = model.Category{}

	// 查询数据库
	c.DB.Debug().Where("id=?", requestCategoryId).First(&requestCategory)
	if requestCategory.ID == 0 {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx, gin.H{"requestCategory":requestCategory,}, "success")

}

