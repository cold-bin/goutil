// @author cold bin
// @date 2022/10/22

package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// CorsMiddleware 跨域问题解决
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有的域,也可以指定某几个特定的域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  // 跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		// 处理请求
		c.Next()
	}
}

// RateLimitMiddleware 令牌桶限流 该桶以每fillInterval时段一个令牌的速率填充，直到桶内令牌数达到cap
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit...
		if bucket.TakeAvailable(1) < 1 {
			ResOpInfo(c, "服务繁忙~ 请稍后重试")
			c.Abort()
			return
		}
		c.Next()
	}
}

// ReqBodySizeLimitMiddleware 限制请求体大小，单位为字节数
func ReqBodySizeLimitMiddleware(limit int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 入参不合法时，不做限制
		if limit <= 0 {
			c.Next()
			return
		}

		// 取出请求体的字节大小
		size, err := strconv.Atoi(c.GetHeader("Content-Length"))
		if err != nil {
			ResParamErr(c)
			c.Abort()
			return
		}

		if size > limit {
			ResOpInfo(c, "请求体大小大于自定义阈值~")
			c.Abort()
			return
		}
		c.Next()
	}
}
