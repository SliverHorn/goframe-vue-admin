package api

import (
	"gf-vue-admin/server/app/api/request"
	"gf-vue-admin/server/app/api/response"
	model "gf-vue-admin/server/app/model/system"
	service "gf-vue-admin/server/app/service/system"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Menu = new(menu)

type menu struct{}

// @Tags SystemMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuTree [post]
func (m *menu) GetBaseMenuTree(r *ghttp.Request) *response.Response {
	menus := service.Menu.GetTree()
	return &response.Response{Data: g.Map{"menus": menus}, MessageCode: response.SuccessOperation}
}

// @Tags SystemMenu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Menu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addBaseMenu [post]
func (m *menu) Create(r *ghttp.Request) *response.Response {
	var info model.Menu
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorAdd}
	}
	if err := service.Menu.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorAdd}
	}
	return &response.Response{MessageCode: response.SuccessAdd}
}

// @Tags SystemMenu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menu/deleteBaseMenu [post]
func (m *menu) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.Menu.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}

}

// @Tags SystemMenu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /menu/updateBaseMenu [post]
func (m *menu) Update(r *ghttp.Request) *response.Response {
	var info request.UpdateMenu
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := service.Menu.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	return &response.Response{MessageCode: response.SuccessUpdated}

}

// @Tags SystemMenu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuById [post]
func (m *menu) First(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorOperation}
	}
	if data, err := service.Menu.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorOperation}
	} else {
		return &response.Response{Data: g.Map{"menu": data}, MessageCode: response.SuccessOperation}
	}
}

// @Tags SystemMenu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func (m *menu) GetList(r *ghttp.Request) *response.Response {
	var info request.PageInfo
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorOperation}
	}
	if list, total, err := service.Menu.GetList(); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorOperation}
	} else {
		return &response.Response{Data: response.PageResult{
			List:     list,
			Total:    total,
			Page:     info.Page,
			PageSize: info.PageSize,
		}, MessageCode: response.SuccessGetList}
	}
}
