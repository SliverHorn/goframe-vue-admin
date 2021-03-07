package service

import (
	"database/sql"
	"errors"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/request"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/server/app/service/system/internal"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

var Menu = new(menu)

type menu struct {
	_menu             model.Menu
	_parameter        model.MenuParameter
	_menusParameters  model.MenusParameters
	_authoritiesMenus model.AuthoritiesMenus
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 添加基础路由
func (m *menu) Create(menu *model.Menu) error {
	_, err := g.DB().Table(m._menu.TableName()).Insert(menu)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 返回当前选中menu
func (m *menu) First(info *request.GetById) (menu *model.Menu, err error) {
	var entity model.Menu
	err = g.DB().Table(m._menu.TableName()).Where(info.Condition()).Struct(&entity)
	var params = internal.Menu.GetMenusParameters(entity.ID)
	if params != nil {
		entity.Parameters = *params
	}
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除基础路由
func (m *menu) Delete(info *request.GetById) error {
	var entity model.Menu
	if errors.Is(g.DB().Table(m._menu.TableName()).Where(g.Map{"parent_id": info.Id}).Struct(&entity), sql.ErrNoRows) {
		return response.ErrorHasChildrenMenu
	}
	var authorities = internal.Menu.GetAuthoritiesMenus(entity.ID)
	if authorities != nil {
		entity.Authorities = *authorities
	}
	if _, err := g.DB().Table(m._parameter.TableName()).Delete(info.Condition()); err != nil {
		return err
	}
	if _, err := g.DB().Table(m._authoritiesMenus.TableName()).Delete(info.Condition()); err != nil {
		return err
	}
	return nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 更新路由
func (m *menu) Update(info *request.UpdateMenu) error {
	return g.DB().Transaction(func(tx *gdb.TX) error {
		var entity model.Menu
		if err := tx.Table(m._menu.TableName()).WherePri(info.ID).Struct(&entity); err != nil {
			return err
		}
		if entity.Name != info.Name {
			if !errors.Is(tx.Table(m._menu.TableName()).Where(g.Map{"`id` <> ?": info.ID, "`name`": info.Name}).Struct(&model.Menu{}), sql.ErrNoRows) {
				return response.ErrorUpdateMenuName
			}
		}
		if _, err := g.DB().Table(m._parameter.TableName()).Delete(g.Map{"id": info.ID}); err != nil {
			return err
		}
		if _, err := tx.Table(m._menu.TableName()).Data(info.Update()).Insert(); err != nil {
			return response.ErrorUpdateMenu
		}
		if _, err := tx.Table(m._menusParameters.TableName()).Unscoped().Delete(g.Map{"menu_id": info.ID}); err != nil {
			return err
		}
		for _, parameter := range info.Parameters {
			if _, err := tx.Table(m._menusParameters.TableName()).Data(g.Map{"menu_id": info.ID, "parameter_id": parameter.ID}).Insert(); err != nil {
				return response.ErrorCreateParameters
			}
		}
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取路由分页
func (m *menu) GetList() (list interface{}, total int, err error) {
	var menus []model.Menu
	var treeMap = internal.Menu.GetTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		internal.Menu.GetChildrenList(&menus[i], treeMap)
	}
	return menus, total, err
}

func (m *menu) GetTree() *[]model.Menu {
	tree := internal.Menu.GetTreeMap()
	menus := tree["0"]
	for i := 0; i < len(menus); i++ {
		internal.Menu.GetChildrenList(&menus[i], tree)
	}
	return &menus
}
