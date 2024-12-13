// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameExpress = "express"

// Express 快递公司表
type Express struct {
	ID         int64                 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;comment:快递公司id" json:"id"` // 快递公司id
	Code       string                `gorm:"column:code;type:varchar(50);not null;comment:快递公司简称" json:"code"`                   // 快递公司简称
	Name       string                `gorm:"column:name;type:varchar(50);not null;comment:快递公司全称" json:"name"`                   // 快递公司全称
	PartnerID  int64                 `gorm:"column:partner_id;type:tinyint(1);not null;comment:是否需要月结账号" json:"partner_id"`      // 是否需要月结账号
	PartnerKey int64                 `gorm:"column:partner_key;type:tinyint(1);not null;comment:是否需要月结密码" json:"partner_key"`    // 是否需要月结密码
	Net        int64                 `gorm:"column:net;type:tinyint(1);not null;comment:是否需要取件网店" json:"net"`                    // 是否需要取件网店
	Account    string                `gorm:"column:account;type:varchar(100);not null;comment:账号" json:"account"`                // 账号
	Password   string                `gorm:"column:password;type:varchar(100);not null;comment:密码" json:"password"`              // 密码
	NetName    string                `gorm:"column:net_name;type:varchar(100);not null;comment:网点名称" json:"net_name"`            // 网点名称
	Sort       int64                 `gorm:"column:sort;type:int;not null;comment:排序" json:"sort"`                               // 排序
	IsShow     int64                 `gorm:"column:is_show;type:tinyint(1);not null;comment:是否显示" json:"is_show"`                // 是否显示
	Status     int64                 `gorm:"column:status;type:tinyint(1);not null;comment:是否可用" json:"status"`                  // 是否可用
	CreatedAt  int64                 `gorm:"column:created_at;type:bigint unsigned;not null;comment:创建时间" json:"created_at"`     // 创建时间
	UpdatedAt  int64                 `gorm:"column:updated_at;type:bigint unsigned;not null;comment:修改时间" json:"updated_at"`     // 修改时间
	DeletedAt  soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint unsigned;not null;comment:是否删除" json:"-"`             // 是否删除
}

// TableName Express's table name
func (*Express) TableName() string {
	return TableNameExpress
}
