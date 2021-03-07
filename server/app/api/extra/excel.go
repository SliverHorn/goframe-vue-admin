package api

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/request"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	service "github.com/flipped-aurora/gf-vue-admin/server/app/service/extra"
	"github.com/flipped-aurora/gf-vue-admin/server/library/constant"
	"github.com/flipped-aurora/gf-vue-admin/server/library/utils"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/url"
)

var Excel = new(excel)

type excel struct{}

// @Tags ExtraExcel
// @Summary 导出Excel
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Param data body request.ExcelInfo true "导出Excel文件信息"
// @Success 200
// @Router /excel/exportExcel [post]
func (e *excel) Export(r *ghttp.Request) *response.Response {
	var info request.ExcelInfo
	if err := r.Parse(&info); err != nil {
		return &response.Response{Code: 7, Error: err, Message: "转换Excel失败!"}
	}
	filePath := constant.ExcelDir + info.FileName
	if err := service.Excel.Parse(info.InfoList, filePath); err != nil {
		return &response.Response{Code: 7, Error: err, Message: "转换Excel失败!"}
	}
	r.Response.Header().Add("msg", url.QueryEscape("导出Excel成功!"))
	r.Response.Header().Add("success", "true")
	r.Response.ServeFile(filePath)
	return &response.Response{}
}

// @Tags ExtraExcel
// @Summary 导入Excel文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "导入Excel文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"导入成功"}"
// @Router /excel/importExcel [post]
func (e *excel) Import(r *ghttp.Request) *response.Response {
	upload := r.GetUploadFile("file")
	if filename, err := upload.Save(constant.ExcelDir, false); err != nil {
		return &response.Response{Code: 7, Error: err, Message: "接收文件失败!"}
	} else {
		return &response.Response{Code: 0, Data: g.Map{"filename": filename}, Message: "导入成功!"}
	}
}

// @Tags ExtraExcel
// @Summary 加载Excel数据
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"加载数据成功"}"
// @Router /excel/loadExcel [get]
func (e *excel) Load(r *ghttp.Request) *response.Response {
	if menus, err := service.Excel.Parse2Data(); err != nil {
		return &response.Response{Code: 7, Error: err, Message: "加载数据失败!"}
	} else {
		return &response.Response{Code: 0, Data: response.PageResult{List: menus, Total: len(menus), Page: 1, PageSize: 999}, Message: "加载数据成功!"}
	}
}

// @Tags ExtraExcel
// @Summary 下载模板
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param fileName query string true "模板名称"
// @Success 200
// @Router /excel/downloadTemplate [get]
func (e *excel) Download(r *ghttp.Request) *response.Response {
	fileName := r.GetQueryString("fileName")
	filePath := constant.ExcelDir + fileName
	if ok, err := utils.Directory.PathExists(filePath); !ok || err != nil {
		return &response.Response{Code: 7, Error: err, Message: "文件不存在!"}
	}
	r.Response.Header().Add("msg", url.QueryEscape("下载Excel模板成功!"))
	r.Response.Header().Add("success", "true")
	r.Response.ServeFile(filePath)
	return &response.Response{}
}
