package upload

import (
	"github.com/flipped-aurora/gf-vue-admin/server/library/global"
	"github.com/flipped-aurora/gf-vue-admin/server/library/utils"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
)

var Local = new(local)

type local struct {
	err  error
	out  *os.File
	file multipart.File
}

func (l *local) Upload(file *multipart.FileHeader) (filepath string, key string, err error) {

	ext := path.Ext(file.Filename)                                     // 读取文件后缀
	name := utils.MD5V([]byte(strings.TrimSuffix(file.Filename, ext))) // 读取文件名并加密
	filename := name + "_" + time.Now().Format("20060102150405") + ext // 拼接新文件名
	filepath = global.Config.Local.Path + "/" + filename               // 拼接路径和文件名

	if l.err = os.MkdirAll(global.Config.Local.Path, os.ModePerm); l.err != nil { // 尝试创建此路径
		g.Log().Error("function os.MkdirAll() Failed!", g.Map{"err": l.err})
		return filepath, key, err
	}

	if l.file, l.err = file.Open(); l.err != nil { // 读取文件
		g.Log().Error("function file.Open() Failed!", g.Map{"err": l.err})
		return filepath, key, err
	}

	if l.out, l.err = os.Create(filepath); l.err != nil {
		g.Log().Error("function os.Create() Failed!", g.Map{"err": l.err})
		return filepath, key, err
	}

	defer func() { // multipart.File 对象 defer 关闭
		_ = l.out.Close()
		_ = l.file.Close()
	}()

	if _, l.err = io.Copy(l.out, l.file); l.err != nil {// 传输（拷贝）文件
		g.Log().Error("function io.Copy Failed!", g.Map{"err": l.err})
		return filepath, key, err
	}

	return filepath, filename, nil
}

func (l *local) Delete(key string) error {

	filepath := global.Config.Local.Path + "/" + key

	if strings.Contains(filepath, global.Config.Local.Path) {
		if err := os.Remove(filepath); err != nil {
			g.Log().Error("本地文件删除失败!", g.Map{"err": err})
			return err
		}
	}
	return nil
}
