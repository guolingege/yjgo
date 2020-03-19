package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	"yj-app/app/service/utils/token"
)

// 鉴权中间件，只有登录成功之后才能通过
func JWT(r *ghttp.Request) {
	t := token.Instance()
	if t != nil {
		data, err := t.GetTokenData(r)
		if err != nil {
			r.Response.WriteJsonExit(model.CommonRes{
				Code:  model.UNAUTHORIZED,
				Msg:   err.Error(),
				Data:  nil,
				Btype: 0,
			})
			return
		}

		if data == nil {
			r.Response.WriteJsonExit(model.CommonRes{
				Code:  model.UNAUTHORIZED,
				Msg:   "token无效",
				Data:  nil,
				Btype: 0,
			})
		}

		r.SetParam("token", *data)
		r.Middleware.Next()
	} else {
		r.Response.WriteJsonExit(model.CommonRes{
			Code:  model.UNAUTHORIZED,
			Msg:   "您没有访问权限",
			Data:  nil,
			Btype: 0,
		})
	}
}
