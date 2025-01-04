package validation

import (
	"github.com/shopspring/decimal"
)

type CreateProduct struct {
	ID            int64       `json:"id"`                              // 商品id
	Image         string      `json:"image" binding:"required"`        // 商品图片
	SliderImages  []string    `json:"sliderImages" binding:"required"` // 轮播图
	SliderImage   string      `json:"sliderImage" binding:"required"`  // 轮播图
	StoreName     string      `json:"storeName" binding:"required"`    // 商品名称
	StoreInfo     string      `json:"storeInfo" binding:"required"`    // 商品简介
	Keyword       string      `json:"keyword" binding:"required"`      // 关键字
	CateID        string      `json:"cateId" binding:"required"`       // 分类id 逗号隔开
	UnitName      string      `json:"unitName" binding:"required"`     // 单位名
	Sort          int64       `json:"sort"`                            // 排序
	IsHot         bool        `json:"isHot"`                           // 是否热卖
	IsBenefit     bool        `json:"isBenefit"`                       // 是否优惠
	IsBest        bool        `json:"isBest"`                          // 是否精品
	IsNew         bool        `json:"isNew"`                           // 是否新品
	IsGood        bool        `json:"isGood"`                          // 是否优品推荐
	GiveIntegral  int64       `json:"giveIntegral"`                    // 获得积分
	isSub         bool        `json:"isSub" binding:"required"`        // 是否单独分佣不能为空
	Ficti         int64       `json:"ficti"`                           // 虚拟销量
	TempID        int64       `json:"tempId" binding:"required"`       // 运费模板ID
	SpecType      bool        `json:"specType"`                        // 规格 0单 1多
	Activity      []string    `json:"activity"`                        // 活动显示排序 0=默认，1=秒杀，2=砍价，3=拼团
	AttrList      []Attr      `json:"Attr"`                            // 商品属性
	AttrValueList []AttrValue `json:"attrValue"`                       // 商品属性详情
	Context       string      `json:"context"`                         // 商品描述
	CouponIds     []int64     `json:"couponIds"`                       // 优惠卷id集合
	FlatPattern   string      `json:"flatPattern"`                     // 展示图
}

type AttrValue struct {
	ProductId    int64           `json:"productId" default:"0" binding:"required"` // 商品ID|添加时为0，修改时为商品id
	Stock        int             `json:"stock" binding:"required"`                 // 库存
	Sku          string          `json:"sku"`                                      // sku|活动商品必传
	Price        decimal.Decimal `json:"price" binding:"required"`                 // 规格属性金额
	Image        string          `json:"image" binding:"required"`                 // 商品规格属性图片
	Cost         decimal.Decimal `json:"cost" binding:"required"`                  // 成本价
	OtPrice      decimal.Decimal `json:"otPrice" binding:"required"`               // 原价
	Weight       decimal.Decimal `json:"weight" binding:"required"`                // 重量
	Volume       decimal.Decimal `json:"volume" binding:"required"`                // 体积
	Brokerage    decimal.Decimal `json:"brokerage" binding:"required"`             // 一级返佣
	BrokerageTwo decimal.Decimal `json:"brokerageTwo" binding:"required"`          // 二级返佣
	AttrValue    string          `json:"attrValue"`                                // 属性对应 {"尺码":"2XL","颜色":"DX027白色"}
	BarCode      string          `json:"barCode"`                                  // 商品条码
	Quota        int64           `json:"quota"`                                    // 活动限购数量|活动商品专用字段
	QuotaShow    int64           `json:"quotaShow"`                                // 活动限购数量显示|活动商品专用字段,添加时不传
	MinPrice     decimal.Decimal `json:"minPrice"`                                 // 砍价商品最低价|砍价专用
}

type Attr struct {
	AttrName   string `json:"attrName"`
	AttrValues string `json:"attrValues"`
	ID         int    `json:"id"`
}
