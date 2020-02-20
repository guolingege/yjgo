package menu

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	menuModel "yj-app/app/model/system/menu"
	menuService "yj-app/app/service/system/menu"
	userService "yj-app/app/service/system/user"
	"yj-app/app/service/utils/response"
)

//列表页
func List(r *ghttp.Request) {
	response.WriteTpl(r, "system/menu/list.html")
}

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *menuModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "列表查询", req, model.Buniss_Other, err.Error())
	}
	rows := make([]menuModel.Entity, 0)
	result, err := menuService.SelectListAll(req)

	if err == nil && result != nil {
		rows = *result
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

	response.WriteTpl(r, "system/menu/add.html", g.Map{"menu": pmenu})
}

//新增页面保存
func AddSave(r *ghttp.Request) {
	var req *menuModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "新增菜单", req, model.Buniss_Add, err.Error())
	}

	if menuService.CheckMenuNameUniqueAll(req.MenuName, req.ParentId) == "1" {
		response.ErrorMsg(r, "新增菜单", req, model.Buniss_Add, "菜单名称已存在")
	}

	rid, err := menuService.AddSave(req, r.Session)

	if err != nil || rid <= 0 {
		response.ErrorAdd(r, "新增菜单", req)
	}
	response.SucessDataAdd(r, "新增菜单", req, rid)
}

//修改页面
func Edit(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	if id <= 0 {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "参数错误",
		})
		return
	}

	menu, err := menuService.SelectRecordById(id)

	if err != nil || menu == nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "菜单不存在",
		})
		return
	}

	response.WriteTpl(r, "system/menu/edit.html", g.Map{
		"menu": menu,
	})
}

//修改页面保存
func EditSave(r *ghttp.Request) {
	var req *menuModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "修改菜单", req, model.Buniss_Add, err.Error())
	}

	if menuService.CheckMenuNameUnique(req.MenuName, req.MenuId, req.ParentId) == "1" {
		response.ErrorMsg(r, "修改菜单", req, model.Buniss_Add, "菜单名称已存在")
	}

	rs, err := menuService.EditSave(req, r.Session)

	if err != nil || rs <= 0 {
		response.ErrorAdd(r, "修改菜单", req)
	}
	response.SucessDataAdd(r, "修改菜单", req, rs)
}

//删除数据
func Remove(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	rs := menuService.DeleteRecordById(id)

	if rs {
		response.SucessDataDel(r, "删除菜单", g.Map{"id": id}, rs)
	} else {
		response.ErrorDataMsg(r, "删除菜单", g.Map{"id": id}, model.Buniss_Del, 0, "未删除数据")
	}
}

//选择菜单树
func SelectMenuTree(r *ghttp.Request) {
	menuId := r.GetQueryInt64("menuId")
	menu, err := menuService.SelectRecordById(menuId)
	if err != nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "参数错误",
		})
		return
	}
	response.WriteTpl(r, "system/menu/tree.html", g.Map{
		"menu": menu,
	})
}

//加载所有菜单列表树
func MenuTreeData(r *ghttp.Request) {
	user := userService.GetProfile(r.Session)
	if user == nil {
		response.ErrorMsg(r, "加载菜单列表树", g.Map{"userId": user.UserId,}, model.Buniss_Other, "登陆超时")
	}
	ztrees, err := menuService.MenuTreeData(user.UserId)
	if err != nil {
		response.ErrorMsg(r, "加载菜单列表树", g.Map{"userId": user.UserId,}, model.Buniss_Other, err.Error())
	}
	r.Response.WriteJsonExit(ztrees)
}

//选择图标
func Icon(r *ghttp.Request) {
	response.WriteTpl(r, "system/menu/icon.html")
}

//加载角色菜单列表树
func RoleMenuTreeData(r *ghttp.Request) {
	roleId := r.GetQueryInt64("roleId")

	user := userService.GetProfile(r.Session)
	if user == nil || user.UserId <= 0 {
		response.ErrorMsg(r, "菜单树", g.Map{"roleId": roleId}, model.Buniss_Other, "登陆超时")
	}

	result, err := menuService.RoleMenuTreeData(roleId, user.UserId)

	if err != nil {
		response.ErrorMsg(r, "菜单树", g.Map{"roleId": roleId}, model.Buniss_Other, err.Error())
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
