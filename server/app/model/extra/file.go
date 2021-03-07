package model

import "github.com/flipped-aurora/gf-vue-admin/server/library/global"

type File struct {
	global.Model
	Url  string `json:"url" gorm:"comment:文件地址"`
	Tag  string `json:"tag" gorm:"comment:文件标签"`
	Key  string `json:"key" gorm:"comment:编号"`
	Name string `json:"name" gorm:"comment:文件名"`
}

func (f *File) TableName() string {
	return "files"
}

