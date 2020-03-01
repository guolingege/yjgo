package modal

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/utils/response"
)

func Dialog(r *ghttp.Request) {
	response.BuildTpl(r, "demo/modal/dialog.html").WriteTplExtend()
}

func Form(r *ghttp.Request) {
	response.BuildTpl(r, "demo/modal/form.html").WriteTplExtend()
}

func Layer(r *ghttp.Request) {
	response.BuildTpl(r, "demo/modal/layer.html").WriteTplExtend()
}

func Table(r *ghttp.Request) {
	response.BuildTpl(r, "demo/modal/table.html").WriteTplExtend()
}

func Check(r *ghttp.Request) {
	response.BuildTpl(r, "demo/modal/table/check.html").WriteTplExtend()
}

func Parent(r *ghttp.Request) {
	response.BuildTpl(r, "demo/modal/table/parent.html").WriteTplExtend()
}

func Radio(r *ghttp.Request) {
	response.BuildTpl(r, "demo/modal/table/radio.html").WriteTplExtend()
}
