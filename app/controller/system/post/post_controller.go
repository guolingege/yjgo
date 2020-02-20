package post

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	postModel "yj-app/app/model/system/post"
	postService "yj-app/app/service/system/post"
	"yj-app/app/service/utils/response"
)

//列表页
func List(r *ghttp.Request) {
	response.WriteTpl(r, "system/post/list.html")
}

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *postModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "列表查询", req, model.Buniss_Other, err.Error())
	}
	rows := make([]postModel.Entity, 0)
	result, page, err := postService.SelectListByPage(req)

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
	response.WriteTpl(r, "system/post/add.html")
}

//新增页面保存
func AddSave(r *ghttp.Request) {
	var req *postModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "新增岗位", req, model.Buniss_Add, err.Error())
	}

	if postService.CheckPostNameUniqueAll(req.PostName) == "1" {
		response.ErrorMsg(r, "新增岗位", req, model.Buniss_Add, "岗位名称已存在")
	}

	if postService.CheckPostCodeUniqueAll(req.PostCode) == "1" {
		response.ErrorMsg(r, "新增岗位", req, model.Buniss_Add, "岗位编码已存在")
	}

	rid, err := postService.AddSave(req, r.Session)

	if err != nil || rid <= 0 {
		response.ErrorAdd(r, "新增岗位", req)
	}
	response.SucessDataAdd(r, "新增岗位", req, rid)
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

	post, err := postService.SelectRecordById(id)

	if err != nil || post == nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "岗位不存在",
		})
		return
	}

	response.WriteTpl(r, "system/post/edit.html", g.Map{
		"post": post,
	})
}

//修改页面保存
func EditSave(r *ghttp.Request) {
	var req *postModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "修改岗位", req, model.Buniss_Add, err.Error())
	}

	if postService.CheckPostNameUnique(req.PostName, req.PostId) == "1" {
		response.ErrorMsg(r, "修改岗位", req, model.Buniss_Add, "岗位名称已存在")
	}

	if postService.CheckPostCodeUnique(req.PostCode, req.PostId) == "1" {
		response.ErrorMsg(r, "修改岗位", req, model.Buniss_Add, "岗位编码已存在")
	}

	rs, err := postService.EditSave(req, r.Session)

	if err != nil || rs <= 0 {
		response.ErrorAdd(r, "修改岗位", req)
	}
	response.SucessDataAdd(r, "修改岗位", req, rs)
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorEdit(r, "删除岗位", req)
	}

	rs := postService.DeleteRecordByIds(req.Ids)

	if rs > 0 {

		response.SucessDataDel(r, "删除岗位", req, rs)
	} else {
		response.ErrorDataMsg(r, "删除岗位", req, model.Buniss_Del, 0, "未删除数据")
	}
}

//导出
func Export(r *ghttp.Request) {
	var req *postModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorOther(r, "导出Excel", req)
	}
	url, err := postService.Export(req)

	if err != nil {
		response.ErrorMsg(r, "导出Excel", req, model.Buniss_Other, err.Error())
	}

	response.SucessMsg(r, "导出Excel", req, model.Buniss_Other, url)
}

//检查岗位名称是否已经存在不包括本岗位
func CheckPostNameUnique(r *ghttp.Request) {
	var req *postModel.CheckPostNameReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := postService.CheckPostNameUnique(req.PostName, req.PostId)

	r.Response.WritefExit(result)
}

//检查岗位名称是否已经存在
func CheckPostNameUniqueAll(r *ghttp.Request) {
	var req *postModel.CheckPostNameALLReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := postService.CheckPostNameUniqueAll(req.PostName)

	r.Response.WritefExit(result)
}

//检查岗位编码是否已经存在不包括本岗位
func CheckPostCodeUnique(r *ghttp.Request) {
	var req *postModel.CheckPostCodeReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := postService.CheckPostCodeUnique(req.PostCode, req.PostId)

	r.Response.WritefExit(result)
}

//检查岗位编码是否已经存在
func CheckPostCodeUniqueAll(r *ghttp.Request) {
	var req *postModel.CheckPostCodeALLReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := postService.CheckPostCodeUniqueAll(req.PostCode)

	r.Response.WritefExit(result)
}
