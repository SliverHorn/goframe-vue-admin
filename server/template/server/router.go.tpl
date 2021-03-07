package router

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	"github.com/flipped-aurora/gf-vue-admin/server/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type {{.Abbreviation}} struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func New{{.StructName}}Router(router *ghttp.RouterGroup) interfaces.Router {
	return &{{.Abbreviation}}{router: router, response: &response.Handler{}}
}

func (r *{{.Abbreviation}}) Init() {
	group := r.router.Group("/{{.Abbreviation}}")
	{
		group.POST("create", r.response.Handler()(api.{{.StructName}}.Create))             // 新建{{.StructName}}
		group.GET("first", r.response.Handler()(api.{{.StructName}}.First))                // 根据ID获取{{.StructName}}
		group.PUT("update", r.response.Handler()(api.{{.StructName}}.Update))              // 更新{{.StructName}}
		group.DELETE("delete", r.response.Handler()(api.{{.StructName}}.Delete))           // 删除{{.StructName}}
		group.DELETE("deletes", r.response.Handler()(api.{{.StructName}}.Deletes))         // 批量删除{{.StructName}}
		group.GET("getList", r.response.Handler()(api.{{.StructName}}.GetList))            // 获取{{.StructName}}列表
	}
}