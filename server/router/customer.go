package router

import (
	v1 "server/app/api/v1"
	"server/app/middleware"

	"github.com/gogf/gf/frame/g"
)

// InitCustomerRouter 注册客户路由
func InitCustomerRouter() {
	CustomerRouter := g.Server().Group("customer").Middleware(
		middleware.JwtAuth,
		middleware.CasbinMiddleware,
	)
	{
		CustomerRouter.POST("customer", v1.CreateCustomer)     // 创建客户
		CustomerRouter.PUT("customer", v1.UpdateCustomer)      // 更新客户
		CustomerRouter.DELETE("customer", v1.DeleteCustomer)   // 删除客户
		CustomerRouter.GET("customer", v1.FindCustomer)        // 获取单一客户信息
		CustomerRouter.GET("customerList", v1.GetCustomerList) // 获取客户列表
	}
}
