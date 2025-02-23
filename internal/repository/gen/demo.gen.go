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

func newDemo(db *gorm.DB, opts ...gen.DOOption) demo {
	_demo := demo{}

	_demo.demoDo.UseDB(db, opts...)
	_demo.demoDo.UseModel(&model.Demo{})

	tableName := _demo.demoDo.TableName()
	_demo.ALL = field.NewAsterisk(tableName)
	_demo.ID = field.NewInt64(tableName, "id")
	_demo.Name = field.NewString(tableName, "name")
	_demo.Phone = field.NewString(tableName, "phone")
	_demo.CreatedAt = field.NewInt64(tableName, "created_at")
	_demo.CreatedBy = field.NewInt64(tableName, "created_by")
	_demo.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_demo.UpdatedBy = field.NewInt64(tableName, "updated_by")

	_demo.fillFieldMap()

	return _demo
}

// demo demo测试表
type demo struct {
	demoDo demoDo

	ALL       field.Asterisk
	ID        field.Int64  // 自增ID
	Name      field.String // 名称
	Phone     field.String // 手机号
	CreatedAt field.Int64  // 创建时间
	CreatedBy field.Int64  // 创建人
	UpdatedAt field.Int64  // 更新时间
	UpdatedBy field.Int64  // 更新人

	fieldMap map[string]field.Expr
}

func (d demo) Table(newTableName string) *demo {
	d.demoDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d demo) As(alias string) *demo {
	d.demoDo.DO = *(d.demoDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *demo) updateTableName(table string) *demo {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt64(table, "id")
	d.Name = field.NewString(table, "name")
	d.Phone = field.NewString(table, "phone")
	d.CreatedAt = field.NewInt64(table, "created_at")
	d.CreatedBy = field.NewInt64(table, "created_by")
	d.UpdatedAt = field.NewInt64(table, "updated_at")
	d.UpdatedBy = field.NewInt64(table, "updated_by")

	d.fillFieldMap()

	return d
}

func (d *demo) WithContext(ctx context.Context) IDemoDo { return d.demoDo.WithContext(ctx) }

func (d demo) TableName() string { return d.demoDo.TableName() }

func (d demo) Alias() string { return d.demoDo.Alias() }

func (d demo) Columns(cols ...field.Expr) gen.Columns { return d.demoDo.Columns(cols...) }

func (d *demo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *demo) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 7)
	d.fieldMap["id"] = d.ID
	d.fieldMap["name"] = d.Name
	d.fieldMap["phone"] = d.Phone
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["created_by"] = d.CreatedBy
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["updated_by"] = d.UpdatedBy
}

func (d demo) clone(db *gorm.DB) demo {
	d.demoDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d demo) replaceDB(db *gorm.DB) demo {
	d.demoDo.ReplaceDB(db)
	return d
}

type demoDo struct{ gen.DO }

type IDemoDo interface {
	gen.SubQuery
	Debug() IDemoDo
	WithContext(ctx context.Context) IDemoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDemoDo
	WriteDB() IDemoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDemoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDemoDo
	Not(conds ...gen.Condition) IDemoDo
	Or(conds ...gen.Condition) IDemoDo
	Select(conds ...field.Expr) IDemoDo
	Where(conds ...gen.Condition) IDemoDo
	Order(conds ...field.Expr) IDemoDo
	Distinct(cols ...field.Expr) IDemoDo
	Omit(cols ...field.Expr) IDemoDo
	Join(table schema.Tabler, on ...field.Expr) IDemoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDemoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDemoDo
	Group(cols ...field.Expr) IDemoDo
	Having(conds ...gen.Condition) IDemoDo
	Limit(limit int) IDemoDo
	Offset(offset int) IDemoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDemoDo
	Unscoped() IDemoDo
	Create(values ...*model.Demo) error
	CreateInBatches(values []*model.Demo, batchSize int) error
	Save(values ...*model.Demo) error
	First() (*model.Demo, error)
	Take() (*model.Demo, error)
	Last() (*model.Demo, error)
	Find() ([]*model.Demo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Demo, err error)
	FindInBatches(result *[]*model.Demo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Demo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDemoDo
	Assign(attrs ...field.AssignExpr) IDemoDo
	Joins(fields ...field.RelationField) IDemoDo
	Preload(fields ...field.RelationField) IDemoDo
	FirstOrInit() (*model.Demo, error)
	FirstOrCreate() (*model.Demo, error)
	FindByPage(offset int, limit int) (result []*model.Demo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDemoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d demoDo) Debug() IDemoDo {
	return d.withDO(d.DO.Debug())
}

func (d demoDo) WithContext(ctx context.Context) IDemoDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d demoDo) ReadDB() IDemoDo {
	return d.Clauses(dbresolver.Read)
}

func (d demoDo) WriteDB() IDemoDo {
	return d.Clauses(dbresolver.Write)
}

func (d demoDo) Session(config *gorm.Session) IDemoDo {
	return d.withDO(d.DO.Session(config))
}

func (d demoDo) Clauses(conds ...clause.Expression) IDemoDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d demoDo) Returning(value interface{}, columns ...string) IDemoDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d demoDo) Not(conds ...gen.Condition) IDemoDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d demoDo) Or(conds ...gen.Condition) IDemoDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d demoDo) Select(conds ...field.Expr) IDemoDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d demoDo) Where(conds ...gen.Condition) IDemoDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d demoDo) Order(conds ...field.Expr) IDemoDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d demoDo) Distinct(cols ...field.Expr) IDemoDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d demoDo) Omit(cols ...field.Expr) IDemoDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d demoDo) Join(table schema.Tabler, on ...field.Expr) IDemoDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d demoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDemoDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d demoDo) RightJoin(table schema.Tabler, on ...field.Expr) IDemoDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d demoDo) Group(cols ...field.Expr) IDemoDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d demoDo) Having(conds ...gen.Condition) IDemoDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d demoDo) Limit(limit int) IDemoDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d demoDo) Offset(offset int) IDemoDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d demoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDemoDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d demoDo) Unscoped() IDemoDo {
	return d.withDO(d.DO.Unscoped())
}

func (d demoDo) Create(values ...*model.Demo) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d demoDo) CreateInBatches(values []*model.Demo, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d demoDo) Save(values ...*model.Demo) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d demoDo) First() (*model.Demo, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Demo), nil
	}
}

func (d demoDo) Take() (*model.Demo, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Demo), nil
	}
}

func (d demoDo) Last() (*model.Demo, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Demo), nil
	}
}

func (d demoDo) Find() ([]*model.Demo, error) {
	result, err := d.DO.Find()
	return result.([]*model.Demo), err
}

func (d demoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Demo, err error) {
	buf := make([]*model.Demo, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d demoDo) FindInBatches(result *[]*model.Demo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d demoDo) Attrs(attrs ...field.AssignExpr) IDemoDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d demoDo) Assign(attrs ...field.AssignExpr) IDemoDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d demoDo) Joins(fields ...field.RelationField) IDemoDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d demoDo) Preload(fields ...field.RelationField) IDemoDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d demoDo) FirstOrInit() (*model.Demo, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Demo), nil
	}
}

func (d demoDo) FirstOrCreate() (*model.Demo, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Demo), nil
	}
}

func (d demoDo) FindByPage(offset int, limit int) (result []*model.Demo, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d demoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d demoDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d demoDo) Delete(models ...*model.Demo) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *demoDo) withDO(do gen.Dao) *demoDo {
	d.DO = *do.(*gen.DO)
	return d
}
