package api

import (
	"gf-vue-admin/server/app/api/request"
	"gf-vue-admin/server/app/api/response"
	service "gf-vue-admin/server/app/service/extra"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// @Tags ExtraFile
// @Summary 断点续传到服务器
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "an example for breakpoint resume, 断点续传示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"切片创建成功"}"
// @Router /fileUploadAndDownload/breakpointContinue [post]
func (f *file) BreakpointContinue(r *ghttp.Request) *response.Response {
	var info request.BreakpointContinue
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: f.err, MessageCode: response.ErrorCreateChunk}
	}
	if _, f.header, f.err = r.Request.FormFile("file"); f.err != nil {
		return &response.Response{Error: f.err, MessageCode: response.ErrorFormFile}
	}
	if err := service.BreakpointContinue().BreakpointContinue(&info, f.header); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreateChunk}
	}
	return &response.Response{MessageCode: response.SuccessCreateChunk}
}

// @Tags ExtraFile
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "Find the file, 查找文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func (f *file) FindFile(r *ghttp.Request) *response.Response {
	var info request.BreakpointContinue
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: f.err, MessageCode: response.ErrorFind}
	}
	data, err := service.BreakpointContinue().FindOrCreateFile(&info)
	if err != nil {
		return &response.Response{Error: err, Data: g.Map{"file": data}, MessageCode: response.ErrorFind}
	}
	return &response.Response{Data: g.Map{"file": data}, MessageCode: response.SuccessFind}
}

// @Tags ExtraFile
// @Summary 创建文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件完成"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"file uploaded, 文件创建成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func (f *file) BreakpointContinueFinish(r *ghttp.Request) *response.Response {
	var info request.BreakpointContinueFinish
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: f.err, MessageCode: response.ErrorFinish}
	}
	var filePath, err = service.BreakpointContinue().BreakpointContinueFinish(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFinish}
	}
	return &response.Response{Data: g.Map{"filePath": filePath}, MessageCode: response.SuccessFinish}
}

// @Tags ExtraFile
// @Summary 删除切片
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "删除缓存切片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"缓存切片删除成功"}"
// @Router /fileUploadAndDownload/removeChunk [post]
func (f *file) RemoveChunk(r *ghttp.Request) *response.Response {
	var info request.BreakpointContinue
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: f.err, MessageCode: response.ErrorFinish}
	}
	if err := service.BreakpointContinue().DeleteFileChunk(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorRemoveChunk}
	}
	return &response.Response{Data: g.Map{"filePath": info.FilePath}, MessageCode: response.SuccessRemoveChunk}
}
