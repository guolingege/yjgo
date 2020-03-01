package system

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/controller/system/menu"
	"yj-app/app/service/middleware"
)

//加载路由
func init() {
	s := g.Server()
	// 角色路由
	s.Group("/system/menu", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.GET("/", menu.List)
		group.POST("/list", menu.ListAjax)
		group.GET("/add", menu.Add)
		group.POST("/add", menu.AddSave)
		group.ALL("/remove", menu.Remove)
		group.GET("/edit", menu.Edit)
		group.POST("/edit", menu.EditSave)
		group.GET("/icon", menu.Icon)
		group.GET("/selectMenuTree", menu.SelectMenuTree)
		group.ALL("/roleMenuTreeData", menu.RoleMenuTreeData)
		group.ALL("/menuTreeData", menu.MenuTreeData)
		group.POST("/checkMenuNameUnique", menu.CheckMenuNameUnique)
		group.POST("/checkMenuNameUniqueAll", menu.CheckMenuNameUniqueAll)
	})
}
