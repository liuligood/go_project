package upload

import (
	"bytes"
	"context"
	"crmeb_go/config"
	"crmeb_go/utils/md5"
	"errors"
	"github.com/minio/minio-go/v7"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

var MinioClient *Minio // 优化性能，但是不支持动态配置

type Minio struct {
	Client  *minio.Client
	bucket  string
	baseOss BaseOss
}

func GetMinio(conf config.Conf, logger *zap.Logger) (*Minio, error) {
	if MinioClient != nil {
		return MinioClient, nil
	}

	// Initialize minio client object.
	minioClient, err := minio.New(conf.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.Minio.AccessKeyId, conf.Minio.AccessKeySecret, ""),
		Secure: conf.Minio.UseSSL, // Set to true if using https
	})

	if err != nil {
		return nil, err
	}

	// 尝试创建bucket
	err = minioClient.MakeBucket(context.Background(), conf.Minio.BucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(context.Background(), conf.Minio.BucketName)
		if errBucketExists == nil && exists {
			// log.Printf("We already own %s\n", bucketName)
		} else {
			return nil, err
		}
	}

	MinioClient = &Minio{Client: minioClient, bucket: conf.Minio.BucketName, baseOss: BaseOss{Conf: conf, Logger: logger}}

	return MinioClient, nil
}

func (m *Minio) UploadFile(file *multipart.FileHeader) (filePathres, key string, uploadErr error) {
	f, openError := file.Open()
	// mutipart.File to os.File
	if openError != nil {
		m.baseOss.Logger.Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}

	filecontent := bytes.Buffer{}

	_, err := io.Copy(&filecontent, f)
	if err != nil {
		m.baseOss.Logger.Error("读取文件失败", zap.Any("err", err.Error()))
		return "", "", errors.New("读取文件失败, err:" + err.Error())
	}
	f.Close() // 创建文件 defer 关闭

	// 对文件名进行加密存储
	ext := filepath.Ext(file.Filename)
	filename := md5.MD5V([]byte(strings.TrimSuffix(file.Filename, ext))) + ext

	if m.baseOss.Conf.Minio.BasePath == "" {
		filePathres = "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + filename
	} else {
		filePathres = m.baseOss.Conf.Minio.BasePath + "/" + time.Now().Format("2006-01-02") + "/" + filename
	}

	// 设置超时10分钟
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	// Upload the file with PutObject   大文件自动切换为分片上传
	info, err := m.Client.PutObject(ctx, m.baseOss.Conf.Minio.BucketName, filePathres, &filecontent, file.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		m.baseOss.Logger.Error("上传文件到minio失败", zap.Any("err", err.Error()))
		return "", "", errors.New("上传文件到minio失败, err:" + err.Error())
	}
	return m.baseOss.Conf.Minio.BucketUrl + "/" + info.Key, filePathres, nil
}

func (m *Minio) DeleteFile(key string) error {
	// Delete the object from MinIO
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	err := m.Client.RemoveObject(ctx, m.bucket, key, minio.RemoveObjectOptions{})
	return err
}
