package online

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strings"
	onlineModel "yj-app/app/model/monitor/online"
	onlineService "yj-app/app/service/monitor/online"
	userService "yj-app/app/service/system/user"
	"yj-app/app/utils/response"
)

//列表页
func List(r *ghttp.Request) {
	var sessinIdArr = make([]string, 0)
	if userService.SessionList != nil {
		for k, v := range userService.SessionList.Map() {
			tmp := v.(*ghttp.Session)
			if tmp.Get("uid").(int64) > 0 {
				sessinIdArr = append(sessinIdArr, k.(string))
			}
		}
	}
	if len(sessinIdArr) > 0 {
		onlineService.DeleteRecordNotInIds(sessinIdArr)
	}

	response.BuildTpl(r, "monitor/online/list.html").WriteTpl()
}

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *onlineModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).WriteJsonExit()
	}
	rows := make([]onlineModel.Entity, 0)
	result, page, err := onlineService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response.BuildTable(r, page.Total, rows).WriteJsonExit()
}

//用户强退
func ForceLogout(r *ghttp.Request) {
	sessionId := r.GetFormString("sessionId")
	if sessionId == "" {
		response.ErrorResp(r).SetMsg("参数错误").Log("用户强退", g.Map{"sessionId": sessionId}).WriteJsonExit()
	}

	err := userService.ForceLogout(sessionId)
	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("用户强退", g.Map{"sessionId": sessionId}).WriteJsonExit()
	}
	response.SucessResp(r).Log("用户强退", g.Map{"sessionId": sessionId}).WriteJsonExit()
}

//批量强退
func BatchForceLogout(r *ghttp.Request) {
	ids := r.GetFormString("ids")
	if ids == "" {
		response.ErrorResp(r).SetMsg("参数错误").Log("批量强退", g.Map{"ids": ids}).WriteJsonExit()
	}
	ids = strings.ReplaceAll(ids, "[", "")
	ids = strings.ReplaceAll(ids, "]", "")
	ids = strings.ReplaceAll(ids, `"`, "")
	idarr := strings.Split(ids, ",")
	if len(idarr) > 0 {
		for _, sessionId := range idarr {
			if sessionId != "" {
				userService.ForceLogout(sessionId)
			}
		}
	}
	response.SucessResp(r).Log("批量强退", g.Map{"ids": ids}).WriteJsonExit()
}
