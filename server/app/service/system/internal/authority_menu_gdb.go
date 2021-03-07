package internal

import (
	model "gf-vue-admin/server/app/model/system"
	"github.com/gogf/gf/frame/g"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct {
	_authorityMenu model.AuthorityMenu
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取路由总树map
func (a *authorityMenu) GetTreeMap(id string) map[string][]model.AuthorityMenu {
	entities := make([]model.AuthorityMenu, 0, 10)
	if err := g.DB().Table(a._authorityMenu.TableName()).Where(g.Map{"authority_id": id}).Order("sort").Structs(&entities); err != nil {
		g.Log().Error("获取 authority_menu 视图数据失败!", g.Map{"authority_id": id})
	}
	tree := make(map[string][]model.AuthorityMenu, len(entities))
	for _, entity := range entities {
		if parameter := Menu.GetMenusParameters(entity.ID); parameter != nil {
			entity.Parameters = *parameter
		}
		tree[entity.ParentId] = append(tree[entity.ParentId], entity)
	}
	return tree
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取子菜单
func (a *authorityMenu) GetChildrenList(menu *model.AuthorityMenu, treeMap map[string][]model.AuthorityMenu) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		a.GetChildrenList(&menu.Children[i], treeMap)
	}
}
