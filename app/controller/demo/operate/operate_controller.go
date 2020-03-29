package operate

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	"yj-app/app/utils/response"
)

func Add(r *ghttp.Request) {
	response.BuildTpl(r, "demo/operate/add.html").WriteTpl()
}

func Detail(r *ghttp.Request) {
	var tmp us
	tmp.UserId = 1
	tmp.UserName = "测试1"
	tmp.Status = "0"
	tmp.CreateTime = "2020-01-12 02:02:02"
	tmp.UserBalance = 100
	tmp.UserCode = "1000001"
	tmp.UserSex = "0"
	tmp.UserPhone = "15888888888"
	tmp.UserEmail = "111@qq.com"
	response.BuildTpl(r, "demo/operate/detail.html").WriteTpl(g.Map{"user": tmp})
}

type us struct {
	UserId      int64   `json:"userId"`
	UserCode    string  `json:"userCode"`
	UserName    string  `json:"userName"`
	UserSex     string  `json:"userName"`
	UserPhone   string  `json:"userPhone"`
	UserEmail   string  `json:"userEmail"`
	UserBalance float64 `json:"userBalance"`
	Status      string  `json:"status"`
	CreateTime  string  `json:"createTime"`
}

func EditSave(r *ghttp.Request) {
	var tmp us
	tmp.UserId = 1
	tmp.UserName = "测试1"
	tmp.Status = "0"
	tmp.CreateTime = "2020-01-12 02:02:02"
	tmp.UserBalance = 100
	tmp.UserCode = "1000001"
	tmp.UserSex = "0"
	tmp.UserPhone = "15888888888"
	tmp.UserEmail = "111@qq.com"
	response.SucessResp(r).SetData(tmp).Log("demo演示", g.Map{"UserId": 1}).WriteJsonExit()
}

func Edit(r *ghttp.Request) {
	var tmp us
	tmp.UserId = 1
	tmp.UserName = "测试1"
	tmp.Status = "0"
	tmp.CreateTime = "2020-01-12 02:02:02"
	tmp.UserBalance = 100
	tmp.UserCode = "1000001"
	tmp.UserSex = "0"
	tmp.UserPhone = "15888888888"
	tmp.UserEmail = "111@qq.com"
	response.BuildTpl(r, "demo/operate/edit.html").WriteTpl(g.Map{"user": tmp})
}

func Other(r *ghttp.Request) {
	response.BuildTpl(r, "demo/operate/other.html").WriteTpl()
}

func Table(r *ghttp.Request) {
	response.BuildTpl(r, "demo/operate/table.html").WriteTpl()
}

func List(r *ghttp.Request) {
	var rows = make([]us, 0)
	for i := 1; i <= 10; i++ {
		var tmp us
		tmp.UserId = int64(i)
		tmp.UserName = "测试" + string(i)
		tmp.Status = "0"
		tmp.CreateTime = "2020-01-12 02:02:02"
		tmp.UserBalance = 100
		tmp.UserCode = "100000" + string(i)
		tmp.UserSex = "0"
		tmp.UserPhone = "15888888888"
		tmp.UserEmail = "111@qq.com"
		rows = append(rows, tmp)
	}
	r.Response.WriteJsonExit(model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}
