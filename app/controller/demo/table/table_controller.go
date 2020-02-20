package table

import (
	"github.com/gogf/gf/net/ghttp"
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
