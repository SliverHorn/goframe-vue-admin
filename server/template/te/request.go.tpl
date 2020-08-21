package request

type Create{{.StructName}} struct {
    {{- range .Fields}}
        {{- if eq .FieldName "Id" "ID" "CreateAt" "UpdateAt" "DeleteAt"}}{{- else }}
    {{.FieldName}} {{.FieldType}} `p:"{{.FieldJson}}" v:"required|length:1,1000#请输入{{.Comment}}|{{.Comment}}长度为:min到:max位"`
        {{- end }}
    {{- end }}
}

type Update{{.StructName}} struct {
    {{- range .Fields}}
        {{- if eq .FieldName "Id" "ID"}}
    Id {{.FieldType}} `p:"{{.FieldJson}}" v:"required|length:1,1000#请输入{{.Comment}}|{{.Comment}}长度为:min到:max位"`
        {{- else if eq .FieldName "CreateAt" "UpdateAt" "DeleteAt"}}{{- else }}
    {{.FieldName}} {{.FieldType}} `p:"{{.FieldJson}}" v:"required|length:1,1000#请输入{{.Comment}}|{{.Comment}}长度为:min到:max位"`
        {{- end }}
    {{- end }}
}

type Find{{.StructName}} struct {
    {{- range .Fields}}
        {{- if eq .FieldName "Id" "ID"}}
    Id {{.FieldType}} `p:"{{.FieldJson}}" v:"required|length:1,1000#请输入{{.Comment}}|{{.Comment}}长度为:min到:max位"`
        {{- else if eq .FieldName "CreateAt" "UpdateAt" "DeleteAt"}}{{ else }}
    {{.FieldName}} {{.FieldType}} `p:"{{.FieldJson}}" v:"required|length:1,1000#请输入{{.Comment}}|{{.Comment}}长度为:min到:max位"`
        {{- end }}
    {{- end }}
}

type Get{{.StructName}}List struct {
    {{- range .Fields}}
        {{- if eq .FieldName "Id" "ID" "CreateAt" "UpdateAt" "DeleteAt"}}{{ else }}
    {{.FieldName}} {{.FieldType}} `p:"{{.FieldJson}}"`
        {{- end }}
    {{- end }}
	PageInfo
}