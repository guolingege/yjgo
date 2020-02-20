package config

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/middleware"
)

//加载路由
func LoadRounter(s *ghttp.Server) {
	// 参数路由
	s.Group("/system/config", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.GET("/", List)
		group.POST("/list", ListAjax)
		group.GET("/add", Add)
		group.POST("/add", AddSave)
		group.POST("/remove", Remove)
		group.GET("/edit", Edit)
		group.POST("/edit", EditSave)
		group.ALL("/export", Export)
		group.POST("/checkConfigKeyUniqueAll", CheckConfigKeyUniqueAll)
		group.POST("/checkConfigKeyUnique", CheckConfigKeyUnique)
	})
}
