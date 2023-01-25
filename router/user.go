package router

import (
	"github.com/gin-gonic/gin"
	"wobbs/api"
)

func GetUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/info", api.GetUserInfoByID) //获取其他用户信息
	}
}
