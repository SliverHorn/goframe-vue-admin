package router

import (
	api "github.com/flipped-aurora/gf-vue-admin/server/app/api/extra"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	"github.com/flipped-aurora/gf-vue-admin/server/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type workflow struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewWorkflowRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &workflow{router: router, response: &response.Handler{}}
}

func (s *workflow) Init() {
	group := s.router.Group("/workflowProcess")
	{
		group.POST("createWorkflowProcess", s.response.Handler()(api.WorkflowProcess.Create))              // 新建WorkflowProcess
		group.GET("findWorkflowProcess", s.response.Handler()(api.WorkflowProcess.First))                  // 根据ID获取WorkflowProcess
		group.PUT("updateWorkflowProcess", s.response.Handler()(api.WorkflowProcess.Update))               // 更新WorkflowProcess
		group.DELETE("deleteWorkflowProcess", s.response.Handler()(api.WorkflowProcess.Delete))            // 删除WorkflowProcess
		group.DELETE("deleteWorkflowProcessByIds", s.response.Handler()(api.WorkflowProcess.Deletes))      // 批量删除WorkflowProcess
		group.GET("getWorkflowProcessList", s.response.Handler()(api.WorkflowProcess.GetList))             // 获取WorkflowProcess列表
		group.GET("getMyNeed", s.response.Handler()(api.WorkflowProcess.GetMyNeed))                        // 获取我的待办
		group.GET("getMyStated", s.response.Handler()(api.WorkflowProcess.GetMyStated))                    // 获取我发起的工作流
		group.POST("startWorkflow", s.response.Handler()(api.WorkflowProcess.StartWorkflow))               // 开启工作流
		group.GET("findWorkflowStep", s.response.Handler()(api.WorkflowProcess.FindWorkflowStep))          // 根据ID获取工作流步骤
		group.GET("getWorkflowMoveByID", s.response.Handler()(api.WorkflowProcess.GetWorkflowMoveByID))    // 获取我的待办
		group.POST("completeWorkflowMove", s.response.Handler()(api.WorkflowProcess.CompleteWorkflowMove)) // 提交工作流
	}
}
