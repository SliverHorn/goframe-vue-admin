package request

import (
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"
	uuid "github.com/satori/go.uuid"
)

type BaseAdmin struct {
	Username    string `p:"username" v:"required|length:1,30#请输入用户名称|您输入用户名称长度非法"`
	Password    string `p:"password" v:"required|length:6,30#请输入旧密码|旧密码长度为:min到:max位"`
	NewPassword string `p:"new_password" v:"required|length:6,30#请输入新密码|新密码长度为:min到:max位"`
}

type Register struct {
	Avatar      string `p:"headerImg"`
	Username    string `p:"userName"`
	Password    string `p:"passWord"`
	Nickname    string `p:"nickName"`
	AuthorityId string `p:"authorityId"`
}

func (r *Register) Create() *model.Admin {
	return &model.Admin{
		Uuid:        uuid.NewV4().String(),
		Username:    r.Username,
		Password:    r.Password,
		Nickname:    r.Nickname,
		Avatar:      r.Avatar,
		AuthorityId: r.AuthorityId,
	}
}

type ChangePassword struct {
	Uuid        string
	Username    string `p:"username" v:"required|length:1,30#请输入用户名称|您输入用户名称长度非法"`
	Password    string `p:"password" v:"required|length:6,30#请输入旧密码|旧密码长度为:min到:max位"`
	NewPassword string `p:"new_password" v:"required|length:6,30#请输入新密码|新密码长度为:min到:max位"`
}

type SetAuthority struct {
	Uuid        string `p:"uuid" v:"required|length:36,36#请输入管理员UUID|管理员UUID长度为:min到:max位"`
	AuthorityId string `p:"authority_id" v:"required|length:1, 100#请输入角色ID|角色ID长度为:min到:max位"`
}

type UpdateAdmin struct {
	GetById
	Uuid   string
	Avatar string `p:"headerImg" v:"required|length:1, 100#请输入头像链接|头像链接长度为:min到:max位"`
}
