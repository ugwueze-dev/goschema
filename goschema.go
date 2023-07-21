package goschema

import (
	_ "github.com/go-sql-driver/mysql"
)

type GoSchema struct {
	tables []*Table
}

func New() *GoSchema {
	return &GoSchema{}
}

func (g *GoSchema) Table(tableName string, cb func(table *Table)) {
	table := newTable(tableName)
	g.tables = append(g.tables, table)
	cb(table)
}

func (g *GoSchema) Build() string {
	return g.build()
}
