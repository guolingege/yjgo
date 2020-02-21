package table

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	"yj-app/app/service/utils/response"
)

func Button(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/button.html")
}

func Child(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/child.html")
}

func Curd(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/curd.html")
}

func Detail(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/detail.html")
}

func Editable(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/editable.html")
}

func Event(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/event.html")
}

func Export(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/export.html")
}

func FixedColumns(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/fixedColumns.html")
}

func Footer(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/footer.html")
}

func GroupHeader(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/groupHeader.html")
}

func Image(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/image.html")
}

func Multi(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/multi.html")
}

func Other(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/other.html")
}

func PageGo(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/pageGo.html")
}

func Params(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/params.html")
}

func Remember(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/remember.html")
}

func Recorder(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/recorder.html")
}

func Search(r *ghttp.Request) {
	response.WriteTpl(r, "demo/table/search.html")
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


