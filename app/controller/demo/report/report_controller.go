package report

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/utils/response"
)

func Echarts(r *ghttp.Request) {
	response.BuildTpl(r, "demo/report/echarts.html").WriteTplExtend()
}

func Metrics(r *ghttp.Request) {
	response.BuildTpl(r, "demo/report/metrics.html").WriteTplExtend()
}

func Peity(r *ghttp.Request) {
	response.BuildTpl(r, "demo/report/peity.html").WriteTplExtend()
}

func Sparkline(r *ghttp.Request) {
	response.BuildTpl(r, "demo/report/sparkline.html").WriteTplExtend()
}
