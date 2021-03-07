package upload

import (
	"context"
	"fmt"
	"gf-vue-admin/server/integration/upload/internal"
	"gf-vue-admin/server/library/global"
	"mime/multipart"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var Qiniu = new(_qiniu)

type _qiniu struct {
	err   error
	_file multipart.File
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 上传文件到七牛云
func (q *_qiniu) Upload(file *multipart.FileHeader) (path string, key string, err error) {

	putPolicy := storage.PutPolicy{Scope: global.Config.Qiniu.Bucket}
	mac := qbox.NewMac(global.Config.Qiniu.AccessKey, global.Config.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	formUploader := storage.NewFormUploader(internal.Qiniu.Config())
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	if q._file, q.err = file.Open(); q.err != nil {
		g.Log().Error("function file.Open() Failed!", g.Map{"err": q.err})
		return path, key, q.err
	}

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性

	if q.err = formUploader.Put(context.Background(), &ret, upToken, fileKey, q._file, file.Size, &putExtra); q.err != nil {
		g.Log().Error("function formUploader.Put() Failed!", g.Map{"err": q.err})
		return path, key, err
	}

	return global.Config.Qiniu.ImgPath + "/" + ret.Key, ret.Key, nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据key删除七牛云文件
func (q *_qiniu) Delete(key string) error {
	mac := qbox.NewMac(global.Config.Qiniu.AccessKey, global.Config.Qiniu.SecretKey)
	bucketManager := storage.NewBucketManager(mac, internal.Qiniu.Config())
	return bucketManager.Delete(global.Config.Qiniu.AccessKey, key)
}
