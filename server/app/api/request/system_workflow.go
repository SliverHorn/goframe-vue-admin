package request

import model "github.com/flipped-aurora/gf-vue-admin/server/app/model/extra"

type SearchWorkflowProcess struct {
	model.WorkflowProcess
	PageInfo
}
