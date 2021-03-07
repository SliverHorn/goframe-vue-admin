package internal

import (
	"gf-vue-admin/server/library/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var Aliyun = new(aliyun)

type aliyun struct{}

func (a *aliyun) StorageClassType() oss.Option {
	switch global.Config.Aliyun.StorageClassType { // 根据配置文件进行指定存储类型
	case "Standard": // 指定存储类型为标准存储，默认也为标准存储。
		return oss.ObjectStorageClass(oss.StorageStandard)
	case "IA": // 指定存储类型为很少访问存储
		return oss.ObjectStorageClass(oss.StorageIA)
	case "Archive": // 指定存储类型为归档存储。
		return oss.ObjectStorageClass(oss.StorageArchive)
	case "ColdArchive": // 指定存储类型为归档存储。
		return oss.ObjectStorageClass(oss.StorageColdArchive)
	default:
		return oss.ObjectStorageClass(oss.StorageStandard)
	}
}
