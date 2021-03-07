package router

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	api "github.com/flipped-aurora/gf-vue-admin/server/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/server/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/server/router/internal"
	"github.com/gogf/gf/net/ghttp"
)

type authority struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewAuthorityRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &authority{router: router, response: &response.Handler{}}
}

func (a *authority) Init() {
	group := a.router.Group("/authority").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("createAuthority", a.response.Handler()(api.Authority.Create))            // 创建角色
		group.POST("copyAuthority", a.response.Handler()(api.Authority.Copy))                // 拷贝角色
		group.PUT("updateAuthority", a.response.Handler()(api.Authority.Update))             // 更新角色
		group.POST("deleteAuthority", a.response.Handler()(api.Authority.Delete))            // 删除角色
		group.POST("getAuthorityList", a.response.Handler()(api.Authority.GetList))          // 获取角色列表
		group.POST("setDataAuthority", a.response.Handler()(api.Authority.SetDataAuthority)) // 设置角色资源权限
	}
}
