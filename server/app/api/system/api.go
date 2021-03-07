package api

import (
	"gf-vue-admin/server/app/api/request"
	"gf-vue-admin/server/app/api/response"
	service "gf-vue-admin/server/app/service/system"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Api = new(api)

type api struct{}

// @Tags SystemApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/createApi [post]
func (a *api) Create(r *ghttp.Request) *response.Response {
	var info request.CreateApi
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.Api.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// @Tags SystemApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func (a *api) First(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if apis, err := service.Api.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"api": apis}, MessageCode: response.SuccessFirst}
	}
}

// @Tags SystemApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /api/updateApi [post]
func (a *api) Update(r *ghttp.Request) *response.Response {
	var info request.UpdateApi
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := service.Api.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	return &response.Response{Data: g.Map{"api": info}, MessageCode: response.SuccessUpdated}
}

// @Tags SystemApi
// @Summary 删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeleteApi true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteApi [post]
func (a *api) Delete(r *ghttp.Request) *response.Response {
	var info request.DeleteApi
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.Api.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{Data: g.Map{"api": info}, MessageCode: response.SuccessDeleted}
}

// @Tags SystemApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchApi true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func (a *api) GetList(r *ghttp.Request) *response.Response {
	var info request.SearchApi
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	if list, total, err := service.Api.GetList(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	} else {
		return &response.Response{Data: response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, MessageCode: response.SuccessGetList}
	}
}

// @Tags SystemApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
func (a *api) GetAllApis(r *ghttp.Request) *response.Response {
	if apis, err := service.Api.GetAllApi(); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	} else {
		return &response.Response{Data: g.Map{"apis": apis}, MessageCode: response.SuccessUpdated}
	}
}
