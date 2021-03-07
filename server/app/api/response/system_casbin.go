package response

import "github.com/flipped-aurora/gf-vue-admin/server/app/api/request"

type PolicyPath struct {
	Paths []request.CasbinInfo `json:"paths"`
}
