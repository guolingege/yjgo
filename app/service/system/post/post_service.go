package post

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	postModel "yj-app/app/model/system/post"
	userService "yj-app/app/service/system/user"
	"yj-app/app/utils/convert"
	"yj-app/app/utils/excel"
	"yj-app/app/utils/page"
)

//根据主键查询数据
func SelectRecordById(id int64) (*postModel.Entity, error) {
	return postModel.FindOne("post_id", id)
}

//根据主键删除数据
func DeleteRecordById(id int64) bool {
	result, err := postModel.Delete("post_id", id)
	if err == nil {
		affected, _ := result.RowsAffected()
		if affected > 0 {
			return true
		}
	}

	return false
}

//批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := postModel.Delete("post_id in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//添加数据
func AddSave(req *postModel.AddReq, session *ghttp.Session) (int64, error) {
	var post postModel.Entity
	post.PostName = req.PostName
	post.PostCode = req.PostCode
	post.Status = req.Status
	post.PostSort = req.PostSort
	post.Remark = req.Remark
	post.CreateTime = gtime.Now()
	post.CreateBy = ""

	user := userService.GetProfile(session)

	if user != nil {
		post.CreateBy = user.LoginName
	}

	result, err := post.Insert()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil || id <= 0 {
		return 0, err
	}
	return id, nil
}

//修改数据
func EditSave(req *postModel.EditReq, session *ghttp.Session) (int64, error) {

	post, err := postModel.FindOne("post_id=?", req.PostId)

	if err != nil {
		return 0, err
	}

	if post == nil {
		return 0, gerror.New("数据不存在")
	}

	post.PostName = req.PostName
	post.PostCode = req.PostCode
	post.Status = req.Status
	post.Remark = req.Remark
	post.PostSort = req.PostSort
	post.UpdateTime = gtime.Now()
	post.UpdateBy = ""

	user := userService.GetProfile(session)

	if user == nil {
		post.UpdateBy = user.LoginName
	}

	result, err := post.Update()

	if err != nil {
		return 0, err
	}

	rs, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rs, nil
}

//根据条件分页查询角色数据
func SelectListAll(params *postModel.SelectPageReq) ([]postModel.EntityFlag, error) {
	return postModel.SelectListAll(params)
}

//根据条件分页查询角色数据
func SelectListByPage(params *postModel.SelectPageReq) ([]postModel.Entity, *page.Paging, error) {
	return postModel.SelectListByPage(params)
}

// 导出excel
func Export(param *postModel.SelectPageReq) (string, error) {
	result, err := postModel.SelectListExport(param)
	if err != nil {
		return "", err
	}

	head := []string{"岗位序号", "岗位名称", "岗位编码", "岗位排序", "状态"}
	key := []string{"post_id", "post_name", "post_code", "post_sort", "stat"}
	url, err := excel.DownlaodExcel(head, key, result)

	if err != nil {
		return "", err
	}

	return url, nil
}

//根据用户ID查询岗位
func SelectPostsByUserId(userId int64) ([]postModel.EntityFlag, error) {
	var paramsPost *postModel.SelectPageReq
	postAll, err := postModel.SelectListAll(paramsPost)

	if err != nil || postAll == nil {
		return nil, gerror.New("未查询到岗位数据")
	}

	userPost, err := postModel.SelectPostsByUserId(userId)

	if err != nil || userPost == nil {
		return nil, gerror.New("未查询到用户岗位数据")
	} else {
		for i := range postAll {
			for j := range userPost {
				if userPost[j].PostId == postAll[i].PostId {
					postAll[i].Flag = true
					break
				}
			}
		}
	}

	return postAll, nil
}

//检查角色名是否唯一
func CheckPostNameUniqueAll(postName string) string {
	post, err := postModel.CheckPostNameUniqueAll(postName)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 {
		return "1"
	}
	return "0"
}

//检查岗位名称是否唯一
func CheckPostNameUnique(postName string, postId int64) string {
	post, err := postModel.CheckPostNameUnique(postName, postId)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 {
		return "1"
	}
	return "0"
}

//检查岗位编码是否唯一
func CheckPostCodeUniqueAll(postCode string) string {
	post, err := postModel.CheckPostCodeUniqueAll(postCode)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 {
		return "1"
	}
	return "0"
}

//检查岗位编码是否唯一
func CheckPostCodeUnique(postCode string, postId int64) string {
	post, err := postModel.CheckPostCodeUnique(postCode, postId)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 {
		return "1"
	}
	return "0"
}
