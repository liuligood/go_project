package upload

import (
	"context"
	"crmeb_go/utils/izap"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
)

type Qiniu struct {
	baseOss BaseOss
}

func (q *Qiniu) UploadFile(file *multipart.FileHeader) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: q.baseOss.Conf.Qiniu.Bucket}
	mac := qbox.NewMac(q.baseOss.Conf.Qiniu.AccessKey, q.baseOss.Conf.Qiniu.SecretKey)

	upToken := putPolicy.UploadToken(mac)
	cfg := q.qiniuConfig()

	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	f, openError := file.Open()
	if openError != nil {
		izap.Log.Error("function file.Open() failed", zap.Any("err", openError.Error()))

		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}

	defer f.Close()                                                  // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)

	if putErr != nil {
		izap.Log.Error("function formUploader.Put() failed", zap.Any("err", putErr.Error()))
		return "", "", errors.New("function formUploader.Put() failed, err:" + putErr.Error())
	}

	return q.baseOss.Conf.Qiniu.ImgPath + "/" + ret.Key, ret.Key, nil
}

func (q *Qiniu) DeleteFile(key string) error {
	mac := qbox.NewMac(q.baseOss.Conf.Qiniu.AccessKey, q.baseOss.Conf.Qiniu.SecretKey)
	cfg := q.qiniuConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)

	if err := bucketManager.Delete(q.baseOss.Conf.Qiniu.Bucket, key); err != nil {
		izap.Log.Error("function bucketManager.Delete() failed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}
	return nil
}

func (q *Qiniu) qiniuConfig() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      q.baseOss.Conf.Qiniu.UseHTTPS,
		UseCdnDomains: q.baseOss.Conf.Qiniu.UseCdnDomains,
	}
	switch q.baseOss.Conf.Qiniu.Zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
