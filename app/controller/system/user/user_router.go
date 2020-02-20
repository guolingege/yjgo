package user

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/middleware"
)

//加载路由
func LoadRounter(s *ghttp.Server) {
	// 用户管理路由
	s.Group("/system/user", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.GET("/", List)
		group.POST("/list", ListAjax)
		group.GET("/add", Add)
		group.POST("/add", AddSave)
		group.POST("/remove", Remove)
		group.GET("/edit", Edit)
		group.POST("/edit", EditSave)
		group.ALL("/export", Export)
		group.GET("/resetPwd", ResetPwd)
		group.POST("/resetPwd", ResetPwdSave)
	})
	// 个人中心路由
	s.Group("/system/user/profile", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.ALL("/", Profile)
		group.ALL("/avatar", Avatar)
		group.ALL("/resetPwd", EditPwd)
		group.ALL("/update", Update)
		group.ALL("/resetSavePwd", UpdatePassword)
		group.ALL("/checkEmailUnique", CheckEmailUnique)
		group.ALL("/checkPhoneUnique", CheckPhoneUnique)
		group.ALL("/checkLoginNameUnique", CheckLoginNameUnique)
		group.ALL("/checkEmailUniqueAll", CheckEmailUniqueAll)
		group.ALL("/checkPhoneUniqueAll", CheckPhoneUniqueAll)
		group.ALL("/checkPassword", CheckPassword)
		group.ALL("/updateAvatar", UpdateAvatar)
	})
}
