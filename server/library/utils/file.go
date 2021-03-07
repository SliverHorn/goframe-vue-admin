package utils

import (
	"gf-vue-admin/server/library/constant"
	"io/ioutil"
	"os"
	"strconv"
)

var File = new(file)

type file struct {
	file     *os.File
	fileInfo []os.FileInfo
}

// 前端传来文件片与当前片为什么文件的第几片
// 后端拿到以后比较次分片是否上传 或者是否为不完全片
// 前端发送每片多大
// 前端告知是否为最后一片且是否完成

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 断点续传
func (f *file) BreakPointContinue(content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string) (path string, err error) {
	path = constant.BreakpointDir + fileMd5 + "/"
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		return path, err
	}
	return f.makeFileContent(content, fileName, path, contentNumber)
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 检查Md5
func (f *file) CheckMd5(content []byte, chunkMd5 string) bool {
	fileMd5 := MD5V(content)
	if fileMd5 == chunkMd5 { // 可以继续上传
		return true
	} else { // 切片不完整，废弃
		return false
	}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建切片内容
func (f *file) makeFileContent(content []byte, fileName string, fileDir string, contentNumber int) (path string, err error) {
	path = fileDir + fileName + "_" + strconv.Itoa(contentNumber)
	if f.file, err = os.Create(path); err != nil {
		return path, err
	} else {
		if _, err = f.file.Write(content); err != nil {
			return path, err
		}
	}
	defer func() {
		_ = f.file.Close()
	}()
	return path, nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建切片文件
func (f *file) MakeFile(fileName string, FileMd5 string) (path string, err error) {
	path = constant.FinishDir + fileName
	if f.fileInfo, err = ioutil.ReadDir(constant.BreakpointDir + FileMd5); err != nil {
		return path, err
	}
	_ = os.MkdirAll(constant.FinishDir, os.ModePerm)
	if f.file, err = os.OpenFile(constant.FinishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); err != nil {
		return path, err
	}
	defer func() {
		_ = f.file.Close()
	}()
	for k := range f.fileInfo {
		content, _ := ioutil.ReadFile(constant.BreakpointDir + FileMd5 + "/" + fileName + "_" + strconv.Itoa(k))
		if _, err = f.file.Write(content); err != nil {
			_ = os.Remove(constant.FinishDir + fileName)
			return path, err
		}
	}
	return path, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 移除切片
func (f *file) RemoveChunk(FileMd5 string) error {
	return os.RemoveAll(constant.BreakpointDir + FileMd5)
}
