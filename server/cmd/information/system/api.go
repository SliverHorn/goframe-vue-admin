package information

import (
	model "gf-vue-admin/server/app/model/system"
	"gf-vue-admin/server/library/global"
	"github.com/gookit/color"
	"time"

	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

var apis = []model.Api{
	{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/base/login", Description: "用户登录", ApiGroup: "base", Method: "POST"},
	{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/register", Description: "用户注册", ApiGroup: "user", Method: "POST"},
	{Model: global.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/createApi", Description: "创建api", ApiGroup: "api", Method: "POST"},
	{Model: global.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getApiList", Description: "获取api列表", ApiGroup: "api", Method: "POST"},
	{Model: global.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getApiById", Description: "获取api详细信息", ApiGroup: "api", Method: "POST"},
	{Model: global.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/deleteApi", Description: "删除Api", ApiGroup: "api", Method: "POST"},
	{Model: global.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/updateApi", Description: "更新Api", ApiGroup: "api", Method: "POST"},
	{Model: global.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getAllApis", Description: "获取所有api", ApiGroup: "api", Method: "POST"},
	{Model: global.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/createAuthority", Description: "创建角色", ApiGroup: "authority", Method: "POST"},
	{Model: global.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/deleteAuthority", Description: "删除角色", ApiGroup: "authority", Method: "POST"},
	{Model: global.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/getAuthorityList", Description: "获取角色列表", ApiGroup: "authority", Method: "POST"},
	{Model: global.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenu", Description: "获取菜单树", ApiGroup: "menu", Method: "POST"},
	{Model: global.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenuList", Description: "分页获取基础menu列表", ApiGroup: "menu", Method: "POST"},
	{Model: global.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/addBaseMenu", Description: "新增菜单", ApiGroup: "menu", Method: "POST"},
	{Model: global.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getBaseMenuTree", Description: "获取用户动态路由", ApiGroup: "menu", Method: "POST"},
	{Model: global.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/addMenuAuthority", Description: "增加menu和角色关联关系", ApiGroup: "menu", Method: "POST"},
	{Model: global.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenuAuthority", Description: "获取指定角色menu", ApiGroup: "menu", Method: "POST"},
	{Model: global.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/deleteBaseMenu", Description: "删除菜单", ApiGroup: "menu", Method: "POST"},
	{Model: global.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/updateBaseMenu", Description: "更新菜单", ApiGroup: "menu", Method: "POST"},
	{Model: global.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getBaseMenuById", Description: "根据id获取菜单", ApiGroup: "menu", Method: "POST"},
	{Model: global.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/changePassword", Description: "修改密码", ApiGroup: "user", Method: "POST"},
	{Model: global.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/getUserList", Description: "获取用户列表", ApiGroup: "user", Method: "POST"},
	{Model: global.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/setUserAuthority", Description: "修改用户角色", ApiGroup: "user", Method: "POST"},
	{Model: global.Model{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/upload", Description: "文件上传示例", ApiGroup: "fileUploadAndDownload", Method: "POST"},
	{Model: global.Model{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/getFileList", Description: "获取上传文件列表", ApiGroup: "fileUploadAndDownload", Method: "POST"},
	{Model: global.Model{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/updateCasbin", Description: "更改角色api权限", ApiGroup: "casbin", Method: "POST"},
	{Model: global.Model{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/getPolicyPathByAuthorityId", Description: "获取权限列表", ApiGroup: "casbin", Method: "POST"},
	{Model: global.Model{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/deleteFile", Description: "删除文件", ApiGroup: "fileUploadAndDownload", Method: "POST"},
	{Model: global.Model{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/jwt/jsonInBlacklist", Description: "jwt加入黑名单", ApiGroup: "jwt", Method: "POST"},
	{Model: global.Model{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/setDataAuthority", Description: "设置角色资源权限", ApiGroup: "authority", Method: "POST"},
	{Model: global.Model{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/getSystemConfig", Description: "获取配置文件内容", ApiGroup: "system", Method: "POST"},
	{Model: global.Model{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/setSystemConfig", Description: "设置配置文件内容", ApiGroup: "system", Method: "POST"},
	{Model: global.Model{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: "创建客户", ApiGroup: "customer", Method: "POST"},
	{Model: global.Model{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: "更新客户", ApiGroup: "customer", Method: "PUT"},
	{Model: global.Model{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: "删除客户", ApiGroup: "customer", Method: "DELETE"},
	{Model: global.Model{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: "获取单一客户", ApiGroup: "customer", Method: "GET"},
	{Model: global.Model{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customerList", Description: "获取客户列表", ApiGroup: "customer", Method: "GET"},
	{Model: global.Model{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/casbinTest/:pathParam", Description: "RESTFUL模式测试", ApiGroup: "casbin", Method: "GET"},
	{Model: global.Model{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/createTemp", Description: "自动化代码", ApiGroup: "autoCode", Method: "POST"},
	{Model: global.Model{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/updateAuthority", Description: "更新角色信息", ApiGroup: "authority", Method: "PUT"},
	{Model: global.Model{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/copyAuthority", Description: "拷贝角色", ApiGroup: "authority", Method: "POST"},
	{Model: global.Model{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/deleteUser", Description: "删除用户", ApiGroup: "user", Method: "DELETE"},
	{Model: global.Model{ID: 44, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: "新增字典内容", ApiGroup: "sysDictionaryDetail", Method: "POST"},
	{Model: global.Model{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: "删除字典内容", ApiGroup: "sysDictionaryDetail", Method: "DELETE"},
	{Model: global.Model{ID: 46, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: "更新字典内容", ApiGroup: "sysDictionaryDetail", Method: "PUT"},
	{Model: global.Model{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: "根据ID获取字典内容", ApiGroup: "sysDictionaryDetail", Method: "GET"},
	{Model: global.Model{ID: 48, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: "获取字典内容列表", ApiGroup: "sysDictionaryDetail", Method: "GET"},
	{Model: global.Model{ID: 49, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/createSysDictionary", Description: "新增字典", ApiGroup: "sysDictionary", Method: "POST"},
	{Model: global.Model{ID: 50, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/deleteSysDictionary", Description: "删除字典", ApiGroup: "sysDictionary", Method: "DELETE"},
	{Model: global.Model{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/updateSysDictionary", Description: "更新字典", ApiGroup: "sysDictionary", Method: "PUT"},
	{Model: global.Model{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/findSysDictionary", Description: "根据ID获取字典", ApiGroup: "sysDictionary", Method: "GET"},
	{Model: global.Model{ID: 53, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/getSysDictionaryList", Description: "获取字典列表", ApiGroup: "sysDictionary", Method: "GET"},
	{Model: global.Model{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/createSysOperationRecord", Description: "新增操作记录", ApiGroup: "sysOperationRecord", Method: "POST"},
	{Model: global.Model{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/deleteSysOperationRecord", Description: "删除操作记录", ApiGroup: "sysOperationRecord", Method: "DELETE"},
	{Model: global.Model{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/findSysOperationRecord", Description: "根据ID获取操作记录", ApiGroup: "sysOperationRecord", Method: "GET"},
	{Model: global.Model{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/getSysOperationRecordList", Description: "获取操作记录列表", ApiGroup: "sysOperationRecord", Method: "GET"},
	{Model: global.Model{ID: 58, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getTables", Description: "获取数据库表", ApiGroup: "autoCode", Method: "GET"},
	{Model: global.Model{ID: 59, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getDB", Description: "获取所有数据库", ApiGroup: "autoCode", Method: "GET"},
	{Model: global.Model{ID: 60, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getColumn", Description: "获取所选table的所有字段", ApiGroup: "autoCode", Method: "GET"},
	{Model: global.Model{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: "批量删除操作历史", ApiGroup: "sysOperationRecord", Method: "DELETE"},
	{Model: global.Model{ID: 62, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUploader/upload", Description: "插件版分片上传", ApiGroup: "simpleUploader", Method: "POST"},
	{Model: global.Model{ID: 63, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUploader/checkFileMd5", Description: "文件完整度验证", ApiGroup: "simpleUploader", Method: "GET"},
	{Model: global.Model{ID: 64, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUploader/mergeFileMd5", Description: "上传完成合并文件", ApiGroup: "simpleUploader", Method: "GET"},
	{Model: global.Model{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/setUserInfo", Description: "设置用户信息", ApiGroup: "user", Method: "PUT"},
	{Model: global.Model{ID: 66, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/getServerInfo", Description: "获取服务器信息", ApiGroup: "system", Method: "POST"},
	{Model: global.Model{ID: 67, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/email/emailTest", Description: "发送测试邮件", ApiGroup: "email", Method: "POST"},
	{Model: global.Model{ID: 68, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/createWorkflowProcess", Description: "新建工作流", ApiGroup: "workflowProcess", Method: "POST"},
	{Model: global.Model{ID: 69, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/deleteWorkflowProcess", Description: "删除工作流", ApiGroup: "workflowProcess", Method: "DELETE"},
	{Model: global.Model{ID: 70, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/deleteWorkflowProcessByIds", Description: "批量删除工作流", ApiGroup: "workflowProcess", Method: "DELETE"},
	{Model: global.Model{ID: 71, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/updateWorkflowProcess", Description: "更新工作流", ApiGroup: "workflowProcess", Method: "PUT"},
	{Model: global.Model{ID: 72, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/findWorkflowProcess", Description: "根据ID获取工作流", ApiGroup: "workflowProcess", Method: "GET"},
	{Model: global.Model{ID: 73, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/getWorkflowProcessList", Description: "获取工作流", ApiGroup: "workflowProcess", Method: "GET"},
	{Model: global.Model{ID: 74, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/findWorkflowStep", Description: "获取工作流步骤", ApiGroup: "workflowProcess", Method: "GET"},
	{Model: global.Model{ID: 75, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/startWorkflow", Description: "启动工作流", ApiGroup: "workflowProcess", Method: "POST"},
	{Model: global.Model{ID: 76, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/getMyStated", Description: "获取我发起的工作流", ApiGroup: "workflowProcess", Method: "GET"},
	{Model: global.Model{ID: 77, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/getMyNeed", Description: "获取我的待办", ApiGroup: "workflowProcess", Method: "GET"},
	{Model: global.Model{ID: 78, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/getWorkflowMoveByID", Description: "根据id获取当前节点详情和历史", ApiGroup: "workflowProcess", Method: "GET"},
	{Model: global.Model{ID: 79, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/completeWorkflowMove", Description: "提交工作流", ApiGroup: "workflowProcess", Method: "POST"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: apis 表数据初始化
func (a *api) Init() error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 67}).Find(&[]model.Api{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> apis 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&apis).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> apis 表初始数据成功!")
		return nil
	})
}
