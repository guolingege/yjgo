package login

import (
	"encoding/json"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/utils/response"
	"yj-app/app/service/utils/token"
)

// @Summary 登陆
// @Description api测试
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CommonRes
// @Router /api/v1/login [get]
func Login(r *ghttp.Request) {
	gt := token.Instance()

	user := g.Map{
		"userId": "1000",
		"phone":  "18888888888",
	}

	userStr := ""

	if arr, err := json.Marshal(user); err == nil {
		userStr = string(arr)
	}

	token := gt.Login("18888888888", userStr)
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
func Index(r *ghttp.Request) {
	token := r.GetParam("token").(token.TokenContent)
	response.SucessResp(r).SetData(token).WriteJsonExit()
}
