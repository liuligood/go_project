// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/plugin/soft_delete"
)

const TableNameStoreProduct = "store_product"

// StoreProduct 商品表
type StoreProduct struct {
	ID           int64                 `gorm:"column:id;type:mediumint;primaryKey;autoIncrement:true;comment:商品id" json:"id"`                   // 商品id
	MerID        int64                 `gorm:"column:mer_id;type:int unsigned;not null;comment:商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)" json:"mer_id"`   // 商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)
	Image        string                `gorm:"column:image;type:varchar(256);not null;comment:商品图片" json:"image"`                               // 商品图片
	SliderImage  string                `gorm:"column:slider_image;type:varchar(2000);not null;comment:轮播图" json:"slider_image"`                 // 轮播图
	StoreName    string                `gorm:"column:store_name;type:varchar(128);not null;comment:商品名称" json:"store_name"`                     // 商品名称
	StoreInfo    string                `gorm:"column:store_info;type:varchar(256);not null;comment:商品简介" json:"store_info"`                     // 商品简介
	Keyword      string                `gorm:"column:keyword;type:varchar(256);not null;comment:关键字" json:"keyword"`                            // 关键字
	BarCode      string                `gorm:"column:bar_code;type:varchar(15);not null;comment:商品条码（一维码）" json:"bar_code"`                     // 商品条码（一维码）
	CateID       string                `gorm:"column:cate_id;type:varchar(64);not null;comment:分类id" json:"cate_id"`                            // 分类id
	Price        decimal.Decimal       `gorm:"column:price;type:decimal(8,2) unsigned;not null;default:0.00;comment:商品价格" json:"price"`         // 商品价格
	VipPrice     decimal.Decimal       `gorm:"column:vip_price;type:decimal(8,2) unsigned;not null;default:0.00;comment:会员价格" json:"vip_price"` // 会员价格
	OtPrice      decimal.Decimal       `gorm:"column:ot_price;type:decimal(8,2) unsigned;not null;default:0.00;comment:市场价" json:"ot_price"`    // 市场价
	Postage      decimal.Decimal       `gorm:"column:postage;type:decimal(8,2) unsigned;not null;default:0.00;comment:邮费" json:"postage"`       // 邮费
	UnitName     string                `gorm:"column:unit_name;type:varchar(32);not null;comment:单位名" json:"unit_name"`                         // 单位名
	Sort         int64                 `gorm:"column:sort;type:smallint;not null;comment:排序" json:"sort"`                                       // 排序
	Sales        int64                 `gorm:"column:sales;type:mediumint unsigned;not null;comment:销量" json:"sales"`                           // 销量
	Stock        int64                 `gorm:"column:stock;type:mediumint unsigned;not null;comment:库存" json:"stock"`                           // 库存
	IsShow       int64                 `gorm:"column:is_show;type:tinyint(1);not null;default:1;comment:状态（0：未上架，1：上架）" json:"is_show"`         // 状态（0：未上架，1：上架）
	IsHot        int64                 `gorm:"column:is_hot;type:tinyint(1);not null;comment:是否热卖" json:"is_hot"`                               // 是否热卖
	IsBenefit    int64                 `gorm:"column:is_benefit;type:tinyint(1);not null;comment:是否优惠" json:"is_benefit"`                       // 是否优惠
	IsBest       int64                 `gorm:"column:is_best;type:tinyint(1);not null;comment:是否精品" json:"is_best"`                             // 是否精品
	IsNew        int64                 `gorm:"column:is_new;type:tinyint(1);not null;comment:是否新品" json:"is_new"`                               // 是否新品
	AddTime      int64                 `gorm:"column:add_time;type:int unsigned;not null;comment:添加时间" json:"add_time"`                         // 添加时间
	IsPostage    int64                 `gorm:"column:is_postage;type:tinyint unsigned;not null;comment:是否包邮" json:"is_postage"`                 // 是否包邮
	IsDel        int64                 `gorm:"column:is_del;type:tinyint unsigned;not null;comment:是否删除" json:"is_del"`                         // 是否删除
	MerUse       int64                 `gorm:"column:mer_use;type:tinyint unsigned;not null;comment:商户是否代理 0不可代理1可代理" json:"mer_use"`           // 商户是否代理 0不可代理1可代理
	GiveIntegral int64                 `gorm:"column:give_integral;type:int;comment:获得积分" json:"give_integral"`                                 // 获得积分
	Cost         decimal.Decimal       `gorm:"column:cost;type:decimal(8,2) unsigned;not null;default:0.00;comment:成本价" json:"cost"`            // 成本价
	IsSeckill    int64                 `gorm:"column:is_seckill;type:tinyint unsigned;not null;comment:秒杀状态 0 未开启 1已开启" json:"is_seckill"`      // 秒杀状态 0 未开启 1已开启
	IsBargain    int64                 `gorm:"column:is_bargain;type:tinyint unsigned;comment:砍价状态 0未开启 1开启" json:"is_bargain"`                 // 砍价状态 0未开启 1开启
	IsGood       int64                 `gorm:"column:is_good;type:tinyint(1);not null;comment:是否优品推荐" json:"is_good"`                           // 是否优品推荐
	IsSub        int64                 `gorm:"column:is_sub;type:tinyint(1);not null;comment:是否单独分佣" json:"is_sub"`                             // 是否单独分佣
	Ficti        int64                 `gorm:"column:ficti;type:mediumint;default:100;comment:虚拟销量" json:"ficti"`                               // 虚拟销量
	Browse       int64                 `gorm:"column:browse;type:int;comment:浏览量" json:"browse"`                                                // 浏览量
	CodePath     string                `gorm:"column:code_path;type:varchar(64);not null;comment:商品二维码地址(用户小程序海报)" json:"code_path"`            // 商品二维码地址(用户小程序海报)
	SoureLink    string                `gorm:"column:soure_link;type:varchar(255);comment:淘宝京东1688类型" json:"soure_link"`                        // 淘宝京东1688类型
	VideoLink    string                `gorm:"column:video_link;type:varchar(200);not null;comment:主图视频链接" json:"video_link"`                   // 主图视频链接
	TempID       int64                 `gorm:"column:temp_id;type:int;not null;default:1;comment:运费模板ID" json:"temp_id"`                        // 运费模板ID
	SpecType     int64                 `gorm:"column:spec_type;type:tinyint(1);not null;comment:规格 0单 1多" json:"spec_type"`                     // 规格 0单 1多
	Activity     string                `gorm:"column:activity;type:varchar(255);not null;comment:活动显示排序0=默认, 1=秒杀，2=砍价，3=拼团" json:"activity"`   // 活动显示排序0=默认, 1=秒杀，2=砍价，3=拼团
	FlatPattern  string                `gorm:"column:flat_pattern;type:varchar(1000);not null;comment:展示图" json:"flat_pattern"`                 // 展示图
	IsRecycle    int64                 `gorm:"column:is_recycle;type:tinyint(1);not null;comment:是否回收站" json:"is_recycle"`                      // 是否回收站
	CreatedAt    int64                 `gorm:"column:created_at;type:bigint unsigned;not null;comment:创建时间" json:"created_at"`                  // 创建时间
	UpdatedAt    int64                 `gorm:"column:updated_at;type:bigint unsigned;not null;comment:修改时间" json:"updated_at"`                  // 修改时间
	DeletedAt    soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint unsigned;not null;comment:是否删除" json:"-"`                          // 是否删除
}

// TableName StoreProduct's table name
func (*StoreProduct) TableName() string {
	return TableNameStoreProduct
}
