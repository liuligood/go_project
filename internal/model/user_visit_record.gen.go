// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameUserVisitRecord = "user_visit_record"

// UserVisitRecord 用户访问记录表
type UserVisitRecord struct {
	ID        int64                 `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`
	Date      int64                 `gorm:"column:date;type:bigint(20) unsigned;not null;comment:日期" json:"date"`               // 日期
	UID       int64                 `gorm:"column:uid;type:int(10) unsigned;not null;comment:用户uid" json:"uid"`                 // 用户uid
	VisitType int64                 `gorm:"column:visit_type;type:int(10) unsigned;not null;comment:访问类型" json:"visit_type"`    // 访问类型
	CreatedAt int64                 `gorm:"column:created_at;type:bigint(20) unsigned;not null;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt int64                 `gorm:"column:updated_at;type:bigint(20) unsigned;not null;comment:修改时间" json:"updated_at"` // 修改时间
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint(3) unsigned;not null;comment:是否删除" json:"-"`          // 是否删除
}

// TableName UserVisitRecord's table name
func (*UserVisitRecord) TableName() string {
	return TableNameUserVisitRecord
}
