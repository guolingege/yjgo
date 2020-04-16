package user

import (
	"errors"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"yj-app/app/model"
	userModel "yj-app/app/model/system/user"
	"yj-app/app/model/system/user_online"
	"yj-app/app/model/system/user_post"
	"yj-app/app/model/system/user_role"
	"yj-app/app/utils/convert"
	"yj-app/app/utils/excel"
	"yj-app/app/utils/page"
	"yj-app/app/utils/random"
)

var SessionList = gmap.New(true)

//根据主键查询用户信息
func SelectRecordById(id int64) (*userModel.Entity, error) {
	return userModel.FindOne("user_id", id)
}

// 根据条件分页查询用户列表
func SelectRecordList(param *userModel.SelectPageReq) ([]userModel.UserListEntity, *page.Paging, error) {
	return userModel.SelectPageList(param)
}

// 导出excel
func Export(param *userModel.SelectPageReq) (string, error) {
	result, err := userModel.SelectExportList(param)
	if err != nil {
		return "", err
	}

	head := []string{"用户名", "呢称", "Email", "电话号码", "性别", "部门", "领导", "状态", "删除标记", "创建人", "创建时间", "备注"}
	key := []string{"login_name", "user_name", "email", "phonenumber", "sex", "dept_name", "leader", "status", "del_flag", "create_by", "create_time", "remark"}
	url, err := excel.DownlaodExcel(head, key, result)

	if err != nil {
		return "", err
	}

	return url, nil
}

//新增用户
func AddSave(req *userModel.AddReq, session *ghttp.Session) (int64, error) {
	var user userModel.Entity
	user.LoginName = req.LoginName
	user.UserName = req.UserName
	user.Email = req.Email
	user.Phonenumber = req.Phonenumber
	user.Status = req.Status
	user.Sex = req.Sex
	user.DeptId = req.DeptId
	user.Remark = req.Remark

	//生成密码
	newSalt := random.GenerateSubId(6)
	newToken := req.LoginName + req.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken

	user.CreateTime = gtime.Now()

	createUser := GetProfile(session)

	if createUser != nil {
		user.CreateBy = createUser.LoginName
	}

	user.DelFlag = "0"

	var err error
	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Table("sys_user").Insert(user)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	uid, err := result.LastInsertId()

	if err != nil || uid <= 0 {
		return 0, err
	}

	//增加岗位数据
	if req.PostIds != "" {
		postIds := convert.ToInt64Array(req.PostIds, ",")
		userPosts := make([]user_post.Entity, 0)
		for i := range postIds {
			if postIds[i] > 0 {
				var userPost user_post.Entity
				userPost.UserId = uid
				userPost.PostId = postIds[i]
				userPosts = append(userPosts, userPost)
			}
		}
		if len(userPosts) > 0 {
			_, err := tx.Table("sys_user_post").Insert(userPosts)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}

	}

	//增加角色数据
	if req.RoleIds != "" {
		roleIds := convert.ToInt64Array(req.RoleIds, ",")
		userRoles := make([]user_role.Entity, 0)
		for i := range roleIds {
			if roleIds[i] > 0 {
				var userRole user_role.Entity
				userRole.UserId = uid
				userRole.RoleId = roleIds[i]
				userRoles = append(userRoles, userRole)
			}
		}
		if len(userRoles) > 0 {
			_, err := tx.Table("sys_user_role").Insert(userRoles)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}

	return uid, tx.Commit()
}

//新增用户
func EditSave(req *userModel.EditReq, session *ghttp.Session) (int64, error) {

	user, err := userModel.FindOne("user_id=?", req.UserId)
	if err != nil || user == nil {
		return 0, err
	}

	user.UserName = req.UserName
	user.Email = req.Email
	user.Phonenumber = req.Phonenumber
	user.Status = req.Status
	user.Sex = req.Sex
	user.DeptId = req.DeptId
	user.Remark = req.Remark

	user.UpdateTime = gtime.Now()

	updateUser := GetProfile(session)

	if updateUser != nil {
		user.UpdateBy = updateUser.LoginName
	}

	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}

	_, err = tx.Table("sys_user").Update(user, "user_id="+gconv.String(user.UserId))

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	//增加岗位数据
	if req.PostIds != "" {
		postIds := convert.ToInt64Array(req.PostIds, ",")
		userPosts := make([]user_post.Entity, 0)
		for i := range postIds {
			if postIds[i] > 0 {
				var userPost user_post.Entity
				userPost.UserId = user.UserId
				userPost.PostId = postIds[i]
				userPosts = append(userPosts, userPost)
			}
		}
		if len(userPosts) > 0 {
			tx.Table("sys_user_post").Delete("user_id=?", user.UserId)
			_, err := tx.Table("sys_user_post").Insert(userPosts)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}

	}

	//增加角色数据
	if req.RoleIds != "" {
		roleIds := convert.ToInt64Array(req.RoleIds, ",")
		userRoles := make([]user_role.Entity, 0)
		for i := range roleIds {
			if roleIds[i] > 0 {
				var userRole user_role.Entity
				userRole.UserId = user.UserId
				userRole.RoleId = roleIds[i]
				userRoles = append(userRoles, userRole)
			}
		}
		if len(userRoles) > 0 {
			tx.Table("sys_user_role").Delete("user_id=?", user.UserId)
			_, err := tx.Table("sys_user_role").Insert(userRoles)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}

	return 1, tx.Commit()
}

//根据主键删除用户信息
func DeleteRecordById(id int64) bool {
	result, err := userModel.Delete("user_id", id)
	if err == nil {
		affected, _ := result.RowsAffected()
		if affected > 0 {
			return true
		}
	}

	return false
}

//批量删除用户记录
func DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := userModel.Delete("user_id in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//判断是否是系统管理员
func IsAdmin(userId int64) bool {
	if userId == 1 {
		return true
	} else {
		return false
	}
}

// 判断用户是否已经登录
func IsSignedIn(session *ghttp.Session) bool {
	return session.Contains(model.USER_SESSION_MARK)
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func SignIn(loginnName, password string, session *ghttp.Session) (string, error) {
	//查询用户信息
	user, err := userModel.FindOne("login_name=?", loginnName)

	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("用户或密码不正确")
	}

	//校验密码
	token := user.LoginName + password + user.Salt

	token = gmd5.MustEncryptString(token)

	if user.Password != token {
		return "", errors.New("密码错误")
	}

	session.Set("uid", user.UserId)
	session.Set(model.USER_SESSION_MARK, user)
	sessionId := session.Id()
	SessionList.Set(sessionId, session)
	return sessionId, nil
}

//清空用户菜单缓存
func ClearMenuCache(user *userModel.Entity) {
	if IsAdmin(user.UserId) {
		gcache.Remove(model.MENU_CACHE)
	} else {
		gcache.Get(model.MENU_CACHE + gconv.String(user.UserId))
	}
}

// 用户注销
func SignOut(session *ghttp.Session) error {
	user := GetProfile(session)
	if user != nil {
		ClearMenuCache(user)
	}
	sessionId := session.Id()

	SessionList.Remove(sessionId)
	user_online.Delete("sessionId=?", sessionId)

	session.Remove("uid")
	return session.Remove(model.USER_SESSION_MARK)
}

//强退用户
func ForceLogout(sessionId string) error {
	tmp := SessionList.Get(sessionId)
	if tmp != nil {
		session := tmp.(*ghttp.Session)
		if session != nil {
			return SignOut(session)
		}
	}

	return nil
}

// 检查账号是否符合规范,存在返回false,否则true
func CheckPassport(loginName string) bool {
	if i, err := userModel.FindCount("login_name", loginName); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 检查登陆名是否存在,存在返回true,否则false
func CheckNickName(userName string) bool {
	if i, err := userModel.FindCount("user_name", userName); err != nil {
		return false
	} else {
		return i == 1
	}
}

// 检查登陆名是否存在,存在返回true,否则false
func CheckLoginName(loginName string) bool {
	if i, err := userModel.FindCount("login_name", loginName); err != nil {
		return false
	} else {
		return i == 1
	}
}

// 获得用户信息详情
func GetProfile(session *ghttp.Session) (u *userModel.Entity) {
	_ = session.GetStruct(model.USER_SESSION_MARK, &u)
	return
}

//更新用户信息详情
func UpdateProfile(profile *userModel.ProfileReq, session *ghttp.Session) error {
	user := GetProfile(session)

	if profile.UserName != "" {
		user.UserName = profile.UserName
	}

	if profile.Email != "" {
		user.Email = profile.Email
	}

	if profile.Phonenumber != "" {
		user.Phonenumber = profile.Phonenumber
	}

	if profile.Sex != "" {
		user.Sex = profile.Sex
	}

	_, err := user.Update()
	if err != nil {
		return gerror.New("保存数据失败")
	}

	return session.Set(model.USER_SESSION_MARK, user)
}

//更新用户头像
func UpdateAvatar(avatar string, session *ghttp.Session) error {
	user := GetProfile(session)

	if avatar != "" {
		user.Avatar = avatar
	}

	_, err := user.Update()
	if err != nil {
		return gerror.New("保存数据失败")
	}

	return session.Set(model.USER_SESSION_MARK, user)
}

//修改用户密码
func UpdatePassword(profile *userModel.PasswordReq, session *ghttp.Session) error {
	user := GetProfile(session)

	if strings.Compare(profile.OldPassword, "") == 0 {
		return gerror.New("旧密码不能为空")
	}

	if strings.Compare(profile.NewPassword, "") == 0 {
		return gerror.New("新密码不能为空")
	}

	if strings.Compare(profile.Confirm, "") == 0 {
		return gerror.New("确认密码不能为空")
	}

	if strings.Compare(profile.NewPassword, profile.OldPassword) == 0 {
		return gerror.New("新旧密码不能相同")
	}

	if profile.Confirm != profile.NewPassword {
		return gerror.New("确认密码不一致")
	}

	//校验密码
	token := user.LoginName + profile.OldPassword + user.Salt
	token = gmd5.MustEncryptString(token)

	if token != user.Password {
		return gerror.New("原密码不正确")
	}

	//新校验密码
	newSalt := random.GenerateSubId(6)
	newToken := user.LoginName + profile.NewPassword + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken

	_, err := user.Update()
	if err != nil {
		return gerror.New("保存数据失败")
	}

	return session.Set(model.USER_SESSION_MARK, user)
}

//重置用户密码
func ResetPassword(params *userModel.ResetPwdReq) (bool, error) {

	user, err := userModel.FindOne("user_id=?", params.UserId)

	if err != nil {
		return false, gerror.New("用户不存在")
	}

	//新校验密码
	newSalt := random.GenerateSubId(6)
	newToken := user.LoginName + params.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken

	_, err = user.Update()
	if err != nil {
		return false, gerror.New("保存数据失败")
	}

	return true, nil
}

//校验密码是否正确
func CheckPassword(user *userModel.Entity, password string) bool {
	if user == nil || user.UserId <= 0 {
		return false
	}

	//校验密码
	token := user.LoginName + password + user.Salt
	token = gmd5.MustEncryptString(token)

	if strings.Compare(token, user.Password) == 0 {
		return true
	} else {
		return false
	}
}

//检查邮箱是否已使用
func CheckEmailUnique(userId int64, email string) bool {
	return userModel.CheckEmailUnique(userId, email)
}

//检查邮箱是否存在,存在返回true,否则false
func CheckEmailUniqueAll(email string) bool {
	return userModel.CheckEmailUniqueAll(email)
}

//检查手机号是否已使用,存在返回true,否则false
func CheckPhoneUnique(userId int64, phone string) bool {
	return userModel.CheckPhoneUnique(userId, phone)
}

//检查手机号是否已使用 ,存在返回true,否则false
func CheckPhoneUniqueAll(phone string) bool {
	return userModel.CheckPhoneUniqueAll(phone)
}

//根据登陆名查询用户信息
func SelectUserByLoginName(loginName string) (*userModel.Entity, error) {
	return userModel.SelectUserByLoginName(loginName)
}

//根据手机号查询用户信息
func SelectUserByPhoneNumber(phonenumber string) (*userModel.Entity, error) {
	return userModel.SelectUserByPhoneNumber(phonenumber)
}

// 查询已分配用户角色列表
func SelectAllocatedList(roleId int64, loginName, phonenumber string) ([]userModel.Entity, error) {
	return userModel.SelectAllocatedList(roleId, loginName, phonenumber)
}

// 查询未分配用户角色列表
func SelectUnallocatedList(roleId int64, loginName, phonenumber string) ([]userModel.Entity, error) {
	return userModel.SelectUnallocatedList(roleId, loginName, phonenumber)
}
