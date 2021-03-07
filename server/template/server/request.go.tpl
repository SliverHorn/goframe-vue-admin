package request

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/model"
	"github.com/gogf/gf/frame/g"
)

type Search{{.StructName}} struct{
    model.{{.StructName}}
    PageInfo
}

func (s *Search{{.StructName}}) Search() g.Map {
	condition := g.Map{}
	{{- range .Fields}}
                {{- if .FieldSearchType}}
                    {{- if eq .FieldType "string" }}
        if s.{{.FieldName}} != "" {
            condition["`{{.ColumnName}}`{{if eq .FieldSearchType "LIKE"}} like ?{{ end }}"] = {{if eq .FieldSearchType "LIKE"}}"%" +{{ end }} s.{{.FieldName}} {{if eq .FieldSearchType "LIKE"}}+ "%"{{ end }}
        }
                    {{- else if eq .FieldType "bool" }}
        if s.{{.FieldName}} != nil {
            if *s.{{.FieldName}} == true {
                condition["`{{.ColumnName}}`"] = 1
            } else {
                condition["`{{.ColumnName}}`"] = 2
            }
        }
                    {{- else if eq .FieldType "int" }}
        if s.{{.FieldName}} != 0 {
            condition["`{{.ColumnName}}`{{if eq .FieldSearchType "LIKE"}} like ?{{ end }}"] = {{if eq .FieldSearchType "LIKE"}}"%" +{{ end }} s.{{.FieldName}} {{if eq .FieldSearchType "LIKE"}}+ "%"{{ end }}
        }
                    {{- else if eq .FieldType "float64" }}
        if s.{{.FieldName}} != 0 {
            condition["`{{.ColumnName}}`{{if eq .FieldSearchType "LIKE"}} like ?{{ end }}"] = {{if eq .FieldSearchType "LIKE"}}"%" +{{ end }} s.{{.FieldName}} {{if eq .FieldSearchType "LIKE"}}+ "%"{{ end }}
        }
                    {{- else if eq .FieldType "time.Time" }}
        if !s.{{.FieldName}}.IsZero() {
            condition["`{{.ColumnName}}`{{if eq .FieldSearchType "LIKE"}} like ?{{ end }}"] = {{if eq .FieldSearchType "LIKE"}}"%" +{{ end }} s.{{.FieldName}} {{if eq .FieldSearchType "LIKE"}}+ "%"{{ end }}
        }
                    {{- end }}
            {{- end }}
        {{- end }}
    return condition
}