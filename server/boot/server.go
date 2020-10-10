package boot

import (
	"fmt"
	"server/app/middleware"
	"server/router"
	"time"

	"github.com/gogf/gf/frame/g"
)

func InitializeRunServer() {
	s := g.Server()
	s.Use(middleware.Error)
	s.SetReadTimeout(10 * time.Second)
	s.SetWriteTimeout(10 * time.Second)
	s.SetMaxHeaderBytes(1 << 20)
	s.SetIndexFolder(true)
	if g.Cfg("system").GetString("system.OssType") == "local" {
		s.AddStaticPath("/"+g.Cfg("oss").GetString("local.LocalPath"), g.Cfg("oss").GetString("local.LocalPath"))
	}
	s.AddStaticPath("/form-generator", "public/page")
	router.InitializeRouters()
	fmt.Printf(`
	欢迎使用 Gf-Vue-Admin
	当前版本:V1.1.0
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	`, g.Cfg().GetString("server.Address"))
	s.Run()
}
