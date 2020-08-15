package service

import (
	"server/app/api/request"
	"server/app/model/{{.TableName}}"
	"server/library/global"

	"github.com/gogf/gf/frame/g"
)

// Create{{.StructName}} create a {{.StructName}}
// Create{{.StructName}} 创建{{.StructName}}
func Create{{.StructName}}(create *request.Create{{.StructName}}) (err error) {
	insert := {{.TableName}}.Entity{
	{{- range .Fields}}
        {{- if eq .FieldName "Id" "ID" "CreateAt" "UpdateAt" "DeleteAt"}}
        {{ else }}
        {{.FieldName}}: create.{{.FieldName}},
        {{- end}}
    {{- end }}
	}
	_, err = {{.TableName}}.Insert(&insert)
	return err
}

// Delete{{.StructName}} delete {{.StructName}}
// Delete{{.StructName}} 删除 {{.StructName}}
func Delete{{.StructName}}(delete *request.DeleteById) (err error) {
	_, err = {{.TableName}}.Delete(g.Map{"id": delete.Id})
	return err
}

// Delete{{.StructName}} batch deletion {{.StructName}}s
// Delete{{.StructName}} 批量删除 {{.StructName}}s
func Delete{{.StructName}}ByIds(deletes *request.DeleteByIds) (err error) {
	_, err = {{.TableName}}.Delete(g.Map{"id IN(?)": deletes.Ids})
	return err
}

// Update{{.StructName}} update {{.StructName}}
// Update{{.StructName}} 更新 {{.StructName}}
func Update{{.StructName}}(update *request.Update{{.StructName}}) (err error) {
	condition := g.Map{"id": update.Id}
	updateData := g.Map{
    {{- range .Fields}}
        {{- if eq .FieldName "Id" "ID" "CreateAt" "UpdateAt" "DeleteAt"}}
        {{ else }}
        "{{.ColumnName}}": update.{{.FieldName}},
        {{- end}}
    {{- end }}
	}
	_, err = {{.TableName}}.Update(updateData, condition)
	return err
}

// Find{{.StructName}} Gets a single {{.StructName}} based on id
// Find{{.StructName}} 根据id获取单条{{.StructName}}
func Find{{.StructName}}(find *request.Find{{.StructName}}) (data *{{.TableName}}.{{.StructName}}, err error) {
    data = (*{{.TableName}}.{{.StructName}})(nil)
    db := g.DB(global.Db).Table("{{.TableName}}").Safe()
    err = db.Where(g.Map{"id": find.Id}).Struct(&data)
    return
}

// Get{{.StructName}}List Page out the {{.StructName}} list
// Get{{.StructName}}List 分页获取{{.StructName}}列表
{{- range .Fields}}
	    {{- if .FieldSearchType}}
	        {{- if eq .FieldType "bool" }}
func Get{{.StructName}}List(info *request.Get{{.StructName}}List, condition g.Map) (list interface{}, total int, err error) {
            {{- end }}
        {{- end }}
{{- end }}
func Get{{.StructName}}List(info *request.Get{{.StructName}}List) (list interface{}, total int, err error) {
	datalist := ([]*{{.TableName}}.{{.StructName}})(nil)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.DB(global.Db).Table("{{.TableName}}").Safe()
	{{- range .Fields}}
    	    {{- if .FieldSearchType}}
    	        {{- if eq .FieldType "string" }}
    if pageInfo.{{.ColumnName}} != "" {
        condition["`{{.ColumnName}}` {{.FieldSearchType}} ?"] = {{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }}
    }
    	        {{- else if eq .FieldType "int" }}
    if pageInfo.{{.ColumnName}} != 0 {
        condition["`{{.ColumnName}}` {{.FieldSearchType}} ?"] = {{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }}
    }
                {{- else if eq .FieldType "float64" }}
    if info.{{.FieldName}} != 0 {
        condition["`{{.ColumnName}}` {{.FieldSearchType}} ?"] = {{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }}
    }
                {{- else if eq .FieldType "gtime.Time" }}
    if info.{{.FieldName}}.IsZero() {
        condition["`{{.ColumnName}}` {{.FieldSearchType}} ?"] = {{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }}
    }
            {{- end }}
        {{- end }}
    {{- end }}
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Structs(&datalist)
	return datalist, total, err
}
