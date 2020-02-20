package modal

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/utils/response"
)

func Dialog(r *ghttp.Request) {
	response.WriteTpl(r, "demo/modal/dialog.html")
}

func Form(r *ghttp.Request) {
	response.WriteTpl(r, "demo/modal/form.html")
}

func Layer(r *ghttp.Request) {
	response.WriteTpl(r, "demo/modal/layer.html")
}

func Table(r *ghttp.Request) {
	response.WriteTpl(r, "demo/modal/table.html")
}

func Check(r *ghttp.Request) {
	response.WriteTpl(r, "demo/modal/table/check.html")
}

func Parent(r *ghttp.Request) {
	response.WriteTpl(r, "demo/modal/table/parent.html")
}

func Radio(r *ghttp.Request) {
	response.WriteTpl(r, "demo/modal/table/radio.html")
}
