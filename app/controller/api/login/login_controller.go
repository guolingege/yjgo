package login

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/gtoken"
	"yj-app/app/service/utils/response"
)

// @Summary 登陆
// @Description api测试
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CommonRes
// @Router /api/v1/login [get]
func Login(r *ghttp.Request) {
	gt := gtoken.Instance()
	token := gt.Login("18888888888", g.Map{
		"userId": "1000",
		"phone":  "18888888888",
	})
	if token == nil {
		response.ErrorResp(r).SetMsg("生成token失败").WriteJsonExit()
	}
	response.SucessResp(r).SetMsg(token.Token).WriteJsonExit()
}

// @Summary api测试
// @Description api测试
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CommonRes
// @Router /api/v1/loginOut [get]
func LoginOut(r *ghttp.Request) {
	response.SucessResp(r).WriteJsonExit()
}
