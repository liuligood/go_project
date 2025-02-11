// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameSystemGroup = "system_group"

// SystemGroup 组合数据表
type SystemGroup struct {
	ID        int64                 `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true;comment:组合数据ID" json:"id"`      // 组合数据ID
	Name      string                `gorm:"column:name;type:varchar(50);not null;comment:数据组名称" json:"name"`                    // 数据组名称
	Info      string                `gorm:"column:info;type:varchar(256);not null;comment:简介" json:"info"`                      // 简介
	FormID    int64                 `gorm:"column:form_id;type:int(11);not null;comment:form 表单 id" json:"form_id"`             // form 表单 id
	CreatedAt int64                 `gorm:"column:created_at;type:bigint(20) unsigned;not null;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt int64                 `gorm:"column:updated_at;type:bigint(20) unsigned;not null;comment:修改时间" json:"updated_at"` // 修改时间
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint(3) unsigned;not null;comment:是否删除" json:"-"`          // 是否删除
}

// TableName SystemGroup's table name
func (*SystemGroup) TableName() string {
	return TableNameSystemGroup
}
