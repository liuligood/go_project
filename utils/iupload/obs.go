package iupload

import (
	"mime/multipart"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/pkg/errors"
)

var HuaWeiObs = new(Obs)

type Obs struct {
	baseOss BaseOss
}

func (o *Obs) newHuaWeiObsClient() (client *obs.ObsClient, err error) {
	return obs.New(o.baseOss.Conf.HuaWeiObs.AccessKey, o.baseOss.Conf.HuaWeiObs.SecretKey, o.baseOss.Conf.HuaWeiObs.Endpoint)
}

func (o *Obs) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// var open multipart.File
	open, err := file.Open()

	if err != nil {
		return "", "", err
	}
	defer open.Close()

	filename := file.Filename
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: o.baseOss.Conf.HuaWeiObs.Bucket,
				Key:    filename,
			},
			HttpHeader: obs.HttpHeader{
				ContentType: file.Header.Get("content-type"),
			},
		},
		Body: open,
	}

	var client *obs.ObsClient
	client, err = o.newHuaWeiObsClient()
	if err != nil {
		return "", "", errors.Wrap(err, "获取华为对象存储对象失败!")
	}

	_, err = client.PutObject(input)
	if err != nil {
		return "", "", errors.Wrap(err, "文件上传失败!")
	}

	filepath := o.baseOss.Conf.HuaWeiObs.Path + "/" + filename
	return filepath, filename, err
}

func (o *Obs) DeleteFile(key string) error {
	client, err := o.newHuaWeiObsClient()
	if err != nil {
		return errors.Wrap(err, "获取华为对象存储对象失败!")
	}

	input := &obs.DeleteObjectInput{
		Bucket: o.baseOss.Conf.HuaWeiObs.Bucket,
		Key:    key,
	}

	var output *obs.DeleteObjectOutput
	output, err = client.DeleteObject(input)

	if err != nil {
		return errors.Wrapf(err, "删除对象(%s)失败!, output: %v", key, output)
	}
	return nil
}
