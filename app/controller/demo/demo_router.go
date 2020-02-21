package demo

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/controller/demo/form"
	"yj-app/app/controller/demo/icon"
	"yj-app/app/controller/demo/modal"
	"yj-app/app/controller/demo/operate"
	"yj-app/app/controller/demo/report"
	"yj-app/app/controller/demo/table"
)

//加载路由
func LoadRounter(s *ghttp.Server) {
	s.Group("/demo/form", func(group *ghttp.RouterGroup) {

		group.ALL("/autocomplete", form.Autocomplete)
		group.ALL("/basic", form.Basic)
		group.ALL("/button", form.Button)
		group.ALL("/cards", form.Cards)
		group.ALL("/datetime", form.Datetime)
		group.ALL("/duallistbox", form.Duallistbox)
		group.ALL("/grid", form.Grid)

		group.ALL("/jasny", form.Jasny)
		group.ALL("/select", form.Select)
		group.ALL("/sortable", form.Sortable)
		group.ALL("/summernote", form.Summernote)
		group.ALL("/tabs_panels", form.Tabs_panels)

		group.ALL("/timeline", form.Timeline)
		group.ALL("/upload", form.Upload)
		group.ALL("/validate", form.Validate)
		group.ALL("/wizard", form.Wizard)
	})

	s.Group("/demo/icon", func(group *ghttp.RouterGroup) {

		group.ALL("/fontawesome", icon.Fontawesome)
		group.ALL("/glyphicons", icon.Glyphicons)
	})

	s.Group("/demo/modal", func(group *ghttp.RouterGroup) {

		group.ALL("/dialog", modal.Dialog)
		group.ALL("/form", modal.Form)
		group.ALL("/layer", modal.Layer)
		group.ALL("/table", modal.Table)
		group.ALL("/check", modal.Check)
		group.ALL("/parent", modal.Parent)
		group.ALL("/radio", modal.Radio)
	})

	s.Group("/demo/operate", func(group *ghttp.RouterGroup) {

		group.ALL("/add", operate.Add)
		group.ALL("/detail", operate.Detail)
		group.ALL("/edit", operate.Edit)
		group.ALL("/other", operate.Other)
		group.ALL("/table", operate.Table)
	})

	s.Group("/demo/report", func(group *ghttp.RouterGroup) {

		group.ALL("/echarts", report.Echarts)
		group.ALL("/metrics", report.Metrics)
		group.ALL("/peity", report.Peity)
		group.ALL("/sparkline", report.Sparkline)
	})

	s.Group("/demo/table", func(group *ghttp.RouterGroup) {

		group.ALL("/button", table.Button)
		group.ALL("/child", table.Child)
		group.ALL("/curd", table.Curd)
		group.ALL("/detail", table.Detail)
		group.ALL("list", table.List)

		group.ALL("/editable", table.Editable)
		group.ALL("/event", table.Event)
		group.ALL("/export", table.Export)
		group.ALL("/fixedColumns", table.FixedColumns)
		group.ALL("/footer", table.Footer)
		group.ALL("/groupHeader", table.GroupHeader)

		group.ALL("/image", table.Image)
		group.ALL("/multi", table.Multi)
		group.ALL("/other", table.Other)
		group.ALL("/pageGo", table.PageGo)

		group.ALL("/params", table.Params)
		group.ALL("/remember", table.Remember)
		group.ALL("/recorder", table.Recorder)
		group.ALL("/search", table.Search)
	})
}
