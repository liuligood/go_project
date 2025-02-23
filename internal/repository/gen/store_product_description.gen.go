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

func newStoreProductDescription(db *gorm.DB, opts ...gen.DOOption) storeProductDescription {
	_storeProductDescription := storeProductDescription{}

	_storeProductDescription.storeProductDescriptionDo.UseDB(db, opts...)
	_storeProductDescription.storeProductDescriptionDo.UseModel(&model.StoreProductDescription{})

	tableName := _storeProductDescription.storeProductDescriptionDo.TableName()
	_storeProductDescription.ALL = field.NewAsterisk(tableName)
	_storeProductDescription.ID = field.NewInt64(tableName, "id")
	_storeProductDescription.ProductID = field.NewInt64(tableName, "product_id")
	_storeProductDescription.Description = field.NewString(tableName, "description")
	_storeProductDescription.Type = field.NewInt64(tableName, "type")
	_storeProductDescription.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeProductDescription.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeProductDescription.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeProductDescription.fillFieldMap()

	return _storeProductDescription
}

// storeProductDescription 商品描述表
type storeProductDescription struct {
	storeProductDescriptionDo storeProductDescriptionDo

	ALL         field.Asterisk
	ID          field.Int64
	ProductID   field.Int64  // 商品ID
	Description field.String // 商品详情
	Type        field.Int64  // 商品类型
	CreatedAt   field.Int64  // 创建时间
	UpdatedAt   field.Int64  // 修改时间
	DeletedAt   field.Field  // 是否删除

	fieldMap map[string]field.Expr
}

func (s storeProductDescription) Table(newTableName string) *storeProductDescription {
	s.storeProductDescriptionDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeProductDescription) As(alias string) *storeProductDescription {
	s.storeProductDescriptionDo.DO = *(s.storeProductDescriptionDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeProductDescription) updateTableName(table string) *storeProductDescription {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.ProductID = field.NewInt64(table, "product_id")
	s.Description = field.NewString(table, "description")
	s.Type = field.NewInt64(table, "type")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeProductDescription) WithContext(ctx context.Context) IStoreProductDescriptionDo {
	return s.storeProductDescriptionDo.WithContext(ctx)
}

func (s storeProductDescription) TableName() string { return s.storeProductDescriptionDo.TableName() }

func (s storeProductDescription) Alias() string { return s.storeProductDescriptionDo.Alias() }

func (s storeProductDescription) Columns(cols ...field.Expr) gen.Columns {
	return s.storeProductDescriptionDo.Columns(cols...)
}

func (s *storeProductDescription) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeProductDescription) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["id"] = s.ID
	s.fieldMap["product_id"] = s.ProductID
	s.fieldMap["description"] = s.Description
	s.fieldMap["type"] = s.Type
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeProductDescription) clone(db *gorm.DB) storeProductDescription {
	s.storeProductDescriptionDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeProductDescription) replaceDB(db *gorm.DB) storeProductDescription {
	s.storeProductDescriptionDo.ReplaceDB(db)
	return s
}

type storeProductDescriptionDo struct{ gen.DO }

type IStoreProductDescriptionDo interface {
	gen.SubQuery
	Debug() IStoreProductDescriptionDo
	WithContext(ctx context.Context) IStoreProductDescriptionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreProductDescriptionDo
	WriteDB() IStoreProductDescriptionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreProductDescriptionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreProductDescriptionDo
	Not(conds ...gen.Condition) IStoreProductDescriptionDo
	Or(conds ...gen.Condition) IStoreProductDescriptionDo
	Select(conds ...field.Expr) IStoreProductDescriptionDo
	Where(conds ...gen.Condition) IStoreProductDescriptionDo
	Order(conds ...field.Expr) IStoreProductDescriptionDo
	Distinct(cols ...field.Expr) IStoreProductDescriptionDo
	Omit(cols ...field.Expr) IStoreProductDescriptionDo
	Join(table schema.Tabler, on ...field.Expr) IStoreProductDescriptionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductDescriptionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductDescriptionDo
	Group(cols ...field.Expr) IStoreProductDescriptionDo
	Having(conds ...gen.Condition) IStoreProductDescriptionDo
	Limit(limit int) IStoreProductDescriptionDo
	Offset(offset int) IStoreProductDescriptionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductDescriptionDo
	Unscoped() IStoreProductDescriptionDo
	Create(values ...*model.StoreProductDescription) error
	CreateInBatches(values []*model.StoreProductDescription, batchSize int) error
	Save(values ...*model.StoreProductDescription) error
	First() (*model.StoreProductDescription, error)
	Take() (*model.StoreProductDescription, error)
	Last() (*model.StoreProductDescription, error)
	Find() ([]*model.StoreProductDescription, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProductDescription, err error)
	FindInBatches(result *[]*model.StoreProductDescription, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreProductDescription) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreProductDescriptionDo
	Assign(attrs ...field.AssignExpr) IStoreProductDescriptionDo
	Joins(fields ...field.RelationField) IStoreProductDescriptionDo
	Preload(fields ...field.RelationField) IStoreProductDescriptionDo
	FirstOrInit() (*model.StoreProductDescription, error)
	FirstOrCreate() (*model.StoreProductDescription, error)
	FindByPage(offset int, limit int) (result []*model.StoreProductDescription, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreProductDescriptionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeProductDescriptionDo) Debug() IStoreProductDescriptionDo {
	return s.withDO(s.DO.Debug())
}

func (s storeProductDescriptionDo) WithContext(ctx context.Context) IStoreProductDescriptionDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeProductDescriptionDo) ReadDB() IStoreProductDescriptionDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeProductDescriptionDo) WriteDB() IStoreProductDescriptionDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeProductDescriptionDo) Session(config *gorm.Session) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeProductDescriptionDo) Clauses(conds ...clause.Expression) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeProductDescriptionDo) Returning(value interface{}, columns ...string) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeProductDescriptionDo) Not(conds ...gen.Condition) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeProductDescriptionDo) Or(conds ...gen.Condition) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeProductDescriptionDo) Select(conds ...field.Expr) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeProductDescriptionDo) Where(conds ...gen.Condition) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeProductDescriptionDo) Order(conds ...field.Expr) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeProductDescriptionDo) Distinct(cols ...field.Expr) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeProductDescriptionDo) Omit(cols ...field.Expr) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeProductDescriptionDo) Join(table schema.Tabler, on ...field.Expr) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeProductDescriptionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductDescriptionDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeProductDescriptionDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductDescriptionDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeProductDescriptionDo) Group(cols ...field.Expr) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeProductDescriptionDo) Having(conds ...gen.Condition) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeProductDescriptionDo) Limit(limit int) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeProductDescriptionDo) Offset(offset int) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeProductDescriptionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeProductDescriptionDo) Unscoped() IStoreProductDescriptionDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeProductDescriptionDo) Create(values ...*model.StoreProductDescription) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeProductDescriptionDo) CreateInBatches(values []*model.StoreProductDescription, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeProductDescriptionDo) Save(values ...*model.StoreProductDescription) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeProductDescriptionDo) First() (*model.StoreProductDescription, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductDescription), nil
	}
}

func (s storeProductDescriptionDo) Take() (*model.StoreProductDescription, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductDescription), nil
	}
}

func (s storeProductDescriptionDo) Last() (*model.StoreProductDescription, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductDescription), nil
	}
}

func (s storeProductDescriptionDo) Find() ([]*model.StoreProductDescription, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreProductDescription), err
}

func (s storeProductDescriptionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProductDescription, err error) {
	buf := make([]*model.StoreProductDescription, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeProductDescriptionDo) FindInBatches(result *[]*model.StoreProductDescription, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeProductDescriptionDo) Attrs(attrs ...field.AssignExpr) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeProductDescriptionDo) Assign(attrs ...field.AssignExpr) IStoreProductDescriptionDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeProductDescriptionDo) Joins(fields ...field.RelationField) IStoreProductDescriptionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeProductDescriptionDo) Preload(fields ...field.RelationField) IStoreProductDescriptionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeProductDescriptionDo) FirstOrInit() (*model.StoreProductDescription, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductDescription), nil
	}
}

func (s storeProductDescriptionDo) FirstOrCreate() (*model.StoreProductDescription, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductDescription), nil
	}
}

func (s storeProductDescriptionDo) FindByPage(offset int, limit int) (result []*model.StoreProductDescription, count int64, err error) {
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

func (s storeProductDescriptionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeProductDescriptionDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeProductDescriptionDo) Delete(models ...*model.StoreProductDescription) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeProductDescriptionDo) withDO(do gen.Dao) *storeProductDescriptionDo {
	s.DO = *do.(*gen.DO)
	return s
}
