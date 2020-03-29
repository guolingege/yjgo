package form

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/utils/response"
)

func Autocomplete(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/autocomplete.html").WriteTpl()
}

func Basic(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/basic.html").WriteTpl()
}

func Button(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/button.html").WriteTpl()
}

func Cards(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/cards.html").WriteTpl()
}

func Datetime(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/datetime.html").WriteTpl()
}

func Duallistbox(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/duallistbox.html").WriteTpl()
}

func Grid(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/grid.html").WriteTpl()
}

func Jasny(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/jasny.html").WriteTpl()
}

func Select(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/select.html").WriteTpl()
}

func Sortable(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/sortable.html").WriteTpl()
}

func Summernote(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/summernote.html").WriteTpl()
}

func Tabs_panels(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/tabs_panels.html").WriteTpl()
}

func Timeline(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/timeline.html").WriteTpl()
}

func Upload(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/upload.html").WriteTpl()
}

func Validate(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/validate.html").WriteTpl()
}

func Wizard(r *ghttp.Request) {
	response.BuildTpl(r, "demo/form/wizard.html").WriteTpl()
}
