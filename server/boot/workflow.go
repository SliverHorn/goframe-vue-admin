package boot

import model "github.com/flipped-aurora/gf-vue-admin/server/app/model/extra"

var Workflow = new(workflow)

type workflow struct{}

func (w *workflow) model() {
	model.WorkflowBusinessStruct = make(map[string]func() model.Workflow)
	model.WorkflowBusinessStruct["leave"] = func() model.Workflow {
		return new(model.LeaveWorkflow)
	}
}

func (w *workflow) table() {
	model.WorkflowBusinessTable = make(map[string]func() interface{})
	model.WorkflowBusinessTable["leave"] = func() interface{} {
		return new(model.WorkflowLeave)
	}
}

func (w *workflow) Initialize() {
	w.model()
	w.table()
}
