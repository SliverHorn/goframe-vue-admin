package service

import (
	"errors"
	"fmt"
	model "gf-vue-admin/server/app/model/system"
	"gf-vue-admin/server/library/constant"
	"gf-vue-admin/server/library/global"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"reflect"
	"strconv"
)

var Excel = new(_excel)

type _excel struct {
	row  []string
	err  error
	file *excelize.File
	rows *excelize.Rows
}

var fixedHeader = []string{"ID", "路由Name", "路由Path", "是否隐藏", "父节点", "排序", "文件名称"}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 解析数据生成excel文件
func (e *_excel) Parse(infoList []model.Menu, filePath string) error {
	excel := excelize.NewFile()
	if err := excel.SetSheetRow("Sheet1", "A1", &[]string{"ID", "路由Name", "路由Path", "是否隐藏", "父节点", "排序", "文件名称"}); err != nil {
		return err
	}
	for i, menu := range infoList {
		axis := fmt.Sprintf("A%d", i+2)
		if err := excel.SetSheetRow("Sheet1", axis, &[]interface{}{menu.ID, menu.Name, menu.Path, menu.Hidden, menu.ParentId, menu.Sort, menu.Component}); err != nil {
			return err
		}
	}
	return excel.SaveAs(filePath)
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 解析数据生成excel文件
func (e *_excel) Parse2Data() (result []model.Menu, err error) {
	menus := make([]model.Menu, 0)
	skipHeader := true
	if e.file, e.err = excelize.OpenFile(constant.ExcelDir + "ExcelImport.xlsx"); e.err != nil {
		return nil, e.err
	}
	if e.rows, e.err = e.file.Rows("Sheet1"); e.err != nil {
		return nil, e.err
	}
	for e.rows.Next() {
		if e.row, e.err = e.rows.Columns(); e.err != nil {
			return nil, e.err
		}
		if skipHeader {
			if reflect.DeepEqual(e.row, fixedHeader) {
				skipHeader = false
				continue
			} else {
				return nil, errors.New("Excel格式错误! ")
			}
		}
		if len(e.row) != len(fixedHeader) {
			continue
		}
		id, _ := strconv.Atoi(e.row[0])
		hidden, _ := strconv.ParseBool(e.row[3])
		sort, _ := strconv.Atoi(e.row[5])
		menu := model.Menu{Model: global.Model{ID: uint(id)}, Name: e.row[1], Path: e.row[2], Hidden: hidden, ParentId: e.row[4], Sort: sort, Component: e.row[6]}
		menus = append(menus, menu)
	}
	return menus, nil
}
