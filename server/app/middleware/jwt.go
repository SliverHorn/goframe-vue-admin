package middleware

import (
	v1 "server/app/api/v1"
	"server/app/service"
	"server/library/global"

	jwt "github.com/gogf/gf-jwt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// JwtAuth 验证token有效性
func JwtAuth(r *ghttp.Request) {
	v1.GfJWTMiddleware.MiddlewareFunc()(r)
	Token, err := v1.GfJWTMiddleware.ParseToken(r) // 解析token
	if err != nil {
		if err == jwt.ErrExpiredToken {
			global.Result(r, global.ERROR, g.Map{"reload": true}, "授权已过期")
			r.ExitAll()
		}
		global.Result(r, global.ERROR, g.Map{"reload": true}, err.Error())
		r.ExitAll()
	}
	token := Token.Raw
	if service.IsBlacklist(token) {
		global.Result(r, global.ERROR, g.Map{"reload": true}, "您的帐户异地登陆或令牌失效")
		r.ExitAll()
	}
	var claims = gconv.Map(Token.Claims)
	r.SetParam("claims", Token.Claims)
	r.SetParam("admin_authority_id", claims["admin_authority_id"])
	if g.Cfg().GetBool("system.UseMultipoint") {
		if !service.ValidatorRedisToken(gconv.String(claims["admin_uuid"]), token) {
			global.FailWithMessage(r, "Token鉴权失败")
			r.Exit()
		}
	}
	r.Middleware.Next()
}
