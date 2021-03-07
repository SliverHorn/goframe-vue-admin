package request

import model "github.com/flipped-aurora/gf-vue-admin/server/app/model/extra"

type BaseFile struct {
	Url  string `json:"url" gorm:"comment:文件地址"`
	Tag  string `json:"tag" gorm:"comment:文件标签"`
	Key  string `json:"key" gorm:"comment:编号"`
	Name string `json:"name" gorm:"comment:文件名"`
}

type CreateFile struct {
	BaseFile
}

func (c *CreateFile) Create() *model.File {
	return &model.File{
		Url:  c.Url,
		Tag:  c.Tag,
		Key:  c.Key,
		Name: c.Name,
	}
}
