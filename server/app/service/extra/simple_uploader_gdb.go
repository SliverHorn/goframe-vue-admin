package service

import (
	"database/sql"
	"errors"
	"gf-vue-admin/server/app/api/request"
	model "gf-vue-admin/server/app/model/extra"
	"gf-vue-admin/server/library/global"
	"gf-vue-admin/server/library/utils"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

var SimpleUploaderGdb = new(uploaderGdb)

type uploaderGdb struct {
	err       error
	old       multipart.File
	out       *os.File
	file      *os.File
	_entity   model.SimpleUploader
	fileInfos []os.FileInfo

	chunkDir  string
	checkPath string
	chunkPath string
	finishDir string
}

func (u *uploaderGdb) Upload(header *multipart.FileHeader, info *request.Upload) error {
	u.chunkDir = global.Config.Uploader.GetIdentifier(info.Identifier)
	if ok, _ := utils.Directory.PathExists(u.chunkDir); !ok {
		if u.err = utils.Directory.BatchCreate(u.chunkDir); u.err != nil {
			g.Log().Error(`创建目录失败!`, g.Map{"err": u.err})
		}
	}
	u.chunkPath = u.chunkDir + info.Filename + info.ChunkNumber
	if u.err = u.SaveUploadedFile(header, u.chunkPath); u.err != nil {
		return errors.New("切片创建失败! ")
	}
	entity := info.Create(u.chunkPath)
	_, err := g.DB().Table(u._entity.TableName()).Insert(&entity)
	return err
}

func (u *uploaderGdb) CreateChunk(info model.SimpleUploader) error {
	_, err := g.DB().Table(info.TableName()).Insert(&info)
	return err
}

func (u *uploaderGdb) CheckFileMd5(info *request.CheckFileMd5) (uploads *[]model.SimpleUploader, isDone bool, err error) {
	var entities []model.SimpleUploader
	err = g.DB().Table(u._entity.TableName()).Where(g.Map{"identifier": info.Md5, "is_done": true}).Structs(entities)
	isDone = errors.Is(g.DB().Table(u._entity.TableName()).Where(g.Map{"identifier": info.Md5, "is_done": true}).Struct(&model.SimpleUploader{}), sql.ErrNoRows)
	return &entities, !isDone, err
}

func (u *uploaderGdb) MergeFileMd5(info *request.MergeFileMd5) error {
	u.finishDir = global.Config.Uploader.FinishDir
	u.checkPath = global.Config.Uploader.GetCheckPath(info.Md5)

	if !errors.Is(g.DB().Table(u._entity.TableName()).Where(g.Map{"identifier": info.Md5, "is_done": true}).Struct(&model.SimpleUploader{}), sql.ErrNoRows) { //如果文件上传成功 不做后续操作 通知成功即可
		return nil
	}

	if u.fileInfos, u.err = ioutil.ReadDir(u.checkPath); u.err != nil { // 打开切片文件夹
		return u.err
	}
	_ = os.MkdirAll(u.finishDir, os.ModePerm)

	if u.file, u.err = os.OpenFile(u.finishDir+info.Filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); u.err != nil { // 创建目标文件
		return u.err
	}

	defer func() { //关闭文件
		_ = u.file.Close()
	}()

	for k := range u.fileInfos { // 将切片文件按照顺序写入
		content, _ := ioutil.ReadFile(u.checkPath + "/" + info.Filename + strconv.Itoa(k+1))
		if _, u.err = u.file.Write(content); u.err != nil {
			_ = os.Remove(u.finishDir + info.Filename)
		}
	}

	defer func() {
		_ = os.RemoveAll(u.checkPath) //清除切片
	}()

	return g.DB().Transaction(func(tx *gdb.TX) error {
		if _, err := tx.Table(u._entity.TableName()).Delete(g.Map{"identifier": info.Md5, "is_done": false}); err != nil {
			return err
		}
		entity := model.SimpleUploader{IsDone: true, FilePath: u.finishDir + info.Filename, Filename: info.Filename, Identifier: info.Md5}

		_, err := tx.Table(u._entity.TableName()).Insert(&entity)
		return err
	})
}

func (u *uploaderGdb) SaveUploadedFile(file *multipart.FileHeader, chunkPath string) error {
	if u.old, u.err = file.Open(); u.err != nil { // 读取文件
		return errors.New("function file.Open() Filed! err:" + u.err.Error())
	}
	defer func() { // 打开文件 defer 关闭
		_ = u.old.Close()
	}()

	if u.out, u.err = os.Create(chunkPath); u.err != nil {
		return errors.New("function os.Create() Filed! err:" + u.err.Error())
	}
	defer func() { // copy文件 defer 关闭
		_ = u.out.Close()
	}()

	if _, u.err = io.Copy(u.out, u.old); u.err != nil { // 传输（拷贝）文件
		return errors.New("function io.Copy() Filed! err:" + u.err.Error())
	}
	return nil
}
