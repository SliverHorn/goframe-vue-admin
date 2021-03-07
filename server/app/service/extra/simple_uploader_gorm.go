package service

import (
	"errors"
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/request"
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/extra"
	"github.com/flipped-aurora/gf-vue-admin/server/library/global"
	"github.com/flipped-aurora/gf-vue-admin/server/library/utils"
	"github.com/gogf/gf/frame/g"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

var SimpleUploaderGorm = new(uploaderGorm)

type uploaderGorm struct {
	err       error
	old       multipart.File
	out       *os.File
	file      *os.File
	fileInfos []os.FileInfo

	chunkDir  string
	checkPath string
	chunkPath string
	finishDir string
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 上传保存切片文件
func (u *uploaderGorm) Upload(header *multipart.FileHeader, info *request.Upload) error {
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
	return global.Db.Create(&entity).Error
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 保存文件切片路径
func (u *uploaderGorm) CreateChunk(info model.SimpleUploader) error {
	return global.Db.Create(&info).Error
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 检查文件是否已经上传过
func (u *uploaderGorm) CheckFileMd5(info *request.CheckFileMd5) (uploads *[]model.SimpleUploader, isDone bool, err error) {
	var entities []model.SimpleUploader
	err = global.Db.Find(&entities, "identifier = ? AND is_done = ?", info.Md5, false).Error
	isDone = errors.Is(global.Db.First(&model.SimpleUploader{}, "identifier = ? AND is_done = ?", info.Md5, true).Error, gorm.ErrRecordNotFound)
	return &entities, !isDone, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 合并文件
func (u *uploaderGorm) MergeFileMd5(info *request.MergeFileMd5) error {
	u.finishDir = global.Config.Uploader.FinishDir
	u.checkPath = global.Config.Uploader.GetCheckPath(info.Md5)

	if !errors.Is(global.Db.First(&model.SimpleUploader{}, "identifier = ? AND is_done = ?", info.Md5, true).Error, gorm.ErrRecordNotFound) { //如果文件上传成功 不做后续操作 通知成功即可
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

	return global.Db.Transaction(func(tx *gorm.DB) error {
		if u.err = tx.Delete(&model.SimpleUploader{}, "identifier = ? AND is_done = ?", info.Md5, false).Error; u.err != nil { // 删除切片信息
			return u.err
		}
		entity := model.SimpleUploader{IsDone: true, FilePath: u.finishDir + info.Filename, Filename: info.Filename, Identifier: info.Md5}
		if u.err = tx.Create(&entity).Error; u.err != nil { // 添加文件信息
			return u.err
		}
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 保存文件
func (u *uploaderGorm) SaveUploadedFile(file *multipart.FileHeader, chunkPath string) error {
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
