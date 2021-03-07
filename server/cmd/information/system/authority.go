package information

import (
	model "gf-vue-admin/server/app/model/system"
	"gf-vue-admin/server/library/global"
	"github.com/gookit/color"
	"time"

	"gorm.io/gorm"
)

var Authority = new(authority)

type authority struct{}

var authorities = []model.Authority{
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: "普通用户", ParentId: "0", DefaultRouter: "dashboard"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "8881", AuthorityName: "普通用户子角色", ParentId: "888", DefaultRouter: "dashboard"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "9528", AuthorityName: "测试角色", ParentId: "0", DefaultRouter: "dashboard"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: authorities 表数据初始化
func (a *authority) Init() error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ? ", []string{"888", "9528"}).Find(&[]model.Authority{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> authorities 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> authorities 表初始数据成功!")
		return nil
	})
}
