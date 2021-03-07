package model

import "gf-vue-admin/server/library/global"

type Api struct {
	global.Model
	Path        string `orm:"path" json:"path" gorm:"comment:api路径"`
	Method      string `orm:"method" json:"method" gorm:"default:POST" gorm:"comment:方法"`
	ApiGroup    string `orm:"api_group" json:"apiGroup" gorm:"comment:api组"`
	Description string `orm:"description" json:"description" gorm:"comment:api中文描述"`
}

func (a *Api) TableName() string {
	return "apis"
}