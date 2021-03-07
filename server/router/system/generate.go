package router

import (
	"gf-vue-admin/server/app/api/response"
	api "gf-vue-admin/server/app/api/system"
	"gf-vue-admin/server/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type generate struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewGenerateRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &generate{router: router, response: &response.Handler{}}
}

func (g *generate) Init() {
	group := g.router.Group("/autoCode")
	{
		group.GET("getDB", g.response.Handler()(api.Generate.GetDbs))         // 获取数据库
		group.GET("getTables", g.response.Handler()(api.Generate.GetTables))  // 获取对应数据库的表
		group.GET("getColumn", g.response.Handler()(api.Generate.GetColumns)) // 获取指定表所有字段信息
		group.POST("preview", g.response.Handler()(api.Generate.Preview))     // 获取自动创建代码预览
		group.POST("createTemp", g.response.Handler()(api.Generate.Create))   // 创建自动化代码
	}
}
