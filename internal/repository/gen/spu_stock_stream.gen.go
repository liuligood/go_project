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

func newSpuStockStream(db *gorm.DB, opts ...gen.DOOption) spuStockStream {
	_spuStockStream := spuStockStream{}

	_spuStockStream.spuStockStreamDo.UseDB(db, opts...)
	_spuStockStream.spuStockStreamDo.UseModel(&model.SpuStockStream{})

	tableName := _spuStockStream.spuStockStreamDo.TableName()
	_spuStockStream.ALL = field.NewAsterisk(tableName)
	_spuStockStream.ID = field.NewInt64(tableName, "id")
	_spuStockStream.SpuID = field.NewInt64(tableName, "spu_id")
	_spuStockStream.SkuID = field.NewInt64(tableName, "sku_id")
	_spuStockStream.Price = field.NewField(tableName, "price")
	_spuStockStream.ChangedStock = field.NewInt64(tableName, "changed_stock")
	_spuStockStream.State = field.NewInt64(tableName, "state")
	_spuStockStream.Type = field.NewInt64(tableName, "type")
	_spuStockStream.Stock = field.NewInt64(tableName, "stock")
	_spuStockStream.Identifier = field.NewString(tableName, "identifier")
	_spuStockStream.ExtendInfo = field.NewString(tableName, "extend_info")
	_spuStockStream.CreatedBy = field.NewInt64(tableName, "created_by")
	_spuStockStream.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_spuStockStream.CreatedAt = field.NewInt64(tableName, "created_at")
	_spuStockStream.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_spuStockStream.DeletedAt = field.NewField(tableName, "deleted_at")

	_spuStockStream.fillFieldMap()

	return _spuStockStream
}

// spuStockStream 商品库存流水表
type spuStockStream struct {
	spuStockStreamDo spuStockStreamDo

	ALL          field.Asterisk
	ID           field.Int64  // 主键
	SpuID        field.Int64  // 商品id
	SkuID        field.Int64  // 套餐id
	Price        field.Field  // 价格
	ChangedStock field.Int64  // 本次变更数量
	State        field.Int64  // 状态
	Type         field.Int64  // 流水类型
	Stock        field.Int64  // 变更前库存
	Identifier   field.String // 幂等号
	ExtendInfo   field.String // 拓展信息
	CreatedBy    field.Int64  // 创建时间
	UpdatedBy    field.Int64  // 修改时间
	CreatedAt    field.Int64  // 创建时间
	UpdatedAt    field.Int64  // 修改时间
	DeletedAt    field.Field  // 是否删除

	fieldMap map[string]field.Expr
}

func (s spuStockStream) Table(newTableName string) *spuStockStream {
	s.spuStockStreamDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s spuStockStream) As(alias string) *spuStockStream {
	s.spuStockStreamDo.DO = *(s.spuStockStreamDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *spuStockStream) updateTableName(table string) *spuStockStream {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.SpuID = field.NewInt64(table, "spu_id")
	s.SkuID = field.NewInt64(table, "sku_id")
	s.Price = field.NewField(table, "price")
	s.ChangedStock = field.NewInt64(table, "changed_stock")
	s.State = field.NewInt64(table, "state")
	s.Type = field.NewInt64(table, "type")
	s.Stock = field.NewInt64(table, "stock")
	s.Identifier = field.NewString(table, "identifier")
	s.ExtendInfo = field.NewString(table, "extend_info")
	s.CreatedBy = field.NewInt64(table, "created_by")
	s.UpdatedBy = field.NewInt64(table, "updated_by")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *spuStockStream) WithContext(ctx context.Context) ISpuStockStreamDo {
	return s.spuStockStreamDo.WithContext(ctx)
}

func (s spuStockStream) TableName() string { return s.spuStockStreamDo.TableName() }

func (s spuStockStream) Alias() string { return s.spuStockStreamDo.Alias() }

func (s spuStockStream) Columns(cols ...field.Expr) gen.Columns {
	return s.spuStockStreamDo.Columns(cols...)
}

func (s *spuStockStream) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *spuStockStream) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 15)
	s.fieldMap["id"] = s.ID
	s.fieldMap["spu_id"] = s.SpuID
	s.fieldMap["sku_id"] = s.SkuID
	s.fieldMap["price"] = s.Price
	s.fieldMap["changed_stock"] = s.ChangedStock
	s.fieldMap["state"] = s.State
	s.fieldMap["type"] = s.Type
	s.fieldMap["stock"] = s.Stock
	s.fieldMap["identifier"] = s.Identifier
	s.fieldMap["extend_info"] = s.ExtendInfo
	s.fieldMap["created_by"] = s.CreatedBy
	s.fieldMap["updated_by"] = s.UpdatedBy
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s spuStockStream) clone(db *gorm.DB) spuStockStream {
	s.spuStockStreamDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s spuStockStream) replaceDB(db *gorm.DB) spuStockStream {
	s.spuStockStreamDo.ReplaceDB(db)
	return s
}

type spuStockStreamDo struct{ gen.DO }

type ISpuStockStreamDo interface {
	gen.SubQuery
	Debug() ISpuStockStreamDo
	WithContext(ctx context.Context) ISpuStockStreamDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISpuStockStreamDo
	WriteDB() ISpuStockStreamDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISpuStockStreamDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISpuStockStreamDo
	Not(conds ...gen.Condition) ISpuStockStreamDo
	Or(conds ...gen.Condition) ISpuStockStreamDo
	Select(conds ...field.Expr) ISpuStockStreamDo
	Where(conds ...gen.Condition) ISpuStockStreamDo
	Order(conds ...field.Expr) ISpuStockStreamDo
	Distinct(cols ...field.Expr) ISpuStockStreamDo
	Omit(cols ...field.Expr) ISpuStockStreamDo
	Join(table schema.Tabler, on ...field.Expr) ISpuStockStreamDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISpuStockStreamDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISpuStockStreamDo
	Group(cols ...field.Expr) ISpuStockStreamDo
	Having(conds ...gen.Condition) ISpuStockStreamDo
	Limit(limit int) ISpuStockStreamDo
	Offset(offset int) ISpuStockStreamDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISpuStockStreamDo
	Unscoped() ISpuStockStreamDo
	Create(values ...*model.SpuStockStream) error
	CreateInBatches(values []*model.SpuStockStream, batchSize int) error
	Save(values ...*model.SpuStockStream) error
	First() (*model.SpuStockStream, error)
	Take() (*model.SpuStockStream, error)
	Last() (*model.SpuStockStream, error)
	Find() ([]*model.SpuStockStream, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SpuStockStream, err error)
	FindInBatches(result *[]*model.SpuStockStream, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SpuStockStream) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISpuStockStreamDo
	Assign(attrs ...field.AssignExpr) ISpuStockStreamDo
	Joins(fields ...field.RelationField) ISpuStockStreamDo
	Preload(fields ...field.RelationField) ISpuStockStreamDo
	FirstOrInit() (*model.SpuStockStream, error)
	FirstOrCreate() (*model.SpuStockStream, error)
	FindByPage(offset int, limit int) (result []*model.SpuStockStream, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISpuStockStreamDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s spuStockStreamDo) Debug() ISpuStockStreamDo {
	return s.withDO(s.DO.Debug())
}

func (s spuStockStreamDo) WithContext(ctx context.Context) ISpuStockStreamDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s spuStockStreamDo) ReadDB() ISpuStockStreamDo {
	return s.Clauses(dbresolver.Read)
}

func (s spuStockStreamDo) WriteDB() ISpuStockStreamDo {
	return s.Clauses(dbresolver.Write)
}

func (s spuStockStreamDo) Session(config *gorm.Session) ISpuStockStreamDo {
	return s.withDO(s.DO.Session(config))
}

func (s spuStockStreamDo) Clauses(conds ...clause.Expression) ISpuStockStreamDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s spuStockStreamDo) Returning(value interface{}, columns ...string) ISpuStockStreamDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s spuStockStreamDo) Not(conds ...gen.Condition) ISpuStockStreamDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s spuStockStreamDo) Or(conds ...gen.Condition) ISpuStockStreamDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s spuStockStreamDo) Select(conds ...field.Expr) ISpuStockStreamDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s spuStockStreamDo) Where(conds ...gen.Condition) ISpuStockStreamDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s spuStockStreamDo) Order(conds ...field.Expr) ISpuStockStreamDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s spuStockStreamDo) Distinct(cols ...field.Expr) ISpuStockStreamDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s spuStockStreamDo) Omit(cols ...field.Expr) ISpuStockStreamDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s spuStockStreamDo) Join(table schema.Tabler, on ...field.Expr) ISpuStockStreamDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s spuStockStreamDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISpuStockStreamDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s spuStockStreamDo) RightJoin(table schema.Tabler, on ...field.Expr) ISpuStockStreamDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s spuStockStreamDo) Group(cols ...field.Expr) ISpuStockStreamDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s spuStockStreamDo) Having(conds ...gen.Condition) ISpuStockStreamDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s spuStockStreamDo) Limit(limit int) ISpuStockStreamDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s spuStockStreamDo) Offset(offset int) ISpuStockStreamDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s spuStockStreamDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISpuStockStreamDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s spuStockStreamDo) Unscoped() ISpuStockStreamDo {
	return s.withDO(s.DO.Unscoped())
}

func (s spuStockStreamDo) Create(values ...*model.SpuStockStream) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s spuStockStreamDo) CreateInBatches(values []*model.SpuStockStream, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s spuStockStreamDo) Save(values ...*model.SpuStockStream) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s spuStockStreamDo) First() (*model.SpuStockStream, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SpuStockStream), nil
	}
}

func (s spuStockStreamDo) Take() (*model.SpuStockStream, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SpuStockStream), nil
	}
}

func (s spuStockStreamDo) Last() (*model.SpuStockStream, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SpuStockStream), nil
	}
}

func (s spuStockStreamDo) Find() ([]*model.SpuStockStream, error) {
	result, err := s.DO.Find()
	return result.([]*model.SpuStockStream), err
}

func (s spuStockStreamDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SpuStockStream, err error) {
	buf := make([]*model.SpuStockStream, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s spuStockStreamDo) FindInBatches(result *[]*model.SpuStockStream, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s spuStockStreamDo) Attrs(attrs ...field.AssignExpr) ISpuStockStreamDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s spuStockStreamDo) Assign(attrs ...field.AssignExpr) ISpuStockStreamDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s spuStockStreamDo) Joins(fields ...field.RelationField) ISpuStockStreamDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s spuStockStreamDo) Preload(fields ...field.RelationField) ISpuStockStreamDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s spuStockStreamDo) FirstOrInit() (*model.SpuStockStream, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SpuStockStream), nil
	}
}

func (s spuStockStreamDo) FirstOrCreate() (*model.SpuStockStream, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SpuStockStream), nil
	}
}

func (s spuStockStreamDo) FindByPage(offset int, limit int) (result []*model.SpuStockStream, count int64, err error) {
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

func (s spuStockStreamDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s spuStockStreamDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s spuStockStreamDo) Delete(models ...*model.SpuStockStream) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *spuStockStreamDo) withDO(do gen.Dao) *spuStockStreamDo {
	s.DO = *do.(*gen.DO)
	return s
}
