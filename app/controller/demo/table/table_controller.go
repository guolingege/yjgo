package table

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	"yj-app/app/utils/response"
)

func Button(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/button.html").WriteTpl()
}

func Child(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/child.html").WriteTpl()
}

func Curd(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/curd.html").WriteTpl()
}

func Detail(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/detail.html").WriteTpl()
}

func Editable(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/editable.html").WriteTpl()
}

func Event(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/event.html").WriteTpl()
}

func Export(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/export.html").WriteTpl()
}

func FixedColumns(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/fixedColumns.html").WriteTpl()
}

func Footer(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/footer.html").WriteTpl()
}

func GroupHeader(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/groupHeader.html").WriteTpl()
}

func Image(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/image.html").WriteTpl()
}

func Multi(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/multi.html").WriteTpl()
}

func Other(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/other.html").WriteTpl()
}

func PageGo(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/pageGo.html").WriteTpl()
}

func Params(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/params.html").WriteTpl()
}

func Remember(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/remember.html").WriteTpl()
}

func Recorder(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/recorder.html").WriteTpl()
}

func Search(r *ghttp.Request) {
	response.BuildTpl(r, "demo/table/search.html").WriteTpl()
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
