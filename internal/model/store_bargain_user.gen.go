// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/plugin/soft_delete"
)

const TableNameStoreBargainUser = "store_bargain_user"

// StoreBargainUser 用户参与砍价表
type StoreBargainUser struct {
	ID              int64                 `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true;comment:用户参与砍价表ID" json:"id"`          // 用户参与砍价表ID
	UID             int64                 `gorm:"column:uid;type:int(10) unsigned;comment:用户ID" json:"uid"`                                           // 用户ID
	BargainID       int64                 `gorm:"column:bargain_id;type:int(10) unsigned;comment:砍价商品id" json:"bargain_id"`                           // 砍价商品id
	BargainPriceMin decimal.Decimal       `gorm:"column:bargain_price_min;type:decimal(8,2) unsigned;comment:砍价的最低价" json:"bargain_price_min"`        // 砍价的最低价
	BargainPrice    decimal.Decimal       `gorm:"column:bargain_price;type:decimal(8,2);comment:砍价金额" json:"bargain_price"`                           // 砍价金额
	Price           decimal.Decimal       `gorm:"column:price;type:decimal(8,2) unsigned;comment:砍掉的价格" json:"price"`                                 // 砍掉的价格
	Status          int64                 `gorm:"column:status;type:tinyint(3) unsigned;not null;comment:状态 1参与中 2 活动结束参与失败 3活动结束参与成功" json:"status"` // 状态 1参与中 2 活动结束参与失败 3活动结束参与成功
	AddTime         int64                 `gorm:"column:add_time;type:bigint(20) unsigned;comment:参与时间" json:"add_time"`                              // 参与时间
	CreatedAt       int64                 `gorm:"column:created_at;type:bigint(20) unsigned;not null;comment:创建时间" json:"created_at"`                 // 创建时间
	UpdatedAt       int64                 `gorm:"column:updated_at;type:bigint(20) unsigned;not null;comment:修改时间" json:"updated_at"`                 // 修改时间
	DeletedAt       soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint(3) unsigned;not null;comment:是否删除" json:"-"`                          // 是否删除
}

// TableName StoreBargainUser's table name
func (*StoreBargainUser) TableName() string {
	return TableNameStoreBargainUser
}
