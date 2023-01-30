package router

import (
	"github.com/gin-gonic/gin"
	"wobbs-server/middleware"

	"wobbs-server/api"
)

func GetPostRoutes(router *gin.RouterGroup) {
	communityGroup := router.Group("/post")
	{
		communityGroup.GET("/:pid", api.GetPostDetail) //获取帖子详情
		communityGroup.GET("", api.GetPostList)        //获取帖子列表
		communityGroup.Use(middleware.AuthRequired())
		communityGroup.POST("", api.CreatePost) //创建帖子
	}
}
