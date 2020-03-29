package user

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"io"
	"os"
	"strconv"
	"time"
	"yj-app/app/model"
	userModel "yj-app/app/model/system/user"
	userService "yj-app/app/service/system/user"
	"yj-app/app/utils/response"
)

//用户资料页面
func Profile(r *ghttp.Request) {
	user := userService.GetProfile(r.Session)
	response.BuildTpl(r, "system/user/profile/profile.html").WriteTpl(g.Map{
		"user": user,
	})
}

//修改用户信息
func Update(r *ghttp.Request) {
	var req *userModel.ProfileReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
	}

	err := userService.UpdateProfile(req, r.Session)

	if err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
	} else {
		response.SucessResp(r).SetBtype(model.Buniss_Edit).Log("用户管理", req).WriteJsonExit()
	}
}

//修改用户密码
func UpdatePassword(r *ghttp.Request) {
	var req *userModel.PasswordReq
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
	}

	err := userService.UpdatePassword(req, r.Session)

	if err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
	} else {
		response.SucessResp(r).SetBtype(model.Buniss_Edit).Log("修改用户密码", req).WriteJsonExit()
	}
}

//修改头像页面
func Avatar(r *ghttp.Request) {
	user := userService.GetProfile(r.Session)
	response.BuildTpl(r, "system/user/profile/avatar.html").WriteTpl(g.Map{
		"user": user,
	})
}

//修改密码页面
func EditPwd(r *ghttp.Request) {
	user := userService.GetProfile(r.Session)
	response.BuildTpl(r, "system/user/profile/resetPwd.html").WriteTpl(g.Map{
		"user": user,
	})
}

//检查登陆名是否存在
func CheckLoginNameUnique(r *ghttp.Request) {
	var req *userModel.CheckLoginNameReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := userService.CheckLoginName(req.LoginName)

	if result {
		r.Response.WritefExit("1")
	} else {
		r.Response.WritefExit("0")
	}
}

//检查邮箱是否存在
func CheckEmailUnique(r *ghttp.Request) {
	var req *userModel.CheckEmailReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := userService.CheckEmailUnique(req.UserId, req.Email)

	if result {
		r.Response.WritefExit("1")
	} else {
		r.Response.WritefExit("0")
	}
}

//检查邮箱是否存在
func CheckEmailUniqueAll(r *ghttp.Request) {
	var req *userModel.CheckEmailAllReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := userService.CheckEmailUniqueAll(req.Email)

	if result {
		r.Response.WritefExit("1")
	} else {
		r.Response.WritefExit("0")
	}
}

//检查手机号是否存在
func CheckPhoneUnique(r *ghttp.Request) {
	var req *userModel.CheckPhoneReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(model.CommonRes{
			Code: 500,
			Msg:  err.Error(),
		})
	}

	result := userService.CheckPhoneUnique(req.UserId, req.Phonenumber)

	if result {
		r.Response.WritefExit("1")
	} else {
		r.Response.WritefExit("0")
	}

}

//检查手机号是否存在
func CheckPhoneUniqueAll(r *ghttp.Request) {
	var req *userModel.CheckPhoneAllReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(model.CommonRes{
			Code: 500,
			Msg:  err.Error(),
		})
	}

	result := userService.CheckPhoneUniqueAll(req.Phonenumber)

	if result {
		r.Response.WritefExit("1")
	} else {
		r.Response.WritefExit("0")
	}

}

//校验密码是否正确
func CheckPassword(r *ghttp.Request) {
	var req *userModel.CheckPasswordReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(model.CommonRes{
			Code: 500,
			Msg:  err.Error(),
		})
	}

	user := userService.GetProfile(r.Session)

	result := userService.CheckPassword(user, req.Password)

	if result {
		r.Response.WritefExit("true")
	} else {
		r.Response.WritefExit("false")
	}
}

//保存头像
func UpdateAvatar(r *ghttp.Request) {
	user := userService.GetProfile(r.Session)

	curDir, err := os.Getwd()

	if err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("保存头像", g.Map{"userid": user.UserId}).WriteJsonExit()
	}

	saveDir := curDir + "/public/upload/"

	files := r.GetMultipartFiles("avatarfile")

	if files == nil || len(files) < 1 {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg("没有获取到上传文件").Log("保存头像", g.Map{"userid": user.UserId}).WriteJsonExit()
	}

	item := files[0]

	file, err := item.Open()
	if err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("保存头像", g.Map{"userid": user.UserId}).WriteJsonExit()
	}

	defer file.Close()

	curdate := time.Now().UnixNano()
	filename := user.LoginName + strconv.FormatInt(curdate, 10) + ".png"

	f, err := gfile.Create(saveDir + filename)
	defer f.Close()

	if err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("保存头像", g.Map{"userid": user.UserId}).WriteJsonExit()
	}

	if _, err := io.Copy(f, file); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("保存头像", g.Map{"userid": user.UserId}).WriteJsonExit()
	}

	avatar := "/upload/" + filename

	err = userService.UpdateAvatar(avatar, r.Session)

	if err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("保存头像", g.Map{"userid": user.UserId}).WriteJsonExit()
	} else {
		response.SucessResp(r).SetBtype(model.Buniss_Edit).Log("保存头像", g.Map{"userid": user.UserId}).WriteJsonExit()
	}
}
