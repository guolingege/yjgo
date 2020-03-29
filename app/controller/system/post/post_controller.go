package post

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	postModel "yj-app/app/model/system/post"
	postService "yj-app/app/service/system/post"
	"yj-app/app/utils/response"
)

//列表页
func List(r *ghttp.Request) {
	response.BuildTpl(r, "system/post/list.html").WriteTpl()
}

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *postModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
	}
	rows := make([]postModel.Entity, 0)
	result, page, err := postService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response.BuildTable(r, page.Total, rows).WriteJsonExit()
}

//新增页面
func Add(r *ghttp.Request) {
	response.BuildTpl(r, "system/post/add.html").WriteTpl()
}

//新增页面保存
func AddSave(r *ghttp.Request) {
	var req *postModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
	}

	if postService.CheckPostNameUniqueAll(req.PostName) == "1" {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg("岗位名称已存在").Log("岗位管理", req).WriteJsonExit()
	}

	if postService.CheckPostCodeUniqueAll(req.PostCode) == "1" {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg("岗位编码已存在").Log("岗位管理", req).WriteJsonExit()
	}

	pid, err := postService.AddSave(req, r.Session)

	if err != nil || pid <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).Log("岗位管理", req).WriteJsonExit()
	}
	response.ErrorResp(r).SetData(pid).SetBtype(model.Buniss_Add).Log("岗位管理", req).WriteJsonExit()
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

	post, err := postService.SelectRecordById(id)

	if err != nil || post == nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "岗位不存在",
		})
		return
	}

	response.BuildTpl(r, "system/post/edit.html").WriteTpl(g.Map{
		"post": post,
	})
}

//修改页面保存
func EditSave(r *ghttp.Request) {
	var req *postModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
	}

	if postService.CheckPostNameUnique(req.PostName, req.PostId) == "1" {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg("岗位名称已存在").Log("岗位管理", req).WriteJsonExit()
	}

	if postService.CheckPostCodeUnique(req.PostCode, req.PostId) == "1" {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg("岗位编码已存在").Log("岗位管理", req).WriteJsonExit()
	}

	rs, err := postService.EditSave(req, r.Session)

	if err != nil || rs <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).Log("岗位管理", req).WriteJsonExit()
	}
	response.SucessResp(r).SetData(rs).SetBtype(model.Buniss_Edit).Log("岗位管理", req).WriteJsonExit()
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).SetBtype(model.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
	}

	rs := postService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(r).SetBtype(model.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
	}
}

//导出
func Export(r *ghttp.Request) {
	var req *postModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
	}
	url, err := postService.Export(req)

	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
	}
	response.SucessResp(r).SetMsg(url).SetBtype(model.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
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
