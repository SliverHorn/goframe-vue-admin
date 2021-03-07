package api

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/request"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	service "github.com/flipped-aurora/gf-vue-admin/server/app/service/system"
	"github.com/gogf/gf/net/ghttp"
)

var Authority = new(authority)

type authority struct{}

// @Tags SystemAuthority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authority/createAuthority [post]
func (a *authority) Create(r *ghttp.Request) *response.Response {
	var info request.CreateAuthority
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreateAuthority}
	}
	if err := service.Authority.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreateAuthority}
	} else {
		return &response.Response{MessageCode: response.SuccessCreateAuthority}
	}
}

// @Tags SystemAuthority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CopyAuthority true "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拷贝成功"}"
// @Router /authority/copyAuthority [post]
func (a *authority) Copy(r *ghttp.Request) *response.Response {
	var info request.CopyAuthority
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCopyAuthority}
	}
	if err := service.Authority.Copy(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCopyAuthority}
	} else {
		return &response.Response{MessageCode: response.SuccessCopyAuthority}
	}
}

// @Tags SystemAuthority
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authority/updateAuthority [post]
func (a *authority) Update(r *ghttp.Request) *response.Response {
	var info request.UpdateAuthority
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := service.Authority.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	} else {
		return &response.Response{MessageCode: response.SuccessUpdated}
	}
}

// @Tags SystemAuthority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetAuthorityId true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authority/deleteAuthority [post]
func (a *authority) Delete(r *ghttp.Request) *response.Response {
	var info request.GetAuthorityId
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.Authority.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags SystemAuthority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/getAuthorityList [post]
func (a *authority) GetList(r *ghttp.Request) *response.Response {
	var info request.PageInfo
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := service.Authority.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, MessageCode: response.SuccessGetList}
}

// @Tags SystemAuthority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetDataAuthority true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/setDataAuthority [post]
func (a *authority) SetDataAuthority(r *ghttp.Request) *response.Response {
	var info request.SetDataAuthority
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorSetDataAuthority}
	}
	if err := service.Authority.SetDataAuthority(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorSetDataAuthority}
	}
	return &response.Response{MessageCode: response.SuccessSetDataAuthority}
}
