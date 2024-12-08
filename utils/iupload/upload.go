package iupload

import (
	"crmeb_go/config"
	"crmeb_go/utils/izap"
	"mime/multipart"
)

// OSS 对象存储接口
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

type BaseOss struct {
	Conf config.Conf
}

// NewOss OSS的实例化方法
func NewOss(conf config.Conf) OSS {
	switch conf.System.OssType {
	case "local":
		return &Local{BaseOss{Conf: conf}}
	case "qiniu":
		return &Qiniu{BaseOss{Conf: conf}}
	case "tencent-cos":
		return &TencentCOS{BaseOss{Conf: conf}}
	case "aliyun-oss":
		return &AliyunOSS{BaseOss{Conf: conf}}
	case "huawei-obs":
		return &Obs{BaseOss{Conf: conf}}
	case "aws-s3":
		return &AwsS3{BaseOss{Conf: conf}}
	case "cloudflare-r2":
		return &CloudflareR2{BaseOss{Conf: conf}}
	case "minio":
		minioClient, err := GetMinio(conf)
		if err != nil {
			izap.Log.Warn("你配置了使用minio，但是初始化失败，请检查minio可用性或安全配置: " + err.Error())
			panic("minio初始化失败") // 建议这样做，用户自己配置了minio，如果报错了还要把服务开起来，使用起来也很危险
		}
		return minioClient
	default:
		return &Local{}
	}
}
