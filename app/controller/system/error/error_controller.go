package error

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/utils/response"
)

func Unauth(r *ghttp.Request) {
	response.WriteTpl(r, "error/unauth.html")
}

func Error(r *ghttp.Request) {
	response.WriteTpl(r, "error/500.html")
}

func NotFound(r *ghttp.Request) {
	response.WriteTpl(r, "error/404.html")
}
