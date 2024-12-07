package igorm

import (
	"crmeb_go/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type DB struct {
	config config.Conf
}

func NewDB(conf config.Conf) *DB {
	return &DB{
		config: conf,
	}
}

func Gorm(db *DB) *gorm.DB {
	switch db.config.System.DbType {
	case "mysql":
		return db.GormMysql()
	case "pgsql":
		return db.GormPgSql()
	//case "oracle":
	//	return GormOracle()
	//case "mssql":
	//	return GormMssql()
	//case "sqlite":
	//	return GormSqlite()
	default:
		return db.GormMysql()
	}
}

func (db *DB) GormMysql() *gorm.DB {
	m := db.config.Mysql
	if m.Dbname == "" {
		return nil
	}

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), db.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)

		return db
	}

}

// GormPgSql 初始化 Postgresql 数据库
func (db *DB) GormPgSql() *gorm.DB {
	p := db.config.Pgsql
	if p.Dbname == "" {
		return nil
	}

	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), db.Config(p.Prefix, p.Singular)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}

// GormPgSqlByConfig 初始化 Postgresql 数据库 通过参数
func (db *DB) GormPgSqlByConfig(p config.Pgsql) *gorm.DB {
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), db.Config(p.Prefix, p.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}

func (db *DB) Config(prefix string, singular bool) *gorm.Config {
	var general config.GeneralDB
	switch db.config.System.DbType {
	case "mysql":
		general = db.config.Mysql.GeneralDB
	case "pgsql":
		general = db.config.Pgsql.GeneralDB
	//case "oracle":
	//	general = db.config.Pgsql.GeneralDB
	//case "sqlite":
	//	general = db.config.Pgsql.GeneralDB
	//case "mssql":
	//	general = db.config.Pgsql.GeneralDB
	default:
		general = db.config.Mysql.GeneralDB
	}

	return &gorm.Config{
		Logger: logger.New(NewWriter(general, log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      general.LogLevel(),
			Colorful:      true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
