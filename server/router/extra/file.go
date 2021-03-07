package router

import (
	api "github.com/flipped-aurora/gf-vue-admin/server/app/api/extra"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	"github.com/flipped-aurora/gf-vue-admin/server/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type file struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewFileRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &file{router: router, response: &response.Handler{}}
}

func (f *file) Init() {
	group := f.router.Group("/fileUploadAndDownload")
	{
		group.POST("/upload", f.response.Handler()(api.File.UploadFile))                                 // 上传文件
		group.POST("/deleteFile", f.response.Handler()(api.File.Delete))                                 // 删除指定文件
		group.POST("/getFileList", f.response.Handler()(api.File.GetList))                               // 获取上传文件列表
		group.POST("/breakpointContinue", f.response.Handler()(api.File.BreakpointContinue))             // 断点续传
		group.GET("/findFile", f.response.Handler()(api.File.FindFile))                                  // 查询文件上传记录
		group.POST("/breakpointContinueFinish", f.response.Handler()(api.File.BreakpointContinueFinish)) // 查询文件上传成功记录
		group.POST("/removeChunk", f.response.Handler()(api.File.RemoveChunk))                           // 移除文件的切片
	}
}
