// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/plugin/soft_delete"
)

const TableNameSpuStockStream = "spu_stock_stream"

// SpuStockStream 商品库存流水表
type SpuStockStream struct {
	ID           int64                 `gorm:"column:id;type:int(10) unsigned;primaryKey;comment:主键" json:"id"`                          // 主键
	SpuID        int64                 `gorm:"column:spu_id;type:int(10) unsigned;not null;comment:商品id" json:"spu_id"`                  // 商品id
	SkuID        int64                 `gorm:"column:sku_id;type:int(10) unsigned;not null;comment:套餐id" json:"sku_id"`                  // 套餐id
	Price        decimal.Decimal       `gorm:"column:price;type:decimal(10,4) unsigned;not null;default:0.0000;comment:价格" json:"price"` // 价格
	ChangedStock int64                 `gorm:"column:changed_stock;type:int(10) unsigned;not null;comment:本次变更数量" json:"changed_stock"`  // 本次变更数量
	State        int64                 `gorm:"column:state;type:tinyint(3) unsigned;not null;comment:状态" json:"state"`                   // 状态
	Type         int64                 `gorm:"column:type;type:tinyint(3) unsigned;not null;comment:流水类型" json:"type"`                   // 流水类型
	Stock        int64                 `gorm:"column:stock;type:int(10) unsigned;not null;comment:变更前库存" json:"stock"`                   // 变更前库存
	Identifier   string                `gorm:"column:identifier;type:varchar(128);not null;comment:幂等号" json:"identifier"`               // 幂等号
	ExtendInfo   string                `gorm:"column:extend_info;type:varchar(255);not null;comment:拓展信息" json:"extend_info"`            // 拓展信息
	CreatedBy    int64                 `gorm:"column:created_by;type:int(10) unsigned;not null;comment:创建时间" json:"created_by"`          // 创建时间
	UpdatedBy    int64                 `gorm:"column:updated_by;type:int(10) unsigned;not null;comment:修改时间" json:"updated_by"`          // 修改时间
	CreatedAt    int64                 `gorm:"column:created_at;type:bigint(20) unsigned;not null;comment:创建时间" json:"created_at"`       // 创建时间
	UpdatedAt    int64                 `gorm:"column:updated_at;type:bigint(20) unsigned;not null;comment:修改时间" json:"updated_at"`       // 修改时间
	DeletedAt    soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint(3) unsigned;not null;comment:是否删除" json:"-"`                // 是否删除
}

// TableName SpuStockStream's table name
func (*SpuStockStream) TableName() string {
	return TableNameSpuStockStream
}
