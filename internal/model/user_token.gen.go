// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNameUserToken = "user_token"

// UserToken mapped from table <user_token>
type UserToken struct {
	ID          int64                 `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`
	UID         int64                 `gorm:"column:uid;type:int(10) unsigned;not null;comment:用户 id" json:"uid"`                                             // 用户 id
	Token       string                `gorm:"column:token;type:varchar(255);not null;comment:token" json:"token"`                                             // token
	Type        int64                 `gorm:"column:type;type:tinyint(1);default:1;comment:类型，1 公众号， 2 小程序, 3 unionid, 5AppIos,6AppAndroid,7ios" json:"type"` // 类型，1 公众号， 2 小程序, 3 unionid, 5AppIos,6AppAndroid,7ios
	ExpiresTime time.Time             `gorm:"column:expires_time;type:datetime;comment:到期时间" json:"expires_time"`                                             // 到期时间
	LoginIP     string                `gorm:"column:login_ip;type:varchar(32);comment:登录ip" json:"login_ip"`                                                  // 登录ip
	CreatedAt   int64                 `gorm:"column:created_at;type:bigint(20) unsigned;not null;comment:创建时间" json:"created_at"`                             // 创建时间
	UpdatedAt   int64                 `gorm:"column:updated_at;type:bigint(20) unsigned;not null;comment:修改时间" json:"updated_at"`                             // 修改时间
	DeletedAt   soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint(3) unsigned;not null;comment:是否删除" json:"-"`                                      // 是否删除
}

// TableName UserToken's table name
func (*UserToken) TableName() string {
	return TableNameUserToken
}
