package request

import model "gf-vue-admin/server/app/model/extra"

type SearchWorkflowProcess struct {
	model.WorkflowProcess
	PageInfo
}
