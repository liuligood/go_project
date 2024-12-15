package oss

import (
	"crmeb_go/config"
	"crmeb_go/define"
	"crmeb_go/utils/izap"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"io"
	"sync"
)

type Client interface {
	UploadFile(objectName string, filePath string) error
	Download(objectKey, downloadPath string) error
	Delete(srcPath string) error
	BatchDelete(srcPath []string) error
	Rename(srcPath, destPath string) error
	Copy(srcPath, destPath string) error
	GetOnly(path string) (string, error)
	SignURL(ossFilePath, fileName string, expiresInSec int64, category int) (string, error)
	IsExist(filePath string) (bool, error)
	GetFileSize(filePath string) (int64, error)
	UploadFormFile(objectName string, reader io.Reader) error
}

var OssClient Client
var syncOssClient Client

func InitOSS(redisClient redis.UniversalClient, config config.Conf) error {
	// 使用stsToken 上传
	//result, http_err := redisClient.Get(context.Background(), define.OSSSTSTokenKey).Result()
	//if http_err != nil && !errors.Is(http_err, redis.Nil) {
	//	izap.Log.Error("redis http_err:", zap.Error(http_err))
	//
	//	return http_err
	//}
	//
	//if len(result) == 0 {
	//	jsonData, http_err := GenerateSTSToken(config)
	//	if http_err != nil {
	//		izap.Log.Error("GenerateSTSToken http_err: ", zap.Error(http_err))
	//
	//		return http_err
	//	}
	//
	//	redisClient.Set(context.Background(), define.OSSSTSTokenKey, string(jsonData), time.Hour)
	//}

	registerFactory(define.OssAliYun, &AliyunOSSFactory{})

	return nil
}

var once sync.Once

func GetOssClient(c config.Conf) (Client, error) {
	var flag error
	once.Do(func() {
		if c.System.OssType == "" {
			c.System.OssType = define.OssAliYun
		}

		factory, ok := getFactory(c.System.OssType)

		if !ok {
			izap.Log.Error("oss provider not found")
			flag = errors.New("oss provider not found")
			return
		}

		client, err := factory.Create(c)
		if err != nil {
			panic(err)
		}
		OssClient = client
	})

	if flag != nil {
		return nil, flag
	}

	if OssClient == nil {
		return nil, errors.New("oss client is nil")
	}
	return OssClient, nil
}

func GenerateSTSToken(config config.Conf) ([]byte, error) {
	client, err := sts.NewClientWithAccessKey(config.AliyunOSS.Region, config.AliyunOSS.AccessKeyId, config.AliyunOSS.AccessKeySecret)
	if err != nil {
		izap.Log.Error("Error: ", zap.Error(err))
	}
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"
	request.RoleArn = config.AliyunOSS.RoleArn
	request.RoleSessionName = uuid.NewString()
	request.DurationSeconds = define.OssStsTokenDurationSeconds

	// 4. 获取临时凭证
	response, err := client.AssumeRole(request)
	if err != nil {
		fmt.Println("Error getting STS token:", err)
		return nil, err
	}

	var ossInfo = struct {
		OSSAccessKeyId     string `json:"access_key_id"`
		OSSAccessKeySecret string `json:"access_key_secret"`
		OSSBucket          string `json:"bucket"`
		OSSEndpoint        string `json:"endpoint"`
		OSSRegion          string `json:"region"`
		OSSSecurityToken   string `json:"security_token"`
	}{
		OSSAccessKeyId:     response.Credentials.AccessKeyId,
		OSSAccessKeySecret: response.Credentials.AccessKeySecret,
		OSSSecurityToken:   response.Credentials.SecurityToken,
		OSSBucket:          config.AliyunOSS.BucketName,
		OSSEndpoint:        config.AliyunOSS.Endpoint,
		OSSRegion:          config.AliyunOSS.Region,
	}

	izap.Log.Info("GenerateSTSToken OSSConfig:", zap.Any("ossInfo", ossInfo))
	// 将结构体转换为JSON格式的字节切片
	jsonData, err := json.Marshal(ossInfo)
	if err != nil {
		izap.Log.Error("json Marshal http_err: ", zap.Error(err))

		return nil, err
	}

	return jsonData, nil
}
