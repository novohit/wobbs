package router

import (
	"time"
	"wobbs-server/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"wobbs-server/config"
)

func InitRouter() *gin.Engine {
	if err := config.InitTranslator("zh"); err != nil {
		zap.L().Error("validator translator init failed")
	}
	r := gin.New()
	r.Use(config.RecoveryWithZap(zap.L(), false))
	r.Use(config.LoggerWithZap(zap.L(), time.RFC3339, true))
	docs.SwaggerInfo.BasePath = "/api"
	//r.Use(middleware.GlobalErrors())
	v1 := r.Group("/api")
	{
		GetUserRoutes(v1)      //用户相关路由
		GetCommunityRoutes(v1) // 社区相关路由
		GetPostRoutes(v1)      // 帖子相关路由
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	zap.L().Info("init router success")

	return r
}
