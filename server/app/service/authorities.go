package service

import (
	"errors"
	"server/app/api/request"
	"server/app/model"
	"server/app/model/admins"
	"server/app/model/authorities"
	"server/app/model/authority_menu"
	"server/app/model/authority_resources"
	"server/library/global"

	"github.com/gogf/gf/frame/g"
)

// CreateAuthority Create a role
// CreateAuthority 创建一个角色
func CreateAuthority(create *request.CreateAuthority) (authority *authorities.Entity, err error) {
	insert := &authorities.Entity{
		AuthorityId:   create.AuthorityId,
		AuthorityName: create.AuthorityName,
		ParentId:      create.ParentId,
	}
	if !authorities.RecordNotFound(g.Map{"authority_id": insert.AuthorityId}) {
		return insert, errors.New("存在相同角色id")
	}
	if _, err = authorities.Insert(insert); err != nil {
		return insert, errors.New("创建角色失败")
	}
	return insert, nil
}

// CopyAuthority Copy a character
// CopyAuthority 复制一个角色
func CopyAuthority(copyInfo *request.AuthorityCopy) (authority *authorities.Authorities, err error) {
	authority = (*authorities.Authorities)(nil)
	if !authorities.RecordNotFound(g.Map{"authority_id": copyInfo.Authority.AuthorityId}) {
		return authority, errors.New("存在相同角色id")
	}
	// 创建角色信息
	_, err = authorities.Insert(g.Map{
		"authority_id":   copyInfo.Authority.AuthorityId,
		"authority_name": copyInfo.Authority.AuthorityName,
		"parent_id":      copyInfo.Authority.ParentId,
	})
	// 复制旧角色的api权限
	err = CopyCasbins(copyInfo)
	// 复制旧角色的动态菜单
	err = CopyMenus(copyInfo)
	authorityDb := g.DB("default").Table("authorities").Safe()
	err = authorityDb.Where(g.Map{"authority_id": copyInfo.Authority.AuthorityId}).Struct(&authority)
	return authority, err
}

// @title    UpdateAuthority
// @description   更改一个角色
func UpdateAuthority(update *request.UpdateAuthority) (err error) {
	updateData := &authorities.Entity{
		AuthorityId:   update.AuthorityId,
		AuthorityName: update.AuthorityName,
		ParentId:      update.ParentId,
	}
	if _, err = authorities.Update(updateData, g.Map{"authority_id": update.AuthorityId}); err != nil {
		return err
	}
	return
}

// DeleteAuthority Delete role
// DeleteAuthority 删除角色
func DeleteAuthority(auth *request.DeleteAuthority) (err error) {
	var authority *authorities.Entity
	db := g.DB(global.Db).Table("`casbin_rule`").Safe()
	if _, err = admins.FindOne(g.Map{"authority_id": auth.AuthorityId}); err != nil {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if authority, err = authorities.FindOne(g.Map{"parent_id": auth.AuthorityId}); err != nil {
		if authority != nil {
			return errors.New("此角色存在子角色不允许删除")
		}
	}
	if _, err = authorities.Delete(g.Map{"authority_id": auth.AuthorityId}); err != nil {
		return errors.New("角色删除失败")
	}
	if _, err = authority_menu.Delete(g.Map{"authority_id": auth.AuthorityId}); err != nil {
		return errors.New("菜单删除失败")
	}
	_, err = db.Delete(g.Map{"v0": auth.AuthorityId})
	return err
}

// GetInfoList Get data by page
// GetInfoList 分页获取数据
func GetAuthorityInfoList(info *request.PageInfo) (authorityList []*authorities.Authorities, total int, err error) {
	var associated []*authority_resources.Entity
	authorityList = ([]*authorities.Authorities)(nil)
	Children := ([]*authorities.Authorities)(nil)
	authorityDb := g.DB("default").Table("authorities").Safe()
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err = authorityDb.Where(g.Map{"parent_id": "0"}).Limit(limit).Offset(offset).Structs(&authorityList)
	for _, v := range authorityList {
		associated, err = authority_resources.FindAll(g.Map{"authority_id": v.AuthorityId})
		for _, a := range associated {
			if a.ResourcesId != "" {
				err = authorityDb.Where(g.Map{"authority_id": a.ResourcesId}).Structs(&v.DataAuthority) // 资源权限
			}
		}
		err = authorityDb.Where(g.Map{"parent_id": v.AuthorityId}).Structs(&Children)
		for _, c := range Children {
			if v.AuthorityId == c.ParentId {
				v.Children = append(v.Children, c)
			}
		}
	}
	if err != nil {
		return authorityList, total, errors.New("查询失败")
	}
	return authorityList, total, err
}

// GetAuthorityInfo Get all character information
// GetAuthorityInfo 获取所有角色信息
func GetAuthorityInfo(auth *authorities.Authorities) (err error, sa authorities.Authorities) {
	//err = global.GVA_DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	//return err, sa
	return
}

// SetDataAuthority Set role resource permissions
// SetDataAuthority 设置角色资源权限
func SetDataAuthority(auth *request.SetDataAuthority) (err error) {
	_, err = authority_resources.Delete(g.Map{"authority_id": auth.AuthorityId})
	for _, v := range auth.DataAuthority {
		condition := g.Map{
			"authority_id": auth.AuthorityId,
			"resources_id": v.AuthorityId,
		}
		if _, err = authority_resources.Insert(condition); err != nil {
			return err
		}
	}
	return
}

// SetMenuAuthority Menu and character binding
// SetMenuAuthority 菜单与角色绑定
func SetMenuAuthority(auth *authorities.Authorities) (err error) {
	//var s model.SysAuthority
	//global.GVA_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	//err := global.GVA_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus).Error
	//return err
	return
}

// CopyCasbins Copy role api permissions
// CopyCasbins 拷贝角色的api权限
func CopyCasbins(copyInfo *request.AuthorityCopy) (err error) {
	paths := GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	if err = UpdateCasbin(copyInfo.Authority.AuthorityId, paths); err != nil {
		_ = DeleteAuthority(&request.DeleteAuthority{AuthorityId: copyInfo.Authority.AuthorityId})
	}
	return
}

// CopyMenus Copy the dynamic menu of the character
// CopyMenus 拷贝角色的动态菜单
func CopyMenus(copyInfo *request.AuthorityCopy) (err error) {
	var menuList []*model.AuthorityMenu
	info := &request.AuthorityIdInfo{AuthorityId: copyInfo.OldAuthorityId}
	menuList, err = GetMenuAuthority(info)
	for _, v := range menuList {
		menuId := &v.MenuId
		_, err = authority_menu.Insert(g.Map{
			"authority_id": copyInfo.Authority.AuthorityId,
			"menu_id":      menuId,
		})
	}
	return err
}
