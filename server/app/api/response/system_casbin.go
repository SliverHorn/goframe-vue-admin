package response

import "gf-vue-admin/server/app/api/request"

type PolicyPath struct {
	Paths []request.CasbinInfo `json:"paths"`
}
