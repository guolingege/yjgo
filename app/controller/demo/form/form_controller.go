package form

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/utils/response"
)

func Autocomplete(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/autocomplete.html").WriteTplExtend()
}

func Basic(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/basic.html").WriteTplExtend()
}

func Button(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/button.html").WriteTplExtend()
}

func Cards(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/cards.html").WriteTplExtend()
}

func Datetime(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/datetime.html").WriteTplExtend()
}

func Duallistbox(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/duallistbox.html").WriteTplExtend()
}

func Grid(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/grid.html").WriteTplExtend()
}

func Jasny(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/jasny.html").WriteTplExtend()
}

func Select(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/select.html").WriteTplExtend()
}

func Sortable(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/sortable.html").WriteTplExtend()
}

func Summernote(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/summernote.html").WriteTplExtend()
}

func Tabs_panels(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/tabs_panels.html").WriteTplExtend()
}

func Timeline(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/timeline.html").WriteTplExtend()
}

func Upload(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/upload.html").WriteTplExtend()
}

func Validate(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/validate.html").WriteTplExtend()
}

func Wizard(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/wizard.html").WriteTplExtend()
}
