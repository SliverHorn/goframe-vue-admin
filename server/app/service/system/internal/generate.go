package internal

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/constant"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var Generate = new(generate)

type generate struct {
	err    error
	revoke bool
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取全部tpl文件
func (g *generate) GetAllTemplateFile(path string, fileList []string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			if fileList, err = g.GetAllTemplateFile(path+"/"+file.Name(), fileList); err != nil {
				return nil, err
			}
		} else {
			if strings.HasSuffix(file.Name(), ".tpl") {
				fileList = append(fileList, path+"/"+file.Name())
			}
		}
	}
	return fileList, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取需要文件内容
func (g *generate) GetNeedList(autoCode *model.AutoCode) (dataList []model.TemplateData, fileList []string, needMkdir []string, err error) {
	var tplFileList []string
	// 获取 basePath 文件夹下所有tpl文件
	if tplFileList, err = g.GetAllTemplateFile(constant.BasePath, nil); err != nil {
		return dataList, fileList, needMkdir, err
	}
	dataList = make([]model.TemplateData, 0, len(tplFileList))
	fileList = make([]string, 0, len(tplFileList))
	needMkdir = make([]string, 0, len(tplFileList)) // 当文件夹下存在多个tpl文件时，改为map更合理

	for _, value := range tplFileList { // 根据文件路径生成 tplData 结构体，待填充数据
		dataList = append(dataList, model.TemplateData{LocationPath: value})
	}

	for index, value := range dataList { // 生成 *Template, 填充 template 字段
		dataList[index].Template, err = template.ParseFiles(value.LocationPath)
		if err != nil {
			return dataList, fileList, needMkdir, err
		}
	}
	// 生成文件路径，填充 AutoCodePath 字段，readme.txt.tpl不符合规则，需要特殊处理
	// resource/template/web/api.js.tpl -> autoCode/web/autoCode.PackageName/api/autoCode.PackageName.js
	// resource/template/readme.txt.tpl -> autoCode/readme.txt
	for index, value := range dataList {
		trimBase := strings.TrimPrefix(value.LocationPath, constant.BasePath+"/")
		if trimBase == "readme.txt.tpl" {
			dataList[index].AutoCodePath = constant.AutoPath + "readme.txt"
			continue
		}

		if lastSeparator := strings.LastIndex(trimBase, "/"); lastSeparator != -1 {
			origFileName := strings.TrimSuffix(trimBase[lastSeparator+1:], ".tpl")
			firstDot := strings.Index(origFileName, ".")
			if firstDot != -1 {
				dataList[index].AutoCodePath = filepath.Join(constant.AutoPath, trimBase[:lastSeparator], autoCode.PackageName,
					origFileName[:firstDot], autoCode.PackageName+origFileName[firstDot:])
			}
		}

		if lastSeparator := strings.LastIndex(dataList[index].AutoCodePath, string(os.PathSeparator)); lastSeparator != -1 {
			needMkdir = append(needMkdir, dataList[index].AutoCodePath[:lastSeparator])
		}
	}
	for _, value := range dataList {
		fileList = append(fileList, value.AutoCodePath)
	}
	return dataList, fileList, needMkdir, err
}

func (g *generate) AddAutoMoveFile(data *model.TemplateData) {
	dir := filepath.Base(filepath.Dir(data.AutoCodePath))
	base := filepath.Base(data.AutoCodePath)
	fileSlice := strings.Split(data.AutoCodePath, string(os.PathSeparator))
	n := len(fileSlice)
	if n <= 2 {
		return
	}
	if strings.Contains(fileSlice[1], "server") {
		if strings.Contains(fileSlice[n-2], "router") {
			data.AutoMoveFilePath = filepath.Join(dir, base)
		} else if strings.Contains(fileSlice[n-2], "api") {
			data.AutoMoveFilePath = filepath.Join(dir, "v1", base)
		} else if strings.Contains(fileSlice[n-2], "service") {
			data.AutoMoveFilePath = filepath.Join(dir, base)
		} else if strings.Contains(fileSlice[n-2], "model") {
			data.AutoMoveFilePath = filepath.Join(dir, base)
		} else if strings.Contains(fileSlice[n-2], "request") {
			data.AutoMoveFilePath = filepath.Join("model", dir, base)
		}
	} else if strings.Contains(fileSlice[1], "web") {
		if strings.Contains(fileSlice[n-1], "js") {
			data.AutoMoveFilePath = filepath.Join("../", "web", "src", dir, base)
		} else if strings.Contains(fileSlice[n-2], "workflowForm") {
			data.AutoMoveFilePath = filepath.Join("../", "web", "src", "view", filepath.Base(filepath.Dir(filepath.Dir(data.AutoCodePath))), strings.TrimSuffix(base, filepath.Ext(base))+"WorkflowForm.vue")
		} else if strings.Contains(fileSlice[n-2], "form") {
			data.AutoMoveFilePath = filepath.Join("../", "web", "src", "view", filepath.Base(filepath.Dir(filepath.Dir(data.AutoCodePath))), strings.TrimSuffix(base, filepath.Ext(base))+"Form.vue")
		} else if strings.Contains(fileSlice[n-2], "table") {
			data.AutoMoveFilePath = filepath.Join("../", "web", "src", "view", filepath.Base(filepath.Dir(filepath.Dir(data.AutoCodePath))), base)
		}
	}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 文件移动
func (g *generate) FileMove(src string, dst string) error {

	if dst == "" {
		return g.err
	}

	if src, g.err = filepath.Abs(src); g.err != nil {
		return g.err
	}

	if dst, g.err = filepath.Abs(dst); g.err != nil {
		return g.err
	}

	dir := filepath.Dir(dst)

redirect:
	if _, g.err = os.Stat(dir); g.err != nil {
		if g.err = os.MkdirAll(dir, 0755); g.err != nil {
			return g.err
		}
		if !g.revoke {
			g.revoke = true
			goto redirect
		}
	}

	return os.Rename(src, dst)
}
