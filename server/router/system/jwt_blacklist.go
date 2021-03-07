package router

import (
	"gf-vue-admin/server/app/api/response"
	api "gf-vue-admin/server/app/api/system"
	"gf-vue-admin/server/interfaces"
	"gf-vue-admin/server/router/internal"
	"github.com/gogf/gf/net/ghttp"
)

type blacklist struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewJwtBlacklistRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &blacklist{router: router, response: &response.Handler{}}
}

func (j *blacklist) Init() {
	group := j.router.Group("/jwt").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("jsonInBlacklist", j.response.Handler()(api.JwtBlacklist.JwtToBlacklist)) // jwt加入黑名单
	}
}
