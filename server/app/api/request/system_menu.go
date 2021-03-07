package request

import (
	model "gf-vue-admin/server/app/model/system"
	"github.com/gogf/gf/frame/g"
)

type UpdateMenu struct {
	model.Menu
}

func (u *UpdateMenu) Update() g.Map {
	return g.Map{
		"keep_alive":   u.KeepAlive,
		"default_menu": u.DefaultMenu,
		"parent_id":    u.ParentId,
		"path":         u.Path,
		"name":         u.Name,
		"hidden":       u.Hidden,
		"component":    u.Component,
		"title":        u.Title,
		"icon":         u.Icon,
		"sort":         u.Sort,
	}
}
