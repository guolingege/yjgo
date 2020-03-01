package system

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/controller/system/post"
	"yj-app/app/service/middleware"
)

//加载路由
func init() {
	s := g.Server()
	// 岗位路由
	s.Group("/system/post", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.GET("/", post.List)
		group.POST("/list", post.ListAjax)
		group.GET("/add", post.Add)
		group.POST("/add", post.AddSave)
		group.POST("/remove", post.Remove)
		group.GET("/edit", post.Edit)
		group.POST("/edit", post.EditSave)
		group.ALL("/export", post.Export)
		group.POST("/checkPostCodeUniqueAll", post.CheckPostCodeUniqueAll)
		group.POST("/checkPostCodeUnique", post.CheckPostCodeUnique)
		group.POST("/checkPostNameUniqueAll", post.CheckPostNameUniqueAll)
		group.POST("/checkPostNameUnique", post.CheckPostNameUnique)
	})
}
