// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameSpu = "spu"

// Spu 商品表
type Spu struct {
	ID          int64                 `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true;comment:主键" json:"id"`                           // 主键
	Name        string                `gorm:"column:name;type:varchar(64);not null;comment:商品名称" json:"name"`                                      // 商品名称
	Title       string                `gorm:"column:title;type:varchar(255);not null;comment:标题" json:"title"`                                     // 标题
	Description string                `gorm:"column:description;type:varchar(255);not null;comment:商品简介" json:"description"`                       // 商品简介
	ServerID    int64                 `gorm:"column:server_id;type:int(10) unsigned;not null;comment:商家id" json:"server_id"`                       // 商家id
	CategoryID  int64                 `gorm:"column:category_id;type:int(10) unsigned;not null;comment:分类id" json:"category_id"`                   // 分类id
	BrandID     int64                 `gorm:"column:brand_id;type:int(10) unsigned;not null;comment:品牌id" json:"brand_id"`                         // 品牌id
	ModelID     int64                 `gorm:"column:model_id;type:int(10) unsigned;not null;comment:型号id" json:"model_id"`                         // 型号id
	Keyword     string                `gorm:"column:keyword;type:varchar(255);not null;comment:关键词" json:"keyword"`                                // 关键词
	Status      int64                 `gorm:"column:status;type:tinyint(3) unsigned;not null;comment:状态(1:上架中,2:待审核,3:商家下架,4:系统下架)" json:"status"` // 状态(1:上架中,2:待审核,3:商家下架,4:系统下架)
	Sort        int64                 `gorm:"column:sort;type:int(10) unsigned;not null;comment:排序" json:"sort"`                                   // 排序
	Sign        int64                 `gorm:"column:sign;type:int(10) unsigned;not null;comment:spu标志（位运算）" json:"sign"`                           // spu标志（位运算）
	Scene       int64                 `gorm:"column:scene;type:int(10) unsigned;not null;comment:场景（位运算）" json:"scene"`                            // 场景（位运算）
	UnitName    string                `gorm:"column:unit_name;type:varchar(255);not null;comment:单位名称" json:"unit_name"`                           // 单位名称
	Cover       string                `gorm:"column:cover;type:varchar(255);not null;comment:封面图" json:"cover"`                                    // 封面图
	DownReason  string                `gorm:"column:down_reason;type:varchar(255);not null;comment:下架原因" json:"down_reason"`                       // 下架原因
	Version     int64                 `gorm:"column:version;type:int(10) unsigned;not null;comment:版本" json:"version"`                             // 版本
	CreatedAt   int64                 `gorm:"column:created_at;type:bigint(20) unsigned;not null;comment:创建时间" json:"created_at"`                  // 创建时间
	UpdatedAt   int64                 `gorm:"column:updated_at;type:bigint(20) unsigned;not null;comment:修改时间" json:"updated_at"`                  // 修改时间
	DeletedAt   soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint(3) unsigned;not null;comment:是否删除" json:"-"`                           // 是否删除
	CreatedBy   int64                 `gorm:"column:created_by;type:int(10) unsigned;not null;comment:创建时间" json:"created_by"`                     // 创建时间
	UpdatedBy   int64                 `gorm:"column:updated_by;type:int(10) unsigned;not null;comment:修改时间" json:"updated_by"`                     // 修改时间
}

// TableName Spu's table name
func (*Spu) TableName() string {
	return TableNameSpu
}
