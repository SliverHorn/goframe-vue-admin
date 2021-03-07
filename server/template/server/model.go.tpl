package model

import "github.com/flipped-aurora/gf-vue-admin/server/library/global"

// 如果含有time.Time 请自行import time包
type {{.StructName}} struct {
      global.Model {{- range .Fields}}
            {{- if eq .FieldType "bool" }}
      {{.FieldName}}  *{{.FieldType}} `orm:"{{.ColumnName}}" json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- if eq .FieldType "string" -}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}"`
            {{- else }}
      {{.FieldName}}  {{.FieldType}} `orm:"{{.ColumnName}}" json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- if eq .FieldType "string" -}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}"`
            {{- end }} {{- end }}
}

{{ if .TableName }}
func (m *{{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type {{.StructName}}Workflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	{{.StructName}}   `json:"business"`
// }

// func ({{.StructName}}) TableName() string {
// 	return "{{.TableName}}"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["{{.Abbreviation}}"] = func() model.GVA_Workflow {
//   return new(model.{{.StructName}}Workflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["{{.Abbreviation}}"] = func() interface{} {
// 	return new(model.{{.StructName}})
// }