package validation

import (
	"github.com/shopspring/decimal"
)

type CreateProductParam struct {
	ID               int64          `json:"id"`                 // 主键
	Name             string         `json:"name"`               // 商品名称
	Title            string         `json:"title"`              // 标题
	Description      string         `json:"description"`        // 商品简介
	ServerID         int64          `json:"server_id"`          // 商家id
	CategoryID       int64          `json:"category_id"`        // 分类id
	BrandID          int64          `json:"brand_id"`           // 品牌id
	ModelID          int64          `json:"model_id"`           // 型号id
	Keyword          string         `json:"keyword"`            // 关键词
	Status           int64          `json:"status"`             // 状态(1:上架中,2:待审核,3:商家下架,4:系统下架)
	Sort             int64          `json:"sort"`               // 排序
	Sign             int64          `json:"sign"`               // spu标志（位运算）
	Scene            int64          `json:"scene"`              // 场景（位运算）
	UnitName         string         `json:"unit_name"`          // 单位名称
	Cover            string         `json:"cover"`              // 封面图
	DownReason       string         `json:"down_reason"`        // 下架原因
	Version          int64          `json:"version"`            // 版本
	SkuList          []Sku          `json:"sku_list"`           // 套餐信息
	SpuPictureList   []SpuPicture   `json:"spu_picture_list"`   // 商品图片信息
	SpuServiceList   []SpuService   `json:"spu_service_list"`   // 商品服务信息
	SpuLogisticsList []SpuLogistics `json:"spu_logistics_list"` // 商品物流信息

}

type Sku struct {
	ID               int64           `json:"id"`                 // 主键
	SpuID            int64           `json:"spu_id"`             // 商品id
	Stock            int64           `json:"stock"`              // 库存
	Price            decimal.Decimal `json:"price"`              // 价格
	Status           int64           `json:"status"`             // 状态(1:关闭,2:正常)
	Sign             int64           `json:"sign"`               // 标志(位运算)
	Cost             decimal.Decimal `json:"cost"`               // 成本价
	Weight           decimal.Decimal `json:"weight"`             // 重量
	Volume           decimal.Decimal `json:"volume"`             // 体积
	SkuAttributeList []SkuAttribute  `json:"sku_attribute_list"` // 套餐属性
}

type SkuAttribute struct {
	ID    int64  `json:"id"` // 主键
	SkuID int64  `json:"sku_id"`
	Type  int64  `json:"type"`  // 类型
	Title string `json:"title"` // 属性标题
	Value string `json:"value"` // 属性值
	Sort  int64  `json:"sort"`  // 排序
	Sign  int64  `json:"sign"`  // 标志(位运算)
}

type SpuPicture struct {
	ID      int64  `json:"id"`      // 主键
	SpuID   int64  `json:"spu_id"`  // 商品id
	Type    int64  `json:"type"`    // 图片类型(1:封面,2:视频)
	Picture string `json:"picture"` // 图片/视频
	Sort    int64  `json:"sort"`    // 排序
}

type SpuService struct {
	ID        int64  `json:"id"`         // 主键
	SpuID     int64  `json:"spu_id"`     // 商品id
	ServiceID int64  `json:"service_id"` // 服务id
	Type      int64  `json:"type"`       // 类型
	Name      string `json:"name"`       // 服务名称
}

type SpuLogistics struct {
	ID              int64           `json:"id"`                // 主键
	SpuID           int64           `json:"spu_id"`            // 商品id
	Type            int64           `json:"type"`              // 类型(1:寄出,2:寄回)
	Province        string          `json:"province"`          // 省
	City            string          `json:"city"`              // 市
	Region          string          `json:"region"`            // 区/县
	PayType         int64           ` json:"pay_type"`         // 支付类型(1-包邮,2-到付,3-运费模版)
	Price           decimal.Decimal `json:"price"`             // 价格
	TempID          int64           `json:"temp_id"`           // 运费模版id
	ExpressWay      int64           `json:"express_way"`       // 快递方式
	ServerAddressID int64           `json:"server_address_id"` // 商家地址id
	AddService      int64           `json:"add_service"`       // 增值服务(位运算)
}
