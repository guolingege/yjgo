package system

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/controller/system/dept"
	"yj-app/app/service/middleware"
)

//加载路由
func init() {
	s := g.Server()
	// 分组路由注册方式
	s.Group("/system/dept", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.GET("/", dept.List)
		group.POST("/list", dept.ListAjax)
		group.GET("/add", dept.Add)
		group.POST("/add", dept.AddSave)
		group.ALL("/remove", dept.Remove)
		group.GET("/edit", dept.Edit)
		group.POST("/edit", dept.EditSave)
		group.POST("/checkDeptNameUnique", dept.CheckDeptNameUnique)
		group.POST("/checkDeptNameUniqueAll", dept.CheckDeptNameUniqueAll)
		group.GET("/treeData", dept.TreeData)
		group.GET("/selectDeptTree", dept.SelectDeptTree)
		group.GET("/roleDeptTreeData", dept.RoleDeptTreeData)
	})
}
