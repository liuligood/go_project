package model_data

import "crmeb_go/internal/model"

type Condition struct {
	UserID   string
	Nickname string
	Email    string
}

type UserDateResp struct {
	model.User
	EveryDate string `gorm:"column:every_date;type:varchar(32);not null;comment:每天日期" json:"every_date"` // 每天日期
}
