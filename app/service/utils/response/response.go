package response

import (
	"encoding/json"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gview"
	"github.com/gogf/gf/util/gconv"
	"yj-app/app/model"
	operlogService "yj-app/app/service/monitor/operlog"
	configService "yj-app/app/service/system/config"
	"yj-app/app/service/system/dict"
	"yj-app/app/service/system/permission"
)

//输出页面模板
func WriteTpl(r *ghttp.Request, tpl string, params ...gview.Params) error {

	r.GetView().BindFunc("hasPermi", permission.HasPermi)
	r.GetView().BindFunc("getPermiButton", permission.GetPermiButton)
	r.GetView().BindFunc("getDictLabel", dict.GetDictLabel)
	r.GetView().BindFunc("getDictTypeSelect", dict.GetDictTypeSelect)
	r.GetView().BindFunc("getDictTypeRadio", dict.GetDictTypeRadio)
	r.GetView().BindFunc("getDictTypeData", dict.GetDictTypeData)
	ossurl := configService.GetValueByKey("sys.resource.url")
	if ossurl != "" && ossurl != "null" {
		r.Cookie.Set("OssUrl", ossurl)
	} else {
		r.Cookie.Set("OssUrl", "")
	}
	return r.Response.WriteTpl(tpl, params...)
}

//输出json格式的数据到浏览器
func WriteJsonLogExit(r *ghttp.Request, title string, inContent interface{}, outContent model.CommonRes) {
	var inContentStr string
	switch inContent.(type) {
	case string, []byte:
		inContentStr = gconv.String(inContent)
	}
	// Else use json.Marshal function to encode the parameter.
	if b, err := json.Marshal(inContent); err != nil {
		inContentStr = ""
	} else {
		inContentStr = string(b)
	}
	operlogService.Add(r, title, inContentStr, outContent)
	r.Response.WriteJsonExit(outContent)
}

//通用的操作失败响应
func Error(r *ghttp.Request, title string, inContent interface{}, bunissType model.BunissType) {
	responseMsg := model.CommonRes{
		Code:  500,
		Btype: bunissType,
		Msg:   "操作失败",
	}
	WriteJsonLogExit(r, title, inContent, responseMsg)
}

//其它操作失败响应
func ErrorOther(r *ghttp.Request, title string, inContent interface{}) {
	Error(r, title, inContent, model.Buniss_Other)
}

//新增操作失败响应
func ErrorAdd(r *ghttp.Request, title string, inContent interface{}) {
	Error(r, title, inContent, model.Buniss_Add)
}

//修改操作失败响应
func ErrorEdit(r *ghttp.Request, title string, inContent interface{}) {
	Error(r, title, inContent, model.Buniss_Edit)
}

//修改删除失败响应
func ErrorDel(r *ghttp.Request, title string, inContent interface{}) {
	Error(r, title, inContent, model.Buniss_Del)
}

//通用的操作成功响应
func Sucess(r *ghttp.Request, title string, inContent interface{}, bunissType model.BunissType) {
	responseMsg := model.CommonRes{
		Code:  0,
		Btype: bunissType,
		Msg:   "操作成功",
	}
	WriteJsonLogExit(r, title, inContent, responseMsg)
}

//通用的操作成功响应带DATA数据
func SucessData(r *ghttp.Request, title string, inContent interface{}, bunissType model.BunissType, data interface{}) {
	responseMsg := model.CommonRes{
		Code:  0,
		Btype: bunissType,
		Data:  data,
		Msg:   "操作成功",
	}
	WriteJsonLogExit(r, title, inContent, responseMsg)
}

//通用的操作成功响应带MSG数据
func SucessMsg(r *ghttp.Request, title string, inContent interface{}, bunissType model.BunissType, msg string) {
	responseMsg := model.CommonRes{
		Code:  0,
		Btype: bunissType,
		Msg:   msg,
	}
	WriteJsonLogExit(r, title, inContent, responseMsg)
}

//其它操作成功响应
func SucessOther(r *ghttp.Request, title string, inContent interface{}) {
	Sucess(r, title, inContent, model.Buniss_Other)
}

//新增操作成功响应
func SucessAdd(r *ghttp.Request, title string, inContent interface{}) {
	Sucess(r, title, inContent, model.Buniss_Add)
}

//修改操作成功响应
func SucessEdit(r *ghttp.Request, title string, inContent interface{}) {
	Sucess(r, title, inContent, model.Buniss_Edit)
}

//修改删除成功响应
func SucessDel(r *ghttp.Request, title string, inContent interface{}) {
	Sucess(r, title, inContent, model.Buniss_Del)
}

//其它操作成功响应带DATA数据
func SucessDataOther(r *ghttp.Request, title string, inContent, data interface{}) {
	SucessData(r, title, inContent, model.Buniss_Other, data)
}

//新增操作成功响应带DATA数据
func SucessDataAdd(r *ghttp.Request, title string, inContent, data interface{}) {
	SucessData(r, title, inContent, model.Buniss_Add, data)
}

//修改操作成功响应带DATA数据
func SucessDataEdit(r *ghttp.Request, title string, inContent, data interface{}) {
	SucessData(r, title, inContent, model.Buniss_Edit, data)
}

//修改删除成功响应带DATA数据
func SucessDataDel(r *ghttp.Request, title string, inContent, data interface{}) {
	SucessData(r, title, inContent, model.Buniss_Del, data)
}

//通用的操作失败响应
func ErrorDataMsg(r *ghttp.Request, title string, inContent interface{}, bunissType model.BunissType, data interface{}, msg string) {
	responseMsg := model.CommonRes{
		Code:  500,
		Btype: bunissType,
		Data:  data,
		Msg:   msg,
	}
	WriteJsonLogExit(r, title, inContent, responseMsg)
}

//通用的操作失败响应
func ErrorMsg(r *ghttp.Request, title string, inContent interface{}, bunissType model.BunissType, msg string) {
	responseMsg := model.CommonRes{
		Code:  500,
		Btype: bunissType,
		Msg:   msg,
	}
	WriteJsonLogExit(r, title, inContent, responseMsg)
}

func ErrorData(r *ghttp.Request, title string, inContent interface{}, bunissType model.BunissType, data interface{}) {
	ErrorDataMsg(r, title, inContent, bunissType, data, "操作失败")
}

//其它操作失败响应
func ErrorOtherData(r *ghttp.Request, title string, inContent, data interface{}) {
	ErrorData(r, title, inContent, model.Buniss_Other, data)
}

//新增操作失败响应
func ErrorAddData(r *ghttp.Request, title string, inContent, data interface{}) {
	ErrorData(r, title, inContent, model.Buniss_Add, data)
}

//修改操作失败响应
func ErrorEditData(r *ghttp.Request, title string, inContent, data interface{}) {
	ErrorData(r, title, inContent, model.Buniss_Edit, data)
}

//修改删除失败响应
func ErrorDelData(r *ghttp.Request, title string, inContent, data interface{}) {
	ErrorData(r, title, inContent, model.Buniss_Del, data)
}
