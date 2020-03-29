package menu

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"yj-app/app/utils/page"
)

// Entity is the golang structure for table sys_menu.
type EntityExtend struct {
	Entity
	ParentName string         `json:"parentName"` // 父菜单名称
	Children   []EntityExtend `json:"children"`   // 子菜单
}

//检查菜单名称请求参数
type CheckMenuNameReq struct {
	MenuId   int64  `p:"menuId"  v:"required#菜单ID不能为空"`
	ParentId int64  `p:"parentId"  v:"required#父菜单ID不能为空"`
	MenuName string `p:"menuName"  v:"required#菜单名称不能为空"`
}

//检查菜单名称请求参数
type CheckMenuNameALLReq struct {
	ParentId int64  `p:"parentId"  v:"required#父菜单ID不能为空"`
	MenuName string `p:"menuName"  v:"required#菜单名称不能为空"`
}

//分页请求参数
type SelectPageReq struct {
	MenuName  string `p:"menuName"`      //菜单名称
	Visible   string `p:"visible"`       //状态
	BeginTime string `p:"beginTime"`     //开始时间
	EndTime   string `p:"endTime"`       //结束时间
	PageNum   int    `p:"pageNum"`       //当前页码
	PageSize  int    `p:"pageSize"`      //每页数
	SortName  string `p:"orderByColumn"` //排序字段
	SortOrder string `p:"isAsc"`         //排序方式
}

//新增页面请求参数
type AddReq struct {
	ParentId int64  `p:"parentId"  v:"required#父节点不能为空"`
	MenuType string `p:"menuType"  v:"required#菜单类型不能为空"`
	MenuName string `p:"menuName"  v:"required#菜单名称不能为空"`
	OrderNum int    `p:"orderNum" v:"required#显示排序不能为空"`
	Url      string `p:"url"`
	Icon     string `p:"icon"`
	Target   string `p:"target"`
	Perms    string `p:"perms""`
	Visible  string `p:"visible"`
}

//修改页面请求参数
type EditReq struct {
	MenuId   int64  `p:"menuId" v:"required#主键ID不能为空"`
	ParentId int64  `p:"parentId"  v:"required#父节点不能为空"`
	MenuType string `p:"menuType"  v:"required#菜单类型不能为空"`
	MenuName string `p:"menuName"  v:"required#菜单名称不能为空"`
	OrderNum int    `p:"orderNum" v:"required#显示排序不能为空"`
	Url      string `p:"url"`
	Icon     string `p:"icon"`
	Target   string `p:"target"`
	Perms    string `p:"perms""`
	Visible  string `p:"visible"`
}

//根据主键查询数据
func SelectRecordById(id int64) (*EntityExtend, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}
	var result EntityExtend
	model := db.Table("sys_menu t")
	model.Fields("t.menu_id, t.parent_id, t.menu_name, t.order_num, t.url, t.target, t.menu_type, t.visible, t.perms, t.icon, t.remark,(SELECT menu_name FROM sys_menu WHERE menu_id = t.parent_id) parent_name")
	model.Where("menu_id", id)
	err = model.Struct(&result)
	if err != nil {
		return nil, gerror.New("获取数据失败")
	}
	return &result, nil
}

//根据条件分页查询数据
func SelectListPage(param *SelectPageReq) ([]Entity, *page.Paging, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_menu m")

	if param != nil {
		if param.MenuName != "" {
			model.Where("m.menu_name like ?", "%"+param.MenuName+"%")
		}

		if param.Visible != "" {
			model.Where("m.visible = ", param.Visible)
		}

		if param.BeginTime != "" {
			model.Where("date_format(m.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(m.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	total, err := model.Count()

	if err != nil {
		return nil, nil, gerror.New("读取行数失败")
	}

	page := page.CreatePaging(param.PageNum, param.PageSize, total)

	model.Limit(page.StartNum, page.Pagesize)

	var result []Entity

	err = model.Structs(&result)

	if err != nil {
		return nil, nil, gerror.New("读取数据失败")
	} else {
		return result, page, nil
	}
}

//获取所有数据
func SelectListAll(param *SelectPageReq) ([]Entity, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_menu m")

	if param.MenuName != "" {
		model.Where("m.menu_name like ?", "%"+param.MenuName+"%")
	}

	if param.Visible != "" {
		model.Where("m.visible = ", param.Visible)
	}

	if param.BeginTime != "" {
		model.Where("date_format(m.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
	}

	if param.EndTime != "" {
		model.Where("date_format(m.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
	}
	var result []Entity

	err = model.Structs(&result)
	return result, err
}

// 获取管理员菜单数据
func SelectMenuNormalAll() ([]EntityExtend, error) {
	var result []EntityExtend

	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_menu m")
	model.Fields("distinct m.menu_id, m.parent_id, m.menu_name, m.url, m.visible, ifnull(m.perms,'') as perms, m.target, m.menu_type, m.icon, m.order_num, m.create_time")
	model.Where(" m.visible = 0")
	model.Order("m.parent_id, m.order_num")
	model.Structs(&result)

	if err != nil {
		return nil, gerror.New("读取数据失败")
	} else {
		return result, nil
	}
}

//根据用户ID读取菜单数据
func SelectMenusByUserId(userId int64) ([]EntityExtend, error) {
	var result []EntityExtend

	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_menu m")
	model.LeftJoin("sys_role_menu rm", "m.menu_id = rm.menu_id")
	model.LeftJoin("sys_user_role ur", "rm.role_id = ur.role_id")
	model.LeftJoin("sys_role ro", "ur.role_id = ro.role_id")
	model.Fields("distinct m.menu_id, m.parent_id, m.menu_name, m.url, m.visible, ifnull(m.perms,'') as perms, m.target, m.menu_type, m.icon, m.order_num, m.create_time")
	model.Where("ur.user_id = ? and  m.visible = 0  AND ro.status = 0", userId)
	model.Order("m.parent_id, m.order_num")
	model.Structs(&result)

	if err != nil {
		return nil, gerror.New("读取数据失败")
	} else {
		return result, nil
	}
}

//根据角色ID查询菜单
func SelectMenuTree(roleId int64) ([]string, error) {
	db, err := gdb.Instance()

	var result []string

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}
	model := db.Table("sys_menu m")
	model.LeftJoin("sys_role_menu rm", "m.menu_id = rm.menu_id")
	model.Where("rm.role_id = ?", roleId)
	model.Order("m.parent_id, m.order_num")
	model.Fields("concat(m.menu_id, ifnull(m.perms,'')) as perms")
	rs, err := model.All()
	if err != nil {
		return nil, gerror.New("读取数据失败")
	}

	for _, record := range rs {
		if record["perms"].String() != "" {
			result = append(result, record["perms"].String())
		}
	}

	return result, nil
}
