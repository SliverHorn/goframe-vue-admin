package api

import (
	"gf-vue-admin/server/app/api/request"
	"gf-vue-admin/server/app/api/response"
	"gf-vue-admin/server/app/model"
	"gf-vue-admin/server/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var {{.StructName}} = new({{.Abbreviation}})

type {{.Abbreviation}} struct{}

// @Tags {{.StructName}}
// @Summary 创建{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "创建{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /{{.Abbreviation}}/create [post]
func (a *{{.Abbreviation}}) Create(r *ghttp.Request) *response.Response {
	var info model.{{.StructName}}
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.{{.StructName}}.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// @Tags {{.StructName}}
// @Summary 用id查询{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "主键ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /{{.Abbreviation}}/first [get]
func (a *{{.Abbreviation}}) First(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if data, err := service.{{.StructName}}.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"{{.Abbreviation}}": data}, MessageCode: response.SuccessFirst}
	}
}

// @Tags {{.StructName}}
// @Summary 更新{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "{{.Abbreviation}}模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /{{.Abbreviation}}/update [put]
func (a *{{.Abbreviation}}) Update(r *ghttp.Request) *response.Response {
	var info model.{{.StructName}}
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if result, err := service.{{.StructName}}.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	} else {
        return &response.Response{Data: g.Map{"{{.Abbreviation}}": result}, MessageCode: response.SuccessUpdated}
	}
}

// @Tags {{.StructName}}
// @Summary 删除{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "{{.StructName}}模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /{{.Abbreviation}}/delete [delete]
func (a *{{.Abbreviation}}) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.{{.StructName}}.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags {{.StructName}}
// @Summary 批量删除{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /{{.Abbreviation}}/deletes [delete]
func (a *{{.Abbreviation}}) Deletes(r *ghttp.Request) *response.Response {
	var info request.GetByIds
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := service.{{.StructName}}.Deletes(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// @Tags {{.StructName}}
// @Summary 分页获取{{.StructName}}列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Search{{.StructName}} true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /{{.Abbreviation}}/getList [get]
func (a *{{.Abbreviation}}) GetList(r *ghttp.Request) *response.Response {
	var info request.Search{{.StructName}}
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := service.{{.StructName}}.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, MessageCode: response.SuccessGetList}
}
