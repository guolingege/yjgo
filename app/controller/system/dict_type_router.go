package system

import (
	"yj-app/app/controller/system/dict_type"
	"yj-app/app/service/middleware/auth"
	"yj-app/app/service/middleware/router"
)

//加载路由
func init() {
	// 参数路由
	g1 := router.New("admin", "/system/dict", auth.Auth)
	g1.GET("/", "system:dict:view", dict_type.List)
	g1.POST("/list", "system:dict:list", dict_type.ListAjax)
	g1.GET("/add", "system:dict:add", dict_type.Add)
	g1.POST("/add", "system:dict:add", dict_type.AddSave)
	g1.POST("/remove", "system:dict:remove", dict_type.Remove)
	g1.GET("/remove", "system:dict:remove", dict_type.Remove)
	g1.GET("/edit", "system:dict:edit", dict_type.Edit)
	g1.POST("/edit", "system:dict:edit", dict_type.EditSave)
	g1.GET("/detail", "system:dict:detail", dict_type.Detail)
	g1.POST("/export", "system:dict:export", dict_type.Export)
	g1.POST("/checkDictTypeUniqueAll", "system:dict:view", dict_type.CheckDictTypeUniqueAll)
	g1.POST("/checkDictTypeUnique", "system:dict:view", dict_type.CheckDictTypeUnique)
	g1.GET("/selectDictTree", "system:dict:view", dict_type.SelectDictTree)
	g1.GET("/treeData", "system:dict:view", dict_type.TreeData)
}
