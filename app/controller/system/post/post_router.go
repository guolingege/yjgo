package post

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/middleware"
)

//加载路由
func LoadRounter(s *ghttp.Server) {
	// 岗位路由
	s.Group("/system/post", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.GET("/", List)
		group.POST("/list", ListAjax)
		group.GET("/add", Add)
		group.POST("/add", AddSave)
		group.POST("/remove", Remove)
		group.GET("/edit", Edit)
		group.POST("/edit", EditSave)
		group.ALL("/export", Export)
		group.POST("/checkPostCodeUniqueAll", CheckPostCodeUniqueAll)
		group.POST("/checkPostCodeUnique", CheckPostCodeUnique)
		group.POST("/checkPostNameUniqueAll", CheckPostNameUniqueAll)
		group.POST("/checkPostNameUnique", CheckPostNameUnique)
	})
}