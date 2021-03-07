package service

import (
	"gf-vue-admin/server/app/api/response"
	"gf-vue-admin/server/integration"
	"gf-vue-admin/server/library/global"
	"github.com/mojocn/base64Captcha"
)

var (
	Store base64Captcha.Store
	Base = new(base)

)

func init() {
	if global.Config.Captcha.CaptchaInRedis {
		Store = integration.RedisStore
	} else {
		Store = base64Captcha.DefaultMemStore
	}
}
type base struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 生成二维码的信息
func (b *base) Captcha() (result *response.Captcha, err error) {
	var data response.Captcha
	var driver = base64Captcha.NewDriverDigit(global.Config.Captcha.ImageHeight, global.Config.Captcha.ImageWidth, global.Config.Captcha.KeyLong, 0.7, 80) // 字符,公式,验证码配置, 生成默认数字的driver
	var captcha = base64Captcha.NewCaptcha(driver, Store)
	data.Id, data.Path, err = captcha.Generate()
	return &data, err
}
