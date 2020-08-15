package v1

import (
	"fmt"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/service"
	"server/library/global"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Create{{.StructName}} Create {{.StructName}}
// Create{{.StructName}} 创建{{.StructName}}
func Create{{.StructName}}(r *ghttp.Request) {
	var create request.Create{{.StructName}}
	if err := r.Parse(&create); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.Create{{.StructName}}(&create); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("创建失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "创建成功")
}

// Delete{{.StructName}} Delete {{.StructName}}
// Delete{{.StructName}} 删除{{.StructName}}
func Delete{{.StructName}}(r *ghttp.Request) {
	var delete request.DeleteById
	if err := r.Parse(&delete); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.Delete{{.StructName}}(&delete); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// Delete{{.StructName}}ByIds Batch delete {{.StructName}}
// Delete{{.StructName}}ByIds 批量删除{{.StructName}}
func Delete{{.StructName}}ByIds(r *ghttp.Request) {
	var deletes request.DeleteByIds
	if err := r.Parse(&deletes); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.Delete{{.StructName}}ByIds(&deletes); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("批量删除失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "批量删除成功")
}

// Update{{.StructName}} Update {{.StructName}}
// Update{{.StructName}} 更新{{.StructName}}
func Update{{.StructName}}(r *ghttp.Request) {
	var update request.Update{{.StructName}}
	if err := r.Parse(&update); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.Update{{.StructName}}(&update); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("更新失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "更新成功")
}

// Find{{.StructName}} Query {{.StructName}} with id
// Find{{.StructName}} 用id查询{{.StructName}}
func Find{{.StructName}}(r *ghttp.Request) {
	var find request.FindById
	if err := r.Parse(&find); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	data, err := service.Find{{.StructName}}(&find)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, data)
}

// Get{{.StructName}}List Page out the {{.StructName}} list
// Get{{.StructName}}List 分页获取{{.StructName}}列表
func Get{{.StructName}}List(r *ghttp.Request) {
	var pageInfo request.Get{{.StructName}}List
	if err := r.Parse(&pageInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	{{- range .Fields}}
	    {{- if .FieldSearchType}}
	        {{- if eq .FieldType "bool" }}
	condition := g.Map{}
    if pageInfo.{{.ColumnName}} == true || pageInfo.{{.ColumnName}} == false {
        condition["{{.ColumnName}}"] = utils.BoolToInt(pageInfo.{{.ColumnName}})
    }
    if r.GetString("{{.ColumnName}}") == "empty" || r.GetString("{{.ColumnName}}") == "" {
        delete(condition, "{{.ColumnName}}")
    }
	list, total, err := service.Get{{.StructName}}List(&pageInfo, condition)
            {{- end }}
        {{- end }}
    {{- end }}
	list, total, err := service.Get{{.StructName}}List(&pageInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	})
}
