package online

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strings"
	"yj-app/app/model"
	onlineModel "yj-app/app/model/monitor/online"
	onlineService "yj-app/app/service/monitor/online"
	userService "yj-app/app/service/system/user"
	"yj-app/app/service/utils/response"
)

//列表页
func List(r *ghttp.Request) {
	response.WriteTpl(r, "monitor/online/list.html")
}

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *onlineModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "列表查询", req, model.Buniss_Other, err.Error())
	}
	rows := make([]onlineModel.Entity, 0)
	result, page, err := onlineService.SelectListByPage(req)

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

//用户强退
func ForceLogout(r *ghttp.Request) {
	sessionId := r.GetFormString("sessionId")
	if sessionId == "" {
		response.ErrorMsg(r, "用户强退", g.Map{"sessionId": sessionId}, model.Buniss_Other, "参数错误")
	}

	err := userService.ForceLogout(sessionId)
	if err != nil {
		response.ErrorMsg(r, "用户强退", g.Map{"sessionId": sessionId}, model.Buniss_Other, err.Error())
	}
	response.SucessOther(r, "用户强退", g.Map{"sessionId": sessionId})
}

//批量强退
func BatchForceLogout(r *ghttp.Request) {
	ids := r.GetFormString("ids")
	if ids == "" {
		response.ErrorMsg(r, "批量强退", g.Map{"ids": ids}, model.Buniss_Other, "参数错误")
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
	response.SucessOther(r, "批量强退", g.Map{"ids": ids})
}
