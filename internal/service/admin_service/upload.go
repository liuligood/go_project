package admin_service

import (
	"crmeb_go/define"
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
	"crmeb_go/utils/imd5"
	oss "crmeb_go/utils/ioss"
	"crmeb_go/utils/izap"
	"errors"
	"fmt"
	"github.com/h2non/filetype"
	"go.uber.org/zap"
	"strings"
	"time"
)

type UploadServiceImpl interface {
	UploadFile(params *request.UploadParams) (data *response.UploadResp, err error)
}

type UploadService struct {
	svc *server.SvcContext
}

func NewUploadService(svc *server.SvcContext) *UploadService {
	return &UploadService{svc: svc}
}

func (u *UploadService) UploadFile(params *request.UploadParams) (data *response.UploadResp, err error) {
	var uploadFileName string
	if params.ContentLength > (u.svc.Conf.System.FileSize << 20) {
		izap.Log.Error("UploadFile ContentLength [%v]")

		return data, errors.New("文件过大")
	}

	fileData, openError := params.FileHeader.Open()
	if openError != nil {
		izap.Log.Error("function file.Open() Failed", zap.Error(err))

		return data, errors.New("function file.Open() Failed, errorm:" + openError.Error())
	}
	defer fileData.Close() // 创建文件 defer 关闭

	if params.FileHeader.Size == 0 {
		return data, errors.New("文件为空")
	}

	fileSplit := strings.Split(params.FileHeader.Filename, ".")
	if len(fileSplit) < 2 {
		return data, errors.New("获取文件名,文件后缀失败")
	}

	b := make([]byte, params.FileHeader.Size)
	_, err = params.File.Read(b)

	if err != nil {
		izap.Log.Error("read file errorm", zap.Error(err))

		return data, err
	}

	fileType, err := filetype.Get(b)
	if err != nil {
		izap.Log.Error("不支持的文件类型")

		return data, errors.New("不支持的文件类型")
	}

	suffix := fileSplit[len(fileSplit)-1] // 文件后缀
	dir := u.getDirPath(params.PathName)  // 文件路径

	switch {
	case params.FileKey != "" && dir != "":
		keySplit := strings.Split(params.FileKey, ".")
		keyName := keySplit[0]
		uploadFileName = dir + keyName + "." + fileType.Extension
	case params.FileKey != "":
		uploadFileName = params.FileKey
	default:
		tmp := b
		tmp = append(tmp, []byte(params.FileHeader.Filename)...)
		uploadFileName = imd5.MD5V(tmp) + "." + suffix
	}

	// 看情况修改名称
	uploadFileName = u.svc.Conf.AliyunOSS.BasePath + "/" + time.Now().Format(define.SystemTimeFormatDate) + "/" + uploadFileName
	checkData, err := oss.OssClient.IsExist(uploadFileName)

	if err != nil {
		izap.Log.Error("check file exist errorm", zap.Error(err))

		return data, errors.New("检查文件重复上传失败")
	}

	if checkData {
		return data, errors.New("文件重复上传")
	}

	if err := oss.OssClient.UploadFormFile(uploadFileName, fileData); err != nil {
		izap.Log.Error("UploadFile UploadFormFile errorm", zap.Error(err))

		return data, errors.New("文件上传错误")
	}

	return &response.UploadResp{Url: fmt.Sprintf("https://%s.%s/%s", u.svc.Conf.AliyunOSS.BucketName, u.svc.Conf.AliyunOSS.Endpoint, uploadFileName)}, nil
}

func (u *UploadService) getDirPath(pathName string) string {
	if pathName == "" {
		return ""
	}

	dir := fmt.Sprintf("uploads/image/%s/%s/%s/", pathName, time.Now().Format("200601"), time.Now().Format("02"))
	return dir
}
