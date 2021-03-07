package api

import (
	"errors"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/request"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"
	service "github.com/flipped-aurora/gf-vue-admin/server/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/server/library/global"
	"github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"time"
)

var GfJWTMiddleware *jwt.GfJWTMiddleware

func init() {
	GfJWTMiddleware, _ = jwt.New(&jwt.GfJWTMiddleware{
		Realm:           global.Config.Jwt.SigningKey,
		Key:             []byte(global.Config.Jwt.SigningKey),
		Timeout:         time.Duration(global.Config.Jwt.ExpiresAt) * time.Hour * 24,
		MaxRefresh:      time.Duration(global.Config.Jwt.RefreshAt) * time.Hour * 24,
		IdentityKey:     "admin_id",
		TokenLookup:     "header:Authorization, query:token, cookie:jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		LoginResponse:   LoginResponse,
		RefreshResponse: RefreshResponse,
		Unauthorized:    Unauthorized,
		IdentityHandler: IdentityHandler,
		PayloadFunc:     PayloadFunc,
	})
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 设置JWT的身份。
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return  claims[GfJWTMiddleware.IdentityKey]
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 用于定义自定义的未经授权的回调函数。
func Unauthorized(r *ghttp.Request, code int, message string) {
	_ = r.Response.WriteJson(&response.Response{Code: 7, Data: g.Map{"reload": true}, Message: "未登录或非法访问或" + message})
	r.ExitAll()
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 用于定义自定义的登录成功回调函数
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	claims := r.GetParam("admin")
	data, ok := claims.(*model.Admin)
	r.SetParam("claims", data) // 设置参数保存到请求中
	if !ok {
		_ = r.Response.WriteJson(&response.Response{Code: 7, Message: "登录失败!"})
		r.Exit()
	}
	if !global.Config.System.UseMultipoint {
		_ = r.Response.WriteJson(&response.Response{Code: 0, Data: g.Map{"user": data, "token": token, "expiresAt": expire.Unix() * 1000}, Message: "登录成功!"})
		r.Exit()
	}
	redisJwt, err := service.JwtBlacklist.GetRedisJWT(data.Uuid)
	if redisJwt == "" {
		if err = service.JwtBlacklist.SetRedisJWT(data.Uuid, token); err != nil {
			_ = r.Response.WriteJson(&response.Response{Code: 7, Message: "设置登录状态失败!"})
			r.Exit()
		}
		_ = r.Response.WriteJson(&response.Response{Code: 0, Data: g.Map{"user": data, "token": token, "expiresAt": expire.Unix() * 1000}, Message: "登录成功!"})
		r.Exit()
	}
	if err = service.JwtBlacklist.JwtToBlacklist(redisJwt); err != nil {
		_ = r.Response.WriteJson(&response.Response{Code: 7, Message: "jwt作废失败!"})
		r.Exit()
	}
	if err = service.JwtBlacklist.SetRedisJWT(data.Uuid, token); err != nil {
		_ = r.Response.WriteJson(&response.Response{Code: 7, Message: "设置登录状态失败!"})
		r.Exit()
	}
	_ = r.Response.WriteJson(&response.Response{Code: 0, Data: g.Map{"user": data, "token": token, "expiresAt": expire.Unix() * 1000}, Message: "登录成功!"})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 用于获取新令牌，无论当前令牌是否过期。
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	//if service.IsBlacklist(token) {
	//	global.Result(r, global.ERROR, g.Map{"reload": true}, "您的帐户异地登陆或令牌失效")
	//	r.ExitAll()
	//}
	//Token, err := GfJWTMiddleware.ParseToken(r) // 解析token
	//if err != nil {
	//	global.FailWithMessage(r, "Token不正确,更新失败")
	//	r.Exit()
	//}
	//var (
	//	claims   = gconv.Map(Token.Claims)
	//	redisJwt string
	//	admin    *admins.AdminHasOneAuthority
	//)
	//admin, err = service.FindAdmin(gconv.String(claims["admin_uuid"]))
	//if err != nil {
	//	global.FailWithMessage(r, "刷新Token失败")
	//	r.Exit()
	//}
	//if !g.Cfg("system").GetBool("system.UseMultipoint") {
	//	global.OkDetailed(r, response.AdminLogin{User: admin, Token: token, ExpiresAt: expire.Unix() * 1000}, "登录成功!")
	//	r.Exit()
	//}
	//if redisJwt, err = service.GetRedisJWT(admin.Uuid); err != nil {
	//	global.FailWithMessage(r, "刷新Token失败")
	//	r.Exit()
	//}
	//if err == nil && redisJwt != "" {
	//	if err = service.JsonInBlacklist(&jwts.Entity{Jwt: redisJwt}); err != nil {
	//		global.Result(r, code, g.Map{}, "jwt作废失败")
	//		r.Exit()
	//	}
	//	if err := service.SetRedisJWT(admin.Uuid, token); err != nil {
	//		global.Result(r, code, g.Map{}, "设置登录状态失败")
	//		r.Exit()
	//	}
	//}
	//global.OkDetailed(r, response.AdminLogin{User: admin, Token: token, ExpiresAt: expire.Unix() * 1000}, "登录成功!")
	r.ExitAll()
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 用于验证登录参数。 它必须返回用户数据作为用户标识符，并将其存储在Claim Array中。 检查错误（e），以确定适当的错误消息。
func Authenticator(r *ghttp.Request) (interface{}, error) {
	var info request.AdminLogin
	if err := r.Parse(&info); err != nil {
		_ = r.Response.WriteJson(&response.Response{Code: 7, Message: err.Error()})
		r.Exit()
	}
	if !service.Store.Verify(info.CaptchaId, info.Captcha, true) { // 验证码校对
		return nil, errors.New("验证码错误! ")
	}
	if data, err := service.Admin.Login(&info); err != nil {
		_ = r.Response.WriteJson(&response.Response{Code: 7, Error: err})
		r.ExitAll()
		return nil, nil
	} else {
		r.SetParam("admin", data) // 设置参数保存到请求中
		return g.Map{"admin_uuid": data.Uuid, "admin_id": data.ID, "admin_nickname": data.Nickname, "admin_authority_id": data.AuthorityId}, nil
	}
}
