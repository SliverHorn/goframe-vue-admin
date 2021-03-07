package api

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/request"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	service "github.com/flipped-aurora/gf-vue-admin/server/app/service/extra"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"mime/multipart"
)

var SimpleUploader = new(uploader)

type uploader struct {
	err    error
	header *multipart.FileHeader
}

// @Tags ExtraSimpleUploader
// @Summary 断点续传插件版示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "断点续传插件版示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"切片创建成功"}"
// @Router /simpleUploader/upload [post]
func (u *uploader) Upload(r *ghttp.Request) *response.Response {
	var info request.Upload
	if _, u.header, u.err = r.Request.FormFile("file"); u.err != nil {
		return &response.Response{Error: u.err, MessageCode: response.ErrorFormFile}
	}
	if u.err = service.SimpleUploader().Upload(u.header, &info); u.err != nil {
		return &response.Response{Code: 7, Error: u.err, Message: "切片创建失败!"}
	}
	return &response.Response{Code: 0, Message: "切片创建成功!"}
}

// @Tags ExtraSimpleUploader
// @Summary 断点续传插件版示例
// @Security ApiKeyAuth
// @Produce  application/json
// @Param md5 query string true "md5"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /simpleUploader/checkFileMd5 [get]
func (u *uploader) CheckFileMd5(r *ghttp.Request) *response.Response {
	var info request.CheckFileMd5
	chunks, isDone, err := service.SimpleUploader().CheckFileMd5(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCheckFileMd5}
	}
	return &response.Response{Data: g.Map{"chunks": chunks, "isDone": isDone}, MessageCode: response.SuccessCheckFileMd5}
}

// @Tags ExtraSimpleUploader
// @Summary 合并文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param md5 query string true "md5"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"合并成功"}"
// @Router /simpleUploader/mergeFileMd5 [get]
func (u *uploader) MergeFileMd5(r *ghttp.Request) *response.Response {
	var info request.MergeFileMd5
	if err := service.SimpleUploader().MergeFileMd5(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorMergeFileMd5}
	}
	return &response.Response{MessageCode: response.SuccessMergeFileMd5}
}
