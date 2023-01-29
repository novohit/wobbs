package router

import (
	"github.com/gin-gonic/gin"

	"wobbs-server/api"
)

func GetCommunityRoutes(router *gin.RouterGroup) {
	communityGroup := router.Group("/community")
	{
		communityGroup.GET("/category", api.GetCommunityCategory) //获取社区分类
	}
}
