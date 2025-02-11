// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/plugin/soft_delete"
)

const TableNameStoreCoupon = "store_coupon"

// StoreCoupon 优惠券表
type StoreCoupon struct {
	ID               int64                 `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true;comment:优惠券表ID" json:"id"`                 // 优惠券表ID
	Name             string                `gorm:"column:name;type:varchar(64);not null;comment:优惠券名称" json:"name"`                                        // 优惠券名称
	Money            decimal.Decimal       `gorm:"column:money;type:decimal(8,2) unsigned;not null;default:0.00;comment:兑换的优惠券面值" json:"money"`            // 兑换的优惠券面值
	IsLimited        int64                 `gorm:"column:is_limited;type:tinyint(1);comment:是否限量, 默认0 不限量， 1限量" json:"is_limited"`                         // 是否限量, 默认0 不限量， 1限量
	Total            int64                 `gorm:"column:total;type:int(11);not null;comment:发放总数" json:"total"`                                           // 发放总数
	LastTotal        int64                 `gorm:"column:last_total;type:int(11);comment:剩余数量" json:"last_total"`                                          // 剩余数量
	UseType          int64                 `gorm:"column:use_type;type:tinyint(4);not null;default:1;comment:使用类型 1 全场通用, 2 商品券, 3 品类券" json:"use_type"`   // 使用类型 1 全场通用, 2 商品券, 3 品类券
	PrimaryKey       string                `gorm:"column:primary_key;type:varchar(255);not null;comment:所属商品id / 分类id" json:"primary_key"`                 // 所属商品id / 分类id
	MinPrice         decimal.Decimal       `gorm:"column:min_price;type:decimal(8,2) unsigned;not null;default:0.00;comment:最低消费，0代表不限制" json:"min_price"` // 最低消费，0代表不限制
	ReceiveStartTime int64                 `gorm:"column:receive_start_time;type:bigint(20);not null" json:"receive_start_time"`
	ReceiveEndTime   int64                 `gorm:"column:receive_end_time;type:bigint(20)" json:"receive_end_time"`
	IsFixedTime      int64                 `gorm:"column:is_fixed_time;type:tinyint(1);comment:是否固定使用时间, 默认0 否， 1是" json:"is_fixed_time"` // 是否固定使用时间, 默认0 否， 1是
	UseStartTime     int64                 `gorm:"column:use_start_time;type:bigint(20)" json:"use_start_time"`
	UseEndTime       int64                 `gorm:"column:use_end_time;type:bigint(20)" json:"use_end_time"`
	Day              int64                 `gorm:"column:day;type:int(11);comment:天数" json:"day"`                                                 // 天数
	Type             int64                 `gorm:"column:type;type:tinyint(4);not null;default:1;comment:优惠券类型 1 手动领取, 2 新人券, 3 赠送券" json:"type"` // 优惠券类型 1 手动领取, 2 新人券, 3 赠送券
	Sort             int64                 `gorm:"column:sort;type:int(10) unsigned;not null;default:1;comment:排序" json:"sort"`                   // 排序
	Status           int64                 `gorm:"column:status;type:tinyint(3) unsigned;not null;default:1;comment:状态（0：关闭，1：开启）" json:"status"` // 状态（0：关闭，1：开启）
	CreatedAt        int64                 `gorm:"column:created_at;type:bigint(20) unsigned;not null;comment:创建时间" json:"created_at"`            // 创建时间
	UpdatedAt        int64                 `gorm:"column:updated_at;type:bigint(20) unsigned;not null;comment:修改时间" json:"updated_at"`            // 修改时间
	DeletedAt        soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint(3) unsigned;not null;comment:是否删除" json:"-"`                     // 是否删除
}

// TableName StoreCoupon's table name
func (*StoreCoupon) TableName() string {
	return TableNameStoreCoupon
}
