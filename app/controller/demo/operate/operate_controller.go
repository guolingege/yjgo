package operate

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/utils/response"
)

func Add(r *ghttp.Request) {
	response.WriteTpl(r, "demo/operate/add.html")
}

func Detail(r *ghttp.Request) {
	response.WriteTpl(r, "demo/operate/detail.html")
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
	response.SucessDataEdit(r, "demo演示", g.Map{"UserId": 1}, tmp)
}

func Other(r *ghttp.Request) {
	response.WriteTpl(r, "demo/operate/other.html")
}

func Table(r *ghttp.Request) {
	response.WriteTpl(r, "demo/operate/table.html")
}
