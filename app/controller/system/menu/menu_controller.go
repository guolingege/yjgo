package menu

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	menuModel "yj-app/app/model/system/menu"
	menuService "yj-app/app/service/system/menu"
	userService "yj-app/app/service/system/user"
	"yj-app/app/utils/response"
)

//列表页
func List(r *ghttp.Request) {
	response.BuildTpl(r, "system/menu/list.html").WriteTpl()
}

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *menuModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
	}
	rows := make([]menuModel.Entity, 0)
	result, err := menuService.SelectListAll(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	r.Response.WriteJsonExit(rows)
}

//新增页面
func Add(r *ghttp.Request) {
	pid := r.GetQueryInt64("pid")
	var pmenu menuModel.EntityExtend
	pmenu.MenuId = 0
	pmenu.MenuName = "主目录"

	tmp, err := menuService.SelectRecordById(pid)
	if err == nil && tmp != nil && tmp.MenuId > 0 {
		pmenu.MenuId = tmp.MenuId
		pmenu.MenuName = tmp.MenuName
	}
	response.BuildTpl(r, "system/menu/add.html").WriteTpl(g.Map{"menu": pmenu})
}

//新增页面保存
func AddSave(r *ghttp.Request) {
	var req *menuModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
	}

	if menuService.CheckMenuNameUniqueAll(req.MenuName, req.ParentId) == "1" {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg("菜单名称已存在").Log("菜单管理", req).WriteJsonExit()
	}

	id, err := menuService.AddSave(req, r.Session)

	if err != nil || id <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
	}
	response.SucessResp(r).SetBtype(model.Buniss_Add).SetData(id).Log("菜单管理", req).WriteJsonExit()
}

//修改页面
func Edit(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	if id <= 0 {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "参数错误",
		})
		return
	}

	menu, err := menuService.SelectRecordById(id)

	if err != nil || menu == nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "菜单不存在",
		})
		return
	}

	response.BuildTpl(r, "system/menu/edit.html").WriteTpl(g.Map{
		"menu": menu,
	})
}

//修改页面保存
func EditSave(r *ghttp.Request) {
	var req *menuModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
	}

	if menuService.CheckMenuNameUnique(req.MenuName, req.MenuId, req.ParentId) == "1" {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg("菜单名称已存在").Log("菜单管理", req).WriteJsonExit()
	}

	rs, err := menuService.EditSave(req, r.Session)

	if err != nil || rs <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).Log("菜单管理", req).WriteJsonExit()
	}
	response.SucessResp(r).SetBtype(model.Buniss_Edit).SetData(rs).Log("菜单管理", req).WriteJsonExit()
}

//删除数据
func Remove(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	rs := menuService.DeleteRecordById(id)

	if rs {
		response.SucessResp(r).SetBtype(model.Buniss_Del).Log("菜单管理", g.Map{"id": id}).WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).Log("菜单管理", g.Map{"id": id}).WriteJsonExit()
	}
}

//选择菜单树
func SelectMenuTree(r *ghttp.Request) {
	menuId := r.GetQueryInt64("menuId")
	menu, err := menuService.SelectRecordById(menuId)
	if err != nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "菜单不存在",
		})
		return
	}
	response.BuildTpl(r, "system/menu/tree.html").WriteTpl(g.Map{
		"menu": menu,
	})
}

//加载所有菜单列表树
func MenuTreeData(r *ghttp.Request) {
	user := userService.GetProfile(r.Session)
	if user == nil {
		response.ErrorResp(r).SetMsg("登陆超时").Log("菜单管理", g.Map{"userId": user.UserId}).WriteJsonExit()
	}
	ztrees, err := menuService.MenuTreeData(user.UserId)
	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("菜单管理", g.Map{"userId": user.UserId}).WriteJsonExit()
	}
	r.Response.WriteJsonExit(ztrees)
}

//选择图标
func Icon(r *ghttp.Request) {
	response.BuildTpl(r, "system/menu/icon.html").WriteTpl()
}

//加载角色菜单列表树
func RoleMenuTreeData(r *ghttp.Request) {
	roleId := r.GetQueryInt64("roleId")

	user := userService.GetProfile(r.Session)
	if user == nil || user.UserId <= 0 {
		response.ErrorResp(r).SetMsg("登陆超时").Log("菜单管理", g.Map{"roleId": roleId}).WriteJsonExit()
	}

	result, err := menuService.RoleMenuTreeData(roleId, user.UserId)

	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("菜单管理", g.Map{"roleId": roleId}).WriteJsonExit()
	}

	r.Response.WriteJsonExit(result)
}

//检查菜单名是否已经存在不包括本角色
func CheckMenuNameUnique(r *ghttp.Request) {
	var req *menuModel.CheckMenuNameReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := menuService.CheckMenuNameUnique(req.MenuName, req.MenuId, req.ParentId)

	r.Response.WritefExit(result)
}

//检查菜单名是否已经存在
func CheckMenuNameUniqueAll(r *ghttp.Request) {
	var req *menuModel.CheckMenuNameALLReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := menuService.CheckMenuNameUniqueAll(req.MenuName, req.ParentId)

	r.Response.WritefExit(result)
}
