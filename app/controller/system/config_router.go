package system

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/controller/system/config"
	"yj-app/app/service/middleware"
)

//加载路由
func init() {
	s := g.Server()
	// 参数路由
	s.Group("/system/config", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.GET("/", config.List)
		group.POST("/list", config.ListAjax)
		group.GET("/add", config.Add)
		group.POST("/add", config.AddSave)
		group.POST("/remove", config.Remove)
		group.GET("/edit", config.Edit)
		group.POST("/edit", config.EditSave)
		group.ALL("/export", config.Export)
		group.POST("/checkConfigKeyUniqueAll", config.CheckConfigKeyUniqueAll)
		group.POST("/checkConfigKeyUnique", config.CheckConfigKeyUnique)
	})
}
