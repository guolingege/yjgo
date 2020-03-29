package modal

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/utils/response"
)

func Dialog(r *ghttp.Request) {
	response.BuildTpl(r, "demo/modal/dialog.html").WriteTpl()
}

func Form(r *ghttp.Request) {
	response.BuildTpl(r, "demo/modal/form.html").WriteTpl()
}

func Layer(r *ghttp.Request) {
	response.BuildTpl(r, "demo/modal/layer.html").WriteTpl()
}

func Table(r *ghttp.Request) {
	response.BuildTpl(r, "demo/modal/table.html").WriteTpl()
}

func Check(r *ghttp.Request) {
	response.BuildTpl(r, "demo/modal/table/check.html").WriteTpl()
}

func Parent(r *ghttp.Request) {
	response.BuildTpl(r, "demo/modal/table/parent.html").WriteTpl()
}

func Radio(r *ghttp.Request) {
	response.BuildTpl(r, "demo/modal/table/radio.html").WriteTpl()
}
