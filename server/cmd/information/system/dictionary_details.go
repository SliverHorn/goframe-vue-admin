package information

import (
	model "gf-vue-admin/server/app/model/system"
	"gf-vue-admin/server/library/global"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"time"
)

var (
	details          []model.DictionaryDetail
	DictionaryDetail = new(dictionaryDetail)
)

type dictionaryDetail struct{}

func init() {
	_true = new(bool)
	*_true = true
	_false = new(bool)
	*_false = false
	details = []model.DictionaryDetail{
		{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "smallint", Status: _true, Value: 1, Sort: 1, DictionaryID: 2},
		{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "mediumint", Status: _true, Value: 2, Sort: 2, DictionaryID: 2},
		{Model: global.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "int", Status: _true, Value: 3, Sort: 3, DictionaryID: 2},
		{Model: global.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "bigint", Status: _true, Value: 4, Sort: 4, DictionaryID: 2},
		{Model: global.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "date", Status: _true, DictionaryID: 3},
		{Model: global.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "time", Status: _true, Value: 1, Sort: 1, DictionaryID: 3},
		{Model: global.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "year", Status: _true, Value: 2, Sort: 2, DictionaryID: 3},
		{Model: global.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "datetime", Status: _true, Value: 3, Sort: 3, DictionaryID: 3},
		{Model: global.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "timestamp", Status: _true, Value: 5, Sort: 5, DictionaryID: 3},
		{Model: global.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "float", Status: _true, DictionaryID: 4},
		{Model: global.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "double", Status: _true, Value: 1, Sort: 1, DictionaryID: 4},
		{Model: global.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "decimal", Status: _true, Value: 2, Sort: 2, DictionaryID: 4},
		{Model: global.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "char", Status: _true, DictionaryID: 5},
		{Model: global.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "varchar", Status: _true, Value: 1, Sort: 1, DictionaryID: 5},
		{Model: global.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "tinyblob", Status: _true, Value: 2, Sort: 2, DictionaryID: 5},
		{Model: global.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "tinytext", Status: _true, Value: 3, Sort: 3, DictionaryID: 5},
		{Model: global.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "text", Status: _true, Value: 4, Sort: 4, DictionaryID: 5},
		{Model: global.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "blob", Status: _true, Value: 5, Sort: 5, DictionaryID: 5},
		{Model: global.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "mediumblob", Status: _true, Value: 6, Sort: 6, DictionaryID: 5},
		{Model: global.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "mediumtext", Status: _true, Value: 7, Sort: 7, DictionaryID: 5},
		{Model: global.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "longblob", Status: _true, Value: 8, Sort: 8, DictionaryID: 5},
		{Model: global.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "longtext", Status: _true, Value: 9, Sort: 9, DictionaryID: 5},
		{Model: global.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "tinyint", Status: _true, DictionaryID: 6},
	}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: dictionary_details 表数据初始化
func (d *dictionaryDetail) Init() error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 23}).Find(&[]model.DictionaryDetail{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> dictionary_details 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&details).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> dictionary_details 表初始数据成功!")
		return nil
	})
}
