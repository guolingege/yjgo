package response

import (
	"encoding/json"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"yj-app/app/model"
	operlogService "yj-app/app/service/monitor/operlog"
)

// 通用api响应
type ApiResp struct {
	c *model.CommonRes
	r *ghttp.Request
}

//返回一个成功的消息体
func SucessResp(r *ghttp.Request) *ApiResp {
	msg := model.CommonRes{
		Code:  0,
		Btype: model.Buniss_Other,
		Msg:   "操作成功",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

//返回一个错误的消息体
func ErrorResp(r *ghttp.Request) *ApiResp {
	msg := model.CommonRes{
		Code:  500,
		Btype: model.Buniss_Other,
		Msg:   "操作失败",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

//返回一个拒绝访问的消息体
func ForbiddenResp(r *ghttp.Request) *ApiResp {
	msg := model.CommonRes{
		Code:  403,
		Btype: model.Buniss_Other,
		Msg:   "无操作权限",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

//设置消息体的内容
func (resp *ApiResp) SetMsg(msg string) *ApiResp {
	resp.c.Msg = msg
	return resp
}

//设置消息体的编码
func (resp *ApiResp) SetCode(code int) *ApiResp {
	resp.c.Code = code
	return resp
}

//设置消息体的数据
func (resp *ApiResp) SetData(data interface{}) *ApiResp {
	resp.c.Data = data
	return resp
}

//设置消息体的业务类型
func (resp *ApiResp) SetBtype(btype model.BunissType) *ApiResp {
	resp.c.Btype = btype
	return resp
}

//记录操作日志到数据库
func (resp *ApiResp) Log(title string, inParam interface{}) *ApiResp {
	var inContentStr string
	switch inParam.(type) {
	case string, []byte:
		inContentStr = gconv.String(inParam)
	}
	// Else use json.Marshal function to encode the parameter.
	if b, err := json.Marshal(inParam); err != nil {
		inContentStr = ""
	} else {
		inContentStr = string(b)
	}
	operlogService.Add(resp.r, title, inContentStr, resp.c)
	return resp
}

//输出json到客户端
func (resp *ApiResp) WriteJsonExit() {
	resp.r.Response.WriteJsonExit(resp.c)
}
