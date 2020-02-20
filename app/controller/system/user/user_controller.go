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
	"yj-app/app/service/utils/response"
)

//用户列表页
func List(r *ghttp.Request) {
	response.WriteTpl(r, "system/user/list.html")
}

//用户列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *userModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "用户列表查询", req, model.Buniss_Other, err.Error())
	}
	rows := make([]userModel.UserListEntity, 0)
	result, page, err := userService.SelectRecordList(req)

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

//用户新增页面
func Add(r *ghttp.Request) {
	var paramsRole *roleModel.SelectPageReq
	var paramsPost *postModel.SelectPageReq

	roles := make([]roleModel.EntityFlag, 0)
	posts := make([]postModel.EntityFlag, 0)

	rolesP, _ := roleService.SelectRecordAll(paramsRole)

	if rolesP != nil {
		roles = *rolesP
	}

	postP, _ := postService.SelectListAll(paramsPost)

	if postP != nil {
		posts = * postP
	}

	response.WriteTpl(r, "system/user/add.html", g.Map{
		"roles": roles,
		"posts": posts,
	})
}

//保存新增用户数据
func AddSave(r *ghttp.Request) {
	var req *userModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "新增用户", req, model.Buniss_Add, err.Error())
	}

	//判断登陆名是否已注册
	isHadName := userService.CheckLoginName(req.LoginName)
	if isHadName {
		response.ErrorMsg(r, "新增用户", req, model.Buniss_Add, "登陆名已经存在")
	}

	//判断手机号码是否已注册
	isHadPhone := userService.CheckPhoneUniqueAll(req.Phonenumber)
	if isHadPhone {
		response.ErrorMsg(r, "新增用户", req, model.Buniss_Add, "手机号码已经存在")
	}

	//判断邮箱是否已注册
	isHadEmail := userService.CheckEmailUniqueAll(req.Email)
	if isHadEmail {
		response.ErrorMsg(r, "新增用户", req, model.Buniss_Add, "邮箱已经存在")
	}

	uid, err := userService.AddSave(req, r.Session)

	if err != nil || uid <= 0 {
		response.ErrorAdd(r, "新增用户", req)
	}
	response.SucessDataAdd(r, "新增用户", req, uid)
}

//用户修改页面
func Edit(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	if id <= 0 {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "参数错误",
		})
		return
	}

	user, err := userService.SelectRecordById(id)

	if err != nil || user == nil {
		response.WriteTpl(r, "error/error.html", g.Map{
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

	if rolesP != nil {
		roles = *rolesP
	}

	postP, _ := postService.SelectPostsByUserId(id)

	if postP != nil {
		posts = * postP
	}

	response.WriteTpl(r, "system/user/edit.html", g.Map{
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
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "参数错误",
		})
		return
	}

	user, err := userService.SelectRecordById(id)

	if err != nil || user == nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "用户不存在",
		})
		return
	}
	response.WriteTpl(r, "system/user/resetPwd.html", g.Map{
		"user": user,
	})
}

//重置密码保存
func ResetPwdSave(r *ghttp.Request) {
	var req *userModel.ResetPwdReq
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "重置密码", req, model.Buniss_Edit, err.Error())
	}

	result, err := userService.ResetPassword(req)

	if err != nil || !result {
		response.ErrorMsg(r, "重置密码", req, model.Buniss_Edit, err.Error())
	} else {
		response.SucessEdit(r, "重置密码", req)
	}
}

//保存修改用户数据
func EditSave(r *ghttp.Request) {
	var req *userModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorEdit(r, "修改用户", req)
	}

	//判断手机号码是否已注册
	isHadPhone := userService.CheckPhoneUnique(req.UserId, req.Phonenumber)
	if isHadPhone {
		response.ErrorMsg(r, "修改用户", req, model.Buniss_Edit, "手机号码已经存在")
	}

	//判断邮箱是否已注册
	isHadEmail := userService.CheckEmailUnique(req.UserId, req.Email)
	if isHadEmail {
		response.ErrorMsg(r, "修改用户", req, model.Buniss_Edit, "邮箱已经存在")
	}

	uid, err := userService.EditSave(req, r.Session)

	if err != nil || uid <= 0 {
		response.ErrorEdit(r, "修改用户", req)
	}

	response.SucessDataEdit(r, "修改用户", req, uid)
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorEdit(r, "删除用户", req)
	}

	rs := userService.DeleteRecordByIds(req.Ids)

	if rs > 0 {

		response.SucessDataDel(r, "删除用户", req, rs)
	} else {
		response.ErrorDataMsg(r, "删除用户", req, model.Buniss_Del, 0, "未删除数据")
	}
}

//导出
func Export(r *ghttp.Request) {
	var req *userModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(model.CommonRes{
			Code: 500,
			Msg:  err.Error(),
		})
		response.ErrorOther(r, "导出Excel", req)
	}
	url, err := userService.Export(req)

	if err != nil {
		response.ErrorMsg(r, "导出Excel", req, model.Buniss_Other, err.Error())
	}

	response.SucessMsg(r, "导出Excel", req, model.Buniss_Other, url)
}
