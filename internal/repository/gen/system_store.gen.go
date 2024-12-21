// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"crmeb_go/internal/model"
)

func newSystemStore(db *gorm.DB, opts ...gen.DOOption) systemStore {
	_systemStore := systemStore{}

	_systemStore.systemStoreDo.UseDB(db, opts...)
	_systemStore.systemStoreDo.UseModel(&model.SystemStore{})

	tableName := _systemStore.systemStoreDo.TableName()
	_systemStore.ALL = field.NewAsterisk(tableName)
	_systemStore.ID = field.NewInt64(tableName, "id")
	_systemStore.Name = field.NewString(tableName, "name")
	_systemStore.Introduction = field.NewString(tableName, "introduction")
	_systemStore.Phone = field.NewString(tableName, "phone")
	_systemStore.Address = field.NewString(tableName, "address")
	_systemStore.DetailedAddress = field.NewString(tableName, "detailed_address")
	_systemStore.Image = field.NewString(tableName, "image")
	_systemStore.Latitude = field.NewString(tableName, "latitude")
	_systemStore.Longitude = field.NewString(tableName, "longitude")
	_systemStore.ValidTime = field.NewInt64(tableName, "valid_time")
	_systemStore.DayTime = field.NewInt64(tableName, "day_time")
	_systemStore.IsShow = field.NewInt64(tableName, "is_show")
	_systemStore.CreatedAt = field.NewInt64(tableName, "created_at")
	_systemStore.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_systemStore.DeletedAt = field.NewField(tableName, "deleted_at")

	_systemStore.fillFieldMap()

	return _systemStore
}

// systemStore 门店自提
type systemStore struct {
	systemStoreDo systemStoreDo

	ALL             field.Asterisk
	ID              field.Int64
	Name            field.String // 门店名称
	Introduction    field.String // 简介
	Phone           field.String // 手机号码
	Address         field.String // 省市区
	DetailedAddress field.String // 详细地址
	Image           field.String // 门店logo
	Latitude        field.String // 纬度
	Longitude       field.String // 经度
	ValidTime       field.Int64  // 核销有效日期
	DayTime         field.Int64  // 每日营业开关时间
	IsShow          field.Int64  // 是否显示
	CreatedAt       field.Int64  // 创建时间
	UpdatedAt       field.Int64  // 修改时间
	DeletedAt       field.Field  // 是否删除

	fieldMap map[string]field.Expr
}

func (s systemStore) Table(newTableName string) *systemStore {
	s.systemStoreDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s systemStore) As(alias string) *systemStore {
	s.systemStoreDo.DO = *(s.systemStoreDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *systemStore) updateTableName(table string) *systemStore {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Name = field.NewString(table, "name")
	s.Introduction = field.NewString(table, "introduction")
	s.Phone = field.NewString(table, "phone")
	s.Address = field.NewString(table, "address")
	s.DetailedAddress = field.NewString(table, "detailed_address")
	s.Image = field.NewString(table, "image")
	s.Latitude = field.NewString(table, "latitude")
	s.Longitude = field.NewString(table, "longitude")
	s.ValidTime = field.NewInt64(table, "valid_time")
	s.DayTime = field.NewInt64(table, "day_time")
	s.IsShow = field.NewInt64(table, "is_show")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *systemStore) WithContext(ctx context.Context) ISystemStoreDo {
	return s.systemStoreDo.WithContext(ctx)
}

func (s systemStore) TableName() string { return s.systemStoreDo.TableName() }

func (s systemStore) Alias() string { return s.systemStoreDo.Alias() }

func (s systemStore) Columns(cols ...field.Expr) gen.Columns { return s.systemStoreDo.Columns(cols...) }

func (s *systemStore) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *systemStore) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 15)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["introduction"] = s.Introduction
	s.fieldMap["phone"] = s.Phone
	s.fieldMap["address"] = s.Address
	s.fieldMap["detailed_address"] = s.DetailedAddress
	s.fieldMap["image"] = s.Image
	s.fieldMap["latitude"] = s.Latitude
	s.fieldMap["longitude"] = s.Longitude
	s.fieldMap["valid_time"] = s.ValidTime
	s.fieldMap["day_time"] = s.DayTime
	s.fieldMap["is_show"] = s.IsShow
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s systemStore) clone(db *gorm.DB) systemStore {
	s.systemStoreDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s systemStore) replaceDB(db *gorm.DB) systemStore {
	s.systemStoreDo.ReplaceDB(db)
	return s
}

type systemStoreDo struct{ gen.DO }

type ISystemStoreDo interface {
	gen.SubQuery
	Debug() ISystemStoreDo
	WithContext(ctx context.Context) ISystemStoreDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISystemStoreDo
	WriteDB() ISystemStoreDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISystemStoreDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISystemStoreDo
	Not(conds ...gen.Condition) ISystemStoreDo
	Or(conds ...gen.Condition) ISystemStoreDo
	Select(conds ...field.Expr) ISystemStoreDo
	Where(conds ...gen.Condition) ISystemStoreDo
	Order(conds ...field.Expr) ISystemStoreDo
	Distinct(cols ...field.Expr) ISystemStoreDo
	Omit(cols ...field.Expr) ISystemStoreDo
	Join(table schema.Tabler, on ...field.Expr) ISystemStoreDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISystemStoreDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISystemStoreDo
	Group(cols ...field.Expr) ISystemStoreDo
	Having(conds ...gen.Condition) ISystemStoreDo
	Limit(limit int) ISystemStoreDo
	Offset(offset int) ISystemStoreDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemStoreDo
	Unscoped() ISystemStoreDo
	Create(values ...*model.SystemStore) error
	CreateInBatches(values []*model.SystemStore, batchSize int) error
	Save(values ...*model.SystemStore) error
	First() (*model.SystemStore, error)
	Take() (*model.SystemStore, error)
	Last() (*model.SystemStore, error)
	Find() ([]*model.SystemStore, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemStore, err error)
	FindInBatches(result *[]*model.SystemStore, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SystemStore) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISystemStoreDo
	Assign(attrs ...field.AssignExpr) ISystemStoreDo
	Joins(fields ...field.RelationField) ISystemStoreDo
	Preload(fields ...field.RelationField) ISystemStoreDo
	FirstOrInit() (*model.SystemStore, error)
	FirstOrCreate() (*model.SystemStore, error)
	FindByPage(offset int, limit int) (result []*model.SystemStore, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISystemStoreDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s systemStoreDo) Debug() ISystemStoreDo {
	return s.withDO(s.DO.Debug())
}

func (s systemStoreDo) WithContext(ctx context.Context) ISystemStoreDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s systemStoreDo) ReadDB() ISystemStoreDo {
	return s.Clauses(dbresolver.Read)
}

func (s systemStoreDo) WriteDB() ISystemStoreDo {
	return s.Clauses(dbresolver.Write)
}

func (s systemStoreDo) Session(config *gorm.Session) ISystemStoreDo {
	return s.withDO(s.DO.Session(config))
}

func (s systemStoreDo) Clauses(conds ...clause.Expression) ISystemStoreDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s systemStoreDo) Returning(value interface{}, columns ...string) ISystemStoreDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s systemStoreDo) Not(conds ...gen.Condition) ISystemStoreDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s systemStoreDo) Or(conds ...gen.Condition) ISystemStoreDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s systemStoreDo) Select(conds ...field.Expr) ISystemStoreDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s systemStoreDo) Where(conds ...gen.Condition) ISystemStoreDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s systemStoreDo) Order(conds ...field.Expr) ISystemStoreDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s systemStoreDo) Distinct(cols ...field.Expr) ISystemStoreDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s systemStoreDo) Omit(cols ...field.Expr) ISystemStoreDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s systemStoreDo) Join(table schema.Tabler, on ...field.Expr) ISystemStoreDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s systemStoreDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISystemStoreDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s systemStoreDo) RightJoin(table schema.Tabler, on ...field.Expr) ISystemStoreDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s systemStoreDo) Group(cols ...field.Expr) ISystemStoreDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s systemStoreDo) Having(conds ...gen.Condition) ISystemStoreDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s systemStoreDo) Limit(limit int) ISystemStoreDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s systemStoreDo) Offset(offset int) ISystemStoreDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s systemStoreDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemStoreDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s systemStoreDo) Unscoped() ISystemStoreDo {
	return s.withDO(s.DO.Unscoped())
}

func (s systemStoreDo) Create(values ...*model.SystemStore) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s systemStoreDo) CreateInBatches(values []*model.SystemStore, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s systemStoreDo) Save(values ...*model.SystemStore) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s systemStoreDo) First() (*model.SystemStore, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemStore), nil
	}
}

func (s systemStoreDo) Take() (*model.SystemStore, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemStore), nil
	}
}

func (s systemStoreDo) Last() (*model.SystemStore, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemStore), nil
	}
}

func (s systemStoreDo) Find() ([]*model.SystemStore, error) {
	result, err := s.DO.Find()
	return result.([]*model.SystemStore), err
}

func (s systemStoreDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemStore, err error) {
	buf := make([]*model.SystemStore, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s systemStoreDo) FindInBatches(result *[]*model.SystemStore, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s systemStoreDo) Attrs(attrs ...field.AssignExpr) ISystemStoreDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s systemStoreDo) Assign(attrs ...field.AssignExpr) ISystemStoreDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s systemStoreDo) Joins(fields ...field.RelationField) ISystemStoreDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s systemStoreDo) Preload(fields ...field.RelationField) ISystemStoreDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s systemStoreDo) FirstOrInit() (*model.SystemStore, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemStore), nil
	}
}

func (s systemStoreDo) FirstOrCreate() (*model.SystemStore, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemStore), nil
	}
}

func (s systemStoreDo) FindByPage(offset int, limit int) (result []*model.SystemStore, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s systemStoreDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s systemStoreDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s systemStoreDo) Delete(models ...*model.SystemStore) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *systemStoreDo) withDO(do gen.Dao) *systemStoreDo {
	s.DO = *do.(*gen.DO)
	return s
}
