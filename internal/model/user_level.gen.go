// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameUserLevel = "user_level"

// UserLevel 用户等级记录表
type UserLevel struct {
	ID          int64                 `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	UID         int64                 `gorm:"column:uid;type:int;not null;comment:用户uid" json:"uid"`                      // 用户uid
	LevelID     int64                 `gorm:"column:level_id;type:int;not null;comment:等级vip" json:"level_id"`            // 等级vip
	Grade       int64                 `gorm:"column:grade;type:int;not null;comment:会员等级" json:"grade"`                   // 会员等级
	Status      int64                 `gorm:"column:status;type:tinyint(1);not null;comment:0:禁止,1:正常" json:"status"`     // 0:禁止,1:正常
	Mark        string                `gorm:"column:mark;type:varchar(255);not null;comment:备注" json:"mark"`              // 备注
	Remind      int64                 `gorm:"column:remind;type:tinyint(1);not null;comment:是否已通知" json:"remind"`         // 是否已通知
	Discount    int64                 `gorm:"column:discount;type:int;not null;default:100;comment:享受折扣" json:"discount"` // 享受折扣
	ExpiredTime int64                 `gorm:"column:expired_time;type:bigint" json:"expired_time"`
	CreatedAt   int64                 `gorm:"column:created_at;type:bigint unsigned;not null;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt   int64                 `gorm:"column:updated_at;type:bigint unsigned;not null;comment:修改时间" json:"updated_at"` // 修改时间
	DeletedAt   soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint unsigned;not null;comment:是否删除" json:"-"`         // 是否删除
}

// TableName UserLevel's table name
func (*UserLevel) TableName() string {
	return TableNameUserLevel
}
