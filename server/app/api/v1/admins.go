package v1

import (
	"fmt"
	"mime/multipart"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/model/admins"
	"server/app/service"
	"server/library/global"
	"server/library/utils"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/net/ghttp"
)

// ChangePassword "Administrator changes password"
// ChangePassword "管理员修改密码"
func ChangePassword(r *ghttp.Request) {
	var c request.ChangePassword
	if err := r.Parse(&c); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	claims := getAdminClaims(r)
	c.Uuid = claims.AdminUuid
	if err := service.ChangePassword(&c); err == nil {
		global.OkWithMessage(r, "修改失败")
		r.Exit()
	}
	global.OkWithMessage(r, "修改成功")
}

// UploadHeaderImg User uploads profile picture
// UploadHeaderImg 用户上传头像
func UploadHeaderImg(r *ghttp.Request) {
	var (
		err      error
		filePath string
		header   *multipart.FileHeader
		admin    *admins.Entity
	)
	claims := getAdminClaims(r)
	userUuid := claims.AdminUuid
	if _, header, err = r.Request.FormFile("headerImg"); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("上传文件失败，%v", err))
	}
	if filePath, _, err = utils.Upload(header); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("接收返回值失败，%v", err))
	}
	// 修改数据库后得到修改后的user并且返回供前端使用
	admin, err = service.UploadHeaderImg(userUuid, filePath)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("修改数据库链接失败，%v", err))
	} else {
		global.OkDetailed(r, response.AdminResponse{Admin: admin}, "上传成功")
	}
}

// GetAdminList Paging gets the list of users
// GetAdminList 分页获取用户列表
func GetAdminList(r *ghttp.Request) {
	var get request.PageInfo
	if err := r.Parse(&get); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetAdminList(&get)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，err:%v", err))
		r.Exit()
	}
	global.OkDetailed(r, response.PageResult{List: list, Total: total, Page: get.Page, PageSize: get.PageSize}, "获取成功")
}

// SetUserAuthority Set user permissions
// SetUserAuthority 设置用户权限
func SetUserAuthority(r *ghttp.Request) {
	var set request.SetAdminAuthority
	if err := r.Parse(&set); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.SetUserAuthority(&set); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("修改失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "修改成功")
}

// DeleteAdmin Delete user
// DeleteAdmin 删除用户
func DeleteAdmin(r *ghttp.Request) {
	var d request.DeleteById
	if err := r.Parse(&d); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteAdmin(&d); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除成功, err:%v", err))
	}
	global.OkWithMessage(r, "删除成功")
}

// getAdminClaims 获取jwt里含有的管理员信息
func getAdminClaims(r *ghttp.Request) (claims *request.CustomClaims) {
	claims = new(request.CustomClaims)
	claimsMap := r.GetParam("claims")
	if err := gconv.Struct(claimsMap, claims); err != nil {
		g.Log().Errorf("管理员信息失败!, err:%v", err)
		r.ExitAll()
	}
	return
}
