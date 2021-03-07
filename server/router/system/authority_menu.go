package router

import (
	"gf-vue-admin/server/app/api/response"
	api "gf-vue-admin/server/app/api/system"
	"gf-vue-admin/server/interfaces"
	"gf-vue-admin/server/router/internal"
	"github.com/gogf/gf/net/ghttp"
)

type _menu struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewMenuRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &_menu{router: router, response: &response.Handler{}}
}

func (m *_menu) Init() {
	group := m.router.Group("/menu").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("addBaseMenu", m.response.Handler()(api.Menu.Create))                         // 新增菜单
		group.POST("getBaseMenuById", m.response.Handler()(api.Menu.First))                      // 根据id获取菜单
		group.POST("updateBaseMenu", m.response.Handler()(api.Menu.Update))                      // 更新菜单
		group.POST("deleteBaseMenu", m.response.Handler()(api.Menu.Delete))                      // 删除菜单
		group.POST("getMenuList", m.response.Handler()(api.Menu.GetList))                        // 分页获取基础menu列表
		group.POST("getBaseMenuTree", m.response.Handler()(api.Menu.GetBaseMenuTree))            // 获取用户动态路由
		group.POST("getMenu", m.response.Handler()(api.AuthorityMenu.GetMenu))                   // 获取菜单树
		group.POST("addMenuAuthority", m.response.Handler()(api.AuthorityMenu.AddMenuAuthority)) // 增加menu和角色关联关系
		group.POST("getMenuAuthority", m.response.Handler()(api.AuthorityMenu.GetMenuAuthority)) // 获取指定角色menu
	}
}
