package information

import (
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/server/library/global"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"time"
)

var (
	_true        *bool
	_false       *bool
	Dictionary   = new(dictionary)
	dictionaries []model.Dictionary
)

func init() {
	_true = new(bool)
	*_true = true
	_false = new(bool)
	*_false = false
	dictionaries = []model.Dictionary{
		{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "性别", Type: "sex", Status: _true, Desc: "性别字典"},
		{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库int类型", Type: "int", Status: _true, Desc: "int类型对应的数据库类型"},
		{Model: global.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库时间日期类型", Type: "time.Time", Status: _true, Desc: "数据库时间日期类型"},
		{Model: global.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库浮点型", Type: "float64", Status: _true, Desc: "数据库浮点型"},
		{Model: global.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库字符串", Type: "string", Status: _true, Desc: "数据库字符串"},
		{Model: global.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库bool类型", Type: "bool", Status: _true, Desc: "数据库bool类型"},
	}
}

type dictionary struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: dictionaries 表数据初始化
func (d *dictionary) Init() error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 6}).Find(&[]model.Dictionary{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> dictionaries 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&dictionaries).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> dictionaries 表初始数据成功!")
		return nil
	})
}
