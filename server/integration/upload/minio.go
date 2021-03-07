package upload

import (
	"context"
	"gf-vue-admin/server/integration/upload/internal"
	"gf-vue-admin/server/library/global"
	"mime/multipart"

	"github.com/gogf/gf/frame/g"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var Minio = new(_minio)

type _minio struct {
	err    error
	ctx    context.Context
	file   multipart.File
	info   minio.UploadInfo
	client *minio.Client
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 上传文件到Minio
func (m *_minio) Upload(file *multipart.FileHeader) (path string, key string, err error) {

	if m.err = m.client.MakeBucket(m.ctx, global.Config.Minio.Bucket, minio.MakeBucketOptions{Region: ""}); m.err != nil {
		if exists, existsErr := m.client.BucketExists(m.ctx, global.Config.Minio.Bucket); !exists && existsErr != nil { // Check to see if we already own this bucket (which happens if you run this twice)
			g.Log().Error("function client.BucketExists() Failed!", g.Map{"err": m.err})
			return path, key, m.err
		}
		g.Log().Info("We Already Own!", g.Map{"bucket": global.Config.Minio.Bucket})
	} else {
		g.Log().Info("Successfully Created!", g.Map{"bucket": global.Config.Minio.Bucket})
	}

	objectName := internal.Upload.GetObjectName(file.Filename)

	if m.file, m.err = file.Open(); m.err != nil {
		g.Log().Error("function file.Open() Failed!", g.Map{"err": m.err})
		return path, key, m.err
	}

	contentType := file.Header.Get("content-type") // 获取文件类型

	if m.info, m.err = m.client.PutObject(m.ctx, global.Config.Minio.Bucket, objectName, m.file, file.Size, minio.PutObjectOptions{ContentType: contentType}); m.err != nil {
		g.Log().Error("function client.PutObject() Failed!", g.Map{"err": m.err})
		return path, key, m.err
	}

	g.Log().Printf("Successfully uploaded!", g.Map{"ObjectName": objectName, "UploadInfo Size": m.info}, m.info.Size)

	return global.Config.Minio.Path + "/" + global.Config.Minio.Bucket + "/" + objectName, objectName, nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除文件
func (m *_minio) Delete(key string) error {
	options := minio.RemoveObjectOptions{GovernanceBypass: true}
	return m.client.RemoveObject(context.Background(), global.Config.Minio.Bucket, key, options)
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 初始化minio的Client对象
func MinioInit() (result *_minio, err error) {
	var info _minio
	info.ctx = context.Background()
	if info.client, info.err = minio.New(global.Config.Minio.Endpoint, &minio.Options{ // 初始化我的客户端对象。
		Creds:  credentials.NewStaticV4(global.Config.Minio.Id, global.Config.Minio.Secret, global.Config.Minio.Token),
		Secure: global.Config.Minio.UseSsl,
	}); info.err != nil {
		return &info, info.err
	}
	return &info, nil
}
