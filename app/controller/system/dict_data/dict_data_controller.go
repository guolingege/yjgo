package dict_data

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	dictModel "yj-app/app/model/system/dict_data"
	dictService "yj-app/app/service/system/dict_data"
	"yj-app/app/service/utils/response"
)

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *dictModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "列表查询", req, model.Buniss_Other, err.Error())
	}
	rows := make([]dictModel.Entity, 0)
	result, page, err := dictService.SelectListByPage(req)

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
	dictType := r.GetQueryString("dictType")
	response.WriteTpl(r, "system/dict/data/add.html", g.Map{"dictType": dictType})
}

//新增页面保存
func AddSave(r *ghttp.Request) {
	var req *dictModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "新增字典数据", req, model.Buniss_Add, err.Error())
	}

	rid, err := dictService.AddSave(req, r.Session)

	if err != nil || rid <= 0 {
		response.ErrorAdd(r, "新增字典数据", req)
	}
	response.SucessDataAdd(r, "新增字典数据", req, rid)
}

//修改页面
func Edit(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	if id <= 0 {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "字典数据错误",
		})
		return
	}

	entity, err := dictService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "字典数据不存在",
		})
		return
	}

	response.WriteTpl(r, "system/dict/data/edit.html", g.Map{
		"dict": entity,
	})
}

//修改页面保存
func EditSave(r *ghttp.Request) {
	var req *dictModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "字典数据岗位", req, model.Buniss_Add, err.Error())
	}

	rs, err := dictService.EditSave(req, r.Session)

	if err != nil || rs <= 0 {
		response.ErrorAdd(r, "修改字典数据", req)
	}
	response.SucessDataAdd(r, "修改字典数据", req, rs)
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorEdit(r, "删除字典数据", req)
	}

	rs := dictService.DeleteRecordByIds(req.Ids)

	if rs > 0 {

		response.SucessDataDel(r, "删除字典数据", req, rs)
	} else {
		response.ErrorDataMsg(r, "删除字典数据", req, model.Buniss_Del, 0, "未删除数据")
	}
}

//导出
func Export(r *ghttp.Request) {
	var req *dictModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorOther(r, "导出Excel", req)
	}
	url, err := dictService.Export(req)

	if err != nil {
		response.ErrorMsg(r, "导出Excel", req, model.Buniss_Other, err.Error())
	}

	response.SucessMsg(r, "导出Excel", req, model.Buniss_Other, url)
}
