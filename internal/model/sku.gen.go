// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/plugin/soft_delete"
)

const TableNameSku = "sku"

// Sku 套餐表
type Sku struct {
	ID        int64                 `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`
	SpuID     int64                 `gorm:"column:spu_id;type:int(10) unsigned;not null" json:"spu_id"`
	Stock     int64                 `gorm:"column:stock;type:int(10) unsigned;not null;comment:库存" json:"stock"`                      // 库存
	Price     decimal.Decimal       `gorm:"column:price;type:decimal(10,4) unsigned;not null;default:0.0000;comment:价格" json:"price"` // 价格
	Status    int64                 `gorm:"column:status;type:tinyint(3) unsigned;not null;comment:状态(1:关闭,2:正常)" json:"status"`      // 状态(1:关闭,2:正常)
	Sign      int64                 `gorm:"column:sign;type:int(10) unsigned;not null;comment:标志(位运算)" json:"sign"`                   // 标志(位运算)
	Cost      decimal.Decimal       `gorm:"column:cost;type:decimal(10,4) unsigned;not null;default:0.0000;comment:成本价" json:"cost"`  // 成本价
	Weight    decimal.Decimal       `gorm:"column:weight;type:decimal(10,2) unsigned;not null;default:0.00;comment:重量" json:"weight"` // 重量
	Volume    decimal.Decimal       `gorm:"column:volume;type:decimal(10,2) unsigned;not null;default:0.00;comment:体积" json:"volume"` // 体积
	CreatedAt int64                 `gorm:"column:created_at;type:bigint(20) unsigned;not null;comment:创建时间" json:"created_at"`       // 创建时间
	UpdatedAt int64                 `gorm:"column:updated_at;type:bigint(20) unsigned;not null;comment:修改时间" json:"updated_at"`       // 修改时间
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint(3) unsigned;not null;comment:是否删除" json:"-"`                // 是否删除
	CreatedBy int64                 `gorm:"column:created_by;type:int(10) unsigned;not null;comment:创建时间" json:"created_by"`          // 创建时间
	UpdatedBy int64                 `gorm:"column:updated_by;type:int(10) unsigned;not null;comment:修改时间" json:"updated_by"`          // 修改时间
}

// TableName Sku's table name
func (*Sku) TableName() string {
	return TableNameSku
}
