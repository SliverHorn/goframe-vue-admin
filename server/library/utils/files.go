package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/gogf/gf/frame/g"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

// 接收两个参数 一个文件流 一个 bucket 你的七牛云标准空间的名字
func Upload(file *multipart.FileHeader) (path string, key string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: g.Cfg().GetString("qiniu.Bucket"),
	}
	mac := qbox.NewMac(g.Cfg().GetString("qiniu.AccessKey"), g.Cfg().GetString("qiniu.SecretKey"))
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	f, e := file.Open()
	if e != nil {
		fmt.Println(e)
		return "", "", e
	}
	dataLen := file.Size
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	err = formUploader.Put(context.Background(), &ret, upToken, fileKey, f, dataLen, &putExtra)
	if err != nil {
		g.Log().Error("upload file fail:", err)
		return "", "", err
	}
	return g.Cfg().GetString("qiniu.ImgPath") + "/" + ret.Key, ret.Key, err
}

func DeleteFile(key string) error {
	mac := qbox.NewMac(g.Cfg().GetString("qiniu.AccessKey"), g.Cfg().GetString("qiniu.SecretKey"))
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	// cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err := bucketManager.Delete(g.Cfg().GetString("qiniu.AccessKey"), key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// 前端传来文件片与当前片为什么文件的第几片
// 后端拿到以后比较次分片是否上传 或者是否为不完全片
// 前端发送每片多大
// 前端告知是否为最后一片且是否完成

const breakpointDir = "./breakpointDir/"
const finishDir = "./fileDir/"

func BreakPointContinue(content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string) (string, error) {
	path := breakpointDir + fileMd5 + "/"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return path, err
	}
	return makeFileContent(content, fileName, path, contentNumber)
}

func CheckMd5(content []byte, chunkMd5 string) (CanUpload bool) {
	fileMd5 := MD5V(content)
	if fileMd5 == chunkMd5 {
		return true // "可以继续上传"
	} else {
		return false // "切片不完整，废弃"
	}
}

func makeFileContent(content []byte, fileName string, FileDir string, contentNumber int) (string, error) {
	path := FileDir + fileName + "_" + strconv.Itoa(contentNumber)
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return path, err
	} else {
		_, err = f.Write(content)
		if err != nil {
			return path, err
		}
	}
	return path, err
}

func MakeFile(fileName string, FileMd5 string) (string, error) {
	rd, err := ioutil.ReadDir(breakpointDir + FileMd5)
	if err != nil {
		return finishDir + fileName, err
	}
	_ = os.MkdirAll(finishDir, os.ModePerm)
	fd, err := os.OpenFile(finishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return finishDir + fileName, err
	}
	defer fd.Close()
	for k := range rd {
		content, _ := ioutil.ReadFile(breakpointDir + FileMd5 + "/" + fileName + "_" + strconv.Itoa(k))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(finishDir + fileName)
			return finishDir + fileName, err
		}
	}
	return finishDir + fileName, nil
}

func RemoveChunk(FileMd5 string) error {
	err := os.RemoveAll(breakpointDir + FileMd5)
	return err
}
