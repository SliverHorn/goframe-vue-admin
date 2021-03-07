package api

import (
	"errors"
	"gf-vue-admin/server/app/api/response"
	service "gf-vue-admin/server/app/service/system"
	"gf-vue-admin/server/library/global"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Base = new(base)

type base struct{}

// @Tags SystemBase
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func (b *base) Captcha(r *ghttp.Request) *response.Response {
	if result, err := service.Base.Captcha(); err != nil {
		return &response.Response{Data: result, MessageCode: response.ErrorCaptcha}
	} else {
		return &response.Response{Data: result, MessageCode: response.SuccessCaptcha}
	}
}

// @Tags SystemInitDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Param data body request.InitDB true "初始化数据库参数"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"自动创建数据库成功"}"
// @Router /init/initdb [post]
func (b *base) InitDB() *response.Response {
	if global.Db != nil {
		return &response.Response{Error: errors.New("global.Db 不为空! "), Message: "非法访问!"}
	}
	return &response.Response{Message: "自动创建数据库成功!"}
}

// @Tags SystemCheckDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"探测完成"}"
// @Router /init/checkdb [post]
func (b *base) CheckDB() *response.Response {
	if global.Db != nil {
		return &response.Response{Data: g.Map{"needInit": false}, Message: "数据库无需初始化!"}
	}
	return &response.Response{Data: g.Map{"needInit": true}, Message: "前往初始化数据库!"}
}
