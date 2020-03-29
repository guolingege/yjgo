package system

import (
	"yj-app/app/controller/hello"
	errorc "yj-app/app/controller/system/error"
	"yj-app/app/controller/system/index"
	"yj-app/app/service/middleware/auth"
	"yj-app/app/service/middleware/router"
)

func init() {
	// 加载登陆路由
	g1 := router.New("admin", "/")
	g1.ANY("/", "", hello.Hello)
	g1.GET("/login", "", index.Login)
	g1.GET("/captchaImage", "", index.CaptchaImage)
	g1.POST("/checklogin", "", index.CheckLogin)
	g1.GET("/500", "", errorc.Error)
	g1.GET("/404", "", errorc.NotFound)
	g1.GET("/403", "", errorc.Unauth)
	g1.GET("/index", "", index.Index)
	g1.GET("/logout", "", index.Logout)

	// 加载框架路由
	g2 := router.New("admin", "/system", auth.Auth)
	g2.GET("/main", "", index.Main)
	g2.GET("/switchSkin", "", index.SwitchSkin)
	g2.GET("/download", "", index.Download)
}
