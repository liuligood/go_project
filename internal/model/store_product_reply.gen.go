// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameStoreProductReply = "store_product_reply"

// StoreProductReply 评论表
type StoreProductReply struct {
	ID                   int64                 `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:评论ID" json:"id"`                                 // 评论ID
	ProductID            int64                 `gorm:"column:product_id;type:int unsigned;not null;comment:商品id" json:"product_id"`                             // 商品id
	UID                  int64                 `gorm:"column:uid;type:int unsigned;not null;comment:用户ID" json:"uid"`                                           // 用户ID
	Oid                  int64                 `gorm:"column:oid;type:int unsigned;not null;comment:订单ID" json:"oid"`                                           // 订单ID
	Unique               string                `gorm:"column:unique;type:varchar(32);not null;comment:商品唯一id" json:"unique"`                                    // 商品唯一id
	ReplyType            string                `gorm:"column:reply_type;type:varchar(32);not null;default:product;comment:某种商品类型(普通商品、秒杀商品）" json:"reply_type"` // 某种商品类型(普通商品、秒杀商品）
	ProductScore         int64                 `gorm:"column:product_score;type:tinyint unsigned;not null;comment:商品分数" json:"product_score"`                   // 商品分数
	ServiceScore         int64                 `gorm:"column:service_score;type:tinyint unsigned;not null;comment:服务分数" json:"service_score"`                   // 服务分数
	Comment              string                `gorm:"column:comment;type:varchar(512);not null;comment:评论内容" json:"comment"`                                   // 评论内容
	Pics                 string                `gorm:"column:pics;type:text;not null;comment:评论图片" json:"pics"`                                                 // 评论图片
	MerchantReplyContent string                `gorm:"column:merchant_reply_content;type:varchar(300);not null;comment:管理员回复内容" json:"merchant_reply_content"`  // 管理员回复内容
	MerchantReplyTime    int64                 `gorm:"column:merchant_reply_time;type:int unsigned;not null;comment:管理员回复时间" json:"merchant_reply_time"`        // 管理员回复时间
	IsReply              int64                 `gorm:"column:is_reply;type:tinyint unsigned;not null;comment:0未回复1已回复" json:"is_reply"`                         // 0未回复1已回复
	Nickname             string                `gorm:"column:nickname;type:varchar(64);not null;comment:用户名称" json:"nickname"`                                  // 用户名称
	Avatar               string                `gorm:"column:avatar;type:varchar(255);not null;comment:用户头像" json:"avatar"`                                     // 用户头像
	Sku                  string                `gorm:"column:sku;type:varchar(128);not null;comment:商品规格属性值,多个,号隔开" json:"sku"`                                 // 商品规格属性值,多个,号隔开
	CreatedAt            int64                 `gorm:"column:created_at;type:bigint unsigned;not null;comment:创建时间" json:"created_at"`                          // 创建时间
	UpdatedAt            int64                 `gorm:"column:updated_at;type:bigint unsigned;not null;comment:修改时间" json:"updated_at"`                          // 修改时间
	DeletedAt            soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint unsigned;not null;comment:是否删除" json:"-"`                                  // 是否删除
}

// TableName StoreProductReply's table name
func (*StoreProductReply) TableName() string {
	return TableNameStoreProductReply
}
