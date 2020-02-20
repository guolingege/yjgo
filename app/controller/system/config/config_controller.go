package config

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	configModel "yj-app/app/model/system/config"
	configService "yj-app/app/service/system/config"
	"yj-app/app/service/utils/response"
)

//列表页
func List(r *ghttp.Request) {
	response.WriteTpl(r, "system/config/list.html")
}

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *configModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "列表查询", req, model.Buniss_Other, err.Error())
	}
	rows := make([]configModel.Entity, 0)
	result, page, err := configService.SelectListByPage(req)

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
	response.WriteTpl(r, "system/config/add.html")
}

//新增页面保存
func AddSave(r *ghttp.Request) {
	var req *configModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "新增参数", req, model.Buniss_Add, err.Error())
	}

	if configService.CheckConfigKeyUniqueAll(req.ConfigKey) == "1" {
		response.ErrorMsg(r, "新增参数", req, model.Buniss_Add, "参数键名已存在")
	}

	rid, err := configService.AddSave(req, r.Session)

	if err != nil || rid <= 0 {
		response.ErrorAdd(r, "新增参数", req)
	}
	response.SucessDataAdd(r, "新增参数", req, rid)
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

	entity, err := configService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "参数不存在",
		})
		return
	}

	response.WriteTpl(r, "system/config/edit.html", g.Map{
		"config": entity,
	})
}

//修改页面保存
func EditSave(r *ghttp.Request) {
	var req configModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "参数岗位", req, model.Buniss_Add, err.Error())
	}

	if configService.CheckConfigKeyUnique(req.ConfigKey, req.ConfigId) == "1" {
		response.ErrorMsg(r, "新增参数", req, model.Buniss_Add, "参数键名已存在")
	}

	rs, err := configService.EditSave(&req, r.Session)

	if err != nil || rs <= 0 {
		response.ErrorAdd(r, "修改参数", req)
	}
	response.SucessDataAdd(r, "修改参数", req, rs)
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorEdit(r, "删除参数", req)
	}

	rs := configService.DeleteRecordByIds(req.Ids)

	if rs > 0 {

		response.SucessDataDel(r, "删除参数", req, rs)
	} else {
		response.ErrorDataMsg(r, "删除参数", req, model.Buniss_Del, 0, "未删除数据")
	}
}

//导出
func Export(r *ghttp.Request) {
	var req *configModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorOther(r, "导出Excel", req)
	}
	url, err := configService.Export(req)

	if err != nil {
		response.ErrorMsg(r, "导出Excel", req, model.Buniss_Other, err.Error())
	}

	response.SucessMsg(r, "导出Excel", req, model.Buniss_Other, url)
}

//检查参数键名是否已经存在不包括本参数
func CheckConfigKeyUnique(r *ghttp.Request) {
	var req *configModel.CheckConfigKeyReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := configService.CheckConfigKeyUnique(req.ConfigKey, req.ConfigId)

	r.Response.WritefExit(result)
}

//检查参数键名是否已经存在
func CheckConfigKeyUniqueAll(r *ghttp.Request) {
	var req *configModel.CheckPostCodeALLReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := configService.CheckConfigKeyUniqueAll(req.ConfigKey)

	r.Response.WritefExit(result)
}
