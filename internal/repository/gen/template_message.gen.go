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

func newTemplateMessage(db *gorm.DB, opts ...gen.DOOption) templateMessage {
	_templateMessage := templateMessage{}

	_templateMessage.templateMessageDo.UseDB(db, opts...)
	_templateMessage.templateMessageDo.UseModel(&model.TemplateMessage{})

	tableName := _templateMessage.templateMessageDo.TableName()
	_templateMessage.ALL = field.NewAsterisk(tableName)
	_templateMessage.ID = field.NewInt64(tableName, "id")
	_templateMessage.Type = field.NewInt64(tableName, "type")
	_templateMessage.TempKey = field.NewString(tableName, "temp_key")
	_templateMessage.Name = field.NewString(tableName, "name")
	_templateMessage.Content = field.NewString(tableName, "content")
	_templateMessage.TempID = field.NewString(tableName, "temp_id")
	_templateMessage.Status = field.NewInt64(tableName, "status")
	_templateMessage.CreatedAt = field.NewInt64(tableName, "created_at")
	_templateMessage.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_templateMessage.DeletedAt = field.NewField(tableName, "deleted_at")

	_templateMessage.fillFieldMap()

	return _templateMessage
}

// templateMessage 微信模板
type templateMessage struct {
	templateMessageDo templateMessageDo

	ALL       field.Asterisk
	ID        field.Int64  // 模板id
	Type      field.Int64  // 0=订阅消息,1=微信模板消息
	TempKey   field.String // 模板编号
	Name      field.String // 模板名
	Content   field.String // 回复内容
	TempID    field.String // 模板ID
	Status    field.Int64  // 状态
	CreatedAt field.Int64  // 创建时间
	UpdatedAt field.Int64  // 修改时间
	DeletedAt field.Field  // 是否删除

	fieldMap map[string]field.Expr
}

func (t templateMessage) Table(newTableName string) *templateMessage {
	t.templateMessageDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t templateMessage) As(alias string) *templateMessage {
	t.templateMessageDo.DO = *(t.templateMessageDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *templateMessage) updateTableName(table string) *templateMessage {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.Type = field.NewInt64(table, "type")
	t.TempKey = field.NewString(table, "temp_key")
	t.Name = field.NewString(table, "name")
	t.Content = field.NewString(table, "content")
	t.TempID = field.NewString(table, "temp_id")
	t.Status = field.NewInt64(table, "status")
	t.CreatedAt = field.NewInt64(table, "created_at")
	t.UpdatedAt = field.NewInt64(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *templateMessage) WithContext(ctx context.Context) ITemplateMessageDo {
	return t.templateMessageDo.WithContext(ctx)
}

func (t templateMessage) TableName() string { return t.templateMessageDo.TableName() }

func (t templateMessage) Alias() string { return t.templateMessageDo.Alias() }

func (t templateMessage) Columns(cols ...field.Expr) gen.Columns {
	return t.templateMessageDo.Columns(cols...)
}

func (t *templateMessage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *templateMessage) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 10)
	t.fieldMap["id"] = t.ID
	t.fieldMap["type"] = t.Type
	t.fieldMap["temp_key"] = t.TempKey
	t.fieldMap["name"] = t.Name
	t.fieldMap["content"] = t.Content
	t.fieldMap["temp_id"] = t.TempID
	t.fieldMap["status"] = t.Status
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t templateMessage) clone(db *gorm.DB) templateMessage {
	t.templateMessageDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t templateMessage) replaceDB(db *gorm.DB) templateMessage {
	t.templateMessageDo.ReplaceDB(db)
	return t
}

type templateMessageDo struct{ gen.DO }

type ITemplateMessageDo interface {
	gen.SubQuery
	Debug() ITemplateMessageDo
	WithContext(ctx context.Context) ITemplateMessageDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITemplateMessageDo
	WriteDB() ITemplateMessageDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITemplateMessageDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITemplateMessageDo
	Not(conds ...gen.Condition) ITemplateMessageDo
	Or(conds ...gen.Condition) ITemplateMessageDo
	Select(conds ...field.Expr) ITemplateMessageDo
	Where(conds ...gen.Condition) ITemplateMessageDo
	Order(conds ...field.Expr) ITemplateMessageDo
	Distinct(cols ...field.Expr) ITemplateMessageDo
	Omit(cols ...field.Expr) ITemplateMessageDo
	Join(table schema.Tabler, on ...field.Expr) ITemplateMessageDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITemplateMessageDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITemplateMessageDo
	Group(cols ...field.Expr) ITemplateMessageDo
	Having(conds ...gen.Condition) ITemplateMessageDo
	Limit(limit int) ITemplateMessageDo
	Offset(offset int) ITemplateMessageDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITemplateMessageDo
	Unscoped() ITemplateMessageDo
	Create(values ...*model.TemplateMessage) error
	CreateInBatches(values []*model.TemplateMessage, batchSize int) error
	Save(values ...*model.TemplateMessage) error
	First() (*model.TemplateMessage, error)
	Take() (*model.TemplateMessage, error)
	Last() (*model.TemplateMessage, error)
	Find() ([]*model.TemplateMessage, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TemplateMessage, err error)
	FindInBatches(result *[]*model.TemplateMessage, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TemplateMessage) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITemplateMessageDo
	Assign(attrs ...field.AssignExpr) ITemplateMessageDo
	Joins(fields ...field.RelationField) ITemplateMessageDo
	Preload(fields ...field.RelationField) ITemplateMessageDo
	FirstOrInit() (*model.TemplateMessage, error)
	FirstOrCreate() (*model.TemplateMessage, error)
	FindByPage(offset int, limit int) (result []*model.TemplateMessage, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITemplateMessageDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t templateMessageDo) Debug() ITemplateMessageDo {
	return t.withDO(t.DO.Debug())
}

func (t templateMessageDo) WithContext(ctx context.Context) ITemplateMessageDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t templateMessageDo) ReadDB() ITemplateMessageDo {
	return t.Clauses(dbresolver.Read)
}

func (t templateMessageDo) WriteDB() ITemplateMessageDo {
	return t.Clauses(dbresolver.Write)
}

func (t templateMessageDo) Session(config *gorm.Session) ITemplateMessageDo {
	return t.withDO(t.DO.Session(config))
}

func (t templateMessageDo) Clauses(conds ...clause.Expression) ITemplateMessageDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t templateMessageDo) Returning(value interface{}, columns ...string) ITemplateMessageDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t templateMessageDo) Not(conds ...gen.Condition) ITemplateMessageDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t templateMessageDo) Or(conds ...gen.Condition) ITemplateMessageDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t templateMessageDo) Select(conds ...field.Expr) ITemplateMessageDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t templateMessageDo) Where(conds ...gen.Condition) ITemplateMessageDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t templateMessageDo) Order(conds ...field.Expr) ITemplateMessageDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t templateMessageDo) Distinct(cols ...field.Expr) ITemplateMessageDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t templateMessageDo) Omit(cols ...field.Expr) ITemplateMessageDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t templateMessageDo) Join(table schema.Tabler, on ...field.Expr) ITemplateMessageDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t templateMessageDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITemplateMessageDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t templateMessageDo) RightJoin(table schema.Tabler, on ...field.Expr) ITemplateMessageDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t templateMessageDo) Group(cols ...field.Expr) ITemplateMessageDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t templateMessageDo) Having(conds ...gen.Condition) ITemplateMessageDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t templateMessageDo) Limit(limit int) ITemplateMessageDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t templateMessageDo) Offset(offset int) ITemplateMessageDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t templateMessageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITemplateMessageDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t templateMessageDo) Unscoped() ITemplateMessageDo {
	return t.withDO(t.DO.Unscoped())
}

func (t templateMessageDo) Create(values ...*model.TemplateMessage) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t templateMessageDo) CreateInBatches(values []*model.TemplateMessage, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t templateMessageDo) Save(values ...*model.TemplateMessage) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t templateMessageDo) First() (*model.TemplateMessage, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TemplateMessage), nil
	}
}

func (t templateMessageDo) Take() (*model.TemplateMessage, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TemplateMessage), nil
	}
}

func (t templateMessageDo) Last() (*model.TemplateMessage, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TemplateMessage), nil
	}
}

func (t templateMessageDo) Find() ([]*model.TemplateMessage, error) {
	result, err := t.DO.Find()
	return result.([]*model.TemplateMessage), err
}

func (t templateMessageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TemplateMessage, err error) {
	buf := make([]*model.TemplateMessage, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t templateMessageDo) FindInBatches(result *[]*model.TemplateMessage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t templateMessageDo) Attrs(attrs ...field.AssignExpr) ITemplateMessageDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t templateMessageDo) Assign(attrs ...field.AssignExpr) ITemplateMessageDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t templateMessageDo) Joins(fields ...field.RelationField) ITemplateMessageDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t templateMessageDo) Preload(fields ...field.RelationField) ITemplateMessageDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t templateMessageDo) FirstOrInit() (*model.TemplateMessage, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TemplateMessage), nil
	}
}

func (t templateMessageDo) FirstOrCreate() (*model.TemplateMessage, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TemplateMessage), nil
	}
}

func (t templateMessageDo) FindByPage(offset int, limit int) (result []*model.TemplateMessage, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t templateMessageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t templateMessageDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t templateMessageDo) Delete(models ...*model.TemplateMessage) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *templateMessageDo) withDO(do gen.Dao) *templateMessageDo {
	t.DO = *do.(*gen.DO)
	return t
}
