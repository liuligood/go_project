package service_data

import "mime/multipart"

type GetLoginPicParams struct {
	BaseServiceParams
	Temp int64 `json:"temp"`
}

type UploadParams struct {
	BaseServiceParams
	File          multipart.File
	FileHeader    *multipart.FileHeader
	FileKey       string
	PathName      string
	FileType      string
	ContentLength int64
}
