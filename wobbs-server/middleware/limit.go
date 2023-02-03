package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	rl "go.uber.org/ratelimit"
)

const OnceConsume = 1

// RateLimit 创建指定填充速率和容量大小的令牌桶
func RateLimit(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(ctx *gin.Context) {
		//if bucket.Take(OnceConsume) > 0 // take 返回的是还需要多少时间取到令牌
		//bucket.TakeAvailable(OnceConsume) 返回可用的令牌数
		// 没有足够的令牌 可以选择Sleep等待或者直接返回
		if bucket.TakeAvailable(OnceConsume) < OnceConsume {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "rate limit...",
			})
			ctx.Abort()
			return
		}
		// 可以取到令牌
		ctx.Next()
	}
}

// RateLimit2 漏桶一般用于调用方 这里只是做演示
/**
两者使用场景不同的。
令牌桶一般用来保护自己，对调用方限流，防止自己被打垮；
漏桶用来保护被调用方，因为你不知道被调用方是否有什么保护机制，是否只有自己这个调用方，所以必须保证以确定的漏出速度来调用。
*/
func RateLimit2(rate int) func(c *gin.Context) {
	limiter := rl.New(rate) // rate: 速率 一秒多少滴水
	return func(ctx *gin.Context) {
		// Take() 会阻塞至拿到水滴的时间
		limiter.Take()
		// 可以取到水滴
		ctx.Next()
	}
}
