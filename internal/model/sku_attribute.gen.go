// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameSkuAttribute = "sku_attribute"

// SkuAttribute sku属性表
type SkuAttribute struct {
	ID        int64                 `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:主键" json:"id"`          // 主键
	SkuID     int64                 `gorm:"column:sku_id;type:int unsigned;not null;comment:skuId" json:"sku_id"`           // skuId
	Type      int64                 `gorm:"column:type;type:tinyint unsigned;not null;comment:类型" json:"type"`              // 类型
	Title     string                `gorm:"column:title;type:varchar(64);not null;comment:属性标题" json:"title"`               // 属性标题
	Value     string                `gorm:"column:value;type:varchar(128);not null;comment:属性值" json:"value"`               // 属性值
	Sort      int64                 `gorm:"column:sort;type:int unsigned;not null;comment:排序" json:"sort"`                  // 排序
	Sign      int64                 `gorm:"column:sign;type:tinyint unsigned;not null;comment:标志(位运算)" json:"sign"`         // 标志(位运算)
	CreatedAt int64                 `gorm:"column:created_at;type:bigint unsigned;not null;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt int64                 `gorm:"column:updated_at;type:bigint unsigned;not null;comment:修改时间" json:"updated_at"` // 修改时间
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint unsigned;not null;comment:是否删除" json:"-"`         // 是否删除
	CreatedBy int64                 `gorm:"column:created_by;type:int unsigned;not null;comment:创建时间" json:"created_by"`    // 创建时间
	UpdatedBy int64                 `gorm:"column:updated_by;type:int unsigned;not null;comment:修改时间" json:"updated_by"`    // 修改时间
}

// TableName SkuAttribute's table name
func (*SkuAttribute) TableName() string {
	return TableNameSkuAttribute
}
