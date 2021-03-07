package router

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	api "github.com/flipped-aurora/gf-vue-admin/server/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/server/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/server/router/internal"
	"github.com/gogf/gf/net/ghttp"
)

type config struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewConfigRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &config{router: router, response: &response.Handler{}}
}

func (c *config) Init() {
	group := c.router.Group("/system").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("getSystemConfig", c.response.Handler()(api.Config.GetConfig))   // 获取配置文件内容
		group.POST("setSystemConfig", c.response.Handler()(api.Config.SetConfig))   // 设置配置文件内容
		group.POST("getServerInfo", c.response.Handler()(api.Config.GetServerInfo)) // 获取服务器信息
		group.POST("reloadSystem", c.response.Handler()(api.Config.ReloadSystem))   // 重启服务
	}
}
