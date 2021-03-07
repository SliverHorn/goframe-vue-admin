package boot

import (
	"github.com/flipped-aurora/gf-vue-admin/server/boot/internal"
	"github.com/flipped-aurora/gf-vue-admin/server/library/global"
	"github.com/flipped-aurora/gf-vue-admin/server/library/utils"
	"github.com/flipped-aurora/gf-vue-admin/server/router"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

var Server = new(_server)

type _server struct{}

func (s *_server) Initialize() {
	server := g.Server()
	address := g.Cfg().GetString("server.address")
	server.SetIndexFolder(true)
	if global.Config.System.OssType == "local" {
		_ = utils.Directory.BatchCreate(global.Config.Local.Path)
		server.AddStaticPath("/"+global.Config.Local.Path, global.Config.Local.Path)
	}
	server.AddStaticPath("/form-generator", "public/page")
	server.Use(internal.Middleware.Error, internal.Middleware.CORS)
	router.Routers.Init()
	g.Log().Printf(`
	欢迎使用 Gf-Vue-Admin
	当前版本:V2.4.0
	加群方式:微信号：SliverHorn QQ群：1040044540
	默认自动化文档地址:http://127.0.0.1%s/swagger
	默认前端文件运行地址:http://127.0.0.1:8080
	如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-vue-admin.com/docs/coffee
`, address)
	server.Plugin(&swagger.Swagger{})
	server.SetPort()
	server.Run()
}
