package controller

import (
	"demo_items/gin_project/gin_vue_v2/common"
	"demo_items/gin_project/gin_vue_v2/model"
	"demo_items/gin_project/gin_vue_v2/response"
	"demo_items/gin_project/gin_vue_v2/vo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

type IPostController interface {
	ResetController
	MultiplyShows(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func NewPostcontroller() IPostController {
	db := common.GetDB() // 连接数据库
	db.AutoMigrate(&model.Post{}) // 迁移Post数据模型; 与数据库建立模型映射

	return PostController{DB:db,}
}

func (p PostController) Create(ctx *gin.Context) { // 增加文章

	// 声明一个PostRequest结构体对象, 接收并前台的传过来的JSON数据
	var requestPost vo.PostRequest
	if err := ctx.ShouldBind(&requestPost); err != nil { // 校验参数，各字段值是否存在空值
		log.Print(err.Error())
		response.Fail(ctx, nil, "文章信息填写不全")
		return
	}

	// 从上下文管理器中取出 user对应字段的信息
	user, _ := ctx.Get("user")

	// 声明一个Post结构体对象
	var createPost = model.Post{
		UserId: user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:requestPost.Title,
		HeadImg:requestPost.HeadImg,
		Content:requestPost.Content,
	}

	// 将Post结构体对象保存至数据库
	if err := p.DB.Debug().Create(&createPost).Error; err != nil {
		response.Fail(ctx, nil, "保存文章出错")
		return
	}

	// 给前台回应
	response.Success(ctx, gin.H{"createPost": createPost,}, "文章添加成功")

}

func (p PostController) Delete(ctx *gin.Context) { // 删除文章
	// 获取path参数
	postId := ctx.Params.ByName("postId")

	// 判断数据库有无此文章
	var deletePost = model.Post{}
	if err := p.DB.Debug().Where("id=?", postId).First(&deletePost).Error; err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// 进行数据库删除，判断err 返回值是否为 nil
	if err := p.DB.Debug().Delete(&deletePost).Error; err != nil {
		response.Fail(ctx,nil, "文章删除失败")
		return
	}

	// 给前台回应
	response.Success(ctx, gin.H{"title": deletePost.Title,}, "文章删除成功")

}

func (p PostController) Update(ctx *gin.Context) { // 修改文章
	// 获取path参数
	postId := ctx.Params.ByName("postId")

	// 查询数据库是否存在此文章信息
	// 声明一个Post结构体对象
	var posts  = model.Post{}

	// 查询数据库，把查询结构赋值给Post结构体对象, 根据文章的id 判断是否存在于数据库中
	if p.DB.Debug().Where("id=?", postId).First(&posts).RecordNotFound() {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// 声明一个PostRequest结构体对象, 接收并前台的传过来的JSON数据
	var modifyPost vo.PostRequest
	if err := ctx.ShouldBind(&modifyPost); err != nil { // 校验参数，各字段值是否存在空值
		log.Print(err.Error())
		response.Fail(ctx, nil, "文章信息填写不全")
		return
	}

	// 从上下文管理器中取出 user对应字段的信息
	user, _ := ctx.Get("user")

	// 声明一个Post结构体对象
	var updatePost = model.Post{
		UserId: user.(model.User).ID,
		CategoryId: modifyPost.CategoryId,
		Title:modifyPost.Title,
		HeadImg:modifyPost.HeadImg,
		Content:modifyPost.Content,
	}

	// 更新数据库某行对应字段的信息
	if err := p.DB.Debug().Model(&model.Post{}).Where("id=?", postId).Update(updatePost).Error; err != nil {
		response.Fail(ctx, nil, "文章更新错误")
		return
	}

	// 给前台回应
	response.Success(ctx, gin.H{"updatePost": updatePost,}, "文章更新成功")


}

func (p PostController) Show(ctx *gin.Context) { // 查询文章
	// 获取path参数
	postId := ctx.Params.ByName("postId")

	// 声明一个Post结构体对象
	var post  = model.Post{}

	// 查询数据库，把查询结构赋值给Post结构体对象, 根据文章的id 判断是否存在于数据库中
	// Preload 外键查询(Category、 CategoryId 符合对应的外键关系)
	if p.DB.Debug().Preload("Category").Where("id=?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// 给前台回应
	response.Success(ctx, gin.H{"post": post,}, "成功")

}

func (p PostController) MultiplyShows(ctx *gin.Context)  { // 查询多篇文章
	// 获取path参数
	postNumStr := ctx.Query("postNum")
	postNum, _ := strconv.Atoi(postNumStr)
	offsetStr := ctx.Query("offset") // 获取文章数量查询的偏移值
	offset, _ := strconv.Atoi(offsetStr)


	// 定义一个Post结构体类型的切片对象，承接postNum篇文章
	posts := new([]model.Post) // new 返回对应类型的指针

	// 按照查询数据库前postNum篇文章
	if err := p.DB.Debug().Preload("Category").Limit(postNum).Offset(offset * postNum).Find(posts).Error; err != nil {
		response.Fail(ctx, nil, "多篇文章查询失败")
		return
	}

	// 给前台回应
	response.Success(ctx, gin.H{"posts": posts,}, "多篇文章查询成功")


}
