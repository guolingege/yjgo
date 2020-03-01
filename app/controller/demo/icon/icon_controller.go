package icon

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/utils/response"
)

func Fontawesome(r *ghttp.Request) {
	response.BuildTpl(r, "demo/icon/fontawesome.html").WriteTplExtend()
}

func Glyphicons(r *ghttp.Request) {
	response.BuildTpl(r, "demo/icon/glyphicons.html").WriteTplExtend()
}
