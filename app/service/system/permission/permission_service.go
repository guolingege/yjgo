package permission

import (
	"github.com/gogf/gf/util/gconv"
	"strings"
	menuModel "yj-app/app/model/system/menu"
	menuService "yj-app/app/service/system/menu"
	userService "yj-app/app/service/system/user"
)

//根据用户id和权限字符串判断是否输出控制按钮
func GetPermiButton(u interface{}, permission, funcName, text, aclassName, iclassName string) string {

	result := HasPermi(u, permission)

	htmlstr := ""
	if result == "" {
		htmlstr = `<a class="` + aclassName + `" onclick="` + funcName + `" hasPermission="` + permission + `">
                    <i class="` + iclassName + `"></i> ` + text + `
                </a>`
	}

	return htmlstr
}

//根据用户id和权限字符串判断是否有此权限
func HasPermi(u interface{}, permission string) string {
	if u == nil {
		return "disabled"
	}

	uid := gconv.Int64(u)

	if uid <= 0 {
		return "disabled"
	}
	//获取权限信息
	var menus *[]menuModel.EntityExtend
	if userService.IsAdmin(uid) {
		menus, _ = menuService.SelectMenuNormalAll()
	} else {
		menus, _ = menuService.SelectMenusByUserId(gconv.String(uid))
	}

	if menus != nil && len(*menus) > 0 {
		for i := range *menus {
			if strings.EqualFold((*menus)[i].Perms, permission) {
				return ""
			}
		}
	}

	return "disabled"
}
