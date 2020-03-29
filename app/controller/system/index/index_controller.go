package index

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"io/ioutil"
	"os"
	"yj-app/app/model/system/menu"
	configService "yj-app/app/service/system/config"
	menuService "yj-app/app/service/system/menu"
	userService "yj-app/app/service/system/user"
	"yj-app/app/utils/response"
)

//后台框架首页
func Index(r *ghttp.Request) {
	if userService.IsSignedIn(r.Session) {
		user := userService.GetProfile(r.Session)
		loginname := user.LoginName
		username := user.UserName
		avatar := user.Avatar
		if avatar == "" {
			avatar = "/resource/img/profile.jpg"
		}

		var menus []menu.EntityExtend

		//获取菜单数据
		if userService.IsAdmin(user.UserId) {
			tmp, err := menuService.SelectMenuNormalAll()
			fmt.Println(tmp)
			if err == nil {
				menus = tmp
			}

		} else {
			tmp, err := menuService.SelectMenusByUserId(user.UserId)
			fmt.Println(tmp)
			if err == nil {
				menus = tmp
			}
		}

		//获取配置数据
		sideTheme := configService.GetValueByKey("sys.index.sideTheme")
		skinName := configService.GetValueByKey("sys.index.skinName")

		response.BuildTpl(r, "index.html").WriteTpl(g.Map{
			"avatar":    avatar,
			"loginname": loginname,
			"username":  username,
			"menus":     menus,
			"sideTheme": sideTheme,
			"skinName":  skinName,
		})
	} else {
		r.Response.RedirectTo("/login")
	}
}

//后台框架欢迎页面
func Main(r *ghttp.Request) {
	response.BuildTpl(r, "main.html").WriteTpl()
}

//下载文件
func Download(r *ghttp.Request) {
	fileName := r.GetQueryString("fileName")
	delete := r.GetQueryBool("delete")

	if fileName == "" {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "参数错误",
		})

		return
	}

	// 创建路径
	curDir, err := os.Getwd()
	if err != nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "获取目录失败",
		})
		return
	}

	filepath := curDir + "/public/upload/" + fileName
	file, err := os.Open(filepath)

	defer file.Close()

	if err != nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "参数错误",
		})
		return
	}

	b, _ := ioutil.ReadAll(file)

	r.Response.Header().Add("Content-Disposition", "attachment")
	r.Response.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	r.Response.Write(b)

	if delete {
		os.Remove(filepath)
	}

}

//切换皮肤
func SwitchSkin(r *ghttp.Request) {
	response.BuildTpl(r, "skin.html").WriteTpl()
}

//注销
func Logout(r *ghttp.Request) {
	if userService.IsSignedIn(r.Session) {
		userService.SignOut(r.Session)
	}

	r.Response.RedirectTo("/login")
}
