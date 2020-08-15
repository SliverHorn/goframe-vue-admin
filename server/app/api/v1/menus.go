package v1

import (
	"fmt"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/service"
	"server/library/global"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

// GetMenu Gets the user dynamic routing
// GetMenu 获取用户动态路由
func GetMenu(r *ghttp.Request) {
	claims := getAdminClaims(r)
	menus, err := service.GetMenuTree(claims.AdminAuthorityId)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, g.Map{"menus": menus})
}

// GetMenuList Paging gets the base Menu list
// GetMenuList 分页获取基础menu列表
func GetMenuList(r *ghttp.Request) {
	var pageInfo request.PageInfo
	if err := r.Parse(&pageInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetMenuList()
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，err:%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	})
}

// GetBaseMenuTree Gets the user dynamic routing
// GetBaseMenuTree 获取用户动态路由
func GetBaseMenuTree(r *ghttp.Request) {
	menus, err := service.GetBaseMenuTree()
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, g.Map{"menus": menus})
}

// AddMenuAuthority Increases menu and role association
// AddMenuAuthority 增加menu和角色关联关系
func AddMenuAuthority(r *ghttp.Request) {
	var addMenuAuthorityInfo request.AddMenuAuthorityInfo
	if err := r.Parse(&addMenuAuthorityInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	err := service.AddMenuAuthority(&addMenuAuthorityInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("添加失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "添加成功")
}

// GetMenuAuthority Gets the specified role Menu
// GetMenuAuthority 获取指定角色menu
func GetMenuAuthority(r *ghttp.Request) {
	var authorityIdInfo request.AuthorityIdInfo
	if err := r.Parse(&authorityIdInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	menus, err := service.GetMenuAuthority(&authorityIdInfo)
	if err != nil {
		global.FailWithDetailed(r, global.ERROR, g.Map{"menus": menus}, fmt.Sprintf("添加失败，%v", err))
		r.Exit()
	}
	global.Result(r, global.SUCCESS, g.Map{"menus": menus}, "获取成功")
}

// CreateBaseMenu The new menu
// CreateBaseMenu 新增菜单
func CreateBaseMenu(r *ghttp.Request) {
	var create request.CreateBaseMenu
	if err := r.Parse(&create); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.CreateBaseMenu(&create); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("添加失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "添加成功")
}

// DeleteBaseMenu Delete menu
// DeleteBaseMenu 删除菜单
func DeleteBaseMenu(r *ghttp.Request) {
	var deleteInfo request.GetById
	if err := r.Parse(&deleteInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteBaseMenu(&deleteInfo); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除菜单失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// UpdateBaseMenu Update menu
// UpdateBaseMenu 更新菜单
func UpdateBaseMenu(r *ghttp.Request) {
	var updateInfo request.UpdateBaseMenu
	if err := r.Parse(&updateInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.UpdateBaseMenu(&updateInfo); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除菜单失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "修改成功")
}

// GetBaseMenuById Get the menu by ID
// GetBaseMenuById 根据id获取菜单
func GetBaseMenuById(r *ghttp.Request) {
	var idInfo request.GetById
	if err := r.Parse(&idInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	menu, err := service.GetBaseMenuById(&idInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("查询失败，err:%v", err))
		r.Exit()
	}
	global.OkWithData(r, g.Map{"menu": menu})
}
