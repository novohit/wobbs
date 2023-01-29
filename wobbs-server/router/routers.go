package router

import (
	"time"

	"github.com/gin-gonic/gin"
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
	//r.Use(middleware.GlobalErrors())
	v1 := r.Group("/api")
	{
		GetUserRoutes(v1)      //用户相关路由
		GetCommunityRoutes(v1) // 社区相关路由
	}
	zap.L().Info("init router success")
	return r
}
