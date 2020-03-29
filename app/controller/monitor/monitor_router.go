package monitor

import (
	"yj-app/app/controller/monitor/job"
	"yj-app/app/controller/monitor/logininfor"
	"yj-app/app/controller/monitor/online"
	"yj-app/app/controller/monitor/operlog"
	"yj-app/app/controller/monitor/server"
	"yj-app/app/service/middleware/auth"
	"yj-app/app/service/middleware/router"
)

//加载路由
func init() {
	// 服务监控
	g1 := router.New("admin", "/monitor/server", auth.Auth)
	g1.GET("/", "monitor:server:view", server.Server)

	//登陆日志
	g2 := router.New("admin", "/monitor/logininfor", auth.Auth)
	g2.GET("/", "monitor:logininfor:view", logininfor.List)
	g2.POST("/list", "monitor:logininfor:list", logininfor.ListAjax)
	g2.POST("/export", "monitor:logininfor:export", logininfor.Export)
	g2.POST("/clean", "monitor:logininfor:remove", logininfor.Clean)
	g2.POST("/remove", "monitor:logininfor:remove", logininfor.Remove)
	g2.POST("/unlock", "monitor:logininfor:unlock", logininfor.Unlock)

	//操作日志
	g3 := router.New("admin", "/monitor/operlog", auth.Auth)
	g3.GET("/", "monitor:operlog:view", operlog.List)
	g3.POST("/list", "monitor:operlog:list", operlog.ListAjax)
	g3.POST("/export", "monitor:operlog:export", operlog.Export)
	g3.POST("/remove", "monitor:operlog:export", operlog.Remove)
	g3.POST("/clean", "monitor:operlog:export", operlog.Clean)
	g3.GET("/detail", "monitor:operlog:detail", operlog.Detail)

	//在线用户
	g4 := router.New("admin", "/monitor/online", auth.Auth)
	g4.GET("/", "monitor:online:view", online.List)
	g4.POST("/list", "monitor:online:list", online.ListAjax)
	g4.POST("/forceLogout", "monitor:online:forceLogout", online.ForceLogout)
	g4.POST("/batchForceLogout", "monitor:online:batchForceLogout", online.BatchForceLogout)

	//定时任务
	g5 := router.New("admin", "/monitor/job", auth.Auth)
	g5.GET("/", "monitor:job:view", job.List)
	g5.POST("/list", "monitor:job:list", job.ListAjax)
	g5.GET("/add", "monitor:job:add", job.Add)
	g5.POST("/add", "monitor:job:add", job.AddSave)
	g5.GET("/edit", "monitor:job:edit", job.Edit)
	g5.POST("/edit", "monitor:job:edit", job.EditSave)
	g5.POST("/export", "monitor:job:export", job.Export)
	g5.POST("/remove", "monitor:job:remove", job.Remove)
	g5.POST("/detailLog", "monitor:job:detail", job.DetailLog)
	g5.GET("/detail", "monitor:job:detail", job.Detail)
	g5.GET("/jobLog", "monitor:job:view", job.LogList)
	g5.POST("/jobLog", "monitor:job:list", job.LogListAjax)
	g5.POST("/run", "monitor:job:changeStatus", job.Start)
	g5.POST("/start", "monitor:job:changeStatus", job.Start)
	g5.POST("/stop", "monitor:job:changeStatus", job.Stop)

	//定时任务日志
	g6 := router.New("admin", "/monitor/jobLog", auth.Auth)
	g6.GET("/", "monitor:job:view", job.LogList)
	g6.POST("/list", "monitor:job:list", job.LogListAjax)
	g6.POST("/export", "monitor:job:export", job.Export)
	g6.POST("/remove", "monitor:job:remove", job.Remove)
}
