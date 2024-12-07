package upload

import (
	"context"
	"crmeb_go/utils/izap"
	"errors"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
)

type TencentCOS struct {
	baseOss BaseOss
}

// UploadFile upload file to COS
func (t *TencentCOS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client := t.NewClient()

	f, openError := file.Open()
	if openError != nil {
		izap.Log.Error("function file.Open() failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}

	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), t.baseOss.Conf.TencentCOS.PathPrefix+"/"+fileKey, f, nil)
	if err != nil {
		panic(err)
	}

	return t.baseOss.Conf.TencentCOS.BaseURL + "/" + t.baseOss.Conf.TencentCOS.PathPrefix + "/" + fileKey, fileKey, nil
}

// DeleteFile delete file form COS
func (t *TencentCOS) DeleteFile(key string) error {
	client := t.NewClient()

	name := t.baseOss.Conf.TencentCOS.PathPrefix + "/" + key
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		izap.Log.Error("function bucketManager.Delete() failed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}

	return nil
}

// NewClient init COS client
func (t *TencentCOS) NewClient() *cos.Client {
	urlStr, _ := url.Parse("https://" + t.baseOss.Conf.TencentCOS.Bucket + ".cos." + t.baseOss.Conf.TencentCOS.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  t.baseOss.Conf.TencentCOS.SecretID,
			SecretKey: t.baseOss.Conf.TencentCOS.SecretKey,
		},
	})
	return client
}
