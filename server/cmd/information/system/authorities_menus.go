package information

import (
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/server/library/global"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

type AuthorityMenus struct {
	AuthorityId string
	BaseMenuId  uint
}

var authorityMenus = []model.AuthoritiesMenus{
	{AuthorityId: "888", MenuId: 1},
	{AuthorityId: "888", MenuId: 2},
	{AuthorityId: "888", MenuId: 3},
	{AuthorityId: "888", MenuId: 4},
	{AuthorityId: "888", MenuId: 5},
	{AuthorityId: "888", MenuId: 6},
	{AuthorityId: "888", MenuId: 7},
	{AuthorityId: "888", MenuId: 8},
	{AuthorityId: "888", MenuId: 9},
	{AuthorityId: "888", MenuId: 10},
	{AuthorityId: "888", MenuId: 11},
	{AuthorityId: "888", MenuId: 12},
	{AuthorityId: "888", MenuId: 13},
	{AuthorityId: "888", MenuId: 14},
	{AuthorityId: "888", MenuId: 15},
	{AuthorityId: "888", MenuId: 16},
	{AuthorityId: "888", MenuId: 17},
	{AuthorityId: "888", MenuId: 18},
	{AuthorityId: "888", MenuId: 19},
	{AuthorityId: "888", MenuId: 20},
	{AuthorityId: "888", MenuId: 21},
	{AuthorityId: "888", MenuId: 22},
	{AuthorityId: "888", MenuId: 23},
	{AuthorityId: "888", MenuId: 24},
	{AuthorityId: "888", MenuId: 25},
	{AuthorityId: "888", MenuId: 26},
	{AuthorityId: "888", MenuId: 27},
	{AuthorityId: "888", MenuId: 28},
	{AuthorityId: "888", MenuId: 29},
	{AuthorityId: "8881",MenuId: 1},
	{AuthorityId: "8881",MenuId: 2},
	{AuthorityId: "8881",MenuId: 8},
	{AuthorityId: "9528",MenuId: 1},
	{AuthorityId: "9528",MenuId: 2},
	{AuthorityId: "9528",MenuId: 3},
	{AuthorityId: "9528",MenuId: 4},
	{AuthorityId: "9528",MenuId: 5},
	{AuthorityId: "9528",MenuId: 6},
	{AuthorityId: "9528",MenuId: 7},
	{AuthorityId: "9528",MenuId: 8},
	{AuthorityId: "9528",MenuId: 9},
	{AuthorityId: "9528",MenuId: 10},
	{AuthorityId: "9528",MenuId: 11},
	{AuthorityId: "9528",MenuId: 12},
	{AuthorityId: "9528",MenuId: 14},
	{AuthorityId: "9528",MenuId: 15},
	{AuthorityId: "9528",MenuId: 16},
	{AuthorityId: "9528",MenuId: 17},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: authorities_menus 表数据初始化
func (a *authoritiesMenus) Init() error {
	return global.Db.Table("authorities_menus").Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ('888', '8881', '9528')").Find(&[]AuthorityMenus{}).RowsAffected == 48 {
			color.Danger.Println("\n[Mysql] --> authorities_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorityMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> authorities_menus 表初始数据成功!")
		return nil
	})
}
