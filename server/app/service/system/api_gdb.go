package service

import (
	"database/sql"
	"errors"
	"gf-vue-admin/server/app/api/request"
	"gf-vue-admin/server/app/api/response"
	model "gf-vue-admin/server/app/model/system"
	"github.com/gogf/gf/frame/g"
)

var Api = new(api)

type api struct {
	_api model.Api
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 新增基础Api
func (a *api) Create(info *request.CreateApi) error {
	err := g.DB().Table(a._api.TableName()).Where(g.Map{"path": info.Path, "method": info.Method}).Struct(&model.Api{})
	if !errors.Is(err, sql.ErrNoRows) {
		g.Log().Error(response.ErrorSameApi, g.Map{"path": info.Path, "method": info.Method})
		return err
	}
	entity := info.Create()
	_, err = g.DB().Table(a._api.TableName()).Insert(entity)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取api
func (a *api) First(info *request.GetById) (result *model.Api, err error) {
	var entity model.Api
	err = g.DB().Table(a._api.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id更新api
func (a *api) Update(info *request.UpdateApi) error {
	var old model.Api
	if err := g.DB().Table(a._api.TableName()).Where(g.Map{"id": info.Id}).Struct(&old); err != nil {
		return err
	}
	if old.Path == info.Path || old.Method == info.Method {
		return response.ErrorSameApi
	}
	_, err := g.DB().Table(a._api.TableName()).Data(info.Update()).Update(info.Condition())
	err = Casbin.UpdateApi(old.Path, info.Path, old.Method, info.Method)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除基础Api
func (a *api) Delete(info *request.DeleteApi) error {
	_, err := g.DB().Table(a._api.TableName()).Delete(info.Condition())
	Casbin.Clear(info.Path, info.Method)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取数据
func (a *api) GetList(info *request.SearchApi) (list interface{}, total int, err error) {
	var apis []model.Api
	condition := info.Search()
	db := g.DB().Table(a._api.TableName()).Safe()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&apis)
	return apis, total, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取所有api
func (a *api) GetAllApi() (result *[]model.Api, err error) {
	var apis []model.Api
	err = g.DB().Table(a._api.TableName()).Structs(&apis)
	return &apis, err
}
