package system

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/controller/system/role"
	"yj-app/app/service/middleware"
)

//加载路由
func init() {
	s := g.Server()
	// 角色路由
	s.Group("/system/role", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.GET("/", role.List)
		group.POST("/list", role.ListAjax)
		group.GET("/add", role.Add)
		group.POST("/add", role.AddSave)
		group.POST("/remove", role.Remove)
		group.GET("/edit", role.Edit)
		group.POST("/edit", role.EditSave)
		group.ALL("/export", role.Export)
		group.POST("/checkRoleKeyUnique", role.CheckRoleKeyUnique)
		group.POST("/checkRoleNameUniqueAll", role.CheckRoleNameUniqueAll)
		group.POST("/checkRoleNameUnique", role.CheckRoleNameUnique)
		group.POST("/checkRoleKeyUniqueAll", role.CheckRoleKeyUniqueAll)
		group.GET("/authDataScope", role.AuthDataScope)
		group.POST("/authDataScope", role.AuthDataScopeSave)
		group.GET("/authUser", role.AuthUser)
		group.POST("/allocatedList", role.AllocatedList)
		group.GET("/selectUser", role.SelectUser)
		group.POST("/unallocatedList", role.UnallocatedList)
		group.POST("/selectAll", role.SelectAll)
		group.ALL("/cancel", role.Cancel)
		group.ALL("/cancelAll", role.CancelAll)
	})
}
