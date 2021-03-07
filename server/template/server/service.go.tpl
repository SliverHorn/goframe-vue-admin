package service

import (
	"gf-vue-admin/server/app/api/request"
	"gf-vue-admin/server/app/model"
	"github.com/gogf/gf/frame/g"
)

var {{.StructName}} = new({{.Abbreviation}})

type {{.Abbreviation}} struct {
	_{{.Abbreviation}} model.{{.StructName}}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建{{.StructName}}记录
func (s *{{.Abbreviation}}) Create(info *model.{{.StructName}}) error {
	_, err := g.DB().Table(s._{{.Abbreviation}}.TableName()).Insert(info)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取{{.StructName}}记录
func (s *{{.Abbreviation}}) First(info *request.GetById) (result *model.{{.StructName}}, err error) {
	var entity model.{{.StructName}}
	err = g.DB().Table(s._{{.Abbreviation}}.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新{{.StructName}}记录
func (s *{{.Abbreviation}}) Update(info *model.{{.StructName}}) (result *model.{{.StructName}}, err error) {
    _, err = g.DB().Table(s._{{.Abbreviation}}.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除{{.StructName}}记录
func (s *{{.Abbreviation}}) Delete(info *request.GetById) error {
	_, err := g.DB().Table(s._{{.Abbreviation}}.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量{{.StructName}}记录
func (s *{{.Abbreviation}}) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(s._{{.Abbreviation}}.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取{{.StructName}}记录列表
func (s *{{.Abbreviation}}) GetList(info *request.Search{{.StructName}}) (list *[]model.{{.StructName}}, total int, err error) {
	var {{.Abbreviation}}s []model.{{.StructName}}
	db := g.DB().Table(s._{{.Abbreviation}}.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&{{.Abbreviation}}s)
	return &{{.Abbreviation}}s, total, err
}
