package ihttp

import (
	"crmeb_go/utils/ierror"
	"net/http"
)

const (
	SuccessMessage = "操作成功"
	ErrorMessage   = "操作失败"
)

// RespData 结构体，用于API响应
type RespData struct {
	Code    int         `json:"code"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta,omitempty"`
}

type Response interface {
	GetRespData() *RespData
	SetCode(code int) (RespData Response)
	MetaData(meta interface{}) (RespData Response)
	AddMeta(key string, value interface{}) (RespData Response)
	AddComponentMeta(value interface{}) (RespData Response)
}

// ErrMsg 错误类型响应
func ErrMsg(message string, status ...int) *RespData {
	statusCode := http.StatusBadRequest
	if len(status) != 0 {
		statusCode = status[0]
	}
	return &RespData{
		Code:    http.StatusOK,
		Status:  statusCode,
		Message: message,
	}
}

// Error 错误类型响应
func Error(err error, status ...int) Response {
	code := http.StatusInternalServerError
	statusCode := http.StatusBadRequest
	message := err.Error()
	var data interface{}
	if ierror.IsError(err) {
		newError := ierror.UnmarshalError(err)
		statusCode = int(newError.Status)
		code = int(newError.Status)
		message = newError.Message
		data = newError.Metadata
	}
	if len(status) != 0 {
		statusCode = status[0]
	}
	return &RespData{
		Code:    code,
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
}

// Data 数组类型响应
func Data(data interface{}) Response {
	return &RespData{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    data,
	}
}

// List 列表类型响应
func List[T any](data []T) Response {
	if len(data) == 0 {
		data = []T{}
	}
	return &RespData{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    data,
	}
}

// Nil 空响应
func Nil() Response {
	return &RespData{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    nil,
	}
}

// Paginate 分页类型响应
func Paginate[T any](list []T, total, page, pageSize int64) Response {
	var data = make(map[string]interface{})
	data["data"] = map[string]interface{}{
		"count":     len(list),
		"total":     total,
		"page":      page,
		"limit":     pageSize,
		"totalPage": calcTotalPage(total, pageSize),
		"list":      list,
	}
	if len(list) == 0 {
		list = []T{}
	}
	return &RespData{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Status:  0,
		Data:    data,
	}
}

// calcTotalPage 计算总页数的辅助函数
func calcTotalPage(total, pageSize int64) int64 {
	if pageSize == 0 {
		return 0
	}
	totalPage := total / pageSize
	if total%pageSize != 0 {
		totalPage++
	}
	return totalPage
}

// AddMeta 添加元数据
func (r *RespData) AddComponentMeta(value interface{}) (RespData Response) {
	r.AddMeta("component", value)
	return r
}

// AddMeta 添加元数据
func (r *RespData) AddMeta(key string, value interface{}) (RespData Response) {
	if r.Meta == nil {
		var meta = make(map[string]interface{})
		meta[key] = value
		r.Meta = meta
	} else {
		r.Meta.(map[string]any)[key] = value
	}
	return r
}

// MetaData 设置元数据
func (r *RespData) MetaData(meta interface{}) (RespData Response) {
	r.Meta = meta
	return r
}

// SetCode 设置状态码
func (r *RespData) SetCode(code int) (RespData Response) {
	r.Code = code
	return r
}

// GetRespData 获取响应数据
func (r *RespData) GetRespData() *RespData {
	return r
}
