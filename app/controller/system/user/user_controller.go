package user

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	postModel "yj-app/app/model/system/post"
	roleModel "yj-app/app/model/system/role"
	userModel "yj-app/app/model/system/user"
	deptServic "yj-app/app/service/system/dept"
	postService "yj-app/app/service/system/post"
	roleService "yj-app/app/service/system/role"
	userService "yj-app/app/service/system/user"
	"yj-app/app/utils/response"
)

//用户列表页
func List(r *ghttp.Request) {
	response.BuildTpl(r, "system/user/list.html").WriteTpl()
}

//用户列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *userModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
	}
	rows := make([]userModel.UserListEntity, 0)
	result, page, err := userService.SelectRecordList(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	response.BuildTable(r, page.Total, rows).WriteJsonExit()
}

//用户新增页面
func Add(r *ghttp.Request) {
	var paramsRole *roleModel.SelectPageReq
	var paramsPost *postModel.SelectPageReq

	roles := make([]roleModel.EntityFlag, 0)
	posts := make([]postModel.EntityFlag, 0)

	rolesP, _ := roleService.SelectRecordAll(paramsRole)

	if len(rolesP) > 0 {
		roles = rolesP
	}

	postP, _ := postService.SelectListAll(paramsPost)

	if len(postP) > 0 {
		posts = postP
	}
	response.BuildTpl(r, "system/user/add.html").WriteTpl(g.Map{
		"roles": roles,
		"posts": posts,
	})
}

//保存新增用户数据
func AddSave(r *ghttp.Request) {
	var req *userModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("新增用户", req).WriteJsonExit()
	}

	//判断登陆名是否已注册
	isHadName := userService.CheckLoginName(req.LoginName)
	if isHadName {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg("登陆名已经存在").Log("新增用户", req).WriteJsonExit()
	}

	//判断手机号码是否已注册
	isHadPhone := userService.CheckPhoneUniqueAll(req.Phonenumber)
	if isHadPhone {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg("手机号码已经存在").Log("新增用户", req).WriteJsonExit()
	}

	//判断邮箱是否已注册
	isHadEmail := userService.CheckEmailUniqueAll(req.Email)
	if isHadEmail {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg("邮箱已经存在").Log("新增用户", req).WriteJsonExit()
	}

	uid, err := userService.AddSave(req, r.Session)

	if err != nil || uid <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).Log("新增用户", req).WriteJsonExit()
	}
	response.SucessResp(r).SetData(uid).SetBtype(model.Buniss_Add).Log("新增用户", req).WriteJsonExit()
}

//用户修改页面
func Edit(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	if id <= 0 {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "参数错误",
		})
		return
	}

	user, err := userService.SelectRecordById(id)

	if err != nil || user == nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "用户不存在",
		})
		return
	}

	//获取部门信息
	deptName := ""
	if user.DeptId > 0 {
		dept := deptServic.SelectDeptById(user.DeptId)
		if dept != nil {
			deptName = dept.DeptName
		}
	}

	roles := make([]roleModel.EntityFlag, 0)
	posts := make([]postModel.EntityFlag, 0)

	rolesP, _ := roleService.SelectRoleContactVo(id)

	if len(rolesP) > 0 {
		roles = rolesP
	}

	postP, _ := postService.SelectPostsByUserId(id)

	if len(postP) > 0 {
		posts = postP
	}

	response.BuildTpl(r, "system/user/edit.html").WriteTpl(g.Map{
		"user":     user,
		"deptName": deptName,
		"roles":    roles,
		"posts":    posts,
	})
}

//重置密码
func ResetPwd(r *ghttp.Request) {
	id := r.GetQueryInt64("userId")

	if id <= 0 {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "参数错误",
		})
		return
	}

	user, err := userService.SelectRecordById(id)

	if err != nil || user == nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "用户不存在",
		})
		return
	}
	response.BuildTpl(r, "system/user/resetPwd.html").WriteTpl(g.Map{
		"user": user,
	})
}

//重置密码保存
func ResetPwdSave(r *ghttp.Request) {
	var req *userModel.ResetPwdReq
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("重置密码", req).WriteJsonExit()
	}

	result, err := userService.ResetPassword(req)

	if err != nil || !result {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("重置密码", req).WriteJsonExit()
	} else {
		response.SucessResp(r).SetBtype(model.Buniss_Edit).Log("重置密码", req).WriteJsonExit()
	}
}

//保存修改用户数据
func EditSave(r *ghttp.Request) {
	var req *userModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("修改用户", req).WriteJsonExit()
	}

	//判断手机号码是否已注册
	isHadPhone := userService.CheckPhoneUnique(req.UserId, req.Phonenumber)
	if isHadPhone {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg("手机号码已经存在").Log("修改用户", req).WriteJsonExit()
	}

	//判断邮箱是否已注册
	isHadEmail := userService.CheckEmailUnique(req.UserId, req.Email)
	if isHadEmail {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg("邮箱已经存在").Log("修改用户", req).WriteJsonExit()
	}

	uid, err := userService.EditSave(req, r.Session)

	if err != nil || uid <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).Log("修改用户", req).WriteJsonExit()
	}

	response.SucessResp(r).SetData(uid).SetBtype(model.Buniss_Edit).Log("修改用户", req).WriteJsonExit()
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("删除用户", req).WriteJsonExit()
	}

	rs := userService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(r).SetData(rs).SetBtype(model.Buniss_Del).Log("删除用户", req).WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).Log("删除用户", req).WriteJsonExit()
	}
}

//导出
func Export(r *ghttp.Request) {
	var req *userModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("导出Excel", req).WriteJsonExit()
	}
	url, err := userService.Export(req)

	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("导出Excel", req).WriteJsonExit()
	}
	response.SucessResp(r).SetMsg(url).Log("导出Excel", req).WriteJsonExit()
}
