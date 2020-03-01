package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	_ "yj-app/app/controller/api"
	_ "yj-app/app/controller/demo"
	"yj-app/app/controller/hello"
	_ "yj-app/app/controller/module"
	_ "yj-app/app/controller/monitor"
	_ "yj-app/app/controller/system"
	errorc "yj-app/app/controller/system/error"
	"yj-app/app/controller/system/index"
	_ "yj-app/app/controller/tool"
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
}
