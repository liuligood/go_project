// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameUserIntegralRecord = "user_integral_record"

// UserIntegralRecord 用户积分记录表
type UserIntegralRecord struct {
	ID         int64                 `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:记录id" json:"id"`                                   // 记录id
	UID        int64                 `gorm:"column:uid;type:int;not null;comment:用户uid" json:"uid"`                                                     // 用户uid
	LinkID     string                `gorm:"column:link_id;type:varchar(32);not null;default:0;comment:关联id-orderNo,(sign,system默认为0）" json:"link_id"`  // 关联id-orderNo,(sign,system默认为0）
	LinkType   string                `gorm:"column:link_type;type:varchar(32);not null;default:order;comment:关联类型（order,sign,system）" json:"link_type"` // 关联类型（order,sign,system）
	Type       int64                 `gorm:"column:type;type:int;not null;default:1;comment:类型：1-增加，2-扣减" json:"type"`                                  // 类型：1-增加，2-扣减
	Title      string                `gorm:"column:title;type:varchar(64);not null;comment:标题" json:"title"`                                            // 标题
	Integral   int64                 `gorm:"column:integral;type:int;not null;comment:积分" json:"integral"`                                              // 积分
	Balance    int64                 `gorm:"column:balance;type:int;not null;comment:剩余" json:"balance"`                                                // 剩余
	Mark       string                `gorm:"column:mark;type:varchar(512);not null;comment:备注" json:"mark"`                                             // 备注
	Status     int64                 `gorm:"column:status;type:tinyint(1);not null;default:1;comment:状态：1-订单创建，2-冻结期，3-完成，4-失效（订单退款）" json:"status"`    // 状态：1-订单创建，2-冻结期，3-完成，4-失效（订单退款）
	FrozenTime int64                 `gorm:"column:frozen_time;type:int;not null;comment:冻结期时间（天）" json:"frozen_time"`                                  // 冻结期时间（天）
	ThawTime   int64                 `gorm:"column:thaw_time;type:bigint;not null;comment:解冻时间" json:"thaw_time"`                                       // 解冻时间
	CreatedAt  int64                 `gorm:"column:created_at;type:bigint unsigned;not null;comment:创建时间" json:"created_at"`                            // 创建时间
	UpdatedAt  int64                 `gorm:"column:updated_at;type:bigint unsigned;not null;comment:修改时间" json:"updated_at"`                            // 修改时间
	DeletedAt  soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint unsigned;not null;comment:是否删除" json:"-"`                                    // 是否删除
}

// TableName UserIntegralRecord's table name
func (*UserIntegralRecord) TableName() string {
	return TableNameUserIntegralRecord
}
