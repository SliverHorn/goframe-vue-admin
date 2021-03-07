package api

import (
	"errors"
	"fmt"
	"gf-vue-admin/server/app/api/response"
	model "gf-vue-admin/server/app/model/system"
	service "gf-vue-admin/server/app/service/system"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/url"
	"os"
)

var Generate = new(generate)

type generate struct{}

// @Tags SystemGenerate
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getDatabase [get]
func (a *generate) GetDbs(r *ghttp.Request) *response.Response {
	if dbs, err := service.Generate.GetDbs(); err != nil {
		return &response.Response{Code: 7, Error: err}
	} else {
		return &response.Response{Code: 0, Data: g.Map{"dbs": dbs}}
	}
}

// @Tags SystemGenerate
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getTables [get]
func (a *generate) GetTables(r *ghttp.Request) *response.Response {
	db := r.GetString("dbName", "")
	if tables, err := service.Generate.GetTables(db); err != nil {
		return &response.Response{Code: 7, Error: err}
	} else {
		return &response.Response{Code: 0, Data: g.Map{"tables": tables}}
	}
}

// @Tags SystemGenerate
// @Summary 获取当前表所有字段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getColumn [get]
func (a *generate) GetColumns(r *ghttp.Request) *response.Response {
	db := r.GetString("dbName", "")
	table := r.GetString("tableName")
	if columns, err := service.Generate.GetColumns(db, table); err != nil {
		return &response.Response{Code: 7, Error: err}
	} else {
		return &response.Response{Code: 0, Data: g.Map{"columns": columns}}
	}
}

// @Tags SystemGenerate
// @Summary 预览创建后的代码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AutoCode true "预览创建代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/preview [post]
func (a *generate) Preview(r *ghttp.Request) *response.Response {
	var info model.AutoCode
	if err := r.Parse(&info); err != nil {
		return &response.Response{Code: 7, Error: err, Message: "预览代码失败!"}
	}
	if result, err := service.Generate.Preview(&info); err != nil {
		return &response.Response{Code: 7, Error: err, Message: "预览代码失败!"}
	} else {
		return &response.Response{Code: 0, Data: g.Map{"autoCode": result}, Message: "预览成功"}
	}
}

// @Tags SystemGenerate
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AutoCode true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/createTemp [post]
func (a *generate) Create(r *ghttp.Request) *response.Response {
	var info model.AutoCode
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if info.AutoCreateApiToSql {
		if err := service.Generate.AutoCreateApis(&info); err != nil {
			g.Log().Error("自动化创建失败, 请自行清空垃圾数据!", g.Map{"err": err})
			r.Response.Header().Add("success", "false")
			r.Response.Header().Add("msg", url.QueryEscape("自动化创建失败!请自行清空垃圾数据!"))
			return &response.Response{}
		}
	}
	err := service.Generate.Create(&info)
	if err != nil {
		if errors.Is(err, model.ErrorAutoMove) {
			r.Response.Header().Add("success", "false")
			r.Response.Header().Add("msgtype", "success")
			r.Response.Header().Add("msg", url.QueryEscape(err.Error()))
		} else {
			r.Response.Header().Add("success", "false")
			r.Response.Header().Add("msg", url.QueryEscape(err.Error()))
			_ = os.Remove("./generate.zip")
		}

	} else {
		r.Response.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "generate.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		r.Response.Header().Add("Content-Type", "application/json")
		r.Response.Header().Add("success", "true")
		r.Response.ServeFile("./generate.zip")
		_ = os.Remove("./generate.zip")
	}
	return &response.Response{}
}
