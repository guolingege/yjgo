package system

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/controller/system/user"
	"yj-app/app/service/middleware"
)

//加载路由
func init() {
	s := g.Server()
	// 用户管理路由
	s.Group("/system/user", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.GET("/", user.List)
		group.POST("/list", user.ListAjax)
		group.GET("/add", user.Add)
		group.POST("/add", user.AddSave)
		group.POST("/remove", user.Remove)
		group.GET("/edit", user.Edit)
		group.POST("/edit", user.EditSave)
		group.ALL("/export", user.Export)
		group.GET("/resetPwd", user.ResetPwd)
		group.POST("/resetPwd", user.ResetPwdSave)
	})
	// 个人中心路由
	s.Group("/system/user/profile", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.ALL("/", user.Profile)
		group.ALL("/avatar", user.Avatar)
		group.ALL("/resetPwd", user.EditPwd)
		group.ALL("/update", user.Update)
		group.ALL("/resetSavePwd", user.UpdatePassword)
		group.ALL("/checkEmailUnique", user.CheckEmailUnique)
		group.ALL("/checkPhoneUnique", user.CheckPhoneUnique)
		group.ALL("/checkLoginNameUnique", user.CheckLoginNameUnique)
		group.ALL("/checkEmailUniqueAll", user.CheckEmailUniqueAll)
		group.ALL("/checkPhoneUniqueAll", user.CheckPhoneUniqueAll)
		group.ALL("/checkPassword", user.CheckPassword)
		group.ALL("/updateAvatar", user.UpdateAvatar)
	})
}
