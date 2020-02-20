package report

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/utils/response"
)

func Echarts(r *ghttp.Request) {
	response.WriteTpl(r, "demo/report/echarts.html")
}

func Metrics(r *ghttp.Request) {
	response.WriteTpl(r, "demo/report/metrics.html")
}

func Peity(r *ghttp.Request) {
	response.WriteTpl(r, "demo/report/peity.html")
}

func Sparkline(r *ghttp.Request) {
	response.WriteTpl(r, "demo/report/sparkline.html")
}
