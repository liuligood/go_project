package response

import "time"

type GetSystemConfigResult struct {
	ID         uint64    `json:"id"`          // 配置id
	Name       string    ` json:"name"`       // 字段名称
	Title      string    `json:"title"`       // 字段提示文字
	FormID     uint64    `json:"form_id"`     // 表单id
	Value      string    ` json:"value"`      // 值
	Status     uint64    `json:"status"`      // 是否隐藏
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}
