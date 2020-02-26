package monitor

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/controller/monitor/job"
	"yj-app/app/controller/monitor/logininfor"
	"yj-app/app/controller/monitor/online"
	"yj-app/app/controller/monitor/operlog"
	"yj-app/app/controller/monitor/server"
	"yj-app/app/service/middleware"
)

//加载路由
func LoadRounter(s *ghttp.Server) {
	// 服务监控
	s.Group("/monitor/server", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.GET("/", server.Server)
	})

	//登陆日志
	s.Group("/monitor/logininfor", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.GET("/", logininfor.List)
		group.POST("/list", logininfor.ListAjax)
		group.ALL("/export", logininfor.Export)
		group.POST("/clean", logininfor.Clean)
		group.POST("/remove", logininfor.Remove)
		group.POST("/unlock", logininfor.Unlock)
	})

	//操作日志
	s.Group("/monitor/operlog", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.GET("/", operlog.List)
		group.POST("/list", operlog.ListAjax)
		group.ALL("/export", operlog.Export)
		group.POST("/remove", operlog.Remove)
		group.POST("/clean", operlog.Clean)
		group.GET("/detail", operlog.Detail)
	})

	//在线用户
	s.Group("/monitor/online", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.GET("/", online.List)
		group.POST("/list", online.ListAjax)
		group.POST("/forceLogout", online.ForceLogout)
		group.POST("/batchForceLogout", online.BatchForceLogout)
	})

	//定时任务
	s.Group("/monitor/job", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.GET("/", job.List)
		group.POST("/list", job.ListAjax)
		group.GET("/add", job.Add)
		group.POST("/add", job.AddSave)
		group.GET("/edit", job.Edit)
		group.POST("/edit", job.EditSave)
		group.ALL("/export", job.Export)
		group.POST("/remove", job.Remove)
		group.POST("/detailLog", job.DetailLog)
		group.GET("/detail", job.Detail)
		group.GET("/jobLog", job.LogList)
		group.POST("/jobLog", job.LogListAjax)
		group.POST("/run", job.Start)
		group.POST("/start", job.Start)
		group.POST("/stop", job.Stop)
	})

	//定时任务日志
	s.Group("/monitor/jobLog", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.GET("/", job.LogList)
		group.POST("/list", job.LogListAjax)
		group.ALL("/export", job.Export)
		group.POST("/remove", job.Remove)
	})
}
