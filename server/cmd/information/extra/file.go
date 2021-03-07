package information

import (
	model "gf-vue-admin/server/app/model/extra"
	"gf-vue-admin/server/library/global"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"time"
)

var File = new(file)

type file struct{}

var files = []model.File{
	{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Url: "http://qmplusimg.henrongyi.top/gvalogo.png", Tag: "png", Key: "158787308910.png", Name: "10.png"},
	{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Url: "http://qmplusimg.henrongyi.top/1576554439myAvatar.png", Tag: "png", Key: "1587973709logo.png", Name: "logo.png"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: files 表初始化数据
func (f *file) Init() error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.File{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> files 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&files).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> files 表初始数据成功!")
		return nil
	})
}
