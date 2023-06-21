package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRouter 将路由注册到给定的 Gin 引擎中。
func RegisterRouter(r *gin.Engine) {
	r.GET("/a", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	rt := Route{ge: r}
	rt.Register("GET", "/b", b)
}

// RouteProcessor 是处理 Gin 上下文并返回字符串的函数类型。
type RouteProcessor func(*gin.Context) string

// Router 定义了注册路由的接口。
type Router interface {
	Register(RouteProcessor)
}

// Route 包含 Gin 引擎的路由结构。
type Route struct {
	ge *gin.Engine
}

// Register 在 Route 中根据指定的方法和路径注册路由处理函数。
func (r *Route) Register(method string, path string, processor RouteProcessor) {
	switch method {
	case "GET":
		r.ge.GET(path, func(ctx *gin.Context) {
			ctx.String(http.StatusOK, processor(ctx))
		})
	}
}

// b 是一个处理函数，返回字符串 "hello ret"。
func b(c *gin.Context) string {
	return "hello ret"
}
