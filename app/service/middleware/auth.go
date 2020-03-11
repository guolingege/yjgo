package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"strings"
	"yj-app/app/model"
	menuService "yj-app/app/service/system/menu"
	userService "yj-app/app/service/system/user"
)

//框架页面只判断是否登陆不做权限判断
var FramePages = []string{"/index", "/system/main", "/system/download", "/system/switchSkin", "/login", "/logout"}

// 鉴权中间件，只有登录成功之后才能通过
func Auth(r *ghttp.Request) {
	//判断是否登陆
	if userService.IsSignedIn(r.Session) {
		//根据url判断是否有权限
		url := r.Request.URL
		if !IsFramePage(url.Path) {
			//获取用户信息
			user := userService.GetProfile(r.Session)
			//获取用户菜单列表
			menus, err := menuService.SelectMenuNormalByUser(user.UserId)
			if err != nil {
				r.Response.RedirectTo("/500")
				return
			}

			if menus == nil {
				r.Response.RedirectTo("/500")
				return
			}

			hasPermission := false

			for i := range menus {
				if strings.EqualFold(menus[i].Url, url.Path) {
					hasPermission = true
					break
				}
			}

			if !hasPermission {
				ajaxString := r.Request.Header.Get("X-Requested-With")
				if strings.EqualFold(ajaxString, "XMLHttpRequest") {
					r.Response.WriteJsonExit(model.CommonRes{
						Code: 403,
						Msg:  "您没有操作权限",
					})
					return
				} else {
					r.Response.RedirectTo("/403")
					return
				}
			}
		}

		r.Middleware.Next()
	} else {
		r.Response.RedirectTo("/login")
	}
}

//判断是否是框架页面
func IsFramePage(path string) bool {
	for i := range FramePages {
		if strings.EqualFold(FramePages[i], path) {
			return true
		}
	}
	return false
}
