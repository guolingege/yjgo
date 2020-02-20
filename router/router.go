package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/controller/api"
	"yj-app/app/controller/demo"
	"yj-app/app/controller/hello"
	"yj-app/app/controller/monitor"
	"yj-app/app/controller/system/config"
	"yj-app/app/controller/system/dept"
	"yj-app/app/controller/system/dict_data"
	"yj-app/app/controller/system/dict_type"
	errorc "yj-app/app/controller/system/error"
	"yj-app/app/controller/system/index"
	"yj-app/app/controller/system/menu"
	"yj-app/app/controller/system/post"
	"yj-app/app/controller/system/role"
	"yj-app/app/controller/system/user"
	"yj-app/app/controller/tool"
)

func init() {
	s := g.Server()
	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	//s.SetRewrite("/favicon.ico", "/resource/favicon.ico")
	// 加载登陆路由
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", hello.Hello)
		group.ALL("/login", index.Login)
		group.ALL("/captchaImage", index.CaptchaImage)
		group.ALL("/checklogin", index.CheckLogin)
		group.ALL("/500", errorc.Error)
		group.ALL("/404", errorc.NotFound)
		group.ALL("/403", errorc.Unauth)

		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/index", index.Index)
			group.ALL("/logout", index.Logout)
		})
	})

	// 加载框架路由
	s.Group("/system", func(group *ghttp.RouterGroup) {
		group.ALL("/main", index.Main)
		group.ALL("/switchSkin", index.SwitchSkin)
		group.ALL("/download", index.Download)
	})

	//自定义的业务路由在下面附加绑定
	//api相关路由
	api.LoadRounter(s)
	//工具相关路由
	tool.LoadRounter(s)
	//演示案例相关路由
	demo.LoadRounter(s)
	//系统监控相关路由
	monitor.LoadRounter(s)
	//用户管理相关路由
	user.LoadRounter(s)
	//部门管理相关路由
	dept.LoadRounter(s)
	//角色管理相关路由
	role.LoadRounter(s)
	//菜单管理相关路由
	menu.LoadRounter(s)
	//岗位管理相关路由
	post.LoadRounter(s)
	//系统参数管理相关路由
	config.LoadRounter(s)
	//字典数据管理相关路由
	dict_data.LoadRounter(s)
	//字典类别管理相关路由
	dict_type.LoadRounter(s)

}
