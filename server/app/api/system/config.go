package api

import (
	"gf-vue-admin/server/app/api/response"
	service "gf-vue-admin/server/app/service/system"
	"gf-vue-admin/server/library/config"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Config = new(_config)

type _config struct{}

// @Tags SystemConfig
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/getSystemConfig [post]
func (c *_config) GetConfig(r *ghttp.Request) *response.Response {
	return &response.Response{Data: g.Map{"config": service.Config.GetConfig()}, MessageCode: response.SuccessOperation}
}

// @Tags SystemConfig
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body config.Config true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /system/setSystemConfig [post]
func (c *_config) SetConfig(r *ghttp.Request) *response.Response {
	var info config.Config
	if err := service.Config.SetConfig(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorOperation}
	}
	return &response.Response{MessageCode: response.SuccessOperation}
}

// @Tags SystemConfig
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/getServerInfo [post]
func (c *_config) GetServerInfo(r *ghttp.Request) *response.Response {
	if result, err := service.Config.GetServerInfo(); err != nil {
		g.Log().Error("获取失败!", g.Map{"err": err})
		return &response.Response{Error: err, MessageCode: response.ErrorOperation}
	} else {
		return &response.Response{Data: g.Map{"server": result}, MessageCode: response.SuccessOperation}
	}
}

// @Tags SystemConfig
// @Summary 重启服务
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"重启系统成功"}"
// @Router /system/reloadSystem [post]
func (c *_config) ReloadSystem(r *ghttp.Request) *response.Response {
	if err := ghttp.RestartAllServer(); err != nil {
		return &response.Response{Code: 7, Error: err, Message: "重启系统失败!"}
	}
	return &response.Response{Code: 0, Message: "重启系统成功!"}
}
