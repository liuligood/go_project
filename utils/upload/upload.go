package upload

import (
	"crmeb_go/config"
	"go.uber.org/zap"
	"mime/multipart"
)

// OSS 对象存储接口
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

type BaseOss struct {
	Conf   config.Conf
	Logger *zap.Logger
}

// NewOss OSS的实例化方法
func NewOss(conf config.Conf, logger *zap.Logger) OSS {
	switch conf.System.OssType {
	case "local":
		return &Local{BaseOss{Conf: conf, Logger: logger}}
	case "qiniu":
		return &Qiniu{BaseOss{Conf: conf, Logger: logger}}
	case "tencent-cos":
		return &TencentCOS{BaseOss{Conf: conf, Logger: logger}}
	case "aliyun-oss":
		return &AliyunOSS{BaseOss{Conf: conf, Logger: logger}}
	case "huawei-obs":
		return &Obs{BaseOss{Conf: conf, Logger: logger}}
	case "aws-s3":
		return &AwsS3{BaseOss{Conf: conf, Logger: logger}}
	case "cloudflare-r2":
		return &CloudflareR2{BaseOss{Conf: conf, Logger: logger}}
	case "minio":
		minioClient, err := GetMinio(conf, logger)
		if err != nil {
			logger.Warn("你配置了使用minio，但是初始化失败，请检查minio可用性或安全配置: " + err.Error())
			panic("minio初始化失败") // 建议这样做，用户自己配置了minio，如果报错了还要把服务开起来，使用起来也很危险
		}
		return minioClient
	default:
		return &Local{}
	}
}
