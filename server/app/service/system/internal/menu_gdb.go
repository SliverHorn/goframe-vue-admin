package internal

import (
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"
	"github.com/gogf/gf/frame/g"
	"strconv"
)

var Menu = new(menu)

type menu struct {
	_menu             model.Menu
	_parameters       model.MenuParameter
	_menusParameters  model.MenusParameters
	_authoritiesMenus model.AuthoritiesMenus

	_menuParameterMap   map[uint]model.MenuParameter
	_menusParametersMap map[uint][]uint
}

func (m *menu) Init() {
	_menusParametersMap := make(map[uint][]uint)
	entities := make([]model.MenusParameters, 0, 10)
	if err := g.DB().Table(m._menusParameters.TableName()).Structs(&entities); err != nil {
		g.Log().Error("获取 menus_parameters 表数据失败!")
	}
	for _, entity := range entities {
		if value, ok := _menusParametersMap[entity.MenuId]; ok {
			value = append(value, entity.ParameterId)
			_menusParametersMap[entity.MenuId] = value
		} else {
			m1 := make([]uint, 0, 1)
			m1 = append(m1, entity.ParameterId)
			_menusParametersMap[entity.MenuId] = m1
		}
	}

	var parameters []model.MenuParameter
	if err := g.DB().Table(m._parameters.TableName()).Structs(&parameters); err != nil {
		g.Log().Error("获取 menus_parameters 表数据失败!")
	}
	m._menuParameterMap = make(map[uint]model.MenuParameter, len(parameters))
	for _, parameter := range parameters {
		m._menuParameterMap[parameter.MenuID] = parameter
	}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取菜单的Parameters
func (m *menu) GetMenusParameters(id uint) *[]model.MenuParameter {
	m.Init()
	if parameters, ok := m._menusParametersMap[id]; ok {
		p1 := make([]model.MenuParameter, 0, len(parameters))
		for _, parameter := range parameters {
			p1 = append(p1, m._menuParameterMap[parameter])
		}
		return &p1
	} else {
		return nil
	}
}

func (m *menu) GetAuthoritiesMenus(id uint) *[]model.Authority {
	entities := make([]model.AuthoritiesMenus, 0, 10)
	if err := g.DB().Table(m._menusParameters.TableName()).Where(g.Map{"menu_id": id}).Structs(&entities); err != nil {
		g.Log().Error("获取 authorities_menus 表数据失败!", g.Map{"menu_id": id})
		return nil
	}
	authorities := make([]model.Authority, 0, len(entities))
	for _, entity := range entities {
		authorities = append(authorities, Authority().authorityMap[entity.AuthorityId])
	}
	return &authorities
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取菜单的子菜单
func (m *menu) GetChildrenList(menu *model.Menu, treeMap map[string][]model.Menu) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		m.GetChildrenList(&menu.Children[i], treeMap)
	}
	return
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取路由总树map
func (m *menu) GetTreeMap() map[string][]model.Menu {
	var entities []model.Menu
	if err := g.DB().Table(m._menu.TableName()).Structs(&entities); err != nil {
		g.Log().Error("获取 menus 表数据失败!")
		return nil
	}
	tree := make(map[string][]model.Menu, len(entities))
	for _, entity := range entities {
		if parameters := Menu.GetMenusParameters(entity.ID); parameters != nil {
			entity.Parameters = *parameters
		}
		tree[entity.ParentId] = append(tree[entity.ParentId], entity)
	}
	return tree
}
