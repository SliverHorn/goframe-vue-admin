package api

import (
	"gf-vue-admin/server/app/api/response"
	service "gf-vue-admin/server/app/service/system"
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

var JwtBlacklist = new(blacklist)

type blacklist struct{}

// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func (b *blacklist) JwtToBlacklist(r *ghttp.Request) *response.Response {
	var token = r.Request.Header.Get("Authorization")
	var parts = strings.SplitN(token, " ", 2)
	if err := service.JwtBlacklist.JwtToBlacklist(parts[1]); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorJwtBlackList}
	}
	return &response.Response{MessageCode: response.SuccessJwtBlackList}
}
