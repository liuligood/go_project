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

func newUserLevel(db *gorm.DB, opts ...gen.DOOption) userLevel {
	_userLevel := userLevel{}

	_userLevel.userLevelDo.UseDB(db, opts...)
	_userLevel.userLevelDo.UseModel(&model.UserLevel{})

	tableName := _userLevel.userLevelDo.TableName()
	_userLevel.ALL = field.NewAsterisk(tableName)
	_userLevel.ID = field.NewInt64(tableName, "id")
	_userLevel.UID = field.NewInt64(tableName, "uid")
	_userLevel.LevelID = field.NewInt64(tableName, "level_id")
	_userLevel.Grade = field.NewInt64(tableName, "grade")
	_userLevel.Status = field.NewInt64(tableName, "status")
	_userLevel.Mark = field.NewString(tableName, "mark")
	_userLevel.Remind = field.NewInt64(tableName, "remind")
	_userLevel.Discount = field.NewInt64(tableName, "discount")
	_userLevel.ExpiredTime = field.NewInt64(tableName, "expired_time")
	_userLevel.CreatedAt = field.NewInt64(tableName, "created_at")
	_userLevel.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_userLevel.DeletedAt = field.NewField(tableName, "deleted_at")

	_userLevel.fillFieldMap()

	return _userLevel
}

// userLevel 用户等级记录表
type userLevel struct {
	userLevelDo userLevelDo

	ALL         field.Asterisk
	ID          field.Int64
	UID         field.Int64  // 用户uid
	LevelID     field.Int64  // 等级vip
	Grade       field.Int64  // 会员等级
	Status      field.Int64  // 0:禁止,1:正常
	Mark        field.String // 备注
	Remind      field.Int64  // 是否已通知
	Discount    field.Int64  // 享受折扣
	ExpiredTime field.Int64
	CreatedAt   field.Int64 // 创建时间
	UpdatedAt   field.Int64 // 修改时间
	DeletedAt   field.Field // 是否删除

	fieldMap map[string]field.Expr
}

func (u userLevel) Table(newTableName string) *userLevel {
	u.userLevelDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userLevel) As(alias string) *userLevel {
	u.userLevelDo.DO = *(u.userLevelDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userLevel) updateTableName(table string) *userLevel {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.UID = field.NewInt64(table, "uid")
	u.LevelID = field.NewInt64(table, "level_id")
	u.Grade = field.NewInt64(table, "grade")
	u.Status = field.NewInt64(table, "status")
	u.Mark = field.NewString(table, "mark")
	u.Remind = field.NewInt64(table, "remind")
	u.Discount = field.NewInt64(table, "discount")
	u.ExpiredTime = field.NewInt64(table, "expired_time")
	u.CreatedAt = field.NewInt64(table, "created_at")
	u.UpdatedAt = field.NewInt64(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userLevel) WithContext(ctx context.Context) IUserLevelDo {
	return u.userLevelDo.WithContext(ctx)
}

func (u userLevel) TableName() string { return u.userLevelDo.TableName() }

func (u userLevel) Alias() string { return u.userLevelDo.Alias() }

func (u userLevel) Columns(cols ...field.Expr) gen.Columns { return u.userLevelDo.Columns(cols...) }

func (u *userLevel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userLevel) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 12)
	u.fieldMap["id"] = u.ID
	u.fieldMap["uid"] = u.UID
	u.fieldMap["level_id"] = u.LevelID
	u.fieldMap["grade"] = u.Grade
	u.fieldMap["status"] = u.Status
	u.fieldMap["mark"] = u.Mark
	u.fieldMap["remind"] = u.Remind
	u.fieldMap["discount"] = u.Discount
	u.fieldMap["expired_time"] = u.ExpiredTime
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u userLevel) clone(db *gorm.DB) userLevel {
	u.userLevelDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userLevel) replaceDB(db *gorm.DB) userLevel {
	u.userLevelDo.ReplaceDB(db)
	return u
}

type userLevelDo struct{ gen.DO }

type IUserLevelDo interface {
	gen.SubQuery
	Debug() IUserLevelDo
	WithContext(ctx context.Context) IUserLevelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserLevelDo
	WriteDB() IUserLevelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserLevelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserLevelDo
	Not(conds ...gen.Condition) IUserLevelDo
	Or(conds ...gen.Condition) IUserLevelDo
	Select(conds ...field.Expr) IUserLevelDo
	Where(conds ...gen.Condition) IUserLevelDo
	Order(conds ...field.Expr) IUserLevelDo
	Distinct(cols ...field.Expr) IUserLevelDo
	Omit(cols ...field.Expr) IUserLevelDo
	Join(table schema.Tabler, on ...field.Expr) IUserLevelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserLevelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserLevelDo
	Group(cols ...field.Expr) IUserLevelDo
	Having(conds ...gen.Condition) IUserLevelDo
	Limit(limit int) IUserLevelDo
	Offset(offset int) IUserLevelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserLevelDo
	Unscoped() IUserLevelDo
	Create(values ...*model.UserLevel) error
	CreateInBatches(values []*model.UserLevel, batchSize int) error
	Save(values ...*model.UserLevel) error
	First() (*model.UserLevel, error)
	Take() (*model.UserLevel, error)
	Last() (*model.UserLevel, error)
	Find() ([]*model.UserLevel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserLevel, err error)
	FindInBatches(result *[]*model.UserLevel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserLevel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserLevelDo
	Assign(attrs ...field.AssignExpr) IUserLevelDo
	Joins(fields ...field.RelationField) IUserLevelDo
	Preload(fields ...field.RelationField) IUserLevelDo
	FirstOrInit() (*model.UserLevel, error)
	FirstOrCreate() (*model.UserLevel, error)
	FindByPage(offset int, limit int) (result []*model.UserLevel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserLevelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userLevelDo) Debug() IUserLevelDo {
	return u.withDO(u.DO.Debug())
}

func (u userLevelDo) WithContext(ctx context.Context) IUserLevelDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userLevelDo) ReadDB() IUserLevelDo {
	return u.Clauses(dbresolver.Read)
}

func (u userLevelDo) WriteDB() IUserLevelDo {
	return u.Clauses(dbresolver.Write)
}

func (u userLevelDo) Session(config *gorm.Session) IUserLevelDo {
	return u.withDO(u.DO.Session(config))
}

func (u userLevelDo) Clauses(conds ...clause.Expression) IUserLevelDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userLevelDo) Returning(value interface{}, columns ...string) IUserLevelDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userLevelDo) Not(conds ...gen.Condition) IUserLevelDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userLevelDo) Or(conds ...gen.Condition) IUserLevelDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userLevelDo) Select(conds ...field.Expr) IUserLevelDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userLevelDo) Where(conds ...gen.Condition) IUserLevelDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userLevelDo) Order(conds ...field.Expr) IUserLevelDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userLevelDo) Distinct(cols ...field.Expr) IUserLevelDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userLevelDo) Omit(cols ...field.Expr) IUserLevelDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userLevelDo) Join(table schema.Tabler, on ...field.Expr) IUserLevelDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userLevelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserLevelDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userLevelDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserLevelDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userLevelDo) Group(cols ...field.Expr) IUserLevelDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userLevelDo) Having(conds ...gen.Condition) IUserLevelDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userLevelDo) Limit(limit int) IUserLevelDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userLevelDo) Offset(offset int) IUserLevelDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userLevelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserLevelDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userLevelDo) Unscoped() IUserLevelDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userLevelDo) Create(values ...*model.UserLevel) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userLevelDo) CreateInBatches(values []*model.UserLevel, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userLevelDo) Save(values ...*model.UserLevel) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userLevelDo) First() (*model.UserLevel, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserLevel), nil
	}
}

func (u userLevelDo) Take() (*model.UserLevel, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserLevel), nil
	}
}

func (u userLevelDo) Last() (*model.UserLevel, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserLevel), nil
	}
}

func (u userLevelDo) Find() ([]*model.UserLevel, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserLevel), err
}

func (u userLevelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserLevel, err error) {
	buf := make([]*model.UserLevel, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userLevelDo) FindInBatches(result *[]*model.UserLevel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userLevelDo) Attrs(attrs ...field.AssignExpr) IUserLevelDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userLevelDo) Assign(attrs ...field.AssignExpr) IUserLevelDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userLevelDo) Joins(fields ...field.RelationField) IUserLevelDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userLevelDo) Preload(fields ...field.RelationField) IUserLevelDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userLevelDo) FirstOrInit() (*model.UserLevel, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserLevel), nil
	}
}

func (u userLevelDo) FirstOrCreate() (*model.UserLevel, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserLevel), nil
	}
}

func (u userLevelDo) FindByPage(offset int, limit int) (result []*model.UserLevel, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userLevelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userLevelDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userLevelDo) Delete(models ...*model.UserLevel) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userLevelDo) withDO(do gen.Dao) *userLevelDo {
	u.DO = *do.(*gen.DO)
	return u
}
