package report

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/utils/response"
)

func Echarts(r *ghttp.Request) {
	response.BuildTpl(r, "demo/report/echarts.html").WriteTpl()
}

func Metrics(r *ghttp.Request) {
	response.BuildTpl(r, "demo/report/metrics.html").WriteTpl()
}

func Peity(r *ghttp.Request) {
	response.BuildTpl(r, "demo/report/peity.html").WriteTpl()
}

func Sparkline(r *ghttp.Request) {
	response.BuildTpl(r, "demo/report/sparkline.html").WriteTpl()
}
