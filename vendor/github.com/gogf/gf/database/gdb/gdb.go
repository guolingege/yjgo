// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gdb provides ORM features for popular relationship databases.
package gdb

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gogf/gf/os/glog"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/util/grand"
)

// DB is the interface for ORM operations.
type DB interface {
	// Open creates a raw connection object for database with given node configuration.
	// Note that it is not recommended using the this function manually.
	Open(config *ConfigNode) (*sql.DB, error)

	// Query APIs.
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(sql string, args ...interface{}) (sql.Result, error)
	Prepare(sql string, execOnMaster ...bool) (*sql.Stmt, error)

	// Internal APIs for CURD, which can be overwrote for custom CURD implements.
	doQuery(link dbLink, query string, args ...interface{}) (rows *sql.Rows, err error)
	doGetAll(link dbLink, query string, args ...interface{}) (result Result, err error)
	doExec(link dbLink, query string, args ...interface{}) (result sql.Result, err error)
	doPrepare(link dbLink, query string) (*sql.Stmt, error)
	doInsert(link dbLink, table string, data interface{}, option int, batch ...int) (result sql.Result, err error)
	doBatchInsert(link dbLink, table string, list interface{}, option int, batch ...int) (result sql.Result, err error)
	doUpdate(link dbLink, table string, data interface{}, condition string, args ...interface{}) (result sql.Result, err error)
	doDelete(link dbLink, table string, condition string, args ...interface{}) (result sql.Result, err error)

	// Query APIs for convenience purpose.
	GetAll(query string, args ...interface{}) (Result, error)
	GetOne(query string, args ...interface{}) (Record, error)
	GetValue(query string, args ...interface{}) (Value, error)
	GetCount(query string, args ...interface{}) (int, error)
	GetStruct(objPointer interface{}, query string, args ...interface{}) error
	GetStructs(objPointerSlice interface{}, query string, args ...interface{}) error
	GetScan(objPointer interface{}, query string, args ...interface{}) error

	// Master/Slave support.
	Master() (*sql.DB, error)
	Slave() (*sql.DB, error)

	// Ping.
	PingMaster() error
	PingSlave() error

	// Transaction.
	Begin() (*TX, error)

	Insert(table string, data interface{}, batch ...int) (sql.Result, error)
	Replace(table string, data interface{}, batch ...int) (sql.Result, error)
	Save(table string, data interface{}, batch ...int) (sql.Result, error)

	BatchInsert(table string, list interface{}, batch ...int) (sql.Result, error)
	BatchReplace(table string, list interface{}, batch ...int) (sql.Result, error)
	BatchSave(table string, list interface{}, batch ...int) (sql.Result, error)

	Update(table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error)
	Delete(table string, condition interface{}, args ...interface{}) (sql.Result, error)

	// Create model.
	From(tables string) *Model
	Table(tables string) *Model
	Schema(schema string) *Schema

	// Configuration methods.
	SetDebug(debug bool)
	SetSchema(schema string)
	SetLogger(logger *glog.Logger)
	GetLogger() *glog.Logger
	SetMaxIdleConnCount(n int)
	SetMaxOpenConnCount(n int)
	SetMaxConnLifetime(d time.Duration)
	Tables(schema ...string) (tables []string, err error)
	TableFields(table string, schema ...string) (map[string]*TableField, error)

	// Internal methods.
	getCache() *gcache.Cache
	getChars() (charLeft string, charRight string)
	getDebug() bool
	getPrefix() string
	getMaster(schema ...string) (*sql.DB, error)
	getSlave(schema ...string) (*sql.DB, error)
	quoteWord(s string) string
	quoteString(s string) string
	handleTableName(table string) string
	filterFields(schema, table string, data map[string]interface{}) map[string]interface{}
	convertValue(fieldValue []byte, fieldType string) interface{}
	rowsToResult(rows *sql.Rows) (Result, error)
	handleSqlBeforeExec(sql string) string
}

// dbLink is a common database function wrapper interface for internal usage.
type dbLink interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(sql string, args ...interface{}) (sql.Result, error)
	Prepare(sql string) (*sql.Stmt, error)
}

// dbBase is the base struct for database management.
type dbBase struct {
	db               DB            // DB interface object.
	group            string        // Configuration group name.
	debug            *gtype.Bool   // Enable debug mode for the database.
	cache            *gcache.Cache // Cache manager.
	schema           *gtype.String // Custom schema for this object.
	prefix           string        // Table prefix.
	logger           *glog.Logger  // Logger.
	maxIdleConnCount int           // Max idle connection count.
	maxOpenConnCount int           // Max open connection count.
	maxConnLifetime  time.Duration // Max TTL for a connection.
}

// Sql is the sql recording struct.
type Sql struct {
	Sql    string        // SQL string(may contain reserved char '?').
	Args   []interface{} // Arguments for this sql.
	Format string        // Formatted sql which contains arguments in the sql.
	Error  error         // Execution result.
	Start  int64         // Start execution timestamp in milliseconds.
	End    int64         // End execution timestamp in milliseconds.
}

// 表字段结构信息
type TableField struct {
	Index   int         // 用于字段排序(因为map类型是无序的)
	Name    string      // 字段名称
	Type    string      // 字段类型
	Null    bool        // 是否可为null
	Key     string      // 索引信息
	Default interface{} // 默认值
	Extra   string      // 其他信息
	Comment string      // 字段描述
}

// 返回数据表记录值
type Value = *gvar.Var

// 返回数据表记录Map
type Record map[string]Value

// 返回数据表记录List
type Result []Record

// 关联数组，绑定一条数据表记录(使用别名)
type Map = map[string]interface{}

// 关联数组列表(索引从0开始的数组)，绑定多条记录(使用别名)
type List = []Map

const (
	gINSERT_OPTION_DEFAULT      = 0
	gINSERT_OPTION_REPLACE      = 1
	gINSERT_OPTION_SAVE         = 2
	gINSERT_OPTION_IGNORE       = 3
	gDEFAULT_BATCH_NUM          = 10 // Per count for batch insert/replace/save
	gDEFAULT_CONN_MAX_LIFE_TIME = 30 // Max life time for per connection in pool.
)

var (
	// Instance map.
	instances = gmap.NewStrAnyMap(true)
)

// New creates and returns an ORM object with global configurations.
// The parameter <name> specifies the configuration group name,
// which is DEFAULT_GROUP_NAME in default.
func New(name ...string) (db DB, err error) {
	group := configs.defaultGroup
	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}
	configs.RLock()
	defer configs.RUnlock()

	if len(configs.config) < 1 {
		return nil, errors.New("empty database configuration")
	}
	if _, ok := configs.config[group]; ok {
		if node, err := getConfigNodeByGroup(group, true); err == nil {
			base := &dbBase{
				group:           group,
				debug:           gtype.NewBool(),
				cache:           gcache.New(),
				schema:          gtype.NewString(),
				logger:          glog.New(),
				prefix:          node.Prefix,
				maxConnLifetime: gDEFAULT_CONN_MAX_LIFE_TIME,
			}
			switch node.Type {
			case "mysql":
				base.db = &dbMysql{dbBase: base}
			case "pgsql":
				base.db = &dbPgsql{dbBase: base}
			case "mssql":
				base.db = &dbMssql{dbBase: base}
			case "sqlite":
				base.db = &dbSqlite{dbBase: base}
			case "oracle":
				base.db = &dbOracle{dbBase: base}
			default:
				return nil, errors.New(fmt.Sprintf(`unsupported database type "%s"`, node.Type))
			}
			return base.db, nil
		} else {
			return nil, err
		}
	} else {
		return nil, errors.New(fmt.Sprintf(`database configuration node "%s" is not found`, group))
	}
}

// Instance returns an instance for DB operations.
// The parameter <name> specifies the configuration group name,
// which is DEFAULT_GROUP_NAME in default.
func Instance(name ...string) (db DB, err error) {
	group := configs.defaultGroup
	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}
	v := instances.GetOrSetFuncLock(group, func() interface{} {
		db, err = New(group)
		return db
	})
	if v != nil {
		return v.(DB), nil
	}
	return
}

// 获取指定数据库角色的一个配置项，内部根据权重计算负载均衡
func getConfigNodeByGroup(group string, master bool) (*ConfigNode, error) {
	if list, ok := configs.config[group]; ok {
		// 将master, slave集群列表拆分出来
		masterList := make(ConfigGroup, 0)
		slaveList := make(ConfigGroup, 0)
		for i := 0; i < len(list); i++ {
			if list[i].Role == "slave" {
				slaveList = append(slaveList, list[i])
			} else {
				masterList = append(masterList, list[i])
			}
		}
		if len(masterList) < 1 {
			return nil, errors.New("at least one master node configuration's need to make sense")
		}
		if len(slaveList) < 1 {
			slaveList = masterList
		}
		if master {
			return getConfigNodeByWeight(masterList), nil
		} else {
			return getConfigNodeByWeight(slaveList), nil
		}
	} else {
		return nil, errors.New(fmt.Sprintf("empty database configuration for item name '%s'", group))
	}
}

// 按照负载均衡算法(优先级配置)从数据库集群中选择一个配置节点出来使用
// 算法说明举例，
// 1、假如2个节点的priority都是1，那么随机大小范围为[0, 199]；
// 2、那么节点1的权重范围为[0, 99]，节点2的权重范围为[100, 199]，比例为1:1；
// 3、假如计算出的随机数为99;
// 4、那么选择的配置为节点1;
func getConfigNodeByWeight(cg ConfigGroup) *ConfigNode {
	if len(cg) < 2 {
		return &cg[0]
	}
	var total int
	for i := 0; i < len(cg); i++ {
		total += cg[i].Weight * 100
	}
	// 如果total为0表示所有连接都没有配置priority属性，那么默认都是1
	if total == 0 {
		for i := 0; i < len(cg); i++ {
			cg[i].Weight = 1
			total += cg[i].Weight * 100
		}
	}
	// 不能取到末尾的边界点
	r := grand.N(0, total)
	if r > 0 {
		r -= 1
	}
	min := 0
	max := 0
	for i := 0; i < len(cg); i++ {
		max = min + cg[i].Weight*100
		//fmt.Printf("r: %d, min: %d, max: %d\n", r, min, max)
		if r >= min && r < max {
			return &cg[i]
		} else {
			min = max
		}
	}
	return nil
}

// getSqlDb retrieves and returns a underlying database connection object.
// The parameter <master> specifies whether retrieves master node connection if
// master-slave nodes are configured.
func (bs *dbBase) getSqlDb(master bool, schema ...string) (sqlDb *sql.DB, err error) {
	// Load balance.
	node, err := getConfigNodeByGroup(bs.group, master)
	if err != nil {
		return nil, err
	}
	// Default value checks.
	if node.Charset == "" {
		node.Charset = "utf8"
	}
	// Changes the schema.
	nodeSchema := bs.schema.Val()
	if len(schema) > 0 && schema[0] != "" {
		nodeSchema = schema[0]
	}
	if nodeSchema != "" {
		// Value copy.
		n := *node
		n.Name = nodeSchema
		node = &n
	}
	// Cache the underlying connection object by node.
	v := bs.cache.GetOrSetFuncLock(node.String(), func() interface{} {
		sqlDb, err = bs.db.Open(node)
		if err != nil {
			return nil
		}
		if bs.maxIdleConnCount > 0 {
			sqlDb.SetMaxIdleConns(bs.maxIdleConnCount)
		} else if node.MaxIdleConnCount > 0 {
			sqlDb.SetMaxIdleConns(node.MaxIdleConnCount)
		}

		if bs.maxOpenConnCount > 0 {
			sqlDb.SetMaxOpenConns(bs.maxOpenConnCount)
		} else if node.MaxOpenConnCount > 0 {
			sqlDb.SetMaxOpenConns(node.MaxOpenConnCount)
		}

		if bs.maxConnLifetime > 0 {
			sqlDb.SetConnMaxLifetime(bs.maxConnLifetime * time.Second)
		} else if node.MaxConnLifetime > 0 {
			sqlDb.SetConnMaxLifetime(node.MaxConnLifetime * time.Second)
		}
		return sqlDb
	}, 0)
	if v != nil && sqlDb == nil {
		sqlDb = v.(*sql.DB)
	}
	if node.Debug {
		bs.db.SetDebug(node.Debug)
	}
	return
}

// SetSchema changes the schema for this database connection object.
// Importantly note that when schema configuration changed for the database,
// it affects all operations on the database object in the future.
func (bs *dbBase) SetSchema(schema string) {
	bs.schema.Set(schema)
}

// Master creates and returns a connection from master node if master-slave configured.
// It returns the default connection if master-slave not configured.
func (bs *dbBase) Master() (*sql.DB, error) {
	return bs.getSqlDb(true, bs.schema.Val())
}

// Slave creates and returns a connection from slave node if master-slave configured.
// It returns the default connection if master-slave not configured.
func (bs *dbBase) Slave() (*sql.DB, error) {
	return bs.getSqlDb(false, bs.schema.Val())
}

// getMaster acts like function Master but with additional <schema> parameter specifying
// the schema for the connection. It is defined for internal usage.
// Also see Master.
func (bs *dbBase) getMaster(schema ...string) (*sql.DB, error) {
	return bs.getSqlDb(true, schema...)
}

// getSlave acts like function Slave but with additional <schema> parameter specifying
// the schema for the connection. It is defined for internal usage.
// Also see Slave.
func (bs *dbBase) getSlave(schema ...string) (*sql.DB, error) {
	return bs.getSqlDb(false, schema...)
}
