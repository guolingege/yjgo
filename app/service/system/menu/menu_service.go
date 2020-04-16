package menu

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"time"
	"yj-app/app/model"
	menuModel "yj-app/app/model/system/menu"
	userService "yj-app/app/service/system/user"
	"yj-app/app/utils/convert"
	"yj-app/app/utils/page"
)

//根据主键查询数据
func SelectRecordById(id int64) (*menuModel.EntityExtend, error) {
	return menuModel.SelectRecordById(id)
}

//根据条件查询数据
func SelectListAll(params *menuModel.SelectPageReq) ([]menuModel.Entity, error) {
	return menuModel.SelectListAll(params)
}

//根据条件分页查询数据
func SelectListPage(params *menuModel.SelectPageReq) ([]menuModel.Entity, *page.Paging, error) {
	return menuModel.SelectListPage(params)
}

//根据主键删除数据
func DeleteRecordById(id int64) bool {
	result, err := menuModel.Delete("menu_id", id)
	if err == nil {
		affected, _ := result.RowsAffected()
		if affected > 0 {
			return true
		}
	}

	return false
}

//添加数据
func AddSave(req *menuModel.AddReq, session *ghttp.Session) (int64, error) {

	var menu menuModel.Entity
	menu.MenuName = req.MenuName
	menu.Visible = req.Visible
	menu.ParentId = req.ParentId
	menu.Remark = ""
	menu.MenuType = req.MenuType
	menu.Url = req.Url
	menu.Perms = req.Perms
	menu.Target = req.Target
	menu.Icon = req.Icon
	menu.OrderNum = req.OrderNum
	menu.CreateTime = gtime.Now()
	menu.CreateBy = ""

	user := userService.GetProfile(session)

	if user == nil {
		menu.CreateBy = user.LoginName
	}

	var err error
	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Table("sys_menu").Insert(menu)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil || id <= 0 {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

//修改数据
func EditSave(req *menuModel.EditReq, session *ghttp.Session) (int64, error) {

	menu, err := menuModel.FindOne("menu_id=?", req.MenuId)

	if err != nil {
		return 0, err
	}

	if menu == nil {
		return 0, gerror.New("角色不存在")
	}

	menu.MenuName = req.MenuName
	menu.Visible = req.Visible
	menu.ParentId = req.ParentId
	menu.Remark = ""
	menu.MenuType = req.MenuType
	menu.Url = req.Url
	menu.Perms = req.Perms
	menu.Target = req.Target
	menu.Icon = req.Icon
	menu.OrderNum = req.OrderNum
	menu.UpdateTime = gtime.Now()
	menu.UpdateBy = ""

	user := userService.GetProfile(session)

	if user == nil {
		menu.UpdateBy = user.LoginName
	}

	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}

	_, err = tx.Table("sys_menu").Update(menu, "menu_id="+gconv.String(menu.MenuId))

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return 1, tx.Commit()
}

//批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := menuModel.Delete("menu_id in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//加载所有菜单列表树
func MenuTreeData(userId int64) (*[]model.Ztree, error) {
	var result *[]model.Ztree
	menuList, err := SelectMenuNormalByUser(userId)
	if err != nil {
		return nil, err
	}
	result, err = InitZtree(menuList, nil, false)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//获取用户的菜单数据
func SelectMenuNormalByUser(userId int64) ([]menuModel.EntityExtend, error) {
	if userService.IsAdmin(userId) {
		return SelectMenuNormalAll()
	} else {
		return SelectMenusByUserId(userId)
	}
}

//获取管理员菜单数据
func SelectMenuNormalAll() ([]menuModel.EntityExtend, error) {
	//从缓存读取
	cache := gcache.Get(model.MENU_CACHE)
	if cache != nil {
		return cache.([]menuModel.EntityExtend), nil
	}

	//从数据库中读取
	var result []menuModel.EntityExtend
	result, err := menuModel.SelectMenuNormalAll()

	if err != nil {
		return nil, err
	}

	for i := range result {
		chilrens := getMenuChildPerms(result, result[i].MenuId)

		for j := range chilrens {
			chilrens2 := getMenuChildPerms(result, chilrens[j].MenuId)
			chilrens[j].Children = chilrens2

			if chilrens[j].Target == "" {
				chilrens[j].Target = "menuItem"
			}
			if chilrens[j].Url == "" {
				chilrens[j].Url = "#"
			}
		}

		if chilrens != nil {
			result[i].Children = chilrens

			if result[i].ParentId != 0 {
				if result[i].Target == "" {
					result[i].Target = "menuItem"
				}

				if result[i].Url == "" {
					result[i].Url = "#"
				}
			}

		}
	}

	//存入缓存
	gcache.Set(model.MENU_CACHE, result, time.Hour)
	return result, nil
}

//根据用户ID读取菜单数据
func SelectMenusByUserId(userId int64) ([]menuModel.EntityExtend, error) {
	var result []menuModel.EntityExtend

	//从缓存读取
	cache := gcache.Get(model.MENU_CACHE + gconv.String(userId))

	if cache != nil {
		return cache.([]menuModel.EntityExtend), nil
	}

	//从数据库中读取
	result, err := menuModel.SelectMenusByUserId(gconv.Int64(userId))

	if err != nil {
		return nil, err
	}

	for i := range result {
		chilrens := getMenuChildPerms(result, result[i].MenuId)

		for j := range chilrens {
			chilrens2 := getMenuChildPerms(result, chilrens[j].MenuId)
			chilrens[j].Children = chilrens2
			if chilrens[j].Target == "" {
				chilrens[j].Target = "menuItem"
			}
			if chilrens[j].Url == "" {
				chilrens[j].Url = "#"
			}
		}

		if chilrens != nil {
			result[i].Children = chilrens
			if result[i].ParentId != 0 {
				if result[i].Target == "" {
					result[i].Target = "menuItem"
				}

				if result[i].Url == "" {
					result[i].Url = "#"
				}
			} else {
				if result[i].Url == "" || result[i].Url == "#" {
					result[i].Target = ""
				}
				if result[i].Url == "" {
					result[i].Url = "#"
				}
			}
		}
	}

	//存入缓存
	gcache.Set(model.MENU_CACHE+gconv.String(userId), result, time.Hour)
	return result, nil
}

//根据父id获取子菜单
func getMenuChildPerms(menus []menuModel.EntityExtend, parentId int64) []menuModel.EntityExtend {
	if menus == nil {
		return nil
	}

	var result []menuModel.EntityExtend
	//得到一级菜单
	for i := range menus {
		if menus[i].ParentId == parentId && (menus[i].MenuType == "M" || menus[i].MenuType == "C") {
			if menus[i].Target == "" {
				menus[i].Target = "menuItem"
			}

			if menus[i].Url == "" {
				menus[i].Url = "#"
			}

			result = append(result, menus[i])
		}
	}

	return result
}

//检查菜单名是否唯一
func CheckMenuNameUniqueAll(menuName string, parentId int64) string {
	menu, err := menuModel.FindOne("menu_name=? and parent_id=?", menuName, parentId)
	if err == nil && menu != nil && menu.MenuId > 0 {
		return "1"
	}
	return "0"
}

//检查菜单名是否唯一
func CheckMenuNameUnique(menuName string, menuId, parentId int64) string {
	menu, err := menuModel.FindOne("menu_name=? and menu_id <> ? and parent_id=?", menuName, menuId, parentId)
	if err == nil && menu != nil && menu.MenuId > 0 {
		return "1"
	}
	return "0"
}

//检查权限键是否唯一
func CheckPermsUniqueAll(perms string) string {
	menu, err := menuModel.FindOne("perms=?", perms)
	if err == nil && menu != nil && menu.MenuId > 0 {
		return "1"
	}
	return "0"
}

//检查权限键是否唯一
func CheckPermsUnique(perms string, menuId int64) string {
	menu, err := menuModel.FindOne("perms=? and menu_id <> ?", perms, menuId)
	if err == nil && menu != nil && menu.MenuId > 0 {
		return "1"
	}
	return "0"
}

//根据角色ID查询菜单
func RoleMenuTreeData(roleId, userId int64) (*[]model.Ztree, error) {
	var result *[]model.Ztree
	menuList, err := SelectMenuNormalByUser(userId)
	if err != nil {
		return nil, err
	}

	if roleId > 0 {
		roleMenuList, err := menuModel.SelectMenuTree(roleId)
		if err != nil || roleMenuList == nil {
			result, err = InitZtree(menuList, nil, true)
		} else {
			result, err = InitZtree(menuList, &roleMenuList, true)
		}
	} else {
		result, err = InitZtree(menuList, nil, true)
	}

	return result, nil
}

//对象转菜单树
func InitZtree(menuList []menuModel.EntityExtend, roleMenuList *[]string, permsFlag bool) (*[]model.Ztree, error) {
	var result []model.Ztree
	isCheck := false
	if roleMenuList != nil && len(*roleMenuList) > 0 {
		isCheck = true
	}

	for _, obj := range menuList {
		var ztree model.Ztree
		ztree.Title = obj.MenuName
		ztree.Id = obj.MenuId
		ztree.Name = transMenuName(obj.MenuName, permsFlag)
		ztree.Pid = obj.ParentId
		if isCheck {
			tmp := gconv.String(obj.MenuId) + obj.Perms
			tmpcheck := false
			for j := range *roleMenuList {
				if strings.Compare((*roleMenuList)[j], tmp) == 0 {
					tmpcheck = true
					break
				}
			}
			ztree.Checked = tmpcheck
		}
		result = append(result, ztree)
	}

	return &result, nil
}

func transMenuName(menuName string, permsFlag bool) string {
	if permsFlag {
		return "<font color=\"#888\">&nbsp;&nbsp;&nbsp;" + menuName + "</font>"
	} else {
		return menuName
	}
}
