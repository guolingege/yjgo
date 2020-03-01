package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	// 参数路由
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.GET("/index", Index)
	})
}
