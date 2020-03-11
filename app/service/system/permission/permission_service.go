package permission

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"time"
	menuModel "yj-app/app/model/system/menu"
	menuService "yj-app/app/service/system/menu"
	userService "yj-app/app/service/system/user"
)

//根据用户id和权限字符串判断是否输出控制按钮
func GetPermiButton(u interface{}, permission, funcName, text, aclassName, iclassName string) string {
	startT := time.Now()     //计算当前时间
	tc := time.Since(startT) //计算耗时
	result := HasPermi(u, permission)

	htmlstr := ""
	if result == "" {
		htmlstr = `<a class="` + aclassName + `" onclick="` + funcName + `" hasPermission="` + permission + `">
                    <i class="` + iclassName + `"></i> ` + text + `
                </a>`
	}
	fmt.Printf("time cost 0= %v\n", tc)
	return htmlstr
}

//根据用户id和权限字符串判断是否有此权限
func HasPermi(u interface{}, permission string) string {
	startT := time.Now() //计算当前时间
	if u == nil {
		return "disabled"
	}

	uid := gconv.Int64(u)

	if uid <= 0 {
		return "disabled"
	}
	//获取权限信息
	var menus []menuModel.EntityExtend
	if userService.IsAdmin(uid) {
		startT := time.Now() //计算当前时间
		menus, _ = menuService.SelectMenuNormalAll()
		tc := time.Since(startT) //计算耗时
		fmt.Printf("time cost 1= %v\n", tc)
	} else {
		menus, _ = menuService.SelectMenusByUserId(uid)
	}

	tc := time.Since(startT) //计算耗时
	fmt.Printf("time cost hasPermi= %v\n", tc)

	if menus != nil && len(menus) > 0 {
		for i := range menus {
			if strings.EqualFold((menus)[i].Perms, permission) {
				return ""
			}
		}
	}

	return "disabled"
}
