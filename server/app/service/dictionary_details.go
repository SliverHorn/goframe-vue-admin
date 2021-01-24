package service

import (
	"errors"
	"server/app/api/request"
	"server/app/model/dictionary_details"
	"server/library/utils"

	"github.com/gogf/gf/frame/g"
)

// CreateDictionaryDetail create a DictionaryDetail
// CreateDictionaryDetail 创建一个DictionaryDetail
func CreateDictionaryDetail(create *request.CreateDictionaryDetail) (err error) {
	insert := &dictionary_details.Entity{
		Label:        create.Label,
		Value:        create.Value,
		Status:       utils.BoolToInt(create.Status),
		Sort:         create.Sort,
		DictionaryId: create.DictionaryId,
	}
	_, err = dictionary_details.Insert(insert)
	return
}

// DeleteDictionaryDetail delete a DictionaryDetail
// DeleteDictionaryDetail 删除一个DictionaryDetail
func DeleteDictionaryDetail(delete *request.DeleteById) (err error) {
	_, err = dictionary_details.Delete(g.Map{"id": delete.Id})
	return err
}

// UpdateDictionaryDetail Update a DictionaryDetail
// UpdateDictionaryDetail 更新一个DictionaryDetail
func UpdateDictionaryDetail(update *request.UpdateDictionaryDetail) (err error) {
	condition := g.Map{"id": update.Id}
	updateData := g.Map{
		"label":         update.Label,
		"value":         update.Value,
		"status":        update.Status,
		"sort":          update.Sort,
		"dictionary_id": update.DictionaryId,
	}
	if _, err = dictionary_details.FindOne(condition); err != nil {
		return errors.New("记录不存在, 更新失败")
	}
	if _, err = dictionary_details.Update(updateData, condition); err != nil {
		return errors.New("更新失败")
	}
	return
}

// FindDictionaryDetail Query DictionaryDetail with id
// FindDictionaryDetail 用id查询DictionaryDetail
func FindDictionaryDetail(find *request.FindById) (dictionaryDetails *dictionary_details.DictionaryDetails, err error) {
	dictionaryDetails = (*dictionary_details.DictionaryDetails)(nil)
	db := g.DB("default").Table("dictionary_details").Safe()
	err = db.Where(g.Map{"id": find.Id}).Struct(&dictionaryDetails)
	return dictionaryDetails, err
}

// GetDictionaryDetailList Paging to get a list of DictionaryDetails
// GetDictionaryDetailList 分页获取DictionaryDetail列表
func GetDictionaryDetailList(info *request.GetDictionaryDetailList, condition g.Map) (list []*dictionary_details.DictionaryDetails, total int, err error) {
	list = ([]*dictionary_details.DictionaryDetails)(nil)
	limit := info.PageSize
	if info.Label != "" {
		condition["`label` like ?"] = "%" + info.Label + "%"
	}
	if info.Value != 0 {
		condition["`value`"] = info.Value
	}
	if info.DictionaryId != 0 {
		condition["`dictionary_id`"] = info.DictionaryId
	}
	offset := info.PageSize * (info.Page - 1)
	db := g.DB("default").Table("dictionary_details").Safe()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&list)
	return list, total, err
}
