package model

import (
	"gorm.io/gorm"
)

// BaseModel 基础模型
type BaseModel struct {
	db        *gorm.DB
	database  string
	tableName string
}

func NewBaseModel(db *gorm.DB, database, tableName string) *BaseModel {
	return &BaseModel{db, database, tableName}
}

func (m *BaseModel) DB() *gorm.DB {
	return m.db
}

func (m *BaseModel) Connection() string {
	return m.database
}

func (m *BaseModel) Table() string {
	return m.tableName
}
