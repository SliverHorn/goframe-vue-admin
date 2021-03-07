package model

import "gf-vue-admin/server/library/global"

type DictionaryDetail struct {
	global.Model
	Label        string `orm:"label" json:"label" form:"label" gorm:"column:label;comment:展示值"`
	Status       *bool  `orm:"status" json:"status" form:"status" gorm:"column:status;comment:启用状态"`
	Value        int    `orm:"value" json:"value" form:"value" gorm:"column:value;comment:字典值"`
	Sort         int    `orm:"sort" json:"sort" form:"sort" gorm:"column:sort;comment:排序标记"`
	DictionaryID int    `orm:"dictionary_id" json:"sysDictionaryID" gorm:"column:dictionary_id;comment:关联标记"`
}

func (d *DictionaryDetail) TableName() string {
	return "dictionary_details"
}
