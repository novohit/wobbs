package middleware

import (
	"wobbs/common"

	"github.com/gin-gonic/gin"
)

func GlobalErrors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//打印错误堆栈信息
				//debug.PrintStack()
				//封装通用json返回
				common.FailByMsg(ctx, errorToString(r))
				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				ctx.Abort()
			}
		}()
		//加载完 defer recover，继续后续接口调用
		ctx.Next()
	}
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	case string:
		return r.(string)
	default:
		return ""
	}
}
