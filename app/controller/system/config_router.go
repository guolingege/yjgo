package system

import (
	"yj-app/app/controller/system/config"
	"yj-app/app/service/middleware/auth"
	"yj-app/app/service/middleware/router"
)

//加载路由
func init() {
	// 参数路由
	g1 := router.New("admin", "/system/config", auth.Auth)
	g1.GET("/", "system:config:view", config.List)
	g1.POST("/list", "system:config:list", config.ListAjax)
	g1.GET("/add", "system:config:add", config.Add)
	g1.POST("/add", "system:config:add", config.AddSave)
	g1.POST("/remove", "system:config:remove", config.Remove)
	g1.GET("/edit", "system:config:edit", config.Edit)
	g1.POST("/edit", "system:config:edit", config.EditSave)
	g1.POST("/export", "system:config:export", config.Export)
	g1.POST("/checkConfigKeyUniqueAll", "system:config:view", config.CheckConfigKeyUniqueAll)
	g1.POST("/checkConfigKeyUnique", "system:config:view", config.CheckConfigKeyUnique)
}
