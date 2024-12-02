package model

import (
	"gorm.io/gorm"
)

// DBModelImpl 数据库模型接口
type DBModelImpl interface {
	DB() *gorm.DB
	Connection() string
	Table() string
	Model() *gorm.DB
}

func NewDBModel(db *gorm.DB, database, tableName string) *DBModel {
	return &DBModel{*NewBaseModel(db, database, tableName)}
}

type DBModel struct {
	BaseModel
}

func (m *DBModel) Model() *gorm.DB {
	return m.DB().Table(m.Table())
}
