package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"wobbs-server/common"
	"wobbs-server/config"
	"wobbs-server/pkg/jwt"
)

const TokenKey = "Authorization"

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx
		token := ctx.Request.Header.Get(TokenKey)

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "用户未登录",
			})
			// 所有中间件都存在一个切片中，中间件的执行是由gin通过index移动来控制的
			// 所以return并不能终止其他中间件的继续运行
			//return
			ctx.Abort() // 终止原理很简单 将index往后移动到一个不可能达到的值
			return
		}
		parts := strings.SplitN(token, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "token格式不正确",
			})
			ctx.Abort()
			return
		}
		claims, err := jwt.VerifyToken(parts[1])
		if err != nil {
			zap.L().Error(err.Error())
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "登录过期",
			})
			ctx.Abort()
			return
		}
		result, err := config.RDB.Get(context.Background(), common.KeyUserTokenPrefix+strconv.FormatInt(claims.UserID, 10)+":"+ctx.RemoteIP()).Result()
		// key不存在说明未登录 或者header中的token和redis中的token不一样 说明存在同一ip下同一用户有多次登录
		if err != nil || result != parts[1] {
			if err != nil {
				zap.L().Error(err.Error())
			} else {
				zap.L().Info(fmt.Sprintf("用户:[%d] IP:[%s] 同一时间登录多次", claims.UserID, ctx.RemoteIP()))
			}
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "登录过期",
			})
			ctx.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文c上
		ctx.Set("userId", claims.UserID)

		// 继续执行原有的逻辑
		ctx.Next()
	}
}
