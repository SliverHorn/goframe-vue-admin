package request

import model "gf-vue-admin/server/app/model/system"

type AddMenuAuthority struct {
	GetAuthorityId
	Menus       []model.Menu
}
