package dept

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	deptModel "yj-app/app/model/system/dept"
	deptService "yj-app/app/service/system/dept"
	"yj-app/app/service/utils/response"
)

//列表页
func List(r *ghttp.Request) {
	response.WriteTpl(r, "system/dept/list.html")
}

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *deptModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "列表查询", req, model.Buniss_Other, err.Error())
	}
	rows := make([]deptModel.Entity, 0)
	result, err := deptService.SelectListAll(req)

	if err == nil && result != nil {
		rows = *result
	}

	r.Response.WriteJsonExit(rows)
}

//新增页面
func Add(r *ghttp.Request) {
	pid := r.GetQueryInt64("pid")

	if pid == 0 {
		pid = 100
	}

	tmp := deptService.SelectDeptById(pid)

	response.WriteTpl(r, "system/dept/add.html", g.Map{"dept": tmp})
}

//新增页面保存
func AddSave(r *ghttp.Request) {
	var req *deptModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "新增部门", req, model.Buniss_Add, err.Error())
	}

	if deptService.CheckDeptNameUniqueAll(req.DeptName, req.ParentId) == "1" {
		response.ErrorMsg(r, "新增部门", req, model.Buniss_Add, "菜单名称已存在")
	}

	rid, err := deptService.AddSave(req, r.Session)

	if err != nil || rid <= 0 {
		response.ErrorAdd(r, "新增部门", req)
	}
	response.SucessDataAdd(r, "新增部门", req, rid)
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

	dept := deptService.SelectDeptById(id)

	if dept == nil || dept.DeptId <= 0 {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "部门不存在",
		})
		return
	}

	response.WriteTpl(r, "system/dept/edit.html", g.Map{
		"dept": dept,
	})
}

//修改页面保存
func EditSave(r *ghttp.Request) {
	var req *deptModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "修改部门", req, model.Buniss_Add, err.Error())
	}

	if deptService.CheckDeptNameUnique(req.DeptName, req.DeptId, req.ParentId) == "1" {
		response.ErrorMsg(r, "修改部门", req, model.Buniss_Add, "部门名称已存在")
	}

	rs, err := deptService.EditSave(req, r.Session)

	if err != nil || rs <= 0 {
		response.ErrorAdd(r, "修改部门", req)
	}
	response.SucessDataAdd(r, "修改部门", req, rs)
}

//删除数据
func Remove(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	rs := deptService.DeleteDeptById(id)

	if rs > 0 {
		response.SucessDataDel(r, "删除部门", g.Map{"id": id}, rs)
	} else {
		response.ErrorDataMsg(r, "删除部门", g.Map{"id": id}, model.Buniss_Del, 0, "未删除数据")
	}
}

//加载部门列表树结构的数据
func TreeData(r *ghttp.Request) {
	result, _ := deptService.SelectDeptTree(0, "", "")
	r.Response.WriteJsonExit(result)
}

//加载部门列表树选择页面
func SelectDeptTree(r *ghttp.Request) {
	deptId := r.GetQueryInt64("deptId")
	deptPoint := deptService.SelectDeptById(deptId)

	if deptPoint != nil {
		response.WriteTpl(r, "system/dept/tree.html", g.Map{
			"dept": *deptPoint,
		})
	} else {
		response.WriteTpl(r, "system/dept/tree.html")
	}
}

//加载角色部门（数据权限）列表树
func RoleDeptTreeData(r *ghttp.Request) {
	roleId := r.GetQueryInt64("roleId")

	result, err := deptService.RoleDeptTreeData(roleId)

	if err != nil {
		response.ErrorMsg(r, "菜单树", g.Map{"roleId": roleId}, model.Buniss_Other, err.Error())
	}

	r.Response.WriteJsonExit(result)
}

//检查部门名称是否已经存在
func CheckDeptNameUnique(r *ghttp.Request) {
	var req *deptModel.CheckDeptNameReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := deptService.CheckDeptNameUnique(req.DeptName, req.DeptId, req.ParentId)

	r.Response.WritefExit(result)
}

//检查部门名称是否已经存在
func CheckDeptNameUniqueAll(r *ghttp.Request) {
 	var req *deptModel.CheckDeptNameALLReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := deptService.CheckDeptNameUniqueAll(req.DeptName, req.ParentId)

	r.Response.WritefExit(result)
}
