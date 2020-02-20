package icon

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/utils/response"
)

func Fontawesome(r *ghttp.Request) {
	response.WriteTpl(r, "demo/icon/fontawesome.html")
}

func Glyphicons(r *ghttp.Request) {
	response.WriteTpl(r, "demo/icon/glyphicons.html")
}
