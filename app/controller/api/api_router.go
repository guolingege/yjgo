package api

import (
	"github.com/gogf/gf/net/ghttp"
)

//加载路由
func LoadRounter(s *ghttp.Server) {
	// 参数路由
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.GET("/index", Index)
	})
}