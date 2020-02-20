package logininfor

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	logininforModel "yj-app/app/model/monitor/logininfor"
	logininforService "yj-app/app/service/monitor/logininfor"
	"yj-app/app/service/utils/response"
)

//用户列表页
func List(r *ghttp.Request) {
	response.WriteTpl(r, "monitor/logininfor/list.html")
}

//用户列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *logininforModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(model.CommonRes{
			Code: 500,
			Msg:  err.Error(),
		})
	}

	rows := make([]logininforModel.Entity, 0)

	result, page, err := logininforService.SelectPageList(req)

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

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "删除登陆日志", req, model.Buniss_Del, err.Error())
	}

	rs := logininforService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessDataDel(r, "删除登陆日志", req, rs)
	} else {
		response.ErrorDelData(r, "删除登陆日志", req, 0)
	}
}

//清空记录
func Clean(r *ghttp.Request) {

	rs := logininforService.DeleteRecordAll()

	if rs > 0 {
		response.SucessDataDel(r, "清空登陆日志", "all", rs)
	} else {
		response.ErrorDelData(r, "清空登陆日志", "all", 0)
	}
}

//导出
func Export(r *ghttp.Request) {
	var req *logininforModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "导出登陆日志", req, model.Buniss_Other, err.Error())
	}

	url, err := logininforService.Export(req)

	if err != nil {
		response.ErrorMsg(r, "导出登陆日志", req, model.Buniss_Other, err.Error())
	} else {
		response.SucessMsg(r, "", req, model.Buniss_Other, url)
	}
}

//解锁账号
func Unlock(r *ghttp.Request) {
	loginName := r.GetQueryString("loginName")
	if loginName == "" {
		response.ErrorMsg(r, "解锁账号", "loginName", model.Buniss_Other, "参数错误")
	} else {
		logininforService.RemovePasswordCounts(loginName)
		logininforService.Unlock(loginName)
		r.Response.WriteJsonExit(model.CommonRes{
			Code: 0,
			Msg:  "操作成功",
		})
		response.SucessOther(r, "解锁账号", loginName)
	}

}
