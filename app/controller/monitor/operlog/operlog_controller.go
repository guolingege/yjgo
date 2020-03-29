package operlog

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	operlogModel "yj-app/app/model/monitor/oper_log"
	operlogService "yj-app/app/service/monitor/operlog"
	"yj-app/app/utils/response"
)

//用户列表页
func List(r *ghttp.Request) {
	response.BuildTpl(r, "monitor/operlog/list.html").WriteTpl()
}

//用户列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *operlogModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(model.CommonRes{
			Code: 500,
			Msg:  err.Error(),
		})
	}

	rows := make([]operlogModel.Entity, 0)

	result, page, err := operlogService.SelectPageList(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response.BuildTable(r, page.Total, rows).WriteJsonExit()
}

//清空记录
func Clean(r *ghttp.Request) {

	rs := operlogService.DeleteRecordAll()

	if rs > 0 {
		response.SucessResp(r).SetBtype(model.Buniss_Del).SetData(rs).Log("操作日志管理", "all").WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).Log("操作日志管理", "all").WriteJsonExit()
	}
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("操作日志管理", req).WriteJsonExit()
	}

	rs := operlogService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(r).SetBtype(model.Buniss_Del).SetData(rs).Log("操作日志管理", req).WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).Log("操作日志管理", req).WriteJsonExit()
	}
}

//记录详情
func Detail(r *ghttp.Request) {

	id := r.GetQueryInt64("id")

	if id <= 0 {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "参数错误",
		})
		return
	}

	operLog, err := operlogService.SelectRecordById(id)

	if err != nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "数据不存在",
		})
		return
	}

	response.BuildTpl(r, "monitor/operlog/detail.html").WriteTpl(g.Map{
		"operLog": operLog,
	})
}

//导出
func Export(r *ghttp.Request) {
	var req *operlogModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("导出操作日志", req).WriteJsonExit()
	}
	url, err := operlogService.Export(req)

	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("导出操作日志", req).WriteJsonExit()
	} else {
		response.SucessResp(r).SetMsg(url).Log("导出操作日志", req).WriteJsonExit()
	}
}
