package auth

import (
	"github.com/gogf/gf/net/ghttp"
	"strings"
	"yj-app/app/model"
	"yj-app/app/service/middleware/router"
	menuService "yj-app/app/service/system/menu"
	userService "yj-app/app/service/system/user"
)

// 鉴权中间件，只有登录成功之后才能通过
func Auth(r *ghttp.Request) {
	//判断是否登陆
	if userService.IsSignedIn(r.Session) {
		//根据url判断是否有权限
		url := r.Request.URL
		//获取权限标识
		permission := router.FindPermission(url.Path)
		if len(permission) > 0 {
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
				if strings.EqualFold(menus[i].Perms, permission) {
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
