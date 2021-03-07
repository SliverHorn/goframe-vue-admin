package router

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	api "github.com/flipped-aurora/gf-vue-admin/server/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/server/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/server/router/internal"
	"github.com/gogf/gf/net/ghttp"
)

type admin struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewAdminRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &admin{router: router, response: &response.Handler{}}
}

func (a *admin) Init() {
	group := a.router.Group("/user").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("register", a.response.Handler()(api.Admin.Register))             // 新增用户
		group.PUT("setUserInfo", a.response.Handler()(api.Admin.Update))             // 设置用户信息
		group.DELETE("deleteUser", a.response.Handler()(api.Admin.Delete))           // 删除用户
		group.POST("getUserList", a.response.Handler()(api.Admin.GetList))           // 分页获取用户列表
		group.POST("changePassword", a.response.Handler()(api.Admin.ChangePassword)) // 修改密码
		group.POST("setUserAuthority", a.response.Handler()(api.Admin.SetAuthority)) // 设置用户权限
	}
}
