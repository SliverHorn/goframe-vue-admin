package service

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/request"
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"
	"github.com/gogf/gf/frame/g"
)

var Dictionary = new(dictionary)

type dictionary struct {
	_detail     model.DictionaryDetail
	_dictionary model.Dictionary
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建字典数据
func (d *dictionary) Create(info *request.CreateDictionary) error {
	var entity model.Dictionary
	_, err := g.DB().Table(entity.TableName()).Insert(info.Create())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id或者type获取字典单条数据
func (d *dictionary) First(info *request.FirstDictionary) (result *model.Dictionary, err error) {
	var entity model.Dictionary
	err = g.DB().Table(d._dictionary.TableName()).Where("id", info.Id).Or("type", info.Type).Struct(&entity)
	err = g.DB().Table(d._detail.TableName()).Where(g.Map{"dictionary_id": entity.ID}).Structs(&entity.DictionaryDetails)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 更新字典数据
func (d *dictionary) Update(info *request.UpdateDictionary) error {
	_, err := g.DB().Table(d._dictionary.TableName()).Update(info.Update(), info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除字典数据
func (d *dictionary) Delete(info *request.GetById) error {
	//_, err := g.DB().Table(d._detail.TableName()).Delete(g.Map{"dictionary_id": info.Id})
	_, err := g.DB().Table(d._dictionary.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取字典列表
func (d *dictionary) GetList(info *request.SearchDictionary) (list interface{}, total int, err error) {
	var dictionaries []model.Dictionary
	condition := info.Search()
	db := g.DB().Table(d._dictionary.TableName()).Safe()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&dictionaries)
	return dictionaries, total, err
}
