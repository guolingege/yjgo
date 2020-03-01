package system

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/controller/system/dict_type"
	"yj-app/app/service/middleware"
)

//加载路由
func init() {
	s := g.Server()
	// 参数路由
	s.Group("/system/dict", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.GET("/", dict_type.List)
		group.POST("/list", dict_type.ListAjax)
		group.GET("/add", dict_type.Add)
		group.POST("/add", dict_type.AddSave)
		group.POST("/remove", dict_type.Remove)
		group.GET("/edit", dict_type.Edit)
		group.POST("/edit", dict_type.EditSave)
		group.GET("/detail", dict_type.Detail)
		group.ALL("/export", dict_type.Export)
		group.POST("/checkDictTypeUniqueAll", dict_type.CheckDictTypeUniqueAll)
		group.POST("/checkDictTypeUnique", dict_type.CheckDictTypeUnique)
		group.GET("/selectDictTree", dict_type.SelectDictTree)
		group.ALL("/treeData", dict_type.TreeData)
	})
}
