package request

import model "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"

type AddMenuAuthority struct {
	GetAuthorityId
	Menus       []model.Menu
}
