package api

import (
	"gf-vue-admin/server/app/api/internal"
	"gf-vue-admin/server/app/api/request"
	"gf-vue-admin/server/app/api/response"
	service "gf-vue-admin/server/app/service/system"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var AuthorityMenu =  new(authorityMenu)

type authorityMenu struct {}

// @Tags SystemAuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenu [post]
func (a *authorityMenu) GetMenu(r *ghttp.Request) *response.Response {
	if menus, err := service.AuthorityMenu.GetMenuTree(internal.Info.GetUserAuthorityId(r)); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorOperation}
	} else {
		return &response.Response{Data: g.Map{"menus": menus}, MessageCode: response.SuccessOperation}
	}
}

// @Tags SystemAuthorityMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddMenuAuthority true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addMenuAuthority [post]
func (a *authorityMenu) AddMenuAuthority(r *ghttp.Request) *response.Response {
	var info request.AddMenuAuthority
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorAdd}
	}
	if err := service.AuthorityMenu.AddMenuAuthority(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorAdd}
	}
	return &response.Response{MessageCode: response.SuccessAdd}
}

// @Tags SystemAuthorityMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetAuthorityId true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/GetMenuAuthority [post]
func (a *authorityMenu) GetMenuAuthority(r *ghttp.Request) *response.Response {
	var info request.GetAuthorityId
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	if menus, err := service.AuthorityMenu.GetMenuAuthority(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	} else {
		return &response.Response{Data: g.Map{"menus": menus}, MessageCode: response.SuccessGetList}
	}
}