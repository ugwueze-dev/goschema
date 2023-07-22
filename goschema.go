package goschema

import (
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type GoSchema struct {
	tables []*Table
}

func New() *GoSchema {
	return &GoSchema{}
}

func (g *GoSchema) Table(tableName string, cb TableFunc) {
	table := newTable(tableName)
	g.tables = append(g.tables, table)
	cb(table)
}

func (g *GoSchema) GenerateSchema() string {
	b := &strings.Builder{}

	for _, table := range g.tables {
		// TODO handle table references
		table.generateSchema(b)
	}

	return b.String()
}

/**
func (g *GoSchema) Table(tableName string, cb func(table *Table)) {
	table := newTable(tableName)
	g.tables = append(g.tables, table)
	cb(table)
}

func (g *GoSchema) Build() string {
	return g.build()
}
**/
