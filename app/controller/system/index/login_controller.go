package index

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/mojocn/base64Captcha"
	"github.com/mssola/user_agent"
	"yj-app/app/model"
	logininforModel "yj-app/app/model/monitor/logininfor"
	"yj-app/app/model/system/user_online"
	logininforService "yj-app/app/service/monitor/logininfor"
	userService "yj-app/app/service/system/user"
	"yj-app/app/utils/ip"
	"yj-app/app/utils/response"
)

type RegisterReq struct {
	UserName     string `p:"username"  v:"required|length:5,30#请输入账号|账号长度为:min到:max位"`
	Password     string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	ValidateCode string `p:"validateCode" v:"required|length:4,30#请输入验证码|验证码长度不够"`
	IdKey        string `p:"idkey" v:"required|length:4,30#请输入验证码id|验证码id长度不够"`
}

// 登陆页面
func Login(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		r.Response.WriteJson(g.Map{
			"code": 500,
			"msg":  "未登录或登录超时。请重新登录",
		})
		return
	}

	response.BuildTpl(r, "login.html").WriteTpl()
}

// 图形验证码
func CaptchaImage(r *ghttp.Request) {
	//config struct for digits
	//数字验证码配置
	//var configD = base64Captcha.ConfigDigit{
	//	Height:     80,
	//	Width:      240,
	//	MaxSkew:    0.7,
	//	DotCount:   80,
	//	CaptchaLen: 5,
	//}
	//config struct for audio
	//声音验证码配置
	//var configA = base64Captcha.ConfigAudio{
	//	CaptchaLen: 6,
	//	Language:   "zh",
	//}
	//config struct for Character
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	//创建声音验证码
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	//以base64编码
	//base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
	//创建字符公式验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	//创建数字验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	//base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

	r.Response.WriteJsonExit(model.CaptchaRes{
		Code:  0,
		IdKey: idKeyC,
		Data:  base64stringC,
		Msg:   "操作成功",
	})
}

//验证登陆
func CheckLogin(r *ghttp.Request) {
	var req *RegisterReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(model.CommonRes{
			Code: 500,
			Msg:  err.Error(),
		})
	}
	//比对验证码
	verifyResult := base64Captcha.VerifyCaptcha(req.IdKey, req.ValidateCode)

	if !verifyResult {
		r.Response.WriteJsonExit(model.CommonRes{
			Code: 500,
			Msg:  "验证码不正确",
		})
	}

	isLock := logininforService.CheckLock(req.UserName)

	if isLock {
		r.Response.WriteJsonExit(model.CommonRes{
			Code: 500,
			Msg:  "账号已锁定，请30分钟后再试",
		})
	}

	//验证账号密码
	if sessionId, err := userService.SignIn(req.UserName, req.Password, r.Session); err != nil {

		errTimes := logininforService.SetPasswordCounts(req.UserName)

		having := 5 - errTimes

		//记录日志
		var logininfor logininforModel.Entity
		logininfor.LoginName = req.UserName
		logininfor.Ipaddr = r.GetClientIp()

		userAgent := r.Header.Get("User-Agent")
		ua := user_agent.New(userAgent)
		logininfor.Os = ua.OS()
		logininfor.Browser, _ = ua.Browser()
		logininfor.LoginTime = gtime.Now()
		logininfor.LoginLocation = ip.GetCityByIp(logininfor.Ipaddr)
		logininfor.Msg = "账号或密码不正确"
		logininfor.Status = "0"

		logininfor.Insert()

		r.Response.WriteJsonExit(model.CommonRes{
			Code: 500,
			Msg:  "账号或密码不正确,还有" + gconv.String(having) + "次之后账号将锁定",
		})
	} else {
		//保存在线状态

		userAgent := r.Header.Get("User-Agent")
		ua := user_agent.New(userAgent)
		os := ua.OS()
		browser, _ := ua.Browser()
		loginIp := r.GetClientIp()
		loginLocation := ip.GetCityByIp(loginIp)

		var userOnline user_online.Entity
		userOnline.SessionId = sessionId
		userOnline.LoginName = req.UserName
		userOnline.Browser = browser
		userOnline.Os = os
		userOnline.DeptName = ""
		userOnline.Ipaddr = loginIp
		userOnline.ExpireTime = 1440
		userOnline.StartTimestamp = gtime.Now()
		userOnline.LastAccessTime = gtime.Now()
		userOnline.Status = "on_line"
		userOnline.LoginLocation = loginLocation
		userOnline.Delete()
		userOnline.Insert()

		//移除登陆次数记录
		logininforService.RemovePasswordCounts(req.UserName)
		//记录日志
		var logininfor logininforModel.Entity
		logininfor.LoginName = req.UserName
		logininfor.Ipaddr = loginIp

		logininfor.Os = os
		logininfor.Browser = browser
		logininfor.LoginTime = gtime.Now()
		logininfor.LoginLocation = loginLocation
		logininfor.Msg = "登陆成功"
		logininfor.Status = "0"

		logininfor.Insert()

		r.Response.WriteJsonExit(model.CommonRes{
			Code: 0,
			Msg:  "登陆成功",
		})
	}
}
