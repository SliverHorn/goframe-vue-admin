package internal

import (
	"bytes"
	"gf-vue-admin/server/app/api/request"
	"gf-vue-admin/server/app/api/response"
	api "gf-vue-admin/server/app/api/system"
	model "gf-vue-admin/server/app/model/system"
	service "gf-vue-admin/server/app/service/system"
	"gf-vue-admin/server/library/global"
	"gf-vue-admin/server/library/utils"
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"io/ioutil"
	"strconv"
	"time"
)

var Middleware = new(middleware)

type middleware struct {
	id     int
	err    error
	body   []byte
	result *model.Admin
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 验证token有效性
func (m *middleware) JwtAuth(r *ghttp.Request) {
	api.GfJWTMiddleware.MiddlewareFunc()(r)
	_jwt, err := api.GfJWTMiddleware.ParseToken(r) // 解析token
	if err != nil {
		if err == jwt.ErrExpiredToken {
			_ = r.Response.WriteJson(&response.Response{Code: 7, Data: g.Map{"reload": true}, Message: "授权已过期!"})
			r.ExitAll()
		}
		_ = r.Response.WriteJson(&response.Response{Code: 7, Data: g.Map{"reload": true}, Error: err})
		r.ExitAll()
	}
	if _jwt != nil {
		token := _jwt.Raw
		if service.JwtBlacklist.IsBlacklist(token) {
			_ = r.Response.WriteJson(&response.Response{Code: 7, Data: g.Map{"reload": true}, Message: "您的帐户异地登陆或令牌失效!"})
			r.ExitAll()
		}
		var claims = gconv.Map(_jwt.Claims)
		r.SetParam("claims", _jwt.Claims)
		r.SetParam("admin_authority_id", claims["admin_authority_id"])
		if global.Config.System.UseMultipoint {
			if !service.JwtBlacklist.ValidatorRedisToken(gconv.String(claims["admin_uuid"]), token) {
				_ = r.Response.WriteJson(&response.Response{Code: 7, Data: g.Map{"reload": true}, Message: "Token鉴权失败!"})
				r.Exit()
			}
		}
	}

	r.Middleware.Next()
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 拦截器
func (m *middleware) CasbinRbac(r *ghttp.Request) {
	// 获取请求的URI
	obj := r.Request.URL.RequestURI()
	// 获取请求方法
	act := r.Request.Method
	// 获取用户的角色
	sub := r.GetParam("admin_authority_id")
	e := service.Casbin.Casbin()
	// 判断策略中是否存在
	success, _ := e.Enforce(sub, obj, act)
	if global.Config.System.Env == "develop" || success {
		r.Middleware.Next()
	} else {
		_ = r.Response.WriteJson(&response.Response{Code: 7, Data: g.Map{}, Message: "权限不足"})
		r.ExitAll()
	}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 操作记录中间件
func (m *middleware) OperationRecord(r *ghttp.Request) {
	// Request

	if m.body, m.err = ioutil.ReadAll(r.Request.Body); m.err != nil {
		g.Log().Error(g.Map{"err": m.err})
	}

	r.Request.Body = ioutil.NopCloser(bytes.NewBuffer(m.body))

	if token, err := api.GfJWTMiddleware.ParseToken(r); err != nil { // 优先从jwt获取用户信息
		id, _ := strconv.Atoi(r.Request.Header.Get("x-user-id"))
		if m.result, m.err = service.Admin.FindAdminById(&request.GetById{Id: uint(id)}); m.err != nil {
			g.Log().Error(`Function service.Admin.FindAdminById() Failed!`, g.Map{"err": m.err})
		}
	} else {
		claims := gconv.Map(token.Claims)
		uuid := gconv.String(claims["admin_uuid"])
		if m.result, m.err = service.Admin.FindAdmin(&request.GetByUuid{Uuid: uuid}); m.err != nil {
			g.Log().Error(`Function service.Admin.FindAdmin() Failed!`, g.Map{"err": m.err})
		}
		m.id = int(m.result.ID)
	}

	record := request.CreateOperationRecord{
		BaseOperationRecord: request.BaseOperationRecord{
			Ip:      r.GetClientIp(),
			Method:  r.Request.Method,
			Path:    r.Request.URL.Path,
			Agent:   r.Request.UserAgent(),
			Request: string(m.body), UserID: m.id,
		},
	}
	now := time.Now()

	r.Middleware.Next()

	// Response

	latency := time.Now().Sub(now)

	if err := r.GetError(); err != nil {
		record.ErrorMessage = err.Error()
	}

	record.Status = r.Response.Status
	record.Latency = latency.Microseconds()
	record.Response = string(r.Response.Buffer())

	if err := service.OperationRecord.Create(&record); err != nil {
		g.Log().Error("create operation record error:", g.Map{"err": err})
	}

	str := "接收到的请求为" + record.Request + "\n" + "请求方式为" + record.Method + "\n" + "报错信息如下" + record.ErrorMessage + "\n" + "耗时" + latency.String() + "\n"
	if global.Config.System.ErrorToEmail {
		if record.Status != 200 {
			subject := m.result.Username + "" + record.Ip + "调用了" + record.Path + "报错了"
			if err := utils.Email.ErrorToEmail(subject, str); err != nil {
				g.Log().Errorf("Function utils.Email.ErrorToEmail Failed!", g.Map{"err": err})
			}
		}
	}
}
