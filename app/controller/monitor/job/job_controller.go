package job

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcron"
	"yj-app/app/model"
	jobModel "yj-app/app/model/monitor/job"
	jobLogModel "yj-app/app/model/monitor/job_log"
	jobService "yj-app/app/service/monitor/job"
	jobLogService "yj-app/app/service/monitor/job_log"
	userService "yj-app/app/service/system/user"
	"yj-app/app/utils/convert"
	"yj-app/app/utils/response"
)

//列表页
func List(r *ghttp.Request) {
	response.BuildTpl(r, "monitor/job/list.html").WriteTpl()
}

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *jobModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("定时任务管理", req).WriteJsonExit()
	}
	rows := make([]jobModel.Entity, 0)
	result, page, err := jobService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	response.BuildTable(r, page.Total, rows).WriteJsonExit()
}

//列表页
func LogList(r *ghttp.Request) {
	response.BuildTpl(r, "monitor/job/jobLog.html").WriteTpl()
}

//列表分页数据
func LogListAjax(r *ghttp.Request) {
	var req *jobLogModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("任务日志管理", req).WriteJsonExit()
	}
	rows := make([]jobLogModel.Entity, 0)
	result, page, err := jobLogService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
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
	user := userService.GetProfile(r.Session)
	response.BuildTpl(r, "monitor/job/add.html").WriteTpl(g.Map{"loginName": user.LoginName})
}

//新增页面保存
func AddSave(r *ghttp.Request) {
	var req *jobModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("定时任务管理", req).WriteJsonExit()
	}

	id, err := jobService.AddSave(req, r.Session)

	if err != nil || id <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).Log("定时任务管理", req).WriteJsonExit()
	}
	response.SucessResp(r).SetBtype(model.Buniss_Add).SetData(id).Log("定时任务管理", req).WriteJsonExit()
}

//修改页
func Edit(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	if id <= 0 {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "参数错误",
		})
		return
	}

	entity, err := jobService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "数据不存在",
		})
		return
	}

	user := userService.GetProfile(r.Session)

	response.BuildTpl(r, "monitor/job/edit.html").WriteTpl(g.Map{
		"job":       entity,
		"loginName": user.LoginName,
	})
}

//修改页面保存
func EditSave(r *ghttp.Request) {
	var req jobModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("定时任务管理", req).WriteJsonExit()
	}

	rs, err := jobService.EditSave(&req, r.Session)

	if err != nil || rs <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).Log("定时任务管理", req).WriteJsonExit()
	}
	response.SucessResp(r).SetBtype(model.Buniss_Edit).SetData(rs).Log("定时任务管理", req).WriteJsonExit()
}

//详情页
func Detail(r *ghttp.Request) {
	id := r.GetInt64("id")
	if id <= 0 {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "参数错误",
		})
		return
	}
	job, err := jobService.SelectRecordById(id)
	if err != nil || job == nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "数据不存在",
		})
		return
	}
	response.BuildTpl(r, "monitor/job/detail.html").WriteTpl(g.Map{"job": job})
}

//详情页
func DetailLog(r *ghttp.Request) {
	id := r.GetInt64("id")
	if id <= 0 {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "参数错误",
		})
		return
	}
	jobLog, err := jobLogService.SelectRecordById(id)
	if err != nil || jobLog == nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "数据不存在",
		})
		return
	}
	response.BuildTpl(r, "monitor/job/detailLog.html").WriteTpl(g.Map{"jobLog": jobLog})
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("定时任务管理", req).WriteJsonExit()
	}

	idarr := convert.ToInt64Array(req.Ids, ",")
	list, _ := jobModel.FindAll("job_id in (?)", idarr)
	if list != nil && len(list) > 0 {
		for _, j := range list {
			gcron.Remove(j.JobName)
		}
	}

	rs := jobService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(r).SetBtype(model.Buniss_Del).SetData(rs).Log("定时任务管理", req).WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).Log("定时任务管理", req).WriteJsonExit()
	}
}

//导出
func Export(r *ghttp.Request) {
	var req *jobModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("定时任务管理", req).WriteJsonExit()
	}
	url, err := jobService.Export(req)

	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("定时任务管理", req).WriteJsonExit()
	}
	response.SucessResp(r).SetMsg(url).Log("定时任务管理", req).WriteJsonExit()
}

//启动
func Start(r *ghttp.Request) {
	jobId := r.GetFormInt64("jobId")
	if jobId <= 0 {
		response.ErrorResp(r).SetMsg("参数错误").Log("定时任务管理启动", g.Map{"jobId": jobId}).WriteJsonExit()
	}
	job, _ := jobService.SelectRecordById(jobId)
	if job == nil {
		response.ErrorResp(r).SetMsg("任务不存在").Log("定时任务管理启动", g.Map{"jobId": jobId}).WriteJsonExit()
	}

	err := jobService.Start(job)
	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("定时任务管理启动", g.Map{"jobId": jobId}).WriteJsonExit()
	} else {
		response.ErrorResp(r).Log("定时任务管理启动", g.Map{"jobId": jobId}).WriteJsonExit()
	}
}

//停止
func Stop(r *ghttp.Request) {
	jobId := r.GetFormInt64("jobId")
	if jobId <= 0 {
		response.ErrorResp(r).SetMsg("参数错误").Log("定时任务管理停止", g.Map{"jobId": jobId}).WriteJsonExit()
	}
	job, _ := jobService.SelectRecordById(jobId)
	if job == nil {
		response.ErrorResp(r).SetMsg("任务不存在").Log("定时任务管理停止", g.Map{"jobId": jobId}).WriteJsonExit()
	}

	err := jobService.Stop(job)
	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("定时任务管理停止", g.Map{"jobId": jobId}).WriteJsonExit()
	} else {
		response.SucessResp(r).SetMsg("停止成功").Log("定时任务管理停止", g.Map{"jobId": jobId}).WriteJsonExit()
	}
}
