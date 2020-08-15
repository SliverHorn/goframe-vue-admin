package request

import (
	"server/app/model"
)

type CreateAuthority struct {
	AuthorityId   string `p:"authority_id" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
	AuthorityName string `p:"authority_name" v:"required|length:1,1000#请输入角色名字|角色名字长度为:min到:max位"`
	ParentId      string `p:"authority_id" v:"required|length:1,1000#请输入角色父id|角色父id长度为:min到:max位"`
}

type UpdateAuthority struct {
	AuthorityId   string `p:"authority_id" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
	AuthorityName string `p:"authority_name" v:"required|length:1,1000#请输入角色名字|角色名字长度为:min到:max位"`
	ParentId      string `p:"authority_id" v:"required|length:1,1000#请输入角色父id|角色父id长度为:min到:max位"`
}

type DeleteAuthority struct {
	AuthorityId string `p:"authorityId" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
}
type Authorities struct {
	AuthorityId   string           `r:"authorityId"  orm:"authority_id,primary" json:"authority_id"`    // 角色ID
	AuthorityName string           `r:"authorityName" orm:"authority_name"       json:"authority_name"` // 角色名
	ParentId      string           `r:"parentId" orm:"parent_id"            json:"parent_id"`
	BaseMenu      []model.BaseMenu `json:"menus"`
}

type AuthorityCopy struct {
	Authority      Authorities `json:"authority"`
	OldAuthorityId string      `r:"oldAuthorityId" json:"oldAuthorityId"`
}

type SetDataAuthority struct {
	AuthorityId   string         `r:"authorityId" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
	DataAuthority []*Authorities `r:"dataAuthorityId" json:"dataAuthorityId"`
}

type Authority struct {
	AuthorityId   string `r:"authorityId" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
	AuthorityName string `r:"authorityName" v:"required|length:1,1000#请输入角色名|角色名长度为:min到:max位"`
}
