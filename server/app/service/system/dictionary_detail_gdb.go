package service

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/request"
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"
	"github.com/gogf/gf/frame/g"
)

var DictionaryDetail = new(detail)

type detail struct{
	_detail model.DictionaryDetail
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建字典详情数据
func (d *detail) Create(info *request.CreateDictionaryDetail) error {
	_, err := g.DB().Table(d._detail.TableName()).Insert(info.Create())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取字典详情单条数据
func (d *detail) First(info *request.GetById) (result *model.DictionaryDetail, err error) {
	var entity model.DictionaryDetail
	err = g.DB().Table(d._detail.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 更新字典详情数据
func (d *detail) Update(info *request.UpdateDictionaryDetail) error {
	_, err := g.DB().Table(d._detail.TableName()).Update(info.Update(), info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除字典详情数据
func (d *detail) Delete(info *request.GetById) error {
	_, err := g.DB().Table(d._detail.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取字典详情列表
func (d *detail) GetList(info *request.SearchDictionaryDetail) (list interface{}, total int, err error) {
	var details []model.DictionaryDetail
	db := g.DB().Table(d._detail.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&details)
	return details, total, err
}