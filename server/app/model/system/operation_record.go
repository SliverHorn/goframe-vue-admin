package model

import (
	"gf-vue-admin/server/library/global"
)

type OperationRecord struct {
	global.Model
	Ip           string `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`
	Path         string `json:"path" form:"path" gorm:"column:path;comment:请求路径"`
	Agent        string `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`
	Method       string `json:"method" form:"method" gorm:"column:method;comment:请求方法"`
	Request      string `json:"body" form:"body" gorm:"type:longtext;column:body;comment:请求Body"`
	Response     string `json:"resp" form:"resp" gorm:"type:longtext;column:resp;comment:响应Body"`
	ErrorMessage string `json:"error_message" form:"error_message" gorm:"column:error_message;comment:错误信息"`

	Status int `json:"status" form:"status" gorm:"column:status;comment:请求状态"`
	UserID int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`
	Latency int64 `json:"latency" form:"latency" gorm:"column:latency;comment:延迟"`

	Admin Admin `orm:"-" json:"user" gorm:"foreignkey:ID;references:UserID"`
}

func (o *OperationRecord) TableName() string {
	return "operation_record"
}
