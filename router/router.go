package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	_ "yj-app/app/controller/api"
	_ "yj-app/app/controller/demo"
	_ "yj-app/app/controller/module"
	_ "yj-app/app/controller/monitor"
	_ "yj-app/app/controller/system"
	_ "yj-app/app/controller/tool"
	"yj-app/app/service/middleware/router"
)

func init() {
	//注册路由
	if len(router.GroupList) > 0 {
		for _, routerGroup := range router.GroupList {
			serverName := ""
			if routerGroup.ServerName == "api" {
				serverName = "api"
			}
			g.Server(serverName).Group(routerGroup.RelativePath, func(group *ghttp.RouterGroup) {
				group.Middleware(routerGroup.Handlers...)
				for _, r := range routerGroup.Router {
					switch r.Method {
					case "ANY":
						group.ALL(r.RelativePath, r.HandlerFunc)
					case "GET":
						group.GET(r.RelativePath, r.HandlerFunc)
					case "POST":
						group.POST(r.RelativePath, r.HandlerFunc)
					case "PUT":
						group.PUT(r.RelativePath, r.HandlerFunc)
					case "HEAD":
						group.HEAD(r.RelativePath, r.HandlerFunc)
					case "PATCH":
						group.PATCH(r.RelativePath, r.HandlerFunc)
					case "DELETE":
						group.DELETE(r.RelativePath, r.HandlerFunc)
					case "OPTIONS":
						group.OPTIONS(r.RelativePath, r.HandlerFunc)
					case "CONNECT":
						group.CONNECT(r.RelativePath, r.HandlerFunc)
					case "TRACE":
						group.TRACE(r.RelativePath, r.HandlerFunc)
					}
				}
			})
		}
	}
}
