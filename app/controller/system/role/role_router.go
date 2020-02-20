package role

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/middleware"
)

//加载路由
func LoadRounter(s *ghttp.Server) {
	// 角色路由
	s.Group("/system/role", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.GET("/", List)
		group.POST("/list", ListAjax)
		group.GET("/add", Add)
		group.POST("/add", AddSave)
		group.POST("/remove", Remove)
		group.GET("/edit", Edit)
		group.POST("/edit", EditSave)
		group.ALL("/export", Export)
		group.POST("/checkRoleKeyUnique", CheckRoleKeyUnique)
		group.POST("/checkRoleNameUniqueAll", CheckRoleNameUniqueAll)
		group.POST("/checkRoleNameUnique", CheckRoleNameUnique)
		group.POST("/checkRoleKeyUniqueAll", CheckRoleKeyUniqueAll)
		group.GET("/authDataScope", AuthDataScope)
		group.POST("/authDataScope", AuthDataScopeSave)
		group.GET("/authUser", AuthUser)
		group.POST("/allocatedList", AllocatedList)
		group.GET("/selectUser", SelectUser)
		group.POST("/unallocatedList", UnallocatedList)
		group.POST("/selectAll", SelectAll)
		group.ALL("/cancel", Cancel)
		group.ALL("/cancelAll", CancelAll)
	})
}
