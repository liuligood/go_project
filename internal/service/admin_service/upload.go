package admin_service

import (
	"crmeb_go/internal/data/admin_data"
	service_data "crmeb_go/internal/data/sevice_data"
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
	Upload(params service_data.UploadParams) (data admin_data.UploadResp, err error)
}

type UploadService struct {
	svc *server.SvcContext
}

func NewUploadService(svc *server.SvcContext) *UploadService {
	return &UploadService{svc: svc}
}

func (a UploadService) Upload(params service_data.UploadParams) (data admin_data.UploadResp, err error) {
	var uploadFileName string

	if params.ContentLength > (a.svc.Conf.System.FileSize << 20) {
		izap.Log.Error("UploadFile ContentLength [%v]")

		return data, errors.New("文件过大")
	}

	fileData, openError := params.FileHeader.Open()
	if openError != nil {
		izap.Log.Error("function file.Open() Failed", zap.Error(err))

		return data, errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer fileData.Close() // 创建文件 defer 关闭

	if params.FileHeader.Size == 0 {
		return data, errors.New("文件为空")
	}

	split := strings.Split(params.FileHeader.Filename, ".")
	if len(split) < 2 {
		return data, errors.New("获取文件后缀失败")
	}

	b := make([]byte, params.FileHeader.Size)
	_, err = params.File.Read(b)

	if err != nil {
		izap.Log.Error("read file err", zap.Error(err))

		return data, err
	}

	ft, err := filetype.Get(b)
	if err != nil {
		izap.Log.Error("不支持的文件类型")

		return data, errors.New("不支持的文件类型")
	}

	ext := split[len(split)-1]

	dir := a.getDir(params.PathName, params.FileType) // 文件路径
	switch {
	case params.FileKey != "" && dir != "":
		keySplit := strings.Split(params.FileKey, ".")
		keyName := keySplit[0]
		uploadFileName = dir + keyName + "." + ft.Extension
	case params.FileKey != "":
		uploadFileName = params.FileKey
	default:
		tmp := b
		tmp = append(tmp, []byte(params.FileHeader.Filename)...)
		uploadFileName = imd5.MD5V(tmp) + "." + ext
	}

	uploadFileName = a.svc.Conf.AliyunOSS.BasePath + "/" + time.Now().Format("2006-01-02") + "/" + uploadFileName

	checkData, err := oss.OssClient.IsExist(uploadFileName)
	if err != nil {
		izap.Log.Error("check file exist err", zap.Error(err))

		return data, errors.New("检查文件重复上传失败")
	}

	if checkData {
		return data, errors.New("文件重复上传")
	}

	if err := oss.OssClient.UploadFormFile(uploadFileName, fileData); err != nil {
		return data, errors.New("文件上传错误")
	}

	return admin_data.UploadResp{Url: fmt.Sprintf("https://%s.%s/%s", a.svc.Conf.AliyunOSS.BucketName, a.svc.Conf.AliyunOSS.Endpoint, uploadFileName)}, nil
}

func (l *UploadService) getDir(pathName string, fileType string) string {
	if pathName == "" {
		return ""
	}

	t := time.Now()
	dir := fmt.Sprintf("uploads/image/%s/%s/%s/", pathName, t.Format("200601"), t.Format("02"))

	if fileType == "1" {
		dir = fmt.Sprintf("Public/%s", dir)
	}

	return dir
}
