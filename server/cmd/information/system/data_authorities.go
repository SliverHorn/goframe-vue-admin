package information

import (
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/server/library/global"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var DataAuthorities = new(dataAuthorities)

type dataAuthorities struct{}

type DataAuthority struct {
	AuthorityId   string
	DataAuthority string
}

var infos = []model.DataAuthorities{
	{AuthorityId: "888", DataAuthority: "888"},
	{AuthorityId: "888", DataAuthority: "8881"},
	{AuthorityId: "888", DataAuthority: "9528"},
	{AuthorityId: "9528", DataAuthority: "8881"},
	{AuthorityId: "9528", DataAuthority: "9528"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: data_authorities 表数据初始化
func (d *dataAuthorities) Init() error {
	return global.Db.Table("data_authorities").Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ('888', '9528') ").Find(&[]DataAuthority{}).RowsAffected == 5 {
			color.Danger.Println("\n[Mysql] --> data_authorities 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&infos).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> data_authorities 表初始数据成功!")
		return nil
	})
}
