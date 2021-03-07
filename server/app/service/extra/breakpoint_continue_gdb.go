package service

import (
	"database/sql"
	"errors"
	"gf-vue-admin/server/app/api/request"
	model "gf-vue-admin/server/app/model/extra"
	"gf-vue-admin/server/library/utils"
	"github.com/gogf/gf/frame/g"
	"io/ioutil"
	"mime/multipart"
	"strconv"
)

var BreakpointContinueGdb = new(breakpointContinueGdb)

type breakpointContinueGdb struct {
	err     error
	path    string
	file    multipart.File
	_entity model.BreakpointContinue
	entity  *model.BreakpointContinue
	_chunk  model.BreakpointContinueChunk
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 上传文件时检测当前文件属性，如果没有文件则创建，有则返回文件的当前切片
func (b *breakpointContinueGdb) FindOrCreateFile(info *request.BreakpointContinue) (result *model.BreakpointContinue, err error) {
	var entity model.BreakpointContinue
	create := info.Create()
	if errors.Is(g.DB().Table(b._entity.TableName()).Where(g.Map{"file_md5": info.FileMd5, "is_finish": true}).Struct(&entity), sql.ErrNoRows) {
		_, err = g.DB().Table(b._entity.TableName()).Insert(create)
		err = g.DB().Table(b._chunk.TableName()).Where(g.Map{"file_id": entity.ID}).Structs(create.FileChunk)
		return &entity, err
	}
	create.IsFinish = true
	create.FilePath = entity.FilePath
	_, err = g.DB().Table(b._entity.TableName()).Insert(info.Create())
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建文件切片记录
func (b *breakpointContinueGdb) CreateFileChunk(info *request.CreateFileChunk) error {
	entity := info.Create()
	_, err := g.DB().Table(entity.TableName()).Insert(entity)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除文件切片记录
func (b *breakpointContinueGdb) DeleteFileChunk(info *request.BreakpointContinue) error {
	var entity model.BreakpointContinue
	err := g.DB().Table(b._entity.TableName()).Where(g.Map{"file_md5": info.FileMd5, "file_name": info.FileName}).Struct(&entity)
	entity.IsFinish = true
	entity.FilePath = info.FilePath
	_, err = g.DB().Table(b._entity.TableName()).Where(g.Map{"id": entity.ID}).Update(&entity)
	_, err = g.DB().Table(b._chunk.TableName()).Unscoped().Delete(g.Map{"file_id": entity.ID})
	err = utils.File.RemoveChunk(info.FileMd5)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 断点续传到服务器
func (b *breakpointContinueGdb) BreakpointContinue(info *request.BreakpointContinue, header *multipart.FileHeader) error {
	if b.file, b.err = header.Open(); b.err != nil {
		return errors.New("文件读取失败! ")
	}
	content, _ := ioutil.ReadAll(b.file)
	if !utils.File.CheckMd5(content, info.ChunkMd5) {
		return errors.New("检查md5失败! ")
	}
	b.entity, b.err = b.FindOrCreateFile(info)
	chunkNumber, _ := strconv.Atoi(info.ChunkNumber)
	chunkTotal, _ := strconv.Atoi(info.ChunkTotal)
	if b.path, b.err = utils.File.BreakPointContinue(content, info.FileName, chunkNumber, chunkTotal, info.FileMd5); b.err != nil {
		return errors.New("断点续传失败! ")
	}
	create := request.CreateFileChunk{
		GetById:                     request.GetById{Id: b.entity.ID},
		BaseBreakpointContinueChunk: request.BaseBreakpointContinueChunk{FileChunkPath: b.path, FileChunkNumber: chunkNumber},
	}
	if b.err = b.CreateFileChunk(&create); b.err != nil {
		return errors.New("创建文件记录失败! ")
	}
	return nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 上传文件完成
func (b *breakpointContinueGdb) BreakpointContinueFinish(info *request.BreakpointContinueFinish) (filepath string, err error) {
	return utils.File.MakeFile(info.FileName, info.FileMd5)
}
