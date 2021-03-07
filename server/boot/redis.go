package boot

import (
	"context"
	"gf-vue-admin/server/library/global"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/frame/g"
)

var Redis = new(_redis)

type _redis struct{}

func (r *_redis) Initialize() {
	if global.Config.System.UseMultipoint {
		client := redis.NewClient(&redis.Options{
			DB:       global.Config.Redis.DB, // use default DB
			Addr:     global.Config.Redis.Address,
			Password: global.Config.Redis.Password, // no password set
		})
		if pong, err := client.Ping(context.Background()).Result(); err != nil {
			g.Log().Error("redis connect ping failed!", g.Map{"err": err})
		} else {
			g.Log().Error("redis connect ping response ", g.Map{"pong": pong})
			global.Redis = client
		}
	}
}
