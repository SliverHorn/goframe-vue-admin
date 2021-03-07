package request

import (
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"
	"github.com/gogf/gf/frame/g"
)

type BaseAuthority struct {
	ParentId      string `p:"parent_id" v:"required|length:1,1000#请输入角色父id|角色父id长度为:min到:max位"`
	AuthorityId   string `p:"authority_id" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
	AuthorityName string `p:"authority_name" v:"required|length:1,1000#请输入角色名字|角色名字长度为:min到:max位"`
	DefaultRouter string `p:"defaultRouter"`
}

type CreateAuthority struct {
	BaseAuthority
}

func (c *CreateAuthority) Create() *model.Authority {
	return &model.Authority{
		ParentId:      c.ParentId,
		AuthorityId:   c.AuthorityId,
		AuthorityName: c.AuthorityName,
		DefaultRouter: c.DefaultRouter,
	}
}

type CopyAuthority struct {
	Authority      model.Authority `json:"authority"`
	OldAuthorityId string          `json:"oldAuthorityId"`
}

type UpdateAuthority struct {
	BaseAuthority
}

func (u *UpdateAuthority) Update() g.Map {
	return g.Map{"parent_id": u.ParentId, "authority_name": u.AuthorityName, "default_router": u.DefaultRouter}
}

type Authority struct {
	AuthorityId   string `r:"authorityId" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
	AuthorityName string `r:"authorityName" v:"required|length:1,1000#请输入角色名|角色名长度为:min到:max位"`
}

type SetDataAuthority struct {
	BaseAuthority
	DataAuthorityId []model.Authority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
}

type SetMenuAuthority struct {
	BaseAuthority
	DataAuthorityId []model.Authority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
}