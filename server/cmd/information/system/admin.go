package information

import (
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/server/library/global"
	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

var Admin = new(admin)

type admin struct{}

var admins = []model.Admin{
	{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Uuid: uuid.NewV4().String(), Username: "admin", Password: "123456", Nickname: "超级管理员", Avatar: "http://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "888"},
	{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Uuid: uuid.NewV4().String(), Username: "a303176530", Password: "123456", Nickname: "QMPlusUser", Avatar: "http://qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "9528"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: admins 表数据初始化
func (a *admin) Init() error {
	_ = admins[0].EncryptedPassword()
	_ = admins[1].EncryptedPassword()
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.Admin{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> admins 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> admins 表初始数据成功!")
		return nil
	})
}
