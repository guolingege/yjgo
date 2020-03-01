package tool

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/controller/tool/gen"
	"yj-app/app/service/middleware"
)

//加载路由
func init() {
	s := g.Server()
	// 服务监控
	s.Group("/tool", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.GET("/build", Build)
		group.GET("/swagger", Swagger)
		group.GET("/gen", gen.Gen)
		group.POST("/gen/list", gen.GenList)
		group.ALL("/gen/remove", gen.Remove)
		group.GET("/gen/importTable", gen.ImportTable)
		group.POST("/gen/db/list", gen.DataList)
		group.POST("/gen/importTable", gen.ImportTableSave)
		group.GET("/gen/edit", gen.Edit)
		group.POST("/gen/edit", gen.EditSave)
		group.POST("/gen/column/list", gen.ColumnList)
		group.GET("/gen/preview", gen.Preview)
		group.GET("/gen/genCode", gen.GenCode)
	})
}
