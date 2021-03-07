package api

import (
	"gf-vue-admin/server/app/api/request"
	"gf-vue-admin/server/app/api/response"
	service "gf-vue-admin/server/app/service/system"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var DictionaryDetail = new(detail)

type detail struct{}

// @Tags SystemDictionaryDetail
// @Summary 创建DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateDictionaryDetail true "SysDictionaryDetail模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysDictionaryDetail/createSysDictionaryDetail [post]
func (d *detail) Create(r *ghttp.Request) *response.Response {
	var info request.CreateDictionaryDetail
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.DictionaryDetail.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// @Tags SystemDictionaryDetail
// @Summary 用id查询DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用id查询DictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取一条数据成功!"}"
// @Router /sysDictionaryDetail/findSysDictionaryDetail [get]
func (d *detail) First(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if data, err := service.DictionaryDetail.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"resysDictionaryDetail": data}, MessageCode: response.SuccessFirst}
	}
}

// @Tags SystemDictionaryDetail
// @Summary 更新DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateDictionaryDetail true "更新SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功!"}"
// @Router /sysDictionaryDetail/updateSysDictionaryDetail [put]
func (d *detail) Update(r *ghttp.Request) *response.Response {
	var info request.UpdateDictionaryDetail
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := service.DictionaryDetail.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	return &response.Response{MessageCode: response.SuccessUpdated}
}

// @Tags SystemDictionaryDetail
// @Summary 删除DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDictionaryDetail/deleteSysDictionaryDetail [delete]
func (d *detail) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.DictionaryDetail.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessUpdated}

}

// @Tags SystemDictionaryDetail
// @Summary 分页获取DictionaryDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchDictionaryDetail true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionaryDetail/getSysDictionaryDetailList [get]
func (d *detail) GetList(r *ghttp.Request) *response.Response {
	var info request.SearchDictionaryDetail
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	if list, total, err := service.DictionaryDetail.GetList(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	} else {
		return &response.Response{Data: response.PageResult{
			List:     list,
			Total:    total,
			Page:     info.Page,
			PageSize: info.PageSize}, MessageCode: response.SuccessGetList,
		}
	}
}
