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

func newSkuAttribute(db *gorm.DB, opts ...gen.DOOption) skuAttribute {
	_skuAttribute := skuAttribute{}

	_skuAttribute.skuAttributeDo.UseDB(db, opts...)
	_skuAttribute.skuAttributeDo.UseModel(&model.SkuAttribute{})

	tableName := _skuAttribute.skuAttributeDo.TableName()
	_skuAttribute.ALL = field.NewAsterisk(tableName)
	_skuAttribute.ID = field.NewInt64(tableName, "id")
	_skuAttribute.SkuID = field.NewInt64(tableName, "sku_id")
	_skuAttribute.Type = field.NewInt64(tableName, "type")
	_skuAttribute.Title = field.NewString(tableName, "title")
	_skuAttribute.Value = field.NewString(tableName, "value")
	_skuAttribute.Sort = field.NewInt64(tableName, "sort")
	_skuAttribute.Sign = field.NewInt64(tableName, "sign")
	_skuAttribute.CreatedAt = field.NewInt64(tableName, "created_at")
	_skuAttribute.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_skuAttribute.DeletedAt = field.NewField(tableName, "deleted_at")
	_skuAttribute.CreatedBy = field.NewInt64(tableName, "created_by")
	_skuAttribute.UpdatedBy = field.NewInt64(tableName, "updated_by")

	_skuAttribute.fillFieldMap()

	return _skuAttribute
}

// skuAttribute sku属性表
type skuAttribute struct {
	skuAttributeDo skuAttributeDo

	ALL field.Asterisk
	ID  field.Int64 // 主键
	/*
		skuId

	*/
	SkuID     field.Int64
	Type      field.Int64  // 类型
	Title     field.String // 属性标题
	Value     field.String // 属性值
	Sort      field.Int64  // 排序
	Sign      field.Int64  // 标志(位运算)
	CreatedAt field.Int64  // 创建时间
	UpdatedAt field.Int64  // 修改时间
	DeletedAt field.Field  // 是否删除
	CreatedBy field.Int64  // 创建时间
	UpdatedBy field.Int64  // 修改时间

	fieldMap map[string]field.Expr
}

func (s skuAttribute) Table(newTableName string) *skuAttribute {
	s.skuAttributeDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skuAttribute) As(alias string) *skuAttribute {
	s.skuAttributeDo.DO = *(s.skuAttributeDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skuAttribute) updateTableName(table string) *skuAttribute {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.SkuID = field.NewInt64(table, "sku_id")
	s.Type = field.NewInt64(table, "type")
	s.Title = field.NewString(table, "title")
	s.Value = field.NewString(table, "value")
	s.Sort = field.NewInt64(table, "sort")
	s.Sign = field.NewInt64(table, "sign")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.CreatedBy = field.NewInt64(table, "created_by")
	s.UpdatedBy = field.NewInt64(table, "updated_by")

	s.fillFieldMap()

	return s
}

func (s *skuAttribute) WithContext(ctx context.Context) ISkuAttributeDo {
	return s.skuAttributeDo.WithContext(ctx)
}

func (s skuAttribute) TableName() string { return s.skuAttributeDo.TableName() }

func (s skuAttribute) Alias() string { return s.skuAttributeDo.Alias() }

func (s skuAttribute) Columns(cols ...field.Expr) gen.Columns {
	return s.skuAttributeDo.Columns(cols...)
}

func (s *skuAttribute) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skuAttribute) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["id"] = s.ID
	s.fieldMap["sku_id"] = s.SkuID
	s.fieldMap["type"] = s.Type
	s.fieldMap["title"] = s.Title
	s.fieldMap["value"] = s.Value
	s.fieldMap["sort"] = s.Sort
	s.fieldMap["sign"] = s.Sign
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["created_by"] = s.CreatedBy
	s.fieldMap["updated_by"] = s.UpdatedBy
}

func (s skuAttribute) clone(db *gorm.DB) skuAttribute {
	s.skuAttributeDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skuAttribute) replaceDB(db *gorm.DB) skuAttribute {
	s.skuAttributeDo.ReplaceDB(db)
	return s
}

type skuAttributeDo struct{ gen.DO }

type ISkuAttributeDo interface {
	gen.SubQuery
	Debug() ISkuAttributeDo
	WithContext(ctx context.Context) ISkuAttributeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISkuAttributeDo
	WriteDB() ISkuAttributeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISkuAttributeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISkuAttributeDo
	Not(conds ...gen.Condition) ISkuAttributeDo
	Or(conds ...gen.Condition) ISkuAttributeDo
	Select(conds ...field.Expr) ISkuAttributeDo
	Where(conds ...gen.Condition) ISkuAttributeDo
	Order(conds ...field.Expr) ISkuAttributeDo
	Distinct(cols ...field.Expr) ISkuAttributeDo
	Omit(cols ...field.Expr) ISkuAttributeDo
	Join(table schema.Tabler, on ...field.Expr) ISkuAttributeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISkuAttributeDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISkuAttributeDo
	Group(cols ...field.Expr) ISkuAttributeDo
	Having(conds ...gen.Condition) ISkuAttributeDo
	Limit(limit int) ISkuAttributeDo
	Offset(offset int) ISkuAttributeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISkuAttributeDo
	Unscoped() ISkuAttributeDo
	Create(values ...*model.SkuAttribute) error
	CreateInBatches(values []*model.SkuAttribute, batchSize int) error
	Save(values ...*model.SkuAttribute) error
	First() (*model.SkuAttribute, error)
	Take() (*model.SkuAttribute, error)
	Last() (*model.SkuAttribute, error)
	Find() ([]*model.SkuAttribute, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SkuAttribute, err error)
	FindInBatches(result *[]*model.SkuAttribute, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SkuAttribute) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISkuAttributeDo
	Assign(attrs ...field.AssignExpr) ISkuAttributeDo
	Joins(fields ...field.RelationField) ISkuAttributeDo
	Preload(fields ...field.RelationField) ISkuAttributeDo
	FirstOrInit() (*model.SkuAttribute, error)
	FirstOrCreate() (*model.SkuAttribute, error)
	FindByPage(offset int, limit int) (result []*model.SkuAttribute, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISkuAttributeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s skuAttributeDo) Debug() ISkuAttributeDo {
	return s.withDO(s.DO.Debug())
}

func (s skuAttributeDo) WithContext(ctx context.Context) ISkuAttributeDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skuAttributeDo) ReadDB() ISkuAttributeDo {
	return s.Clauses(dbresolver.Read)
}

func (s skuAttributeDo) WriteDB() ISkuAttributeDo {
	return s.Clauses(dbresolver.Write)
}

func (s skuAttributeDo) Session(config *gorm.Session) ISkuAttributeDo {
	return s.withDO(s.DO.Session(config))
}

func (s skuAttributeDo) Clauses(conds ...clause.Expression) ISkuAttributeDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skuAttributeDo) Returning(value interface{}, columns ...string) ISkuAttributeDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skuAttributeDo) Not(conds ...gen.Condition) ISkuAttributeDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skuAttributeDo) Or(conds ...gen.Condition) ISkuAttributeDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skuAttributeDo) Select(conds ...field.Expr) ISkuAttributeDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skuAttributeDo) Where(conds ...gen.Condition) ISkuAttributeDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skuAttributeDo) Order(conds ...field.Expr) ISkuAttributeDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skuAttributeDo) Distinct(cols ...field.Expr) ISkuAttributeDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skuAttributeDo) Omit(cols ...field.Expr) ISkuAttributeDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skuAttributeDo) Join(table schema.Tabler, on ...field.Expr) ISkuAttributeDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skuAttributeDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISkuAttributeDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skuAttributeDo) RightJoin(table schema.Tabler, on ...field.Expr) ISkuAttributeDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skuAttributeDo) Group(cols ...field.Expr) ISkuAttributeDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skuAttributeDo) Having(conds ...gen.Condition) ISkuAttributeDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skuAttributeDo) Limit(limit int) ISkuAttributeDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skuAttributeDo) Offset(offset int) ISkuAttributeDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skuAttributeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISkuAttributeDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skuAttributeDo) Unscoped() ISkuAttributeDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skuAttributeDo) Create(values ...*model.SkuAttribute) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skuAttributeDo) CreateInBatches(values []*model.SkuAttribute, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skuAttributeDo) Save(values ...*model.SkuAttribute) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skuAttributeDo) First() (*model.SkuAttribute, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuAttribute), nil
	}
}

func (s skuAttributeDo) Take() (*model.SkuAttribute, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuAttribute), nil
	}
}

func (s skuAttributeDo) Last() (*model.SkuAttribute, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuAttribute), nil
	}
}

func (s skuAttributeDo) Find() ([]*model.SkuAttribute, error) {
	result, err := s.DO.Find()
	return result.([]*model.SkuAttribute), err
}

func (s skuAttributeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SkuAttribute, err error) {
	buf := make([]*model.SkuAttribute, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skuAttributeDo) FindInBatches(result *[]*model.SkuAttribute, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skuAttributeDo) Attrs(attrs ...field.AssignExpr) ISkuAttributeDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skuAttributeDo) Assign(attrs ...field.AssignExpr) ISkuAttributeDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skuAttributeDo) Joins(fields ...field.RelationField) ISkuAttributeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skuAttributeDo) Preload(fields ...field.RelationField) ISkuAttributeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skuAttributeDo) FirstOrInit() (*model.SkuAttribute, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuAttribute), nil
	}
}

func (s skuAttributeDo) FirstOrCreate() (*model.SkuAttribute, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuAttribute), nil
	}
}

func (s skuAttributeDo) FindByPage(offset int, limit int) (result []*model.SkuAttribute, count int64, err error) {
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

func (s skuAttributeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skuAttributeDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skuAttributeDo) Delete(models ...*model.SkuAttribute) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skuAttributeDo) withDO(do gen.Dao) *skuAttributeDo {
	s.DO = *do.(*gen.DO)
	return s
}
