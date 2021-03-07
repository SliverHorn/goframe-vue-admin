package api

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/request"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/extra"
	service "github.com/flipped-aurora/gf-vue-admin/server/app/service/extra"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"mime/multipart"
)

var File = new(file)

type file struct {
	err    error
	header *multipart.FileHeader
}

// @Tags ExtraFile
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/upload [post]
func (f *file) UploadFile(r *ghttp.Request) *response.Response {
	noSave := r.GetQueryString("noSave", "0")
	if _, header, err := r.Request.FormFile("file"); err != nil {
		return &response.Response{Code: 7, Error: err, Message: "上传文件失败!"}
	} else { // 文件上传后拿到文件路径
		var data *model.File
		if data, err = service.File.UploadFile(header, noSave); err != nil {
			return &response.Response{Error: err, MessageCode: response.ErrorCreated}
		}
		return &response.Response{Data: g.Map{"file": data}, MessageCode: response.SuccessCreated}
	}
}

// @Tags ExtraFile
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.GetById true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileUploadAndDownload/deleteFile [post]
func (f *file) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.File.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags ExtraFile
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileUploadAndDownload/getFileList [post]
func (f *file) GetList(r *ghttp.Request) *response.Response {
	var info request.PageInfo
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	if list, total, err := service.File.GetList(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	} else {
		return &response.Response{Data: response.PageResult{
			List:     list,
			Total:    total,
			Page:     info.Page,
			PageSize: info.PageSize,
		}, MessageCode: response.SuccessGetList}
	}
}
