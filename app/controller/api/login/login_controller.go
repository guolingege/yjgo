package login

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"yj-app/app/utils/response"
	"yj-app/app/utils/token"
)

type user struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// @Summary 登陆
// @Description api测试
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CommonRes
// @Router /api/v1/login [get]
func Login(r *ghttp.Request) {
	u := new(user)
	if err := r.Parse(u); err != nil {
		response.ErrorResp(r).SetData(err.Error()).WriteJsonExit()
		return
	}

	//验证用户名和密码
	if u == nil || u.Username == "" || u.Password == "" {
		response.ErrorResp(r).SetData("用户名或密码不正确").WriteJsonExit()
		return
	}

	//获取用户id
	uid := 10

	token, err := token.New(gconv.String(uid)).CreateToken()

	if err != nil {
		response.ErrorResp(r).SetData("Error while signing the token").WriteJsonExit()
		return
	}

	//返回token
	response.SucessResp(r).SetData(token).WriteJsonExit()
}

// @Summary api测试
// @Description api测试
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CommonRes
// @Router /api/v1/loginOut [get]
func Index(r *ghttp.Request) {
	uid := r.GetParam("uid")
	response.SucessResp(r).SetData(uid).WriteJsonExit()
}
