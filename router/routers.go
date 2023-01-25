package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
	"wobbs/config"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(config.LoggerWithZap(zap.L(), time.RFC3339, true))
	r.Use(config.RecoveryWithZap(zap.L(), true))
	v1 := r.Group("/api")
	{
		GetUserRoutes(v1) //用户路由
	}
	zap.L().Info("init router success")
	return r
}
