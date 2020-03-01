package system

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/controller/system/dict_data"
	"yj-app/app/service/middleware"
)

//加载路由
func init() {
	s := g.Server()
	// 参数路由
	s.Group("/system/dict/data", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.POST("/list", dict_data.ListAjax)
		group.GET("/add", dict_data.Add)
		group.POST("/add", dict_data.AddSave)
		group.POST("/remove", dict_data.Remove)
		group.GET("/edit", dict_data.Edit)
		group.POST("/edit", dict_data.EditSave)
		group.ALL("/export", dict_data.Export)
	})
}
