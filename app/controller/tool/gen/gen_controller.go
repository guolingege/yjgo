package gen

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"os"
	"strings"
	"yj-app/app/model"
	tableModel "yj-app/app/model/tool/table"
	tableColumnModel "yj-app/app/model/tool/table_column"
	userService "yj-app/app/service/system/user"
	tableService "yj-app/app/service/tool/table"
	"yj-app/app/utils/response"
)

//生成代码列表页面
func Gen(r *ghttp.Request) {
	response.BuildTpl(r, "tool/gen/list.html").WriteTpl()
}

func GenList(r *ghttp.Request) {
	var req *tableModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("生成代码", req).WriteJsonExit()
	}
	rows := make([]tableModel.Entity, 0)
	result, page, err := tableService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response.BuildTable(r, page.Total, rows).WriteJsonExit()
}

//导入数据表
func ImportTable(r *ghttp.Request) {
	response.BuildTpl(r, "tool/gen/importTable.html").WriteTpl()
}

//删除数据
func Remove(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).Log("生成代码", req).WriteJsonExit()
	}

	rs := tableService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(r).SetBtype(model.Buniss_Del).Log("生成代码", req).WriteJsonExit()
	} else {
		response.ErrorResp(r).SetBtype(model.Buniss_Del).Log("生成代码", req).WriteJsonExit()
	}
}

//修改数据
func Edit(r *ghttp.Request) {
	id := r.GetQueryInt64("id")

	if id <= 0 {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "参数错误",
		})
		return
	}

	entity, err := tableService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.ErrorTpl(r).WriteTpl(g.Map{
			"desc": "参数不存在",
		})
		return
	}

	goTypeTpl := tableService.GoTypeTpl()
	queryTypeTpl := tableService.QueryTypeTpl()
	htmlTypeTpl := tableService.HtmlTypeTpl()

	response.BuildTpl(r, "tool/gen/edit.html").WriteTpl(g.Map{
		"table":        entity,
		"goTypeTpl":    goTypeTpl,
		"queryTypeTpl": queryTypeTpl,
		"htmlTypeTpl":  htmlTypeTpl,
	})
}

//修改数据保存
func EditSave(r *ghttp.Request) {
	var req tableModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).SetBtype(model.Buniss_Edit).Log("生成代码", g.Map{"tableName": req.TableName}).WriteJsonExit()
	}
	_, err := tableService.SaveEdit(&req, r.Session)
	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).SetBtype(model.Buniss_Edit).Log("生成代码", g.Map{"tableName": req.TableName}).WriteJsonExit()
	}
	response.SucessResp(r).SetBtype(model.Buniss_Edit).Log("生成代码", g.Map{"tableName": req.TableName}).WriteJsonExit()
}

//预览代码
func Preview(r *ghttp.Request) {
	tableId := r.GetQueryInt64("tableId")
	if tableId <= 0 {
		response.ErrorResp(r).SetMsg("参数错误").WriteJsonExit()
	}

	entity, err := tableService.SelectRecordById(tableId)

	if err != nil || entity == nil {
		response.ErrorResp(r).SetMsg("数据不存在").WriteJsonExit()
	}

	tableService.SetPkColumn(entity, entity.Columns)

	addKey := "vm/html/add.html.vm"
	addValue := ""
	editKey := "vm/html/edit.html.vm"
	editValue := ""

	listKey := "vm/html/list.html.vm"
	listValue := ""
	listTmp := "vm/html/list.html"

	treeKey := "vm/html/tree.html.vm"
	treeValue := ""

	if entity.TplCategory == "tree" {
		listTmp = "vm/html/list-tree.html"
	}

	sqlKey := "vm/sql/menu.sql.vm"
	sqlValue := ""
	entityKey := "vm/go/" + entity.BusinessName + "_entity.go.vm"
	entityValue := ""
	modelKey := "vm/go/" + entity.BusinessName + "_model.go.vm"
	modelValue := ""
	extendKey := "vm/go/" + entity.BusinessName + ".go.vm"
	extendValue := ""
	serviceKey := "vm/go/" + entity.BusinessName + "_service.go.vm"
	serviceValue := ""
	routerKey := "vm/go/" + entity.BusinessName + "_router.go.vm"
	routerValue := ""
	controllerKey := "vm/go/" + entity.BusinessName + "_controller.go.vm"
	controllerValue := ""
	if tmpAdd, err := r.Response.ParseTpl("vm/html/add.html", g.Map{"table": entity}); err == nil {
		addValue = tmpAdd
	}

	if tmpEdit, err := r.Response.ParseTpl("vm/html/edit.html", g.Map{"table": entity}); err == nil {
		editValue = tmpEdit
	}

	if tmpList, err := r.Response.ParseTpl(listTmp, g.Map{"table": entity}); err == nil {
		listValue = tmpList
	}

	if entity.TplCategory == "tree" {
		if tmpTree, err := r.Response.ParseTpl("vm/html/tree.html", g.Map{"table": entity}); err == nil {
			treeValue = tmpTree
		}
	}

	if tmpEntity, err := r.Response.ParseTpl("vm/go/entity.html", g.Map{"table": entity}); err == nil {
		entityValue = tmpEntity
	}

	if tmpModel, err := r.Response.ParseTpl("vm/go/model.html", g.Map{"table": entity}); err == nil {
		modelValue = tmpModel
	}

	if tmpExtend, err := r.Response.ParseTpl("vm/go/extend.html", g.Map{"table": entity}); err == nil {
		extendValue = tmpExtend
	}

	if tmpService, err := r.Response.ParseTpl("vm/go/service.html", g.Map{"table": entity}); err == nil {
		serviceValue = tmpService
	}

	if tmpRouter, err := r.Response.ParseTpl("vm/go/router.html", g.Map{"table": entity}); err == nil {
		routerValue = tmpRouter
	}

	if tmpController, err := r.Response.ParseTpl("vm/go/controller.html", g.Map{"table": entity}); err == nil {
		controllerValue = tmpController
	}

	if tmpSql, err := r.Response.ParseTpl("vm/sql/sql.html", g.Map{"table": entity}); err == nil {
		sqlValue = tmpSql
	}

	if entity.TplCategory == "tree" {
		r.Response.WriteJson(model.CommonRes{
			Code:  0,
			Btype: model.Buniss_Other,
			Data: g.Map{
				addKey:        addValue,
				editKey:       editValue,
				listKey:       listValue,
				treeKey:       treeValue,
				sqlKey:        sqlValue,
				entityKey:     entityValue,
				modelKey:      modelValue,
				extendKey:     extendValue,
				serviceKey:    serviceValue,
				routerKey:     routerValue,
				controllerKey: controllerValue,
			},
		})
	} else {
		r.Response.WriteJson(model.CommonRes{
			Code:  0,
			Btype: model.Buniss_Other,
			Data: g.Map{
				addKey:        addValue,
				editKey:       editValue,
				listKey:       listValue,
				sqlKey:        sqlValue,
				entityKey:     entityValue,
				modelKey:      modelValue,
				extendKey:     extendValue,
				serviceKey:    serviceValue,
				routerKey:     routerValue,
				controllerKey: controllerValue,
			},
		})
	}

}

//生成代码
func GenCode(r *ghttp.Request) {
	tableId := r.GetQueryInt64("tableId")
	if tableId <= 0 {
		response.ErrorResp(r).SetMsg("参数错误").WriteJsonExit()
	}

	entity, err := tableService.SelectRecordById(tableId)

	if err != nil || entity == nil {
		response.ErrorResp(r).SetMsg("数据不存在").WriteJsonExit()
	}

	tableService.SetPkColumn(entity, entity.Columns)

	listTmp := "vm/html/list.html"
	if entity.TplCategory == "tree" {
		listTmp = "vm/html/list-tree.html"
	}

	//获取当前运行时目录
	curDir, err := os.Getwd()

	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("生成代码", g.Map{"tableId": tableId}).WriteJsonExit()
	}

	if tmpAdd, err := r.Response.ParseTpl("vm/html/add.html", g.Map{"table": entity}); err == nil {
		fileName := strings.Join([]string{curDir, "/template/", entity.ModuleName, "/", entity.BusinessName, "/add.html"}, "")
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				f.WriteString(tmpAdd)
			}
			f.Close()
		}
	}

	if tmpEdit, err := r.Response.ParseTpl("vm/html/edit.html", g.Map{"table": entity}); err == nil {
		fileName := strings.Join([]string{curDir, "/template/", entity.ModuleName, "/", entity.BusinessName, "/edit.html"}, "")
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				f.WriteString(tmpEdit)
			}
			f.Close()
		}
	}

	if tmpList, err := r.Response.ParseTpl(listTmp, g.Map{"table": entity}); err == nil {
		fileName := strings.Join([]string{curDir, "/template/", entity.ModuleName, "/", entity.BusinessName, "/list.html"}, "")
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				f.WriteString(tmpList)
			}
			f.Close()
		}
	}

	if entity.TplCategory == "tree" {
		if tmpTree, err := r.Response.ParseTpl("vm/html/tree.html", g.Map{"table": entity}); err == nil {
			fileName := strings.Join([]string{curDir, "/template/", entity.ModuleName, "/", entity.BusinessName, "/", "tree.html"}, "")
			if !gfile.Exists(fileName) {
				f, err := gfile.Create(fileName)
				if err == nil {
					f.WriteString(tmpTree)
				}
				f.Close()
			}
		}
	}

	if tmpEntity, err := r.Response.ParseTpl("vm/go/entity.html", g.Map{"table": entity}); err == nil {
		fileName := strings.Join([]string{curDir, "/app/model/", entity.ModuleName, "/", entity.BusinessName, "/", entity.BusinessName, "_entity.go"}, "")
		if gfile.Exists(fileName) {
			gfile.Remove(fileName)
		}

		f, err := gfile.Create(fileName)
		if err == nil {
			f.WriteString(tmpEntity)
		}
		f.Close()
	}

	if tmpModel, err := r.Response.ParseTpl("vm/go/model.html", g.Map{"table": entity}); err == nil {
		fileName := strings.Join([]string{curDir, "/app/model/", entity.ModuleName, "/", entity.BusinessName, "/", entity.BusinessName, "_model.go"}, "")
		if gfile.Exists(fileName) {
			gfile.Remove(fileName)
		}

		f, err := gfile.Create(fileName)
		if err == nil {
			f.WriteString(tmpModel)
		}
		f.Close()
	}

	if tmpExtend, err := r.Response.ParseTpl("vm/go/extend.html", g.Map{"table": entity}); err == nil {
		fileName := strings.Join([]string{curDir, "/app/model/", entity.ModuleName, "/", entity.BusinessName, "/", entity.BusinessName, ".go"}, "")
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				f.WriteString(tmpExtend)
			}
			f.Close()
		}
	}

	if tmpService, err := r.Response.ParseTpl("vm/go/service.html", g.Map{"table": entity}); err == nil {
		fileName := strings.Join([]string{curDir, "/app/service/", entity.ModuleName, "/", entity.BusinessName, "/", entity.BusinessName, "_service.go"}, "")
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				f.WriteString(tmpService)
			}
			f.Close()
		}
	}

	if tmpRouter, err := r.Response.ParseTpl("vm/go/router.html", g.Map{"table": entity}); err == nil {
		fileName := strings.Join([]string{curDir, "/app/controller/", entity.ModuleName, "/", entity.BusinessName, "_router.go"}, "")
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				f.WriteString(tmpRouter)
			}
			f.Close()
		}
	}

	if tmpController, err := r.Response.ParseTpl("vm/go/controller.html", g.Map{"table": entity}); err == nil {
		fileName := strings.Join([]string{curDir, "/app/controller/", entity.ModuleName, "/", entity.BusinessName, "/", entity.BusinessName, "_controller.go"}, "")
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				f.WriteString(tmpController)
			}
			f.Close()
		}
	}

	if tmpSql, err := r.Response.ParseTpl("vm/sql/sql.html", g.Map{"table": entity}); err == nil {
		fileName := strings.Join([]string{curDir, "/document/sql/", entity.ModuleName, "/", entity.BusinessName, "_menu.sql"}, "")
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				f.WriteString(tmpSql)
			}
			f.Close()
		}
	}
	response.SucessResp(r).Log("生成代码", g.Map{"tableId": tableId}).WriteJsonExit()
}

//查询数据库列表
func DataList(r *ghttp.Request) {
	var req *tableModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("生成代码", req).WriteJsonExit()
	}
	rows := make([]tableModel.Entity, 0)
	result, page, err := tableService.SelectDbTableList(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	r.Response.WriteJsonExit(model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: page.Total,
		Rows:  rows,
	})
}

//导入表结构（保存）
func ImportTableSave(r *ghttp.Request) {
	tables := r.GetFormString("tables")
	if tables == "" {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg("参数错误").Log("生成代码", g.Map{"tables": tables}).WriteJsonExit()
	}

	user := userService.GetProfile(r.Session)
	if user == nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg("登陆超时").Log("生成代码", g.Map{"tables": tables}).WriteJsonExit()
	}

	operName := user.LoginName

	tableArr := strings.Split(tables, ",")
	tableList, err := tableService.SelectDbTableListByNames(tableArr)
	if err != nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("生成代码", g.Map{"tables": tables}).WriteJsonExit()
	}

	if tableList == nil {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg("请选择需要导入的表").Log("生成代码", g.Map{"tables": tables}).WriteJsonExit()
	}

	tableService.ImportGenTable(tableList, operName)
	response.SucessResp(r).Log("导入表结构", g.Map{"tables": tables}).WriteJsonExit()
}

//根据table_id查询表列数据
func ColumnList(r *ghttp.Request) {
	tableId := r.GetQueryInt64("tableId")
	//获取参数
	if tableId <= 0 {
		response.ErrorResp(r).SetBtype(model.Buniss_Add).SetMsg("参数错误").Log("生成代码", g.Map{"tableId": tableId}).WriteJsonExit()
	}
	rows := make([]tableColumnModel.Entity, 0)
	result, err := tableService.SelectGenTableColumnListByTableId(tableId)

	if err == nil && len(result) > 0 {
		rows = result
	}

	r.Response.WriteJsonExit(model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}
