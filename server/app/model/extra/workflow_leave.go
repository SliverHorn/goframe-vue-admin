package model

import (
	"gf-vue-admin/server/library/global"
	"time"
)

type WorkflowLeave struct {
	global.Model
	Cause     string    `json:"cause" form:"cause" gorm:"column:cause;comment:"`
	StartTime time.Time `json:"startTime" form:"startTime" gorm:"column:start_time;comment:"`
	EndTime   time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:"`
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 工作流操作结构体
type LeaveWorkflow struct {
	WorkflowBase  `json:"wf"`
	WorkflowLeave `json:"business"`
}

func (e *WorkflowLeave) TableName() string {
	return "exa_wf_leaves"
}
