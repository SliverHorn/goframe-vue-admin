package model

import (
	"github.com/flipped-aurora/gf-vue-admin/server/library/global"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	global.Model
	Uuid        string    `orm:"uuid" json:"uuid" gorm:"comment:用户UUID"`
	Avatar      string    `orm:"avatar" json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Nickname    string    `orm:"nickname" json:"nickName" gorm:"comment:用户登录名"`
	Username    string    `orm:"username" json:"userName" gorm:"default:系统用户;comment:用户昵称" `
	Password    string    `orm:"password" json:"-" gorm:"comment:用户登录密码"`
	AuthorityId string    `orm:"authority_id" json:"authorityId" gorm:"default:888;comment:用户角色ID"`
	Authority   Authority `orm:"-" json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
}

func (a *Admin) TableName() string {
	return "admins"
}

//@author: SliverHorn
//@description: 密码检查(工具类)
//@return: bool(false 校验失败, true 校验成功)
func (a *Admin) CompareHashAndPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

//@author: SliverHorn
//@description: 加密密码(工具类)
func (a *Admin) EncryptedPassword() error {
	if byTes, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost); err != nil { // 加密密码
		return err
	} else {
		a.Password = string(byTes)
		return nil
	}
}
