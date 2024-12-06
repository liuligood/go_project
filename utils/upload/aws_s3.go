package upload

import (
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.uber.org/zap"
)

type AwsS3 struct {
	baseOss BaseOss
}

func (a *AwsS3) UploadFile(file *multipart.FileHeader) (string, string, error) {
	session := a.newSession()
	uploader := s3manager.NewUploader(session)

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	filename := a.baseOss.Conf.AwsS3.PathPrefix + "/" + fileKey
	f, openError := file.Open()
	if openError != nil {
		a.baseOss.Logger.Error("function file.Open() failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(a.baseOss.Conf.AwsS3.Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		a.baseOss.Logger.Error("function uploader.Upload() failed", zap.Any("err", err.Error()))
		return "", "", err
	}

	return a.baseOss.Conf.AwsS3.BaseURL + "/" + filename, fileKey, nil
}

func (a *AwsS3) DeleteFile(key string) error {
	session := a.newSession()
	svc := s3.New(session)
	filename := a.baseOss.Conf.AwsS3.PathPrefix + "/" + key
	bucket := a.baseOss.Conf.AwsS3.Bucket

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		a.baseOss.Logger.Error("function svc.DeleteObject() failed", zap.Any("err", err.Error()))
		return errors.New("function svc.DeleteObject() failed, err:" + err.Error())
	}

	_ = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	return nil
}

func (a *AwsS3) newSession() *session.Session {
	sess, _ := session.NewSession(&aws.Config{
		Region:           aws.String(a.baseOss.Conf.AwsS3.Region),
		Endpoint:         aws.String(a.baseOss.Conf.AwsS3.Endpoint), //minio在这里设置地址,可以兼容
		S3ForcePathStyle: aws.Bool(a.baseOss.Conf.AwsS3.S3ForcePathStyle),
		DisableSSL:       aws.Bool(a.baseOss.Conf.AwsS3.DisableSSL),
		Credentials: credentials.NewStaticCredentials(
			a.baseOss.Conf.AwsS3.SecretID,
			a.baseOss.Conf.AwsS3.SecretKey,
			"",
		),
	})
	return sess
}
