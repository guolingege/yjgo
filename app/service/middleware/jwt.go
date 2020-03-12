package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	"yj-app/app/service/gtoken"
)

// 鉴权中间件，只有登录成功之后才能通过
func JWT(r *ghttp.Request) {
	gt := gtoken.Instance()
	if gt != nil {
		resp := gt.GetTokenData(r)
		if resp.Code == 0 {
			r.Middleware.Next()
		} else {
			r.Response.WriteJsonExit(resp)
		}
	} else {
		r.Response.WriteJsonExit(model.CommonRes{
			Code:  model.UNAUTHORIZED,
			Msg:   "您没有访问权限",
			Data:  nil,
			Btype: 0,
		})
	}
}
