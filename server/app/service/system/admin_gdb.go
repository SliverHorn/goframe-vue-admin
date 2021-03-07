package service

import (
	"database/sql"
	"errors"
	"gf-vue-admin/server/app/api/request"
	"gf-vue-admin/server/app/api/response"
	"gf-vue-admin/server/app/model/system"
	"gf-vue-admin/server/app/service/system/internal"
	"github.com/gogf/gf/frame/g"
)

var Admin = new(admin)

type admin struct {
	_admin model.Admin
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 注册逻辑代码
func (a *admin) Register(info *request.Register) error {
	var entity model.Admin
	if !errors.Is(g.DB().Table(a._admin.TableName()).Where(g.Map{"username": info.Username}).Struct(&entity), sql.ErrNoRows) {
		return response.ErrorUsernameRegistered
	}
	user := info.Create()
	if err := user.EncryptedPassword(); err != nil {
		return response.ErrorEncryptedPassword
	}
	_, err := g.DB().Table(a._admin.TableName()).Data(user).Insert()
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 修改管理员密码
func (a *admin) ChangePassword(info *request.ChangePassword) error {
	var entity model.Admin
	if err := g.DB().Table(a._admin.TableName()).Where("username", info.Username).Struct(&entity); err != nil {
		return response.ErrorUserNoExist
	}
	if !entity.CompareHashAndPassword(info.Password) {
		return response.ErrorWrongPassword
	}
	entity.Password = info.NewPassword
	if err := entity.EncryptedPassword(); err != nil {
		return err
	}
	if _, err := g.DB().Table(a._admin.TableName()).Where("username", info.Username).Data(g.Map{"password": entity.Password}).Update(); err != nil {
		return err
	}
	return nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取用户列表
func (a *admin) GetAdminList(info *request.PageInfo) (list interface{}, total int, err error) {
	var admins []model.Admin
	var db = g.DB().Table(a._admin.TableName()).Safe()
	limit, offset := info.Paginate()
	total, err = db.Count()
	err = db.Limit(limit).Offset(offset).Structs(&admins)
	for i, entity := range admins {
		admins[i].Authority = internal.Authority().First(entity.AuthorityId)
	}
	return admins, total, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 用于刷新token,根据uuid返回admin信息
func (a *admin) FindAdmin(info *request.GetByUuid) (result *model.Admin, err error) {
	var entity model.Admin
	var db = g.DB().Table(a._admin.TableName()).Safe()
	err = db.Where(g.Map{"uuid": info.Uuid}).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 用于刷新token,根据uuid返回admin信息
func (a *admin) FindAdminById(info *request.GetById) (result *model.Admin, err error) {
	var entity model.Admin
	err = g.DB().Table(a._admin.TableName()).Where(g.Map{"id": info.Id}).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 设置用户权限
func (a *admin) SetUserAuthority(info *request.SetAuthority) error {
	_, err := g.DB().Table(a._admin.TableName()).Update(g.Map{"authority_id": info.AuthorityId}, g.Map{"uuid": info.Uuid})
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除管理员
func (a *admin) Delete(info *request.GetById) error {
	_, err := g.DB().Table(a._admin.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 设置管理员信息
func (a *admin) SetAdminInfo(info *request.UpdateAdmin) (result *model.Admin, err error) {
	_, err = g.DB().Table(a._admin.TableName()).Update(g.Map{"avatar": info.Avatar}, g.Map{"uuid": info.Uuid})
	return a.FindAdmin(&request.GetByUuid{Uuid: info.Uuid})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 设置管理员信息
func (a *admin) Login(info *request.AdminLogin) (result *model.Admin, err error) {
	var entity model.Admin
	if err = g.DB().Table(a._admin.TableName()).Where(g.Map{"username": info.Username}).Scan(&entity); err != nil {
		return &entity, response.ErrorUserNoExist
	}
	entity.Authority = internal.Authority().First(entity.AuthorityId)
	if !entity.CompareHashAndPassword(info.Password) {
		return &entity, response.ErrorWrongPassword
	}
	return &entity, nil
}
