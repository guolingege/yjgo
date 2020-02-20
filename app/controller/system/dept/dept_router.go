package dept

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/middleware"
)

//加载路由
func LoadRounter(s *ghttp.Server) {
	// 分组路由注册方式
	s.Group("/system/dept", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.GET("/", List)
		group.POST("/list", ListAjax)
		group.GET("/add", Add)
		group.POST("/add", AddSave)
		group.ALL("/remove", Remove)
		group.GET("/edit", Edit)
		group.POST("/edit", EditSave)
		group.POST("/checkDeptNameUnique", CheckDeptNameUnique)
		group.POST("/checkDeptNameUniqueAll", CheckDeptNameUniqueAll)
		group.GET("/treeData", TreeData)
		group.GET("/selectDeptTree", SelectDeptTree)
		group.GET("/roleDeptTreeData", RoleDeptTreeData)
	})
}
