package role

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	roleModel "yj-app/app/model/system/role"
	userModel "yj-app/app/model/system/user"
	roleService "yj-app/app/service/system/role"
	userService "yj-app/app/service/system/user"
	"yj-app/app/service/utils/response"
)

//列表页
func List(r *ghttp.Request) {
	response.WriteTpl(r, "system/role/list.html")
}

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *roleModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "列表查询", req, model.Buniss_Other, err.Error())
	}
	rows := make([]roleModel.Entity, 0)
	result, page, err := roleService.SelectRecordPage(req)

	if err == nil && result != nil {
		rows = *result
	}

	r.Response.WriteJsonExit(model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: page.Total,
		Rows:  rows,
	})
}

//新增页面
func Add(r *ghttp.Request) {
	response.WriteTpl(r, "system/role/add.html")
}

//新增页面保存
func AddSave(r *ghttp.Request) {
	var req *roleModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "新增角色", req, model.Buniss_Add, err.Error())
	}

	if roleService.CheckRoleNameUniqueAll(req.RoleName) == "1" {
		response.ErrorMsg(r, "新增角色", req, model.Buniss_Add, "角色名称已存在")
	}

	if roleService.CheckRoleKeyUniqueAll(req.RoleKey) == "1" {
		response.ErrorMsg(r, "新增角色", req, model.Buniss_Add, "角色权限已存在")
	}

	rid, err := roleService.AddSave(req, r.Session)

	if err != nil || rid <= 0 {
		response.ErrorAdd(r, "新增角色", req)
	}
	response.SucessDataAdd(r, "新增角色", req, rid)
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

	role, err := roleService.SelectRecordById(id)

	if err != nil || role == nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "角色不存在",
		})
		return
	}

	response.WriteTpl(r, "system/role/edit.html", g.Map{
		"role": role,
	})
}

//修改页面保存
func EditSave(r *ghttp.Request) {
	var req *roleModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "修改角色", req, model.Buniss_Add, err.Error())
	}

	if roleService.CheckRoleNameUnique(req.RoleName, req.RoleId) == "1" {
		response.ErrorMsg(r, "修改角色", req, model.Buniss_Add, "角色名称已存在")
	}

	if roleService.CheckRoleKeyUnique(req.RoleKey, req.RoleId) == "1" {
		response.ErrorMsg(r, "修改角色", req, model.Buniss_Add, "角色权限已存在")
	}

	rs, err := roleService.EditSave(req, r.Session)

	if err != nil || rs <= 0 {
		response.ErrorAdd(r, "修改角色", req)
	}
	response.SucessDataAdd(r, "修改角色", req, rs)
}

//分配用户添加
func SelectUser(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	if id <= 0 {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "参数错误",
		})
		return
	}

	role, err := roleService.SelectRecordById(id)

	if err != nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "参数错误",
		})
	} else {
		response.WriteTpl(r, "system/role/selectUser.html", g.Map{
			"role": role,
		})
	}
}

//获取用户列表
func UnallocatedList(r *ghttp.Request) {
	roleId := r.GetFormInt64("roleId")
	loginName := r.GetFormString("loginName")
	phonenumber := r.GetFormString("phonenumber")
	var rows []userModel.Entity
	userList, err := userService.SelectUnallocatedList(roleId, loginName, phonenumber)

	if err == nil && userList != nil {
		rows = *userList
	}

	r.Response.WriteJsonExit(model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorEdit(r, "删除角色", req)
	}

	rs := roleService.DeleteRecordByIds(req.Ids)

	if rs > 0 {

		response.SucessDataDel(r, "删除角色", req, rs)
	} else {
		response.ErrorDataMsg(r, "删除角色", req, model.Buniss_Del, 0, "未删除数据")
	}
}

//导出
func Export(r *ghttp.Request) {
	var req *roleModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorOther(r, "导出Excel", req)
	}
	url, err := roleService.Export(req)

	if err != nil {
		response.ErrorMsg(r, "导出Excel", req, model.Buniss_Other, err.Error())
	}

	response.SucessMsg(r, "导出Excel", req, model.Buniss_Other, url)
}

//数据权限
func AuthDataScope(r *ghttp.Request) {
	roleId := r.GetQueryInt64("id")
	role, err := roleService.SelectRecordById(roleId)
	if err != nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "参数错误",
		})
	} else {
		response.WriteTpl(r, "system/role/dataScope.html", g.Map{
			"role": role,
		})
	}
}

//数据权限保存
func AuthDataScopeSave(r *ghttp.Request) {
	var req *roleModel.DataScopeReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorOther(r, "数据权限保存", req)
	}
	if !roleService.CheckRoleAllowed(req.RoleId) {
		response.ErrorMsg(r, "数据权限保存", req, model.Buniss_Other, "不允许操作超级管理员角色")
	}

	rs, err := roleService.AuthDataScope(req, r.Session)
	if err != nil || rs <= 0 {
		response.ErrorMsg(r, "数据权限保存", req, model.Buniss_Other, "保存数据失败")
	} else {
		response.SucessOther(r, "数据权限保存", req)
	}
}

//分配用户
func AuthUser(r *ghttp.Request) {
	roleId := r.GetQueryInt64("id")
	role, err := roleService.SelectRecordById(roleId)
	if err != nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "参数错误",
		})
	} else {
		response.WriteTpl(r, "system/role/authUser.html", g.Map{
			"role": role,
		})
	}
}

//查询已分配用户角色列表
func AllocatedList(r *ghttp.Request) {
	roleId := r.GetFormInt64("roleId")
	loginName := r.GetFormString("loginName")
	phonenumber := r.GetFormString("phonenumber")
	var rows []userModel.Entity
	userList, err := userService.SelectAllocatedList(roleId, loginName, phonenumber)

	if err == nil && userList != nil {
		rows = *userList
	}

	r.Response.WriteJsonExit(model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}

//保存角色选择
func SelectAll(r *ghttp.Request) {
	roleId := r.GetFormInt64("roleId")
	userIds := r.GetFormString("userIds")

	if roleId <= 0 {
		response.ErrorMsg(r, "保存分配用户", g.Map{
			"roleId":  roleId,
			"userIds": userIds,
		}, model.Buniss_Add, "参数错误1")
	}
	if userIds == "" {
		response.ErrorMsg(r, "保存分配用户", g.Map{
			"roleId":  roleId,
			"userIds": userIds,
		}, model.Buniss_Add, "参数错误2")
	}

	rs := roleService.InsertAuthUsers(roleId, userIds)
	if rs > 0 {
		response.SucessAdd(r, "保存分配用户", g.Map{
			"roleId":  roleId,
			"userIds": userIds,
		})
	} else {
		response.ErrorMsg(r, "保存分配用户", g.Map{
			"roleId":  roleId,
			"userIds": userIds,
		}, model.Buniss_Add, "保存失败")
	}

}

//取消用户角色授权
func CancelAll(r *ghttp.Request) {
	roleId := r.GetFormInt64("roleId")
	userIds := r.GetFormString("userIds")
	if roleId > 0 && userIds != "" {
		roleService.DeleteUserRoleInfos(roleId, userIds)
		response.SucessDel(r, "取消用户角色授权", g.Map{
			"roleId":  roleId,
			"userIds": userIds,
		})
	} else {
		response.ErrorMsg(r, "取消用户角色授权", g.Map{
			"roleId":  roleId,
			"userIds": userIds,
		}, model.Buniss_Del, "参数错误")
	}
}

//批量取消用户角色授权
func Cancel(r *ghttp.Request) {
	roleId := r.GetFormInt64("roleId")
	userId := r.GetFormInt64("userId")
	if roleId > 0 && userId > 0 {
		roleService.DeleteUserRoleInfo(userId, roleId)
		response.SucessDel(r, "取消用户角色授权", g.Map{
			"roleId": roleId,
			"userId": userId,
		})
	} else {
		response.ErrorMsg(r, "取消用户角色授权", g.Map{
			"roleId": roleId,
			"userId": userId,
		}, model.Buniss_Del, "参数错误")
	}
}

//检查角色是否已经存在不包括本角色
func CheckRoleNameUnique(r *ghttp.Request) {
	var req *roleModel.CheckRoleNameReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := roleService.CheckRoleNameUnique(req.RoleName, req.RoleId)

	r.Response.WritefExit(result)
}

//检查角色是否已经存在
func CheckRoleNameUniqueAll(r *ghttp.Request) {
	var req *roleModel.CheckRoleNameALLReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := roleService.CheckRoleNameUniqueAll(req.RoleName)

	r.Response.WritefExit(result)
}

//检查角色是否已经存在不包括本角色
func CheckRoleKeyUnique(r *ghttp.Request) {
	var req *roleModel.CheckRoleKeyReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := roleService.CheckRoleKeyUnique(req.RoleKey, req.RoleId)

	r.Response.WritefExit(result)
}

//检查角色是否已经存在
func CheckRoleKeyUniqueAll(r *ghttp.Request) {
	var req *roleModel.CheckRoleKeyALLReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := roleService.CheckRoleKeyUniqueAll(req.RoleKey)

	r.Response.WritefExit(result)
}
