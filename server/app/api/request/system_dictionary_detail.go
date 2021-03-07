package request

import (
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"
	"github.com/gogf/gf/frame/g"
)

type BaseDictionaryDetail struct {
	Label string `p:"label" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"` // 展示值

	Status *bool `p:"status" v:"boolean@required#请输入启用状态"` // 启用状态

	Value        int `p:"value" v:"value@required|length:1,1000#请输入字典值|展示值长度为:min到:max位"`            // 字典值
	Sort         int `p:"sort" v:"label@required|length:1,1000#请输入排序标记|展示值长度为:min到:max位"`            // 排序标记
	DictionaryId int `p:"sysDictionaryID" v:"label@required|length:1,1000#请输入关联标记|展示值长度为:min到:max位"` // 关联标记
}

type CreateDictionaryDetail struct {
	BaseDictionaryDetail
}

func (c *CreateDictionaryDetail) Create() *model.DictionaryDetail {
	return &model.DictionaryDetail{
		Label:        c.Label,
		Status:       c.Status,
		Value:        c.Value,
		Sort:         c.Sort,
		DictionaryID: c.DictionaryId,
	}
}

type UpdateDictionaryDetail struct {
	GetById
	BaseDictionaryDetail
}

func (u *UpdateDictionaryDetail) Update() g.Map {
	return g.Map{"label": u.Label, "status": u.Status, "value": u.Value, "sort": u.Status, "dictionary_id": u.DictionaryId}
}

type SearchDictionaryDetail struct {
	Label        string `p:"label"`           // 展示值
	Status       *bool  `p:"status"`          // 启用状态
	Value        int    `p:"value"`           // 字典值
	Sort         int    `p:"sort"`            // 排序标记
	DictionaryId int    `p:"sysDictionaryID"` // 关联标记
	PageInfo
}

func (s *SearchDictionaryDetail) Search() g.Map {
	condition := make(g.Map, 4)
	if s.Label != "" {
		condition["`label` like ?"] = "%" + s.Label + "%"
	}
	if s.Status != nil {
		if *s.Status == true {
			condition["`status`"] = 1
		} else {
			condition["`status`"] = 2
		}
	}
	if s.Value != 0 {
		condition["`value`"] = s.Value
	}
	if s.DictionaryId != 0 {
		condition["`dictionary_id`"] = s.DictionaryId
	}
	return condition
}