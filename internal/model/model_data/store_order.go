package model_data

import "crmeb_go/internal/model"

type DateCondition struct {
	Start int64
	End   int64
}

type StoreOrderDateResp struct {
	model.StoreOrder
	EveryDate string `gorm:"column:every_date;type:varchar(32);not null;comment:每天日期" json:"every_date"` // 每天日期
}
