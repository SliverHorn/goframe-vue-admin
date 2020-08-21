package service

import (
	"errors"
	"server/app/api/request"
	"server/app/model"
	"server/app/model/authority_menu"
	"server/app/model/menus"
	"server/library/global"
	"server/library/utils"
	"strconv"

	"github.com/gogf/gf/frame/g"
)

// getMenuTreeMap Gets the route total tree Map
// getMenuTreeMap 获取路由总树map
func getMenuTreeMap(authorityId string) (treeMap map[string][]*model.AuthorityMenu, err error) {
	authorityMenus := ([]*model.AuthorityMenu)(nil)
	treeMap = make(map[string][]*model.AuthorityMenu)
	err = g.DB("default").Table("menus m").Safe().RightJoin("authority_menu a", "m.id=a.menu_id").Where(g.Map{"authority_id": authorityId}).Structs(&authorityMenus)
	for _, v := range authorityMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

// getChildrenList Get submenu
// getChildrenList 获取子菜单
func getChildrenList(menu *model.AuthorityMenu, treeMap map[string][]*model.AuthorityMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(menu.Children[i], treeMap)
	}
	return err
}

// getBaseMenuTreeMap 获取路由总树map
func getBaseMenuTreeMap() (treeMap map[string][]*model.BaseMenu, err error) {
	allMenus := ([]*model.BaseMenu)(nil)
	treeMap = make(map[string][]*model.BaseMenu)
	db := g.DB("default").Table("menus").Safe()
	err = db.Order("sort").Structs(&allMenus)
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

// GetMenuTree Gets the dynamic menu tree
// GetMenuTree 获取动态菜单树
func GetMenuTree(authorityId string) (menus []*model.AuthorityMenu, err error) {
	var menuTree map[string][]*model.AuthorityMenu
	menuTree, err = getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(menus[i], menuTree)
	}
	return menus, err
}

// getBaseChildrenList get children of menu
// getBaseChildrenList 获取菜单的子菜单
func getBaseChildrenList(menu *model.BaseMenu, treeMap map[string][]*model.BaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.Id))]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(menu.Children[i], treeMap)
	}
	return err
}

// GetMenuList Get routing pages
// GetMenuList 获取路由分页
func GetMenuList() (list interface{}, total int, err error) {
	var (
		treeMap  map[string][]*model.BaseMenu
		menuList []*model.BaseMenu
	)
	treeMap, err = getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(menuList[i], treeMap)
	}
	return menuList, total, err
}

// GetBaseMenuTree A basic routing tree
// GetBaseMenuTree 获取基础路由树
func GetBaseMenuTree() (menu []*model.BaseMenu, err error) {
	var treeMap map[string][]*model.BaseMenu
	treeMap, err = getBaseMenuTreeMap()
	menu = treeMap["0"]
	for i := 0; i < len(menu); i++ {
		err = getBaseChildrenList(menu[i], treeMap)
	}
	return menu, err
}

// AddMenuAuthority Menus are bound to roles
// AddMenuAuthority 菜单与角色绑定
func AddMenuAuthority(insert *request.AddMenuAuthorityInfo) (err error) {
	_, err = authority_menu.Delete(g.Map{"authority_id": insert.AuthorityId})
	if err != nil {
		return err
	}
	for _, v2 := range insert.Menus {
		_, err = authority_menu.Insert(g.Map{"authority_id": insert.AuthorityId, "menu_id": v2.Id})
	}
	return err
}

// GetMenuAuthority View the current role tree
// GetMenuAuthority 查看当前角色树
func GetMenuAuthority(info *request.AuthorityIdInfo) (authorityMenus []*model.AuthorityMenu, err error) {
	authorityMenus = ([]*model.AuthorityMenu)(nil)
	err = g.DB(global.Db).Table("menus m").Safe().RightJoin("authority_menu a", "m.id=a.menu_id").Where(g.Map{"authority_id": info.AuthorityId}).Structs(&authorityMenus)
	return authorityMenus, err
}

// CreateBaseMenu Increase based routing
// CreateBaseMenu 增加基础路由
func CreateBaseMenu(create *request.CreateBaseMenu) (err error) {
	var menu *menus.Entity
	if menu, err = menus.FindOne(g.Map{"name": create.Name}); err != nil {
		return err
	}
	if menu != nil {
		return errors.New("存在重复name，请修改name")
	}
	insert := &menus.Entity{
		MenuLevel: 0,
		ParentId:  create.ParentId,
		Path:      create.Path,
		Name:      create.Name,
		Hidden:    utils.BoolToInt(create.Hidden),
		Component: create.Component,
		Sort:      create.Sort,
	}
	_, err = menus.Insert(insert)
	return err
}

// DeleteBaseMenu Delete the underlying route
// DeleteBaseMenu 删除基础路由
func DeleteBaseMenu(delete *request.GetById) (err error) {
	var menu *menus.Entity
	if menu, err = menus.FindOne(g.Map{"parent_id": delete.Id}); err != nil {
		return err
	}
	if menu != nil {
		return errors.New("此菜单存在子菜单不可删除")
	}
	_, err = menus.Delete(g.Map{"id": delete.Id})
	db := g.DB(global.Db).Table("authority_menu").Safe()
	_, err = db.Where(g.Map{"menu_id": delete.Id}).Delete()
	return err
}

// UpdateBaseMenu Update the routing
// UpdateBaseMenu 更新路由
func UpdateBaseMenu(update *request.UpdateBaseMenu) (err error) {
	condition := g.Map{"id": update.Id}
	updateDate := g.Map{
		"keep_alive":   update.KeepAlive,
		"default_menu": update.DefaultMenu,
		"parent_id":    update.ParentId,
		"path":         update.Path,
		"name":         update.Name,
		"hidden":       utils.BoolToInt(update.Hidden),
		"component":    update.Component,
		"title":        update.Title,
		"icon":         update.Icon,
		"sort":         update.Sort,
	}
	if menus.RecordNotFound(g.Map{"name": update.Name}) {
		return errors.New("更新失败")
	}
	_, err = menus.Update(updateDate, condition)
	return err
}

// GetBaseMenuById get current menus
// GetBaseMenuById 返回当前选中menu
func GetBaseMenuById(idInfo *request.GetById) (menu *model.BaseMenu, err error) {
	menu = (*model.BaseMenu)(nil)
	db := g.DB(global.Db).Table("menus").Safe()
	err = db.Where(g.Map{"id": idInfo.Id}).Struct(&menu)
	return menu, err
}
