// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
// Note:
// 1. It needs manually import: _ "github.com/lib/pq"
// 2. It does not support Save/Replace features.
// 3. It does not support LastInsertId.

package gdb

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gogf/gf/text/gregex"
)

type dbPgsql struct {
	*dbBase
}

func (db *dbPgsql) Open(config *ConfigNode) (*sql.DB, error) {
	var source string
	if config.LinkInfo != "" {
		source = config.LinkInfo
	} else {
		source = fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			config.User, config.Pass, config.Host, config.Port, config.Name,
		)
	}
	if db, err := sql.Open("postgres", source); err == nil {
		return db, nil
	} else {
		return nil, err
	}
}

func (db *dbPgsql) getChars() (charLeft string, charRight string) {
	return "\"", "\""
}

func (db *dbPgsql) handleSqlBeforeExec(sql string) string {
	index := 0
	sql, _ = gregex.ReplaceStringFunc("\\?", sql, func(s string) string {
		index++
		return fmt.Sprintf("$%d", index)
	})
	sql, _ = gregex.ReplaceString(` LIMIT (\d+),\s*(\d+)`, ` LIMIT $1 OFFSET $2`, sql)
	return sql
}

// TODO
func (db *dbPgsql) Tables(schema ...string) (tables []string, err error) {
	return
}

func (db *dbPgsql) TableFields(table string, schema ...string) (fields map[string]*TableField, err error) {
	table, _ = gregex.ReplaceString("\"", "", table)
	checkSchema := db.schema.Val()
	if len(schema) > 0 && schema[0] != "" {
		checkSchema = schema[0]
	}
	v := db.cache.GetOrSetFunc(
		fmt.Sprintf(`pgsql_table_fields_%s_%s`, table, checkSchema), func() interface{} {
			var result Result
			var link *sql.DB
			link, err = db.getSlave(checkSchema)
			if err != nil {
				return nil
			}
			result, err = db.doGetAll(link, fmt.Sprintf(`
			SELECT a.attname AS field, t.typname AS type FROM pg_class c, pg_attribute a 
	        LEFT OUTER JOIN pg_description b ON a.attrelid=b.objoid AND a.attnum = b.objsubid,pg_type t
	        WHERE c.relname = '%s' and a.attnum > 0 and a.attrelid = c.oid and a.atttypid = t.oid 
			ORDER BY a.attnum`, strings.ToLower(table)))
			if err != nil {
				return nil
			}

			fields = make(map[string]*TableField)
			for i, m := range result {
				fields[m["field"].String()] = &TableField{
					Index: i,
					Name:  m["field"].String(),
					Type:  m["type"].String(),
				}
			}
			return fields
		}, 0)
	if err == nil {
		fields = v.(map[string]*TableField)
	}
	return
}
