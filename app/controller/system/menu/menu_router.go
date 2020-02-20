package menu

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/middleware"
)

//加载路由
func LoadRounter(s *ghttp.Server) {
	// 角色路由
	s.Group("/system/menu", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.GET("/", List)
		group.POST("/list", ListAjax)
		group.GET("/add", Add)
		group.POST("/add", AddSave)
		group.ALL("/remove", Remove)
		group.GET("/edit", Edit)
		group.POST("/edit", EditSave)
		group.GET("/icon", Icon)
		group.GET("/selectMenuTree", SelectMenuTree)
		group.ALL("/roleMenuTreeData", RoleMenuTreeData)
		group.ALL("/menuTreeData", MenuTreeData)
		group.POST("/checkMenuNameUnique", CheckMenuNameUnique)
		group.POST("/checkMenuNameUniqueAll", CheckMenuNameUniqueAll)
	})
}
