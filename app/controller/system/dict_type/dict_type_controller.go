package dict_type

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	dictTypeModel "yj-app/app/model/system/dict_type"
	dictTypeService "yj-app/app/service/system/dict_type"
	"yj-app/app/utils/response"
)

//列表页
func List(r *ghttp.Request) {
	response.BuildTpl(r, "system/dict/type/list.html").WriteTpl()
}

//列表分页数据
func ListAjax(r *ghttp.Request) {
	var req *dictTypeModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("字典类型管理", req).WriteJsonExit()
	}
	rows := make([]dictTypeModel.Entity, 0)
	result, page, err := dictTypeService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response.BuildTable(r, page.Total, rows).WriteJsonExit()
}

//新增页面
func Add(r *ghttp.Request) {
	response.BuildTpl(r, "system/dict/type/add.html").WriteTpl()
}

//新增页面保存
func AddSave(r *ghttp.Request) {
	var req *dictTypeModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
	}

	if dictTypeService.CheckDictTypeUniqueAll(req.DictType) == "1" {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg("字典类型已存在").Log("字典管理", req).WriteJsonExit()
	}

	rid, err := dictTypeService.AddSave(req, r.Session)

	if err != nil || rid <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).Log("字典管理", req).WriteJsonExit()
	}
	response.SucessResp(r).SetData(rid).Log("字典管理", req).WriteJsonExit()
}

//修改页面
func Edit(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	if id <= 0 {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "字典类型错误",
		})
		return
	}

	entity, err := dictTypeService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "字典类型不存在",
		})
		return
	}

	response.BuildTpl(r, "system/dict/type/edit.html").WriteTpl(g.Map{
		"dict": entity,
	})
}

//修改页面保存
func EditSave(r *ghttp.Request) {
	var req *dictTypeModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("字典类型管理", req).WriteJsonExit()
	}

	if dictTypeService.CheckDictTypeUnique(req.DictType, req.DictId) == "1" {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).SetMsg("字典类型已存在").Log("字典类型管理", req).WriteJsonExit()
	}

	rs, err := dictTypeService.EditSave(req, r.Session)

	if err != nil || rs <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Edit).Log("字典类型管理", req).WriteJsonExit()
	}
	response.SucessResp(r).Log("字典类型管理", req).WriteJsonExit()
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
	}

	rs := dictTypeService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(r).SetBtype(model.Buniss_Del).Log("字典管理", req).WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).Log("字典管理", req).WriteJsonExit()
	}
}

//数据详情
func Detail(r *ghttp.Request) {
	dictId := r.GetQueryInt64("dictId")
	if dictId <= 0 {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "参数错误",
		})
		return
	}
	dict, _ := dictTypeService.SelectRecordById(dictId)

	if dict == nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "字典类别不存在",
		})
		return
	}

	dictList, _ := dictTypeService.SelectListAll(nil)
	if dictList == nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "参数错误2",
		})
		return
	}

	response.BuildTpl(r, "system/dict/data/list.html").WriteTpl(g.Map{
		"dict":     dict,
		"dictList": dictList,
	})
}

//选择字典树
func SelectDictTree(r *ghttp.Request) {
	columnId := r.GetQueryInt64("columnId")
	dictType := r.GetQueryString("dictType")
	if columnId <= 0 || dictType == "" {
		response.ErrorTpl(r).WriteTpl(g.Map{
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

	response.BuildTpl(r, "system/dict/type/tree.html").WriteTpl(g.Map{
		"columnId": columnId,
		"dict":     dict,
	})
}

//导出
func Export(r *ghttp.Request) {
	var req *dictTypeModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
	}
	url, err := dictTypeService.Export(req)

	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
	}

	response.SucessResp(r).SetMsg(url).Log("导出Excel", req).WriteJsonExit()
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
