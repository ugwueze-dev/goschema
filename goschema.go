package goschema

import (
	_ "github.com/go-sql-driver/mysql"
)

// GoSchema
type GoSchema struct {
	tables []*Table
}

func New() *GoSchema {
	return &GoSchema{}
}

// Table creates a new database table
func (g *GoSchema) Table(tableName string, cb TableFunc) {
	table := newTable(tableName)
	g.tables = append(g.tables, table)
	cb(table)
}

func (g *GoSchema) GenerateSchema() string {
	return g.buildSchema()
}
