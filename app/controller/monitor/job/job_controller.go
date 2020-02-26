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
	"yj-app/app/service/utils/convert"
	"yj-app/app/service/utils/response"
)

//列表页
func List(r *ghttp.Request) {
	response.WriteTpl(r, "monitor/job/list.html")
}

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *jobModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "列表查询", req, model.Buniss_Other, err.Error())
	}
	rows := make([]jobModel.Entity, 0)
	result, page, err := jobService.SelectListByPage(req)

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

//列表页
func LogList(r *ghttp.Request) {
	response.WriteTpl(r, "monitor/job/jobLog.html")
}

//列表分页数据
func LogListAjax(r *ghttp.Request) {
	var req *jobLogModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "列表查询", req, model.Buniss_Other, err.Error())
	}
	rows := make([]jobLogModel.Entity, 0)
	result, page, err := jobLogService.SelectListByPage(req)

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
	user := userService.GetProfile(r.Session)
	response.WriteTpl(r, "monitor/job/add.html", g.Map{"loginName": user.LoginName})
}

//新增页面保存
func AddSave(r *ghttp.Request) {
	var req *jobModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "新增定时任务调度", req, model.Buniss_Add, err.Error())
	}

	rid, err := jobService.AddSave(req, r.Session)

	if err != nil || rid <= 0 {
		response.ErrorAdd(r, "新增定时任务调度", req)
	}
	response.SucessDataAdd(r, "新增定时任务调度", req, rid)
}

//修改页
func Edit(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	if id <= 0 {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "参数错误",
		})
		return
	}

	entity, err := jobService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "数据不存在",
		})
		return
	}

	user := userService.GetProfile(r.Session)

	response.WriteTpl(r, "monitor/job/edit.html", g.Map{
		"job":       entity,
		"loginName": user.LoginName,
	})
}

//修改页面保存
func EditSave(r *ghttp.Request) {
	var req jobModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "定时任务调度岗位", req, model.Buniss_Add, err.Error())
	}

	rs, err := jobService.EditSave(&req, r.Session)

	if err != nil || rs <= 0 {
		response.ErrorAdd(r, "修改定时任务调度", req)
	}
	response.SucessDataAdd(r, "修改定时任务调度", req, rs)
}

//详情页
func Detail(r *ghttp.Request) {
	id := r.GetInt64("id")
	if id <= 0 {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "参数错误",
		})
		return
	}
	job, err := jobService.SelectRecordById(id)
	if err != nil || job == nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "数据不存在",
		})
		return
	}
	response.WriteTpl(r, "monitor/job/detail.html", g.Map{"job": job})
}

//详情页
func DetailLog(r *ghttp.Request) {
	var jobLog jobLogModel.Entity
	response.WriteTpl(r, "monitor/job/detailLog.html", g.Map{"jobLog": jobLog})
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorEdit(r, "删除定时任务调度", req)
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
		response.SucessDataDel(r, "删除定时任务调度", req, rs)
	} else {
		response.ErrorDataMsg(r, "删除定时任务调度", req, model.Buniss_Del, 0, "未删除数据")
	}
}

//导出
func Export(r *ghttp.Request) {
	var req *jobModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorOther(r, "导出Excel", req)
	}
	url, err := jobService.Export(req)

	if err != nil {
		response.ErrorMsg(r, "导出Excel", req, model.Buniss_Other, err.Error())
	}

	response.SucessMsg(r, "导出Excel", req, model.Buniss_Other, url)
}

//启动
func Start(r *ghttp.Request) {
	jobId := r.GetFormInt64("jobId")
	if jobId <= 0 {
		response.ErrorMsg(r, "", g.Map{"jobId": jobId}, model.Buniss_Other, "参数错误")
	}
	job, _ := jobService.SelectRecordById(jobId)
	if job == nil {
		response.ErrorMsg(r, "", g.Map{"jobId": jobId}, model.Buniss_Other, "任务不存在")
	}
	err := jobService.Start(job)
	if err != nil {
		response.ErrorMsg(r, "", g.Map{"jobId": jobId}, model.Buniss_Other, err.Error())
	} else {
		response.SucessMsg(r, "", g.Map{"jobId": jobId}, model.Buniss_Other, "启动成功")
	}
}

//停止
func Stop(r *ghttp.Request) {
	jobId := r.GetFormInt64("jobId")
	if jobId <= 0 {
		response.ErrorMsg(r, "", g.Map{"jobId": jobId}, model.Buniss_Other, "参数错误")
	}
	job, _ := jobService.SelectRecordById(jobId)
	if job == nil {
		response.ErrorMsg(r, "", g.Map{"jobId": jobId}, model.Buniss_Other, "任务不存在")
	}

	err := jobService.Stop(job)
	if err != nil {
		response.ErrorMsg(r, "", g.Map{"jobId": jobId}, model.Buniss_Other, err.Error())
	} else {
		response.SucessMsg(r, "", g.Map{"jobId": jobId}, model.Buniss_Other, "停止成功")
	}
}
