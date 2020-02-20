package dict_type

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	dictTypeModel "yj-app/app/model/system/dict_type"
	dictTypeService "yj-app/app/service/system/dict_type"
	"yj-app/app/service/utils/response"
)

//列表页
func List(r *ghttp.Request) {
	response.WriteTpl(r, "system/dict/type/list.html")
}

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *dictTypeModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "列表查询", req, model.Buniss_Other, err.Error())
	}
	rows := make([]dictTypeModel.Entity, 0)
	result, page, err := dictTypeService.SelectListByPage(req)

	if err == nil && result != nil {
		rows = *result
	}

	r.Response.WriteJsonExit(model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: page.Total,
		Rows:  rows,
	})
}

//新增页面
func Add(r *ghttp.Request) {
	response.WriteTpl(r, "system/dict/type/add.html")
}

//新增页面保存
func AddSave(r *ghttp.Request) {
	var req *dictTypeModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "新增字典类型", req, model.Buniss_Add, err.Error())
	}

	if dictTypeService.CheckDictTypeUniqueAll(req.DictType) == "1" {
		response.ErrorMsg(r, "新增字典类型", req, model.Buniss_Add, "字典类型已存在")
	}

	rid, err := dictTypeService.AddSave(req, r.Session)

	if err != nil || rid <= 0 {
		response.ErrorAdd(r, "新增字典类型", req)
	}
	response.SucessDataAdd(r, "新增字典类型", req, rid)
}

//修改页面
func Edit(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	if id <= 0 {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "字典类型错误",
		})
		return
	}

	entity, err := dictTypeService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "字典类型不存在",
		})
		return
	}

	response.WriteTpl(r, "system/dict/type/edit.html", g.Map{
		"dict": entity,
	})
}

//修改页面保存
func EditSave(r *ghttp.Request) {
	var req *dictTypeModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorMsg(r, "参数字典类型", req, model.Buniss_Add, err.Error())
	}

	if dictTypeService.CheckDictTypeUnique(req.DictType, req.DictId) == "1" {
		response.ErrorMsg(r, "新增字典类型", req, model.Buniss_Add, "字典类型已存在")
	}

	rs, err := dictTypeService.EditSave(req, r.Session)

	if err != nil || rs <= 0 {
		response.ErrorAdd(r, "修改字典类型", req)
	}
	response.SucessDataAdd(r, "修改字典类型", req, rs)
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorEdit(r, "删除字典类型", req)
	}

	rs := dictTypeService.DeleteRecordByIds(req.Ids)

	if rs > 0 {

		response.SucessDataDel(r, "删除字典类型", req, rs)
	} else {
		response.ErrorDataMsg(r, "删除字典类型", req, model.Buniss_Del, 0, "未删除数据")
	}
}

//数据详情
func Detail(r *ghttp.Request) {
	dictId := r.GetQueryInt64("dictId")
	if dictId <= 0 {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "参数错误",
		})
		return
	}
	dict, _ := dictTypeService.SelectRecordById(dictId)

	if dict == nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "字典类别不存在",
		})
		return
	}

	dictList, _ := dictTypeService.SelectListAll(nil)
	if dictList == nil {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "参数错误2",
		})
		return
	}

	response.WriteTpl(r, "system/dict/data/list.html", g.Map{
		"dict":     dict,
		"dictList": dictList,
	})
}

//选择字典树
func SelectDictTree(r *ghttp.Request) {
	columnId := r.GetQueryInt64("columnId")
	dictType := r.GetQueryString("dictType")
	if columnId <= 0 || dictType == "" {
		response.WriteTpl(r, "error/error.html", g.Map{
			"desc": "参数错误",
		})
		return
	}

	if dictType == "-" {
		dictType = "-"
	}

	var dict dictTypeModel.Entity
	rs := dictTypeService.SelectDictTypeByType(dictType)
	if rs != nil {
		dict = *rs
	}

	response.WriteTpl(r, "system/dict/type/tree.html", g.Map{
		"columnId": columnId,
		"dict":     dict,
	})
}

//导出
func Export(r *ghttp.Request) {
	var req *dictTypeModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorOther(r, "导出Excel", req)
	}
	url, err := dictTypeService.Export(req)

	if err != nil {
		response.ErrorMsg(r, "导出Excel", req, model.Buniss_Other, err.Error())
	}

	response.SucessMsg(r, "导出Excel", req, model.Buniss_Other, url)
}

//检查字典类型是否唯一不包括本参数
func CheckDictTypeUnique(r *ghttp.Request) {
	var req *dictTypeModel.CheckDictTypeReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := dictTypeService.CheckDictTypeUnique(req.DictType, req.DictId)

	r.Response.WritefExit(result)
}

//检查字典类型是否唯一
func CheckDictTypeUniqueAll(r *ghttp.Request) {
	var req *dictTypeModel.CheckDictTypeALLReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteflnExit("1")
	}

	result := dictTypeService.CheckDictTypeUniqueAll(req.DictType)

	r.Response.WritefExit(result)
}

//加载部门列表树结构的数据
func TreeData(r *ghttp.Request) {
	result := dictTypeService.SelectDictTree(nil)
	r.Response.WriteJsonExit(result)
}
