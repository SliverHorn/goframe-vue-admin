package model

import "gf-vue-admin/server/library/global"

type Menu struct {
	global.Model
	Path      string `json:"path" gorm:"comment:路由path"`
	Name      string `json:"name" gorm:"comment:路由name"`
	ParentId  string `json:"parentId" gorm:"comment:父菜单ID"`
	Component string `json:"component" gorm:"comment:对应前端文件路径"`
	Sort      int    `json:"sort" gorm:"comment:排序标记"`
	MenuLevel uint   `json:"-"`
	Hidden    bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`

	Meta `json:"meta" gorm:"comment:附加属性"`

	Children    []Menu          `orm:"-" json:"children" gorm:"-"`
	Parameters  []MenuParameter `orm:"-" json:"parameters" gorm:"many2many:menus_parameters;foreignKey:ID;joinForeignKey:MenuID;References:MenuID;JoinReferences:ParameterID"`
	Authorities []Authority     `orm:"-" json:"authoritys" gorm:"many2many:authorities_menus;foreignKey:ID;joinForeignKey:MenuID;References:AuthorityId;JoinReferences:AuthorityId"`
}

func (b *Menu) TableName() string {
	return "menus"
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"`
	Icon        string `json:"icon" gorm:"comment:菜单图标"`
	Title       string `json:"title" gorm:"comment:菜单名"`
}

type MenuParameter struct {
	global.Model
	Key    string `orm:"key" json:"key" gorm:"comment:地址栏携带参数的key"`
	Type   string `orm:"type" json:"type" gorm:"comment:地址栏携带参数为params还是query"`
	Value  string `orm:"value" json:"value" gorm:"comment:地址栏携带参数的值"`
	MenuID uint   `orm:"menu_id" json:"menuId" gorm:"comment:menu的id"`
}

func (b *MenuParameter) TableName() string {
	return "menu_parameter"
}

type MenusParameters struct {
	MenuId      uint
	ParameterId uint
}

func (m *MenusParameters) TableName() string {
	return "menus_parameters"
}
