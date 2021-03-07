package router

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	api "github.com/flipped-aurora/gf-vue-admin/server/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/server/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/server/router/internal"
	"github.com/gogf/gf/net/ghttp"
)

type casbin struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewCasbinRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &casbin{router: router, response: &response.Handler{}}
}

func (c *casbin) Init() {
	group := c.router.Group("/casbin").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("updateCasbin", c.response.Handler()(api.Casbin.Update))
		group.POST("getPolicyPathByAuthorityId", c.response.Handler()(api.Casbin.GetPolicyPath))
	}
}
