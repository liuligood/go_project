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

func newStoreProductAttrValue(db *gorm.DB, opts ...gen.DOOption) storeProductAttrValue {
	_storeProductAttrValue := storeProductAttrValue{}

	_storeProductAttrValue.storeProductAttrValueDo.UseDB(db, opts...)
	_storeProductAttrValue.storeProductAttrValueDo.UseModel(&model.StoreProductAttrValue{})

	tableName := _storeProductAttrValue.storeProductAttrValueDo.TableName()
	_storeProductAttrValue.ALL = field.NewAsterisk(tableName)
	_storeProductAttrValue.ID = field.NewInt64(tableName, "id")
	_storeProductAttrValue.ProductID = field.NewInt64(tableName, "product_id")
	_storeProductAttrValue.Suk = field.NewString(tableName, "suk")
	_storeProductAttrValue.Stock = field.NewInt64(tableName, "stock")
	_storeProductAttrValue.Sales = field.NewInt64(tableName, "sales")
	_storeProductAttrValue.Price = field.NewField(tableName, "price")
	_storeProductAttrValue.Image = field.NewString(tableName, "image")
	_storeProductAttrValue.Unique = field.NewString(tableName, "unique")
	_storeProductAttrValue.Cost = field.NewField(tableName, "cost")
	_storeProductAttrValue.BarCode = field.NewString(tableName, "bar_code")
	_storeProductAttrValue.OtPrice = field.NewField(tableName, "ot_price")
	_storeProductAttrValue.Weight = field.NewField(tableName, "weight")
	_storeProductAttrValue.Volume = field.NewField(tableName, "volume")
	_storeProductAttrValue.Brokerage = field.NewField(tableName, "brokerage")
	_storeProductAttrValue.BrokerageTwo = field.NewField(tableName, "brokerage_two")
	_storeProductAttrValue.Type = field.NewInt64(tableName, "type")
	_storeProductAttrValue.Quota = field.NewInt64(tableName, "quota")
	_storeProductAttrValue.QuotaShow = field.NewInt64(tableName, "quota_show")
	_storeProductAttrValue.AttrValue = field.NewString(tableName, "attr_value")
	_storeProductAttrValue.IsDel = field.NewInt64(tableName, "is_del")
	_storeProductAttrValue.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeProductAttrValue.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeProductAttrValue.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeProductAttrValue.fillFieldMap()

	return _storeProductAttrValue
}

// storeProductAttrValue 商品属性值表
type storeProductAttrValue struct {
	storeProductAttrValueDo storeProductAttrValueDo

	ALL          field.Asterisk
	ID           field.Int64  // 主键
	ProductID    field.Int64  // 商品ID
	Suk          field.String // 商品属性索引值 (attr_value|attr_value[|....])
	Stock        field.Int64  // 属性对应的库存
	Sales        field.Int64  // 销量
	Price        field.Field  // 属性金额
	Image        field.String // 图片
	Unique       field.String // 唯一值
	Cost         field.Field  // 成本价
	BarCode      field.String // 商品条码
	OtPrice      field.Field  // 原价
	Weight       field.Field  // 重量
	Volume       field.Field  // 体积
	Brokerage    field.Field  // 一级返佣
	BrokerageTwo field.Field  // 二级返佣
	Type         field.Int64  // 活动类型 0=商品，1=秒杀，2=砍价，3=拼团
	Quota        field.Int64  // 活动限购数量
	QuotaShow    field.Int64  // 活动限购数量显示
	AttrValue    field.String // attr_values 创建更新时的属性对应
	IsDel        field.Int64  // 是否删除,0-否，1-是
	CreatedAt    field.Int64  // 创建时间
	UpdatedAt    field.Int64  // 修改时间
	DeletedAt    field.Field  // 是否删除

	fieldMap map[string]field.Expr
}

func (s storeProductAttrValue) Table(newTableName string) *storeProductAttrValue {
	s.storeProductAttrValueDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeProductAttrValue) As(alias string) *storeProductAttrValue {
	s.storeProductAttrValueDo.DO = *(s.storeProductAttrValueDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeProductAttrValue) updateTableName(table string) *storeProductAttrValue {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.ProductID = field.NewInt64(table, "product_id")
	s.Suk = field.NewString(table, "suk")
	s.Stock = field.NewInt64(table, "stock")
	s.Sales = field.NewInt64(table, "sales")
	s.Price = field.NewField(table, "price")
	s.Image = field.NewString(table, "image")
	s.Unique = field.NewString(table, "unique")
	s.Cost = field.NewField(table, "cost")
	s.BarCode = field.NewString(table, "bar_code")
	s.OtPrice = field.NewField(table, "ot_price")
	s.Weight = field.NewField(table, "weight")
	s.Volume = field.NewField(table, "volume")
	s.Brokerage = field.NewField(table, "brokerage")
	s.BrokerageTwo = field.NewField(table, "brokerage_two")
	s.Type = field.NewInt64(table, "type")
	s.Quota = field.NewInt64(table, "quota")
	s.QuotaShow = field.NewInt64(table, "quota_show")
	s.AttrValue = field.NewString(table, "attr_value")
	s.IsDel = field.NewInt64(table, "is_del")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeProductAttrValue) WithContext(ctx context.Context) IStoreProductAttrValueDo {
	return s.storeProductAttrValueDo.WithContext(ctx)
}

func (s storeProductAttrValue) TableName() string { return s.storeProductAttrValueDo.TableName() }

func (s storeProductAttrValue) Alias() string { return s.storeProductAttrValueDo.Alias() }

func (s storeProductAttrValue) Columns(cols ...field.Expr) gen.Columns {
	return s.storeProductAttrValueDo.Columns(cols...)
}

func (s *storeProductAttrValue) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeProductAttrValue) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 23)
	s.fieldMap["id"] = s.ID
	s.fieldMap["product_id"] = s.ProductID
	s.fieldMap["suk"] = s.Suk
	s.fieldMap["stock"] = s.Stock
	s.fieldMap["sales"] = s.Sales
	s.fieldMap["price"] = s.Price
	s.fieldMap["image"] = s.Image
	s.fieldMap["unique"] = s.Unique
	s.fieldMap["cost"] = s.Cost
	s.fieldMap["bar_code"] = s.BarCode
	s.fieldMap["ot_price"] = s.OtPrice
	s.fieldMap["weight"] = s.Weight
	s.fieldMap["volume"] = s.Volume
	s.fieldMap["brokerage"] = s.Brokerage
	s.fieldMap["brokerage_two"] = s.BrokerageTwo
	s.fieldMap["type"] = s.Type
	s.fieldMap["quota"] = s.Quota
	s.fieldMap["quota_show"] = s.QuotaShow
	s.fieldMap["attr_value"] = s.AttrValue
	s.fieldMap["is_del"] = s.IsDel
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeProductAttrValue) clone(db *gorm.DB) storeProductAttrValue {
	s.storeProductAttrValueDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeProductAttrValue) replaceDB(db *gorm.DB) storeProductAttrValue {
	s.storeProductAttrValueDo.ReplaceDB(db)
	return s
}

type storeProductAttrValueDo struct{ gen.DO }

type IStoreProductAttrValueDo interface {
	gen.SubQuery
	Debug() IStoreProductAttrValueDo
	WithContext(ctx context.Context) IStoreProductAttrValueDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreProductAttrValueDo
	WriteDB() IStoreProductAttrValueDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreProductAttrValueDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreProductAttrValueDo
	Not(conds ...gen.Condition) IStoreProductAttrValueDo
	Or(conds ...gen.Condition) IStoreProductAttrValueDo
	Select(conds ...field.Expr) IStoreProductAttrValueDo
	Where(conds ...gen.Condition) IStoreProductAttrValueDo
	Order(conds ...field.Expr) IStoreProductAttrValueDo
	Distinct(cols ...field.Expr) IStoreProductAttrValueDo
	Omit(cols ...field.Expr) IStoreProductAttrValueDo
	Join(table schema.Tabler, on ...field.Expr) IStoreProductAttrValueDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductAttrValueDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductAttrValueDo
	Group(cols ...field.Expr) IStoreProductAttrValueDo
	Having(conds ...gen.Condition) IStoreProductAttrValueDo
	Limit(limit int) IStoreProductAttrValueDo
	Offset(offset int) IStoreProductAttrValueDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductAttrValueDo
	Unscoped() IStoreProductAttrValueDo
	Create(values ...*model.StoreProductAttrValue) error
	CreateInBatches(values []*model.StoreProductAttrValue, batchSize int) error
	Save(values ...*model.StoreProductAttrValue) error
	First() (*model.StoreProductAttrValue, error)
	Take() (*model.StoreProductAttrValue, error)
	Last() (*model.StoreProductAttrValue, error)
	Find() ([]*model.StoreProductAttrValue, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProductAttrValue, err error)
	FindInBatches(result *[]*model.StoreProductAttrValue, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreProductAttrValue) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreProductAttrValueDo
	Assign(attrs ...field.AssignExpr) IStoreProductAttrValueDo
	Joins(fields ...field.RelationField) IStoreProductAttrValueDo
	Preload(fields ...field.RelationField) IStoreProductAttrValueDo
	FirstOrInit() (*model.StoreProductAttrValue, error)
	FirstOrCreate() (*model.StoreProductAttrValue, error)
	FindByPage(offset int, limit int) (result []*model.StoreProductAttrValue, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreProductAttrValueDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeProductAttrValueDo) Debug() IStoreProductAttrValueDo {
	return s.withDO(s.DO.Debug())
}

func (s storeProductAttrValueDo) WithContext(ctx context.Context) IStoreProductAttrValueDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeProductAttrValueDo) ReadDB() IStoreProductAttrValueDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeProductAttrValueDo) WriteDB() IStoreProductAttrValueDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeProductAttrValueDo) Session(config *gorm.Session) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeProductAttrValueDo) Clauses(conds ...clause.Expression) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeProductAttrValueDo) Returning(value interface{}, columns ...string) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeProductAttrValueDo) Not(conds ...gen.Condition) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeProductAttrValueDo) Or(conds ...gen.Condition) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeProductAttrValueDo) Select(conds ...field.Expr) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeProductAttrValueDo) Where(conds ...gen.Condition) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeProductAttrValueDo) Order(conds ...field.Expr) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeProductAttrValueDo) Distinct(cols ...field.Expr) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeProductAttrValueDo) Omit(cols ...field.Expr) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeProductAttrValueDo) Join(table schema.Tabler, on ...field.Expr) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeProductAttrValueDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductAttrValueDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeProductAttrValueDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductAttrValueDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeProductAttrValueDo) Group(cols ...field.Expr) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeProductAttrValueDo) Having(conds ...gen.Condition) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeProductAttrValueDo) Limit(limit int) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeProductAttrValueDo) Offset(offset int) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeProductAttrValueDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeProductAttrValueDo) Unscoped() IStoreProductAttrValueDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeProductAttrValueDo) Create(values ...*model.StoreProductAttrValue) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeProductAttrValueDo) CreateInBatches(values []*model.StoreProductAttrValue, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeProductAttrValueDo) Save(values ...*model.StoreProductAttrValue) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeProductAttrValueDo) First() (*model.StoreProductAttrValue, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductAttrValue), nil
	}
}

func (s storeProductAttrValueDo) Take() (*model.StoreProductAttrValue, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductAttrValue), nil
	}
}

func (s storeProductAttrValueDo) Last() (*model.StoreProductAttrValue, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductAttrValue), nil
	}
}

func (s storeProductAttrValueDo) Find() ([]*model.StoreProductAttrValue, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreProductAttrValue), err
}

func (s storeProductAttrValueDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProductAttrValue, err error) {
	buf := make([]*model.StoreProductAttrValue, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeProductAttrValueDo) FindInBatches(result *[]*model.StoreProductAttrValue, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeProductAttrValueDo) Attrs(attrs ...field.AssignExpr) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeProductAttrValueDo) Assign(attrs ...field.AssignExpr) IStoreProductAttrValueDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeProductAttrValueDo) Joins(fields ...field.RelationField) IStoreProductAttrValueDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeProductAttrValueDo) Preload(fields ...field.RelationField) IStoreProductAttrValueDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeProductAttrValueDo) FirstOrInit() (*model.StoreProductAttrValue, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductAttrValue), nil
	}
}

func (s storeProductAttrValueDo) FirstOrCreate() (*model.StoreProductAttrValue, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductAttrValue), nil
	}
}

func (s storeProductAttrValueDo) FindByPage(offset int, limit int) (result []*model.StoreProductAttrValue, count int64, err error) {
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

func (s storeProductAttrValueDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeProductAttrValueDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeProductAttrValueDo) Delete(models ...*model.StoreProductAttrValue) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeProductAttrValueDo) withDO(do gen.Dao) *storeProductAttrValueDo {
	s.DO = *do.(*gen.DO)
	return s
}
