package v1

import (
	"fmt"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/service"
	"server/library/global"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

// CreateCustomer Create Customer
// CreateCustomer 创建Customer
func CreateCustomer(r *ghttp.Request) {
	var create request.CreateCustomer
	if err := r.Parse(&create); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	claims := getAdminClaims(r)
	create.SysUserId = uint(claims.AdminId)
	create.SysUserAuthorityId = claims.AdminAuthorityId
	if err := service.CreateCustomers(&create); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("创建失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "创建成功")
}

// DeleteCustomer Delete Customers
// DeleteCustomer 删除Customers
func DeleteCustomer(r *ghttp.Request) {
	var delete request.DeleteById
	if err := r.Parse(&delete); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteCustomers(&delete); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// DeleteCustomersByIds Batch delete Customers
// DeleteCustomersByIds 批量删除Customers
func DeleteCustomersByIds(r *ghttp.Request) {
	var deletes request.DeleteByIds
	if err := r.Parse(&deletes); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteCustomersByIds(&deletes); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("批量删除失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "批量删除成功")
}

// UpdateCustomers Update Customer
// UpdateCustomers 更新Customer
func UpdateCustomer(r *ghttp.Request) {
	var update request.UpdateCustomer
	if err := r.Parse(&update); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.UpdateCustomers(&update); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("更新失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "更新成功")
}

// FindCustomer Query Customer with id
// FindCustomer 用id查询Customer
func FindCustomer(r *ghttp.Request) {
	var find request.FindById
	if err := r.Parse(&find); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	data, err := service.FindCustomers(&find)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, g.Map{"customer": data})
}

// GetCustomerList Page out the Customers list
// GetCustomerList 分页获取Customers列表
func GetCustomerList(r *ghttp.Request) {
	var pageInfo request.GetCustomerList
	if err := r.Parse(&pageInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	claims := getAdminClaims(r)
	pageInfo.SysUserAuthorityId = claims.AdminAuthorityId
	list, total, err := service.GetCustomersList(&pageInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	})
}
