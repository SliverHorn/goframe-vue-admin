package api

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/internal"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/request"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/extra"
	service "github.com/flipped-aurora/gf-vue-admin/server/app/service/extra"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var WorkflowProcess = new(process)

type process struct{}

// @Tags ExtraWorkflowProcess
// @Summary 创建WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "创建WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/createWorkflowProcess [post]
func (p *process) Create(r *ghttp.Request) *response.Response {
	var info model.WorkflowProcess
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.CreateWorkflowProcess(info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// @Tags ExtraWorkflowProcess
// @Summary 用id查询WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "用id查询WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /workflowProcess/findWorkflowProcess [get]
func (p *process) First(r *ghttp.Request) *response.Response {
	var info model.WorkflowProcess
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if reslut, err := service.GetWorkflowProcess(info.ID); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"reworkflowProcess": reslut}, MessageCode: response.ErrorFirst}
	}
}

// @Tags ExtraWorkflowProcess
// @Summary 更新WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "更新WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /workflowProcess/updateWorkflowProcess [put]
func (p *process) Update(r *ghttp.Request) *response.Response {
	var info model.WorkflowProcess
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := service.UpdateWorkflowProcess(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	return &response.Response{MessageCode: response.SuccessUpdated}
}

// @Tags ExtraWorkflowProcess
// @Summary 删除WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "删除WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /workflowProcess/deleteWorkflowProcess [delete]
func (p *process) Delete(r *ghttp.Request) *response.Response {
	var info model.WorkflowProcess
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.DeleteWorkflowProcess(info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags ExtraWorkflowProcess
// @Summary 批量删除WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /workflowProcess/deleteWorkflowProcessByIds [delete]
func (p *process) Deletes(r *ghttp.Request) *response.Response {
	var info request.GetByIds
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := service.DeleteWorkflowProcessByIds(info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// @Tags ExtraWorkflowProcess
// @Summary 用id查询工作流步骤
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "用id查询WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /workflowProcess/findWorkflowStep [get]
func (p *process) FindWorkflowStep(r *ghttp.Request) *response.Response {
	var info model.WorkflowProcess
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, Message: "查询失败!"}
	}
	if result, err := service.FindWorkflowStep(info.ID); err != nil {
		return &response.Response{Error: err, Message: "查询失败!"}
	} else {
		return &response.Response{Data: g.Map{"workflow": result}, Message: "查询成功!"}
	}
}

// @Tags ExtraWorkflowProcess
// @Summary 分页获取WorkflowProcess列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchWorkflowProcess true "分页获取WorkflowProcess列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/getWorkflowProcessList [get]
func (p *process) GetList(r *ghttp.Request) *response.Response {
	var info request.SearchWorkflowProcess
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, Message: "获取数据失败!"}
	}
	if list, total, err := service.GetWorkflowProcessInfoList(info); err != nil {
		return &response.Response{Error: err, Message: "获取数据失败!"}
	} else {
		return &response.Response{Data: response.PageResult{
			List:     list,
			Total:    int(total),
			Page:     info.Page,
			PageSize: info.PageSize,
		}, Message: "获取数据成功!"}
	}
}

// @Tags ExtraWorkflowProcess
// @Summary 开启工作流
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/startWorkflow [post]
func (p *process) StartWorkflow(r *ghttp.Request) *response.Response {
	business := r.GetQueryString("businessType")
	info := model.WorkflowBusinessStruct[business]()
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, Message: "获取数据失败!"}
	}
	if err := service.StartWorkflow(info); err != nil {
		return &response.Response{Error: err, Message: "开启失败!"}
	}
	return &response.Response{Message: "启动成功!"}
}

// @Tags ExtraWorkflowProcess
// @Summary 提交工作流
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/completeWorkflowMove [post]
func (p *process) CompleteWorkflowMove(r *ghttp.Request) *response.Response {
	business := r.GetQueryString("businessType")
	info := model.WorkflowBusinessStruct[business]()
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, Message: "获取数据失败!"}
	}
	if err := service.CompleteWorkflowMove(info); err != nil {
		return &response.Response{Error: err, Message: "提交失败!"}
	}
	return &response.Response{Message: "提交成功!"}
}

// @Tags ExtraWorkflowProcess
// @Summary 我发起的工作流
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/getMyStated [get]
func (p *process) GetMyStated(r *ghttp.Request) *response.Response {
	if result, err := service.GetMyStated(internal.Info.GetAdminID(r)); err != nil {
		return &response.Response{Error: err, Message: "获取失败!"}
	} else {
		return &response.Response{Data: g.Map{"wfms": result}, Message: "获取失败!"}
	}
}

// @Tags ExtraWorkflowProcess
// @Summary 我的待办
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/getMyNeed [get]
func (p *process) GetMyNeed(r *ghttp.Request) *response.Response {
	claims := internal.Info.GetAdminClaims(r)
	if result, err := service.GetMyNeed(claims.AdminId, claims.AdminAuthorityId); err != nil {
		return &response.Response{Error: err, Message: "获取失败!"}
	} else {
		return &response.Response{Data: g.Map{"wfms": result}, Message: "获取失败!"}
	}
}

// @Tags ExtraWorkflowProcess
// @Summary 根据id获取当前节点详情和历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取当前节点详情和过往"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/getWorkflowMoveByID [get]
func (p *process) GetWorkflowMoveByID(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, Message: "获取数据失败!"}
	}
	if move, moves, business, err := service.GetWorkflowMoveByID(&info); err != nil {
		return &response.Response{Error: err, Message: "获取数据失败!"}
	} else {
		return &response.Response{Data: g.Map{"move": move, "moves": moves, "business": business}, Message: "获取数据成功!"}
	}
}
