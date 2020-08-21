package service

import (
	"errors"
	"server/app/api/request"
	"server/app/model/admins"
	"server/library/global"
	"server/library/utils"

	"github.com/gogf/gf/frame/g"
)

// ChangePassword Change administrator password
// ChangePassword 修改管理员密码
func ChangePassword(change *request.ChangePassword) (err error) {
	var admin *admins.Entity
	if admin, err = admins.FindOne(g.Map{"username": change.Username}); err != nil {
		return errors.New("用户不存在, 修改密码失败")
	}
	if utils.CompareHashAndPassword(admin.Password, change.Password) {
		if admin.Password, err = utils.EncryptedPassword(change.NewPassword); err != nil {
			return errors.New("修改密码失败")
		}
		if _, err = admins.Save(admin); err != nil {
			return errors.New("修改密码失败")
		}
		return
	}
	return errors.New("旧密码输入有误")
}

// UploadHeaderImg User uploads profile picture
// UploadHeaderImg 用户上传头像
func UploadHeaderImg(userUuid string, filePath string) (adminInfo *admins.Entity, err error) {
	if _, err := admins.Update(g.Map{"header_img": filePath}, g.Map{"uuid": userUuid}); err != nil {
		return adminInfo, errors.New("")
	}
	adminInfo, err = admins.FindOne(g.Map{"uuid": userUuid})
	return adminInfo, err
}

// GetAdminList Paging gets the list of users
// GetAdminList 分页获取用户列表
func GetAdminList(info *request.PageInfo) (list interface{}, total int, err error) {
	var adminList []*admins.Admin
	adminDb := g.DB(global.Db).Table("admins").Safe()
	authorityDb := g.DB(global.Db).Table("authorities").Safe()
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	total, err = adminDb.Count()
	err = adminDb.Limit(limit).Offset(offset).Structs(&adminList)
	for _, v := range adminList {
		err = authorityDb.Where(g.Map{"authority_id": v.AuthorityId}).Struct(&v.Authority)
	}
	return adminList, total, err
}

// FindAdmin Used to refresh token and return admin information according to uuid
// FindAdmin 用于刷新token,根据uuid返回admin信息
func FindAdmin(adminUUID string) (admin *admins.Admin, err error) {
	admin = (*admins.Admin)(nil)
	db := g.DB(global.Db).Table("admins").Safe()
	authorityDb := g.DB(global.Db).Table("authorities").Safe()
	err = db.Where(g.Map{"uuid": adminUUID}).Struct(&admin)
	err = authorityDb.Where(g.Map{"authority_id": admin.AuthorityId}).Struct(&admin.Authority)
	return admin, err
}

// SetUserAuthority Set user permissions
// SetUserAuthority 设置用户权限
func SetUserAuthority(set *request.SetAdminAuthority) (err error) {
	_, err = admins.Update(g.Map{"authority_id": set.AuthorityId}, g.Map{"uuid": set.Uuid})
	return err
}

// DeleteAdmin Delete administrator
// DeleteAdmin 删除管理员
func DeleteAdmin(delete *request.DeleteById) (err error) {
	_, err = admins.Delete(g.Map{"id": delete.Id})
	return err
}
