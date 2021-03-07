package integration

import (
	"context"
	"github.com/flipped-aurora/gf-vue-admin/server/library/global"
	"github.com/gogf/gf/frame/g"
	"github.com/mojocn/base64Captcha"
	"time"
)

var RedisStore = NewRedisStore()

type _redis struct {
	PrefixKey  string
	Expiration time.Duration
}

func NewRedisStore() base64Captcha.Store {
	return &_redis{PrefixKey: "Captcha_", Expiration: 10 * time.Minute}
}

func (r *_redis) Set(id string, value string) {
	if err := global.Redis.Set(context.Background(), r.PrefixKey+id, value, r.Expiration).Err(); err != nil {
		g.Log().Error("设置验证码数据到redis失败!", err)
	}
}

func (r *_redis) Get(id string, clear bool) string {
	var key = r.PrefixKey + id
	if value, err := global.Redis.Get(context.Background(), key).Result(); err != nil {
		g.Log().Error("从redis获取验证码数据失败!", err)
		return value
	} else {
		if clear {
			if err = global.Redis.Del(context.Background(), key).Err(); err != nil {
				g.Log().Error("清空redis中验证码数据失败!", err)
				return ""
			}
		}
		return value
	}
}

func (r *_redis) Verify(id, answer string, clear bool) bool {
	return r.Get(r.PrefixKey + id, clear) == answer
}
