package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"wobbs-server/pkg/jwt"
)

const TokenKey = "Authorization"

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx
		token := c.Request.Header.Get(TokenKey)

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "用户未登录",
			})
			// 所有中间件都存在一个切片中，中间件的执行是由gin通过index移动来控制的
			// 所以return并不能终止其他中间件的继续运行
			//return
			c.Abort() // 终止原理很简单 将index往后移动到一个不可能达到的值
			return
		}
		parts := strings.SplitN(token, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "token格式不正确",
			})
			c.Abort()
			return
		}
		claims, err := jwt.VerifyToken(parts[1])
		if err != nil {
			zap.L().Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "token非法",
			})
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("userId", claims.UserID)

		// 继续执行原有的逻辑
		c.Next()
	}
}
