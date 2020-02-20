package operate

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/utils/response"
)

func Add(r *ghttp.Request) {
	response.WriteTpl(r, "demo/operate/add.html")
}

func Detail(r *ghttp.Request) {
	response.WriteTpl(r, "demo/operate/detail.html")
}

func Edit(r *ghttp.Request) {
	response.WriteTpl(r, "demo/operate/edit.html")
}

func Other(r *ghttp.Request) {
	response.WriteTpl(r, "demo/operate/other.html")
}

func Table(r *ghttp.Request) {
	response.WriteTpl(r, "demo/operate/table.html")
}
