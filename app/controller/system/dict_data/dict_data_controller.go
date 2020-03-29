package dict_data

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	dictModel "yj-app/app/model/system/dict_data"
	dictService "yj-app/app/service/system/dict_data"
	"yj-app/app/utils/response"
)

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *dictModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
	}
	rows := make([]dictModel.Entity, 0)
	result, page, err := dictService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response.BuildTable(r, page.Total, rows).WriteJsonExit()
}

//新增页面
func Add(r *ghttp.Request) {
	dictType := r.GetQueryString("dictType")
	response.BuildTpl(r, "system/dict/data/add.html").WriteTpl(g.Map{"dictType": dictType})
}

//新增页面保存
func AddSave(r *ghttp.Request) {
	var req *dictModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
	}

	rid, err := dictService.AddSave(req, r.Session)

	if err != nil || rid <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).Log("字典数据管理", req).WriteJsonExit()
	}
	response.SucessResp(r).SetData(rid).SetBtype(model.Buniss_Add).Log("字典数据管理", req).WriteJsonExit()
}

//修改页面
func Edit(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	if id <= 0 {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "字典数据错误",
		})
		return
	}

	entity, err := dictService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "字典数据不存在",
		})
		return
	}

	response.BuildTpl(r, "system/dict/data/edit.html").WriteTpl(g.Map{
		"dict": entity,
	})
}

//修改页面保存
func EditSave(r *ghttp.Request) {
	var req *dictModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
	}

	rs, err := dictService.EditSave(req, r.Session)

	if err != nil || rs <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).Log("字典数据管理", req).WriteJsonExit()
	}
	response.SucessResp(r).SetBtype(model.Buniss_Edit).SetData(rs).Log("字典数据管理", req).WriteJsonExit()
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
	}

	rs := dictService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(r).SetBtype(model.Buniss_Del).SetData(rs).Log("字典数据管理", req).WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).Log("字典数据管理", req).WriteJsonExit()
	}
}

//导出
func Export(r *ghttp.Request) {
	var req *dictModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("字典数据导出", req).WriteJsonExit()
	}
	url, err := dictService.Export(req)

	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("字典数据导出", req).WriteJsonExit()
	}
	response.SucessResp(r).SetMsg(url).Log("导出Excel", req).WriteJsonExit()
}
