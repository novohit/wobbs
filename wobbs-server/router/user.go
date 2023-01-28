package router

import (
	"github.com/gin-gonic/gin"
	"wobbs-server/middleware"

	"wobbs-server/api"
)

func GetUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/info", middleware.AuthRequired(), api.GetUserInfoByID) //获取其他用户信息
		userGroup.POST("/register", api.Register)                              // 用户注册
		userGroup.POST("/login", api.Login)                                    // 用户登录
	}
}
