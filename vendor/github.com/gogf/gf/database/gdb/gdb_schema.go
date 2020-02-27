// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gdb

// Schema is a schema object from which it can then create a Model.
type Schema struct {
	db     DB
	tx     *TX
	schema string
}

// Schema creates and returns a schema.
func (bs *dbBase) Schema(schema string) *Schema {
	return &Schema{
		db:     bs.db,
		schema: schema,
	}
}

// Schema creates and returns a initialization model from schema,
// from which it can then create a Model.
func (tx *TX) Schema(schema string) *Schema {
	return &Schema{
		tx:     tx,
		schema: schema,
	}
}

// Table creates and returns a new ORM model.
// The parameter <tables> can be more than one table names, like :
// "user", "user u", "user, user_detail", "user u, user_detail ud"
func (s *Schema) Table(table string) *Model {
	var m *Model
	if s.tx != nil {
		m = s.tx.Table(table)
	} else {
		m = s.db.Table(table)
	}
	m.schema = s.schema
	return m
}

// Model is alias of dbBase.Table.
// See dbBase.Table.
func (s *Schema) Model(table string) *Model {
	return s.Table(table)
}
