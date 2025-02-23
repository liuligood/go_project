// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameSystemUserLevel = "system_user_level"

// SystemUserLevel 普通会员等级
type SystemUserLevel struct {
	ID         int64                 `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`
	Name       string                `gorm:"column:name;type:varchar(255);not null;comment:会员名称" json:"name"`                    // 会员名称
	Experience int64                 `gorm:"column:experience;type:int(11);not null;comment:达到多少升级经验" json:"experience"`         // 达到多少升级经验
	IsShow     int64                 `gorm:"column:is_show;type:tinyint(1);not null;comment:是否显示 1=显示,0=隐藏" json:"is_show"`      // 是否显示 1=显示,0=隐藏
	Grade      int64                 `gorm:"column:grade;type:int(11);not null;comment:会员等级" json:"grade"`                       // 会员等级
	Discount   int64                 `gorm:"column:discount;type:int(11);not null;default:100;comment:享受折扣" json:"discount"`     // 享受折扣
	Icon       string                `gorm:"column:icon;type:varchar(255);not null;comment:会员图标" json:"icon"`                    // 会员图标
	CreatedAt  int64                 `gorm:"column:created_at;type:bigint(20) unsigned;not null;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt  int64                 `gorm:"column:updated_at;type:bigint(20) unsigned;not null;comment:修改时间" json:"updated_at"` // 修改时间
	DeletedAt  soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint(3) unsigned;not null;comment:是否删除" json:"-"`          // 是否删除
}

// TableName SystemUserLevel's table name
func (*SystemUserLevel) TableName() string {
	return TableNameSystemUserLevel
}
