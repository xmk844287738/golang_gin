package controller

import "github.com/gin-gonic/gin"

type ResetController interface { // 抽离共性代码
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
}