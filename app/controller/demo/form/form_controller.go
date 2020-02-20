package form

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/utils/response"
)

func Autocomplete(r *ghttp.Request) {
	response.WriteTpl(r, "demo/form/autocomplete.html")
}

func Basic(r *ghttp.Request) {
	response.WriteTpl(r, "demo/form/basic.html")
}

func Button(r *ghttp.Request) {
	response.WriteTpl(r, "demo/form/button.html")
}

func Cards(r *ghttp.Request) {
	response.WriteTpl(r, "demo/form/cards.html")
}

func Datetime(r *ghttp.Request) {
	response.WriteTpl(r, "demo/form/datetime.html")
}

func Duallistbox(r *ghttp.Request) {
	response.WriteTpl(r, "demo/form/duallistbox.html")
}

func Grid(r *ghttp.Request)  {
	response.WriteTpl(r, "demo/form/grid.html")
}

func Jasny(r *ghttp.Request)  {
	response.WriteTpl(r, "demo/form/jasny.html")
}

func Select(r *ghttp.Request)  {
	response.WriteTpl(r, "demo/form/select.html")
}

func Sortable(r *ghttp.Request)  {
	response.WriteTpl(r, "demo/form/sortable.html")
}

func Summernote(r *ghttp.Request)  {
	response.WriteTpl(r, "demo/form/summernote.html")
}

func Tabs_panels(r *ghttp.Request)  {
	response.WriteTpl(r, "demo/form/tabs_panels.html")
}

func Timeline(r *ghttp.Request)  {
	response.WriteTpl(r, "demo/form/timeline.html")
}

func Upload(r *ghttp.Request)  {
	response.WriteTpl(r, "demo/form/upload.html")
}

func Validate(r *ghttp.Request)  {
	response.WriteTpl(r, "demo/form/validate.html")
}

func Wizard(r *ghttp.Request)  {
	response.WriteTpl(r, "demo/form/wizard.html")
}