package logininfor

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	logininforModel "yj-app/app/model/monitor/logininfor"
	logininforService "yj-app/app/service/monitor/logininfor"
	"yj-app/app/utils/response"
)

//用户列表页
func List(r *ghttp.Request) {
	response.BuildTpl(r, "monitor/logininfor/list.html").WriteTpl()
}

//用户列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *logininforModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).WriteJsonExit()
	}

	rows := make([]logininforModel.Entity, 0)

	result, page, err := logininforService.SelectPageList(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	response.BuildTable(r, page.Total, rows).WriteJsonExit()
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("登陆日志管理", req).WriteJsonExit()
	}

	rs := logininforService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(r).SetBtype(model.Buniss_Del).SetData(rs).Log("登陆日志管理", req).WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).Log("登陆日志管理", req).WriteJsonExit()
	}
}

//清空记录
func Clean(r *ghttp.Request) {

	rs := logininforService.DeleteRecordAll()

	if rs > 0 {
		response.SucessResp(r).SetBtype(model.Buniss_Del).SetData(rs).Log("登陆日志管理", "all").WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).Log("登陆日志管理", "all").WriteJsonExit()
	}
}

//导出
func Export(r *ghttp.Request) {
	var req *logininforModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("导出登陆日志", req).WriteJsonExit()
	}

	url, err := logininforService.Export(req)

	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("导出登陆日志", req).WriteJsonExit()
	} else {
		response.SucessResp(r).SetMsg(url).Log("导出登陆日志", req).WriteJsonExit()
	}
}

//解锁账号
func Unlock(r *ghttp.Request) {
	loginName := r.GetQueryString("loginName")
	if loginName == "" {
		response.ErrorResp(r).SetMsg("参数错误").Log("解锁账号", "loginName="+loginName).WriteJsonExit()
	} else {
		logininforService.RemovePasswordCounts(loginName)
		logininforService.Unlock(loginName)
		response.SucessResp(r).Log("解锁账号", "loginName="+loginName).WriteJsonExit()
	}

}
