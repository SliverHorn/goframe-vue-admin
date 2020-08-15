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

// CreateAuthority Create the role
// CreateAuthority 创建角色
func CreateAuthority(r *ghttp.Request) {
	var c request.CreateAuthority
	if err := r.Parse(&c); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	authority, err := service.CreateAuthority(&c)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("创建失败，%v", err))
		r.Exit()
	}
	global.OkDetailed(r, g.Map{"authority": authority}, "创建成功")
}

// CopyAuthority Copy the role
// CopyAuthority 拷贝角色
func CopyAuthority(r *ghttp.Request) {
	var copyInfo request.AuthorityCopy
	if err := r.Parse(&copyInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	authority, err := service.CopyAuthority(&copyInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("拷贝失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, g.Map{"authority": authority})
}

// DeleteAuthority Delete the role
// DeleteAuthority 删除角色
func DeleteAuthority(r *ghttp.Request) {
	var d request.DeleteAuthority
	if err := r.Parse(&d); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteAuthority(&d); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// UpdateAuthority Set the role resource permissions
// UpdateAuthority 设置角色资源权限
func UpdateAuthority(r *ghttp.Request) {
	var u request.UpdateAuthority
	if err := r.Parse(&u); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.UpdateAuthority(&u); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("更改失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "更改成功")
}

// GetAuthorityList Paging gets the list of roles
// GetAuthorityList 分页获取角色列表
func GetAuthorityList(r *ghttp.Request) {
	var pageInfo request.PageInfo
	if err := r.Parse(&pageInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetAuthorityInfoList(&pageInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，err:%v", err))
		r.Exit()
	}
	global.OkDetailed(r, response.PageResult{List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize}, "获取成功")
}

// SetDataAuthority Set the role resource permissions
// SetDataAuthority 设置角色资源权限
func SetDataAuthority(r *ghttp.Request) {
	var auth *request.SetDataAuthority
	if err := r.Parse(&auth); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	err := service.SetDataAuthority(auth)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("设置关联失败，%v", err))
		r.Exit()
	}
	global.Ok(r)
}
