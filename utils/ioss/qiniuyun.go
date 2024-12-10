package oss

import (
	"crmeb_go/config"
	"crmeb_go/utils/izap"
	"errors"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
	"io"
)

type QiniuyunOSSClient struct {
	bucket *storage.BucketManager
}

type QiniuyunOSSFactory struct{}

func (f *QiniuyunOSSFactory) Create(c config.Conf) (Client, error) {
	mac := qbox.NewMac(c.QiniuyunOSS.AccessKey, c.QiniuyunOSS.SecretKey)

	// 配置七牛云存储客户端
	zone := &storage.ZoneHuadong
	cfg := storage.Config{
		// 根据存储空间所在区域选择对应的 Zone
		Zone:          zone,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}

	// 创建 BucketManager 用于文件管理操作
	bucketManager := storage.NewBucketManager(mac, &cfg)
	return &QiniuyunOSSClient{bucket: bucketManager}, nil
}

func (c *QiniuyunOSSClient) UploadFile(objectName string, filePath string) error {
	izap.Log.Info("UploadFile", zap.String("objectName", objectName), zap.String("filePath", filePath))
	return nil
}

func (c *QiniuyunOSSClient) UploadFormFile(objectName string, reader io.Reader) error {
	izap.Log.Info("UploadFile", zap.String("objectName", objectName), zap.Any("reader", reader))
	
	return nil
}

func (c *QiniuyunOSSClient) Download(objectKey, downloadPath string) error {
	izap.Log.Info("Download", zap.String("objectKey", objectKey), zap.String("downloadPath", downloadPath))
	if len(objectKey) == 0 || len(downloadPath) == 0 {
		return errors.New("objectKey downloadPath is empty")
	}
	return nil
}

func (c *QiniuyunOSSClient) Delete(srcPath string) error {
	izap.Log.Info("Delete", zap.String("srcPath", srcPath))

	return nil
}

func (c *QiniuyunOSSClient) BatchDelete(srcPath []string) error {
	izap.Log.Info("BatchDelete", zap.Strings("srcPath", srcPath))
	return nil
}

func (c *QiniuyunOSSClient) Rename(srcPath, destPath string) error {
	izap.Log.Info("Rename", zap.String("srcPath", srcPath), zap.String("srcPath", destPath))
	return nil
}

func (c *QiniuyunOSSClient) Copy(srcPath, destPath string) error {
	izap.Log.Info("Copy ", zap.String("srcPath", srcPath), zap.String("srcPath", destPath))
	return nil
}

func (c *QiniuyunOSSClient) GetOnly(path string) (string, error) {
	izap.Log.Info("GetOnly ", zap.String("path", path))

	return "", nil
}

func (c *QiniuyunOSSClient) SignURL(ossFilePath, fileName string, expiresInSec int64, category int) (string, error) {
	izap.Log.Info("SignURL", zap.String("ossFilePath", ossFilePath), zap.String("fileName", fileName), zap.Int64("expiresInSec", expiresInSec), zap.Int("category", category))

	return "objectURL", nil
}

func (c *QiniuyunOSSClient) IsExist(filePath string) (bool, error) {
	izap.Log.Info("IsExist", zap.String("filePath", filePath))
	return false, nil
}

func (c *QiniuyunOSSClient) GetFileSize(filePath string) (int64, error) {
	izap.Log.Info("GetFileSize", zap.String("filePath", filePath))
	return 0, nil
}
