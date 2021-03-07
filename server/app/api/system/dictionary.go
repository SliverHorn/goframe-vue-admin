package api

import (
	"gf-vue-admin/server/app/api/request"
	"gf-vue-admin/server/app/api/response"
	service "gf-vue-admin/server/app/service/system"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Dictionary = new(dictionary)

type dictionary struct{}

// @Tags SystemDictionary
// @Summary 创建Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateDictionary true "Dictionary模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysDictionary/createDictionary [post]
func (d *dictionary) Create(r *ghttp.Request) *response.Response {
	var info request.CreateDictionary
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.Dictionary.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{Data: g.Map{"dictionary": info}, MessageCode: response.SuccessCreated}
}

// @Tags SystemDictionary
// @Summary 用id查询Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.FirstDictionary true "ID或字典英名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysDictionary/findDictionary [get]
func (d *dictionary) First(r *ghttp.Request) *response.Response {
	var info request.FirstDictionary
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if data, err := service.Dictionary.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"resysDictionary": data}, MessageCode: response.SuccessFirst}
	}
}

// @Tags SystemDictionary
// @Summary 更新Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateDictionary true "Dictionary模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysDictionary/updateDictionary [put]
func (d *dictionary) Update(r *ghttp.Request) *response.Response {
	var info request.UpdateDictionary
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := service.Dictionary.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	return &response.Response{MessageCode: response.SuccessUpdated}
}

// @Tags SystemDictionary
// @Summary 删除Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "Dictionary模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDictionary/deleteDictionary [delete]
func (d *dictionary) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.Dictionary.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags SystemDictionary
// @Summary 分页获取Dictionary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchDictionary true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionary/getDictionaryList [get]
func (d *dictionary) GetList(r *ghttp.Request) *response.Response {
	var info request.SearchDictionary
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	if list, total, err := service.Dictionary.GetList(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	} else {
		return &response.Response{Data: response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, MessageCode: response.SuccessGetList}
	}
}
