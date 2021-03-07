package model

import "github.com/flipped-aurora/gf-vue-admin/server/library/global"

type Dictionary struct {
	global.Model
	Name              string             `orm:"name"  json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`
	Type              string             `orm:"type" json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`
	Status            *bool              `orm:"status" json:"status" form:"status" gorm:"column:status;comment:状态"`
	Desc              string             `orm:"desc" json:"desc" form:"desc" gorm:"column:desc;comment:描述"`
	DictionaryDetails []DictionaryDetail `orm:"-" json:"sysDictionaryDetails" form:"sysDictionaryID" gorm:"foreignKey:DictionaryID;references:ID"`
}

func (d *Dictionary) TableName() string {
	return "dictionaries"
}
