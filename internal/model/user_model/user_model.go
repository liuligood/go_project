package user_model

import (
	"crmeb_go/internal/model"
	"gorm.io/gorm"
)

type UserModel struct {
	model.DBModel
	Id uint64 `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT;comment:主键" json:"id"`
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{DBModel: *model.NewDBModel(db, "crmeb", "eb_user")}
}
