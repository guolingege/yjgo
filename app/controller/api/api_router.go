package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/controller/api/login"
	"yj-app/app/service/middleware"
)

func init() {
	s := g.Server("api")
	// 不需要鉴权
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.POST("/login", login.Login)
	})
	// 需要鉴权
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.JWT)
		group.GET("/index", login.Index)
	})
}
