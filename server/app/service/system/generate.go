package service

import (
	"database/sql"
	"errors"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/response"
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/server/app/service/system/internal"
	"github.com/flipped-aurora/gf-vue-admin/server/library/constant"
	"github.com/flipped-aurora/gf-vue-admin/server/library/utils"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var Generate = new(generate)

type generate struct {
	err  error
	file *os.File
	data []model.TemplateData // 模板数据切片

	content   []byte   // 生成数据
	files     []string // 全部模板文件名
	directory []string // 需要创建的文件夹/目录
	extension string   // 文件后缀名
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取所有的数据库名
func (s *generate) GetDbs() (result []*response.Dbs, err error) {
	err = g.DB().GetStructs(&result, "SELECT SCHEMA_NAME AS `database` FROM INFORMATION_SCHEMA.SCHEMATA;")
	return result, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取数据库的所有表名
func (s *generate) GetTables(db string) (result []response.Tables, err error) {
	err = g.DB().GetStructs(&result, "select table_name as table_name from information_schema.tables where table_schema = ?", db)
	return result, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取指定数据库与指定表名的表字段
func (s *generate) GetColumns(db string, table string) (result []*response.Columns, err error) {
	err = g.DB().GetStructs(&result, "SELECT COLUMN_NAME column_name,DATA_TYPE data_type,CASE DATA_TYPE WHEN 'longtext' THEN c.CHARACTER_MAXIMUM_LENGTH WHEN 'varchar' THEN c.CHARACTER_MAXIMUM_LENGTH WHEN 'double' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) WHEN 'decimal' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) WHEN 'int' THEN c.NUMERIC_PRECISION WHEN 'bigint' THEN c.NUMERIC_PRECISION ELSE '' END AS data_type_long,COLUMN_COMMENT column_comment FROM INFORMATION_SCHEMA.COLUMNS c WHERE table_name = ? AND table_schema = ?", table, db)
	return result, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建代码
func (s *generate) Create(info *model.AutoCode) error {

	if s.data, s.files, s.directory, s.err = internal.Generate.GetNeedList(info); s.err != nil { // 获取需要文件内容
		return s.err
	}

	if s.err = utils.Directory.BatchCreate(s.directory...); s.err != nil { // 写入文件前，先创建文件夹
		return s.err
	}

	// 生成文件
	for _, value := range s.data {
		if s.file, s.err = os.OpenFile(value.AutoCodePath, os.O_CREATE|os.O_WRONLY, 0755); s.err != nil {
			return s.err
		}
		if s.err = value.Template.Execute(s.file, info); s.err != nil {
			return s.err
		}
		_ = s.file.Close()
	}

	defer func() { // 移除中间文件
		_ = os.RemoveAll(constant.AutoPath)
	}()

	if info.AutoMoveFile { // 判断是否需要自动转移

		for index := range s.data {
			internal.Generate.AddAutoMoveFile(&s.data[index])
		}

		for _, value := range s.data { // 移动文件
			if err := internal.Generate.FileMove(value.AutoCodePath, value.AutoMoveFilePath); err != nil {
				return err
			}
		}

		return model.ErrorAutoMove
	} else { // 打包

		if err := utils.ZipFiles("./generate.zip", s.files, ".", "."); err != nil {
			return err
		}

		return nil
	}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 预览创建代码
func (s *generate) Preview(info *model.AutoCode) (result map[string]string, err error) {

	result = make(map[string]string, len(s.data)) // 初始化map

	if s.data, _, s.directory, s.err = internal.Generate.GetNeedList(info); s.err != nil { // 获取需要文件内容
		return result, s.err
	}

	if s.err = utils.Directory.BatchCreate(s.directory...); s.err != nil { // 写入文件前，先创建文件夹
		return result, s.err
	}

	for _, value := range s.data {
		if s.extension = filepath.Ext(value.AutoCodePath); s.extension == ".txt" {
			continue
		}

		if s.file, s.err = os.OpenFile(value.AutoCodePath, os.O_CREATE|os.O_WRONLY, 0755); s.err != nil {
			return result, s.err
		}

		if s.err = value.Template.Execute(s.file, info); s.err != nil {
			return result, s.err
		}

		_ = s.file.Close()

		if s.file, s.err = os.OpenFile(value.AutoCodePath, os.O_CREATE|os.O_RDONLY, 0755); s.err != nil {
			return result, s.err
		}

		builder := strings.Builder{}
		builder.WriteString("```")

		if s.extension != "" && strings.Contains(s.extension, ".") {
			builder.WriteString(strings.Replace(s.extension, ".", "", -1))
		}

		builder.WriteString("\n\n")

		if s.content, s.err = ioutil.ReadAll(s.file); s.err != nil {
			return result, s.err
		}

		builder.Write(s.content)
		builder.WriteString("\n\n```")
		pathArr := strings.Split(value.AutoCodePath, string(os.PathSeparator))
		result[pathArr[1]+"-"+pathArr[3]] = builder.String()
		_ = s.file.Close()
	}

	defer func() { // 移除中间文件
		_ = os.RemoveAll(constant.AutoPath)
	}()

	return result, nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 自动创建api数据
func (s *generate) AutoCreateApis(info *model.AutoCode) error {
	apis := []model.Api{
		{Path: "/" + info.Abbreviation + "/" + "create", Method: "POST", ApiGroup: info.Abbreviation, Description: "新增" + info.Description},
		{Path: "/" + info.Abbreviation + "/" + "first", Method: "GET", ApiGroup: info.Abbreviation, Description: "根据ID获取" + info.Description},
		{Path: "/" + info.Abbreviation + "/" + "update", Method: "PUT", ApiGroup: info.Abbreviation, Description: "更新" + info.Description},
		{Path: "/" + info.Abbreviation + "/" + "delete", Method: "DELETE", ApiGroup: info.Abbreviation, Description: "删除" + info.Description},
		{Path: "/" + info.Abbreviation + "/" + "deletes", Method: "DELETE", ApiGroup: info.Abbreviation, Description: "批量删除" + info.Description},
		{Path: "/" + info.Abbreviation + "/" + "getList", Method: "GET", ApiGroup: info.Abbreviation, Description: "获取" + info.Description + "列表"},
	}
	return g.DB().Transaction(func(tx *gdb.TX) error {
		for _, v := range apis {
			var entity model.Api
			if errors.Is(tx.Table(entity.TableName()).Where(g.Map{"path": v.Path, "method": v.Method}).Struct(&entity), sql.ErrNoRows) {
				if _, err := tx.Table(entity.TableName()).Insert(&v); err != nil { // 遇到错误时回滚事务
					return err
				}
			}
		}
		return nil
	})
}
