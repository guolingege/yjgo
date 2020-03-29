package role

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	roleModel "yj-app/app/model/system/role"
	userModel "yj-app/app/model/system/user"
	roleService "yj-app/app/service/system/role"
	userService "yj-app/app/service/system/user"
	"yj-app/app/utils/response"
)

//列表页
func List(r *ghttp.Request) {
	response.BuildTpl(r, "system/role/list.html").WriteTpl()
}

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *roleModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
	}
	rows := make([]roleModel.Entity, 0)
	result, page, err := roleService.SelectRecordPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	response.BuildTable(r, page.Total, rows).WriteJsonExit()
}

//新增页面
func Add(r *ghttp.Request) {
	response.BuildTpl(r, "system/role/add.html").WriteTpl()
}

//新增页面保存
func AddSave(r *ghttp.Request) {
	var req *roleModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
	}

	if roleService.CheckRoleNameUniqueAll(req.RoleName) == "1" {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg("角色名称已存在").Log("角色管理", req).WriteJsonExit()
	}

	if roleService.CheckRoleKeyUniqueAll(req.RoleKey) == "1" {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg("角色权限已存在").Log("角色管理", req).WriteJsonExit()
	}

	rid, err := roleService.AddSave(req, r.Session)

	if err != nil || rid <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).Log("角色管理", req).WriteJsonExit()
	}
	response.SucessResp(r).SetData(rid).SetBtype(model.Buniss_Add).Log("角色管理", req).WriteJsonExit()
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

	role, err := roleService.SelectRecordById(id)

	if err != nil || role == nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "角色不存在",
		})
		return
	}

	response.BuildTpl(r, "system/role/edit.html").WriteTpl(g.Map{
		"role": role,
	})
}

//修改页面保存
func EditSave(r *ghttp.Request) {
	var req *roleModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
	}

	if roleService.CheckRoleNameUnique(req.RoleName, req.RoleId) == "1" {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg("角色名称已存在").Log("角色管理", req).WriteJsonExit()
	}

	if roleService.CheckRoleKeyUnique(req.RoleKey, req.RoleId) == "1" {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg("角色权限已存在").Log("角色管理", req).WriteJsonExit()
	}

	rs, err := roleService.EditSave(req, r.Session)

	if err != nil || rs <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).Log("角色管理", req).WriteJsonExit()
	}
	response.SucessResp(r).SetBtype(model.Buniss_Edit).SetData(rs).Log("角色管理", req).WriteJsonExit()
}

//分配用户添加
func SelectUser(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	if id <= 0 {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "参数错误",
		})
		return
	}

	role, err := roleService.SelectRecordById(id)

	if err != nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "角色不存在",
		})
	} else {
		response.BuildTpl(r, "system/role/selectUser.html").WriteTpl(g.Map{
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

	if err == nil && len(userList) > 0 {
		rows = userList
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
		response.ErrorResp(r).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
	}

	rs := roleService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(r).SetBtype(model.Buniss_Del).SetData(rs).Log("角色管理", req).WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).Log("角色管理", req).WriteJsonExit()
	}
}

//导出
func Export(r *ghttp.Request) {
	var req *roleModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
	}
	url, err := roleService.Export(req)

	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
	}
	response.SucessResp(r).SetMsg(url).Log("角色管理", req).WriteJsonExit()
}

//数据权限
func AuthDataScope(r *ghttp.Request) {
	roleId := r.GetQueryInt64("id")
	role, err := roleService.SelectRecordById(roleId)
	if err != nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "角色不存在",
		})
	} else {
		response.BuildTpl(r, "system/role/dataScope.html").WriteTpl(g.Map{
			"role": role,
		})
	}
}

//数据权限保存
func AuthDataScopeSave(r *ghttp.Request) {
	var req *roleModel.DataScopeReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
	}
	if !roleService.CheckRoleAllowed(req.RoleId) {
		response.ErrorResp(r).SetMsg("不允许操作超级管理员角色").Log("角色管理", req).WriteJsonExit()
	}

	rs, err := roleService.AuthDataScope(req, r.Session)
	if err != nil || rs <= 0 {
		response.ErrorResp(r).SetMsg("保存数据失败").SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
	} else {
		response.SucessResp(r).Log("角色管理", req).WriteJsonExit()
	}
}

//分配用户
func AuthUser(r *ghttp.Request) {
	roleId := r.GetQueryInt64("id")
	role, err := roleService.SelectRecordById(roleId)
	if err != nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "角色不存在",
		})
	} else {
		response.BuildTpl(r, "system/role/authUser.html").WriteTpl(g.Map{
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

	if err == nil && len(userList) > 0 {
		rows = userList
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
		response.ErrorResp(r).SetMsg("参数错误1").SetBtype(model.Buniss_Add).Log("角色管理", g.Map{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	}
	if userIds == "" {
		response.ErrorResp(r).SetMsg("参数错误2").SetBtype(model.Buniss_Add).Log("角色管理", g.Map{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	}

	rs := roleService.InsertAuthUsers(roleId, userIds)
	if rs > 0 {
		response.SucessResp(r).SetBtype(model.Buniss_Add).Log("角色管理", g.Map{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).Log("角色管理", g.Map{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	}

}

//取消用户角色授权
func CancelAll(r *ghttp.Request) {
	roleId := r.GetFormInt64("roleId")
	userIds := r.GetFormString("userIds")
	if roleId > 0 && userIds != "" {
		roleService.DeleteUserRoleInfos(roleId, userIds)
		response.SucessResp(r).SetBtype(model.Buniss_Del).Log("角色管理", g.Map{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).SetMsg("参数错误").Log("角色管理", g.Map{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	}
}

//批量取消用户角色授权
func Cancel(r *ghttp.Request) {
	roleId := r.GetFormInt64("roleId")
	userId := r.GetFormInt64("userId")
	if roleId > 0 && userId > 0 {
		roleService.DeleteUserRoleInfo(userId, roleId)
		response.SucessResp(r).SetBtype(model.Buniss_Del).Log("角色管理", g.Map{
			"roleId": roleId,
			"userId": userId,
		}).WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).SetMsg("参数错误").Log("角色管理", g.Map{
			"roleId": roleId,
			"userId": userId,
		}).WriteJsonExit()
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
